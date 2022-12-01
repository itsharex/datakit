// Unless explicitly stated otherwise all files in this repository are licensed
// under the MIT License.
// This product includes software developed at Guance Cloud (https://www.guance.com/).
// Copyright 2021-present Guance, Inc.

package man

import (
	"embed"
)

//go:embed manuals/*.md
var docs embed.FS

var OtherDocs = map[string]bool{
	"apis":                               true,
	"changelog":                          true,
	"confd":                              true,
	"container-log":                      true,
	"datakit-batch-deploy":               true,
	"datakit-conf":                       true,
	"datakit-daemonset-deploy":           true,
	"datakit-dql-how-to":                 true,
	"datakit-filter":                     true,
	"datakit-input-conf":                 true,
	"datakit-install":                    true,
	"datakit-logging":                    true,
	"datakit-logging-how":                true,
	"datakit-monitor":                    true,
	"datakit-offline-install":            true,
	"datakit-on-public":                  true,
	"datakit-pl-global":                  true,
	"datakit-pl-how-to":                  true,
	"datakit-refer-table":                true,
	"datakit-service-how-to":             true,
	"datakit-sink-dataway":               true,
	"datakit-sink-dev":                   true,
	"datakit-sink-guide":                 true,
	"datakit-sink-influxdb":              true,
	"datakit-sink-logstash":              true,
	"datakit-sink-m3db":                  true,
	"datakit-sink-otel-jaeger":           true,
	"datakit-tools-how-to":               true,
	"datakit-tracing":                    true,
	"datakit-tracing-introduction":       true,
	"datakit-tracing-struct":             true,
	"datakit-update":                     true,
	"datatypes":                          true,
	"dataway":                            true,
	"dca":                                true,
	"ddtrace-cpp":                        true,
	"ddtrace-golang":                     true,
	"ddtrace-java":                       true,
	"ddtrace-nodejs":                     true,
	"ddtrace-php":                        true,
	"ddtrace-python":                     true,
	"ddtrace-ruby":                       true,
	"development":                        true,
	"dialtesting_json":                   true,
	"doc-logging":                        true,
	"election":                           true,
	"git-config-how-to":                  true,
	"k8s-config-how-to":                  true,
	"kubernetes-crd":                     true,
	"kubernetes-prom":                    true,
	"kubernetes-prometheus-operator-crd": true,
	"logfwd":                             true,
	"logging-pipeline-bench":             true,
	"logging_socket":                     true,
	"opentelemetry-go":                   true,
	"opentelemetry-java":                 true,
	"pipeline":                           true,
	"profile-ddtrace":                    true,
	"profile-java-async-profiler":        true,
	"prometheus":                         true,
	"python-profiling":                   true,
	"rum":                                true,
	"sec-checker":                        true,
	"snmp":                               true,
	"telegraf":                           true,
	"why-no-data":                        true,
}
