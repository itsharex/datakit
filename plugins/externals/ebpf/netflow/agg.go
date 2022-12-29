//go:build (linux && amd64 && ebpf) || (linux && arm64 && ebpf)
// +build linux,amd64,ebpf linux,arm64,ebpf

package netflow

import (
	"math"
	"strconv"
	"time"

	client "github.com/influxdata/influxdb1-client/v2"
	"github.com/spf13/cast"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/externals/ebpf/k8sinfo"
)

type aggKey struct {
	sAddr string
	dAddr string

	sPort uint32
	dPort uint32

	sType string
	dType string

	pid int

	family    string
	transport string
	direction string
}

type aggValue struct {
	bytesRead    int64
	bytesWritten int64

	retransmits    int64
	rtt            []int64
	rttVar         []int64
	tcpClosed      int64
	tcpEstablished int64

	count int
}

func calLatency(l []int64) int64 {
	if len(l) == 0 {
		return 0
	} else {
		t := int64(0)
		for _, v := range l {
			t += v
		}
		return t / int64(len(l))
	}
}

func kv2point(key *aggKey, value *aggValue, pTime time.Time,
	addTags map[string]string, k8sNetInfo *k8sinfo.K8sNetInfo,
	pidMap map[int][2]string,
) (*client.Point, error) {
	tags := map[string]string{
		"family": key.family,

		"direction": key.direction,
		"transport": key.transport,

		"src_ip": key.sAddr,
		"dst_ip": key.dAddr,

		"src_ip_type": key.sType,
		"dst_ip_type": key.dType,

		"pid": strconv.FormatInt(int64(key.pid), 10),
	}

	if procName, ok := pidMap[key.pid]; ok {
		tags["process_name"] = procName[0]
	} else {
		tags["process_name"] = NoValue
	}

	if key.sPort == math.MaxUint32 {
		tags["src_port"] = "*"
	} else {
		tags["src_port"] = cast.ToString(key.sPort)
	}

	if key.dPort == math.MaxUint32 {
		tags["dst_port"] = "*"
	} else {
		tags["dst_port"] = cast.ToString(key.dPort)
	}

	if dnsRecord != nil {
		tags["dst_domain"] = dnsRecord.LookupAddr(key.dAddr)
	}

	for k, v := range addTags {
		if _, ok := tags[k]; !ok {
			tags[k] = v
		}
	}

	var fields map[string]any

	if key.transport == transportTCP {
		fields = map[string]any{
			"bytes_read":      value.bytesRead,
			"bytes_written":   value.bytesWritten,
			"retransmits":     value.retransmits,
			"rtt":             calLatency(value.rtt),
			"rtt_var":         calLatency(value.rttVar),
			"tcp_closed":      value.tcpClosed,
			"tcp_established": value.tcpEstablished,
		}
	} else {
		fields = map[string]any{
			"bytes_read":    value.bytesRead,
			"bytes_written": value.bytesWritten,
		}
	}

	tags = AddK8sTags2Map(k8sNetInfo, key.sAddr, key.dAddr,
		key.sPort, key.dPort, key.transport, tags)
	return client.NewPoint(srcNameM, tags, fields, pTime)
}

type FlowAgg struct {
	data map[aggKey]*aggValue
}

func (agg *FlowAgg) Len() int {
	return len(agg.data)
}

func (agg *FlowAgg) Append(info ConnectionInfo, stats ConnFullStats) error {
	if !ConnNotNeedToFilter(&info, &stats) {
		return nil
	}

	if agg.data == nil {
		agg.data = map[aggKey]*aggValue{}
	}

	var key aggKey

	// family
	isV6 := !ConnAddrIsIPv4(info.Meta)

	if info.Saddr[0] == 0 && info.Saddr[1] == 0 &&
		info.Daddr[0] == 0 && info.Daddr[1] == 0 {
		if info.Saddr[2] == 0xffff0000 && info.Daddr[2] == 0xffff0000 {
			isV6 = false
		} else if info.Saddr[2] == 0 && info.Daddr[2] == 0 &&
			info.Saddr[3] > 1 && info.Daddr[3] > 1 {
			isV6 = false
		}
	}

	// ip type
	if isV6 {
		key.sType = ConnIPv6Type(info.Saddr)
		key.dType = ConnIPv6Type(info.Daddr)
		key.family = "IPv6"
	} else {
		key.sType = ConnIPv4Type(info.Saddr[3])
		key.dType = ConnIPv4Type(info.Daddr[3])
		key.family = "IPv4"
	}

	// saddr, daddr
	key.sAddr = U32BEToIP(info.Saddr, isV6).String()
	key.dAddr = U32BEToIP(info.Daddr, isV6).String()

	// sport, dport
	key.sPort = info.Sport
	key.dPort = info.Dport

	// transport
	if ConnProtocolIsTCP(info.Meta) {
		key.transport = transportTCP
	} else {
		key.transport = transportUDP
	}

	// direction
	key.direction = ConnDirection2Str(stats.Stats.Direction)

	if IsIncomingFromK8s(k8sNetInfo, key.sAddr,
		key.sPort, key.transport) {
		key.direction = DirectionIncoming
	}

	switch key.direction {
	case DirectionOutgoing:
		if IsEphemeralPort(key.sPort) {
			key.sPort = math.MaxUint32
		}
	case DirectionIncoming:
		if IsEphemeralPort(key.dPort) {
			key.dPort = math.MaxUint32
		}
	}

	// pid
	key.pid = int(info.Pid)

	var value *aggValue
	// agg latency and count ++
	if v, ok := agg.data[key]; ok {
		v.count++
		value = v
	} else {
		value = &aggValue{
			count: 1,
		}
		agg.data[key] = value
	}

	value.bytesRead += int64(stats.Stats.RecvBytes)
	value.bytesWritten += int64(stats.Stats.SentBytes)

	if key.transport == transportTCP {
		value.rtt = append(value.rtt, int64(stats.TCPStats.Rtt))
		value.rttVar = append(value.rttVar, int64(stats.TCPStats.RttVar))
		value.retransmits += int64(stats.TCPStats.Retransmits)
		value.tcpClosed += stats.TotalClosed
		value.tcpEstablished += stats.TotalEstablished
	}

	return nil
}

func (agg *FlowAgg) ToPoint(tags map[string]string,
	k8sInfo *k8sinfo.K8sNetInfo, pidMap map[int][2]string,
) []*client.Point {
	var result []*client.Point

	pTime := time.Now()
	for k, v := range agg.data {
		if pt, err := kv2point(&k, v, pTime, tags, k8sInfo, pidMap); err != nil {
			l.Debug(err)
		} else {
			result = append(result, pt)
		}
	}

	return result
}

func (agg *FlowAgg) Clean() {
	agg.data = make(map[aggKey]*aggValue)
}
