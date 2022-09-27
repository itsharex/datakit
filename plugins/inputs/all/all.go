// Unless explicitly stated otherwise all files in this repository are licensed
// under the MIT License.
// This product includes software developed at Guance Cloud (https://www.guance.com/).
// Copyright 2021-present Guance, Inc.

// Package inputs wraps all inputs implements
package inputs

import (
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/apache"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/beats_output"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/clickhousev1"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/cloudprober"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/consul"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/container"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/coredns"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/cpu"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/ddtrace"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/demo"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/dialtesting"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/disk"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/diskio"

	// nolint:typecheck
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/ebpf"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/elasticsearch"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/etcd"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/external"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/flinkv1"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/gitlab"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/hostdir"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/hostobject"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/iis"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/influxdb"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/jaeger"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/jenkins"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/jvm"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/kafka"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/kafkamq"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/logfwdserver"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/logging"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/logstreaming"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/mem"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/memcached"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/mongodb"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/mysql"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/net"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/netstat"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/nginx"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/nsq"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/nvidiasmi"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/opentelemetry"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/oracle"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/postgresql"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/process"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/profile"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/prom"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/promremote"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/promtail"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/proxy"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/pythond"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/rabbitmq"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/redis"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/rum"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/self"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/sensors"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/skywalking"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/smart"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/socket"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/solr"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/sqlserver"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/ssh"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/statsd"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/swap"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/system"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/tdengine"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/tomcat"

	// only windows.
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/winevent"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/zipkin"
	// deprecated.
)
