//go:build linux && with_dke_test
// +build linux,with_dke_test

package run

import "testing"

func TestDKE(t *testing.T) {
	runCmd(nil, &Flag{
		DataKitAPIServer: "0.0.0.0:9529",
		// Log:              "/dev/stdout",
		LogLevel:  "info",
		PprofHost: "0.0.0.0",
		PprofPort: "6267",
		Service:   "ebpf",

		Enabled: []string{"bpf-netlog"},

		EBPFNet: FlagNet{
			L7NetEnabled: []string{"httpflow"},
		},

		BPFNetLog: FlagBPFNetLog{
			EnableLog:      true,
			EnableMetric:   true,
			L7LogProtocols: []string{"http"},
		},

		EBPFTrace: FlagTrace{
			TraceServer:  "0.0.0.0:9529",
			TraceAllProc: true,
			TraceEnvList: []string{"DKE_SERVICE", "DK_BPFTRACE_SERVICE", "DD_SERVICE", "OTEL_SERVICE_NAME"},
		},
		PIDFile: "/tmp/ebpf.pid",
	})
}
