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
		"cshark":     false,
		"mock":       false,
		"qyt_all":    false,
		"secureexec": false,
		"demo":       false,

		"activemqlog":            true,
		"aliyunactiontrail":      true,
		"aliyuncdn":              true,
		"aliyuncms":              true,
		"aliyuncost":             true,
		"aliyunddos":             true,
		"aliyunfc":               true,
		"aliyunlog":              true,
		"aliyunobject":           true,
		"aliyunprice":            true,
		"aliyunrdsslowlog":       true,
		"aliyunsecurity":         true,
		"ansible":                true,
		"awsbill":                true,
		"awscloudtrail":          true,
		"azure_monitor":          true,
		"baiduIndex":             true,
		"binlog":                 true,
		"cloudflare":             true,
		"confluence":             true,
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
		"elasticsearch":          true,
		"envoy":                  true,
		"etcd":                   true,
		"expressjs":              true,
		"external":               true,
		"file_collector":         true,
		"flink":                  true,
		"fluentd":                true,
		"ginlog":                 true,
		"gitlab":                 true,
		"goruntime":              true,
		"harborMonitor":          true,
		"host_processes":         true,
		"hostobject":             true,
		"httpPacket":             true,
		"httpProb":               true,
		"httpjson":               true,
		"httpstat":               true,
		"huaweiyunces":           true,
		"huaweiyunobject":        true,
		"jira":                   true,
		"jvm":                    true,
		"k8sobject":              true,
		"kafka":                  true,
		"kafkalog":               true,
		"kong":                   true,
		"lighttpd":               true,
		"mem":                    true,
		"mongodb_oplog":          true,
		"mongodb":                true,
		"mysql":                  true,
		"mysqlog":                true,
		"neo4j":                  true,
		"net":                    true,
		"nfsstat":                true,
		"nginx":                  true,
		"nginx_plus":             true,
		"nginx_plus_api":         true,
		"nginx_upstream_check":   true,
		"nginx_vts":              true,
		"nginxlog":               true,
		"oraclemonitor":          true,
		"oracle":                 true,
		"postgresql_replication": true,
		"processes":              true,
		"prom":                   true,
		"proxy":                  true,
		"puppetagent":            true,
		"rabbitmq":               true,
		"redis":                  true,
		"redislog":               true,
		"rum":                    true,
		"scanport":               true,
		"self":                   true,
		"squid":                  true,
		"ssh":                    true,
		"statsd":                 true,
		"swap":                   true,
		"system":                 true,
		"systemd":                true,
		"tailf":                  true,
		"tencentcms":             true,
		"tencentcost":            true,
		"tencentobject":          true,
		"tidb":                   true,
		"timezone":               true,
		"traceJaeger":            true,
		"traceSkywalking":        true,
		"traceZipkin":            true,
		"tracerouter":            true,
		"traefik":                true,
		"ucloud_monitor":         true,
		"wechatminiprogram":      true,
		"wmi":                    true,
		"yarn":                   true,
		"zabbix":                 true,
		"zaplog":                 true,
	}
)
