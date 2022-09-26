// Unless explicitly stated otherwise all files in this repository are licensed
// under the MIT License.
// This product includes software developed at Guance Cloud (https://www.guance.com/).
// Copyright 2021-present Guance, Inc.

package datakit

func Enabled(name string) bool {
	if enabled, ok := allInputs[name]; !ok {
		return false // not exist
	} else {
		if EnableUncheckInputs {
			return true
		} else {
			return enabled
		}
	}
}

var (
	EnableUncheckInputs = false

	allInputs = map[string]bool{
		"activemqlog":            true,
		"ansible":                true,
		"apache":                 true,
		"awsbill":                true,
		"awscloudtrail":          true,
		"azure_monitor":          true,
		"baiduIndex":             true,
		"beats_output":           true,
		"binlog":                 true,
		"clickhousev1":           true,
		"cloudflare":             true,
		"cloudprober":            true,
		"confluence":             true,
		"consul":                 true,
		"container":              true,
		"containerd":             true,
		"coredns":                true,
		"cpu":                    true,
		"csvmetric":              true,
		"csvobject":              true,
		"ddtrace":                true,
		"dialtesting":            true,
		"disk":                   true,
		"diskio":                 true,
		"docker":                 true,
		"docker_containers":      true,
		"dockerlog":              true,
		"druid":                  true,
		"ebpf":                   true,
		"elasticsearch":          true,
		"envoy":                  true,
		"etcd":                   true,
		"expressjs":              true,
		"external":               true,
		"file_collector":         true,
		"flinkv1":                true,
		"fluentd":                true,
		"ginlog":                 true,
		"gitlab":                 true,
		"goruntime":              true,
		"harborMonitor":          true,
		"host_processes":         true,
		"hostdir":                true,
		"hostobject":             true,
		"httpPacket":             true,
		"httpProb":               true,
		"httpjson":               true,
		"httpstat":               true,
		"huaweiyunces":           true,
		"huaweiyunobject":        true,
		"iis":                    true,
		"influxdb":               true,
		"jaeger":                 true,
		"jenkins":                true,
		"jira":                   true,
		"jvm":                    true,
		"k8sobject":              true,
		"kafka":                  true,
		"kafkalog":               true,
		"kong":                   true,
		"kubernetes":             true,
		"lighttpd":               true,
		"logfwdserver":           true,
		"logging":                true,
		"logstreaming":           true,
		"mem":                    true,
		"memcached":              true,
		"mongodb":                true,
		"mongodb_oplog":          true,
		"mysql":                  true,
		"mysqlog":                true,
		"neo4j":                  true,
		"net":                    true,
		"netstat":                true,
		"nfsstat":                true,
		"nginx":                  true,
		"nginx_plus":             true,
		"nginx_plus_api":         true,
		"nginx_upstream_check":   true,
		"nginx_vts":              true,
		"nginxlog":               true,
		"nsq":                    true,
		"nvidia_smi":             true,
		"opentelemetry":          true,
		"oracle":                 true,
		"oraclemonitor":          true,
		"postgresql":             true,
		"postgresql_replication": true,
		"processes":              true,
		"profile":                true,
		"prom":                   true,
		"prom_remote_write":      true,
		"proxy":                  true,
		"puppetagent":            true,
		"pythond":                true,
		"rabbitmq":               true,
		"redis":                  true,
		"redislog":               true,
		"scanport":               true,
		"self":                   true,
		"sensors":                true,
		"skywalking":             true,
		"smart":                  true,
		"socket":                 true,
		"solr":                   true,
		"sqlserver":              true,
		"squid":                  true,
		"ssh":                    true,
		"statsd":                 true,
		"swap":                   true,
		"system":                 true,
		"systemd":                true,
		"tailf":                  true,
		"tdengine":               true,
		"tencentcms":             true,
		"tencentcost":            true,
		"tencentobject":          true,
		"tidb":                   true,
		"timezone":               true,
		"tomcat":                 true,
		"tracerouter":            true,
		"traefik":                true,
		"ucloud_monitor":         true,
		"wechatminiprogram":      true,
		"windows_event":          true,
		"wmi":                    true,
		"yarn":                   true,
		"zabbix":                 true,
		"zaplog":                 true,
		"rum":                    true,
		"zipkin":                 true,
	}
)
