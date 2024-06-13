//go:build linux && with_pcap
// +build linux,with_pcap

package protodec

import (
	"errors"
	"fmt"
	"io"
	"net"
	"testing"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/internal/plugins/externals/ebpf/internal/l7flow/comm"
)

func TestMySQLProto(t *testing.T) {
	type filePort struct {
		fp          string
		localIP     net.IP
		foreignIP   net.IP
		localPort   uint16
		foreignPort uint16
	}

	filePorts := []filePort{
		{
			fp:          "./pcapdata/mysql.pcapng",
			localIP:     net.ParseIP("127.0.0.1"),
			foreignIP:   net.ParseIP("127.0.0.1"),
			localPort:   uint16(51968),
			foreignPort: uint16(3306),
		},
		{
			fp:          "./pcapdata/mysql_2.pcapng",
			localIP:     net.ParseIP("127.0.0.1"),
			foreignIP:   net.ParseIP("127.0.0.1"),
			localPort:   uint16(55716),
			foreignPort: uint16(3306),
		},
		{
			fp:          "./pcapdata/mysql_3.pcapng",
			localIP:     net.ParseIP("127.0.0.1"),
			foreignIP:   net.ParseIP("127.0.0.1"),
			localPort:   uint16(55045),
			foreignPort: uint16(3306),
		},
	}

	for _, fp := range filePorts {
		t.Run(fp.fp, func(t *testing.T) {
			stream := &netStream{}
			if err := stream.Open(fp.fp); err != nil {
				t.Fatal(err)
			}
			defer stream.Close()

			cases := [][2]any{}
			for {
				if dir, netdata, ok, err := stream.Get(fp.localIP, fp.foreignIP, fp.localPort, fp.foreignPort); err != nil {
					if !errors.Is(err, io.EOF) {
						t.Fatal(err)
					}
					break
				} else if ok {
					cases = append(cases, [2]any{dir, netdata})
				}
			}

			var impl ProtoDecPipe
			for _, c := range cases {
				netdata := c[1].(*comm.NetwrkData)
				_, impl, _ = MysqlProtoDetect(netdata.Payload, netdata.ActSize)
				if impl != nil {
					break
				}
			}

			if impl == nil {
				t.Fatal("not found")
			}

			for _, c := range cases {
				netdata := c[1].(*comm.NetwrkData)
				dir := c[0].(comm.NICDirection)

				if dir == comm.NICDIngress {
					netdata.Fn = comm.FnSysRead
				} else {
					netdata.Fn = comm.FnSysWrite
				}
				if len(netdata.Payload) == 0 {
					continue
				}
				impl.Decode(dir, netdata, 0, nil)
			}

			// impl.ConnClose()
			for _, v := range impl.Export(true) {
				t.Log(v.KVs.Pretty())
			}
		})
	}
}

type netStream struct {
	handle *pcap.Handle
}

func (s *netStream) Open(filePath string) error {
	h, err := pcap.OpenOffline(filePath)
	if err != nil {
		return err
	}
	s.handle = h
	return nil
}

func (s *netStream) Get(localIP, foreignIP net.IP, localPort, foreignPort uint16) (comm.NICDirection, *comm.NetwrkData, bool, error) {
	if s.handle == nil {
		return comm.NICDUnknown, nil, false, fmt.Errorf("handle is nil")
	}

	data, _, err := s.handle.ReadPacketData()
	if err != nil {
		return comm.NICDUnknown, nil, false, err
	}
	if len(data) > 4 {
		if data[0] == 0x02 && data[1] == 0x00 &&
			data[2] == 0x00 && data[3] == 0x00 {
			// linktype is null, such as pcap packet start with Null/Loopback
			data = append(make([]byte, 14), data[4:]...)
			data[12] = 0x08
		}
	}

	decoder := NewPktDecoder()
	layerLi := make([]gopacket.LayerType, 0, 10)

	_ = decoder.pktDecode.DecodeLayers(data, &layerLi)
	srcIP := decoder.ipv4.SrcIP
	srcPort := decoder.tcp.SrcPort
	dstIP := decoder.ipv4.DstIP
	dstPort := decoder.tcp.DstPort

	data = decoder.tcp.Payload

	// 对于一个 tcp 连接的两端来说，如果两端位于不同网络（命名）空间或主机，
	// 使用 wireshark 或创建 raw socket 在两个网卡抓取的网络报文（如果中间设备没有变更数据包）通常一致；
	// 但两端使用 loopback，且使用 raw socket 抓包可能会出现报文重复的情况，需要先根据 packet type 过滤后再保存为 pacp 文件，
	// 参考 https://github.com/GuanceCloud/gopacket/commit/0f93dc491d52f704c7e9705004109ddc3ffef887 和
	// https://man7.org/linux/man-pages/man7/packet.7.html 的 sll_pkttype

	// 数据包的 src 代表发送者 ip，dst 代表接收者 ip，local 代表被采集进程/网卡的 ip，foreign 代表网络对端进程的 ip
	// direction 代表是 local 发送数据还是接收数据

	// pkt_src| pkt_dst| --> local| foreigin| direction
	// ip1      ip2      --> ip1    ip2      tx
	// ip2      ip1      --> ip1    ip2      rx
	//
	// ----------- 以下切换 ip2 为被观测进程的地址 -----
	//
	// ip1      ip2      --> ip2    ip1      rx
	// ip2      ip1      --> ip2    ip1      tx
	//
	if srcIP.Equal(localIP) && srcPort == layers.TCPPort(localPort) &&
		dstIP.Equal(foreignIP) && dstPort == layers.TCPPort(foreignPort) {
		return comm.NICDEgress, wrapPacket(
			data, comm.NICDEgress, decoder.tcp.Seq), true, nil
	} else if srcIP.Equal(foreignIP) && srcPort == layers.TCPPort(foreignPort) &&
		dstIP.Equal(localIP) && dstPort == layers.TCPPort(localPort) {
		return comm.NICDIngress, wrapPacket(
			data, comm.NICDIngress, decoder.tcp.Seq), true, nil
	}

	return comm.NICDUnknown, nil, false, nil
}

func wrapPacket(buf []byte, dir comm.NICDirection, tcpSeq uint32) *comm.NetwrkData {
	netdata := &comm.NetwrkData{
		ActSize: len(buf),
		Payload: buf,
		TCPSeq:  tcpSeq,
	}
	return netdata
}

func (s *netStream) Close() {
	if s.handle != nil {
		s.handle.Close()
	}
}

type pktDecoder struct {
	pktDecode   *gopacket.DecodingLayerParser
	vxlanDecode *gopacket.DecodingLayerParser

	eth  *layers.Ethernet
	ipv4 *layers.IPv4
	ipv6 *layers.IPv6
	tcp  *layers.TCP
	udp  *layers.UDP

	// vxlan
	vxlan *layers.VXLAN
}

func NewPktDecoder() *pktDecoder {
	var eth layers.Ethernet
	var ipv4 layers.IPv4
	var ipv6 layers.IPv6
	var tcp layers.TCP
	var udp layers.UDP

	var vxlan layers.VXLAN

	l := []gopacket.DecodingLayer{
		&eth,
		&ipv4, &ipv6,
		&udp,
		&tcp,
	}

	vxlanLi := []gopacket.DecodingLayer{
		&vxlan,

		&eth,
		&ipv4, &ipv6,
		&udp,
		&tcp,
	}

	return &pktDecoder{
		gopacket.NewDecodingLayerParser(layers.LayerTypeEthernet, l...),
		gopacket.NewDecodingLayerParser(layers.LayerTypeVXLAN, vxlanLi...),
		&eth,
		&ipv4, &ipv6,
		&tcp,
		&udp,

		&vxlan,
	}
}
