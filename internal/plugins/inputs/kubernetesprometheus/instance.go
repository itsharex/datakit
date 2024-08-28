// Unless explicitly stated otherwise all files in this repository are licensed
// under the MIT License.
// This product includes software developed at Guance Cloud (https://www.guance.com/).
// Copyright 2021-present Guance, Inc.

package kubernetesprometheus

import (
	"context"
	"regexp"
	"sync/atomic"
	"time"

	"gitlab.jiagouyun.com/cloudcare-tools/datakit/internal/datakit"
	dkio "gitlab.jiagouyun.com/cloudcare-tools/datakit/internal/io"
	dknet "gitlab.jiagouyun.com/cloudcare-tools/datakit/internal/net"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
)

type Role string

const (
	RoleNode      Role = "node"
	RoleService   Role = "service"
	RoleEndpoints Role = "endpoints"
	RolePod       Role = "pod"
)

const matchedScrape = "true"

var supportedRoles = []Role{RoleNode, RoleService, RoleEndpoints, RolePod}

const (
	MateInstanceTag = "__kubernetes_mate_instance"
	MateHostTag     = "__kubernetes_mate_host"
)

type (
	Instance struct {
		Role       string   `toml:"role"`
		Namespaces []string `toml:"namespaces"`
		Selector   string   `toml:"selector"`
		Scrape     string   `toml:"scrape"`

		Target
		Custom `toml:"custom"`
		Auth   `toml:"auth"`

		validator *resourceValidator
	}

	Target struct {
		Scheme   string        `toml:"scheme"`
		Address  string        `toml:"_"` // private
		Port     string        `toml:"port"`
		Path     string        `toml:"path"`
		Params   string        `toml:"params"` // Does not support matches.
		Interval time.Duration `toml:"interval"`
	}

	Custom struct {
		Measurement      string            `toml:"measurement"`
		JobAsMeasurement bool              `toml:"job_as_measurement"`
		Tags             map[string]string `toml:"tags"`
	}

	Auth struct {
		BearerTokenFile string                 `toml:"bearer_token_file"`
		TLSConfig       *dknet.TLSClientConfig `toml:"tls_config"`
	}

	basePromConfig struct {
		urlstr      string
		measurement string
		tags        map[string]string
		// Only used on Endpoints.
		nodeName string
	}
)

func (ins *Instance) setDefault() {
	switch ins.Role {
	case string(RoleNode):
		ins.Address = "__kubernetes_node_address_InternalIP"
	case string(RoleService):
		ins.Address = "__kubernetes_endpoints_address_ip" // Because the Service does not have an IP.
	case string(RoleEndpoints):
		ins.Address = "__kubernetes_endpoints_address_ip"
	case string(RolePod):
		ins.Address = "__kubernetes_pod_ip"
	default:
		// unreachable
	}

	if ins.Scrape == "" {
		ins.Scrape = "true"
	}
	if ins.Scheme == "" {
		ins.Scheme = "http"
	}
	if ins.Path == "" {
		ins.Path = "/metrics"
	}
	if ins.Interval <= 0 {
		ins.Interval = time.Second * 30
	}
}

type InstanceManager struct {
	Instances []*Instance `toml:"instances"`
}

func (im InstanceManager) Run(
	ctx context.Context,
	clientset *kubernetes.Clientset,
	informerFactory informers.SharedInformerFactory,
	feeder dkio.Feeder,
) {
	roleInstances := make([][]*Instance, len(supportedRoles))

	for _, ins := range im.Instances {
		role := Role(ins.Role)

		switch role {
		case RoleNode, RoleService, RoleEndpoints, RolePod:
			// nil
		default:
			klog.Warnf("unexpetced role %s, only supported %v", role, supportedRoles)
			continue
		}

		// set default values
		ins.setDefault()

		v, err := newResourceValidator(ins.Namespaces, ins.Selector)
		if err != nil {
			klog.Warnf("cannot parse selector %s err %s", ins.Selector, err)
			continue
		}
		ins.validator = v

		for idx, r := range supportedRoles {
			if r == role {
				roleInstances[idx] = append(roleInstances[idx], ins)
			}
		}
	}

	for idx, instances := range roleInstances {
		if len(instances) == 0 {
			continue
		}

		var launcher roleLauncher
		var err error

		switch supportedRoles[idx] {
		case RoleNode:
			launcher, err = NewNode(informerFactory, instances, feeder)

		case RoleService:
			launcher, err = NewService(clientset, informerFactory, instances, feeder)

		case RoleEndpoints:
			launcher, err = NewEndpoints(informerFactory, instances, feeder)

		case RolePod:
			launcher, err = NewPod(informerFactory, instances, feeder)
		}

		if err != nil {
			klog.Warn(err)
			continue
		}

		managerGo.Go(func(_ context.Context) error {
			launcher.Run(ctx)
			return nil
		})
	}
}

type roleLauncher interface {
	Run(context.Context)
}

type keyMatcher struct {
	key string
	re  *regexp.Regexp
}

func newKeyMatcher(key string) keyMatcher                  { return keyMatcher{key: key} }
func newKeyMatcherWithRegexp(re *regexp.Regexp) keyMatcher { return keyMatcher{re: re} }

func (k keyMatcher) matches(str string) (matched bool, args []string) {
	if k.key != "" {
		return k.key == str, nil
	}
	if k.re.MatchString(str) {
		args := k.re.FindStringSubmatch(str)
		if len(args) >= 2 {
			return true, args[1:]
		}
	}
	return false, nil
}

func maxConcurrent(nodeLocal bool) int {
	if nodeLocal {
		return datakit.AvailableCPUs + 1
	}
	return datakit.AvailableCPUs*3 + 1
}

type ctxKey int

const (
	ctxKeyNodeName ctxKey = iota + 1
	ctxKeyNodeLocal
	ctxKeyPause
)

func withNodeName(ctx context.Context, nodeName string) context.Context {
	return context.WithValue(ctx, ctxKeyNodeName, nodeName)
}

func withNodeLocal(ctx context.Context, nodeLocal bool) context.Context {
	return context.WithValue(ctx, ctxKeyNodeLocal, nodeLocal)
}

func withPause(ctx context.Context, pause *atomic.Bool) context.Context {
	return context.WithValue(ctx, ctxKeyPause, pause)
}

func nodeNameFrom(ctx context.Context) (string, bool) {
	nodeName, ok := ctx.Value(ctxKeyNodeName).(string)
	return nodeName, ok
}

func nodeLocalFrom(ctx context.Context) bool {
	nodeLocal, ok := ctx.Value(ctxKeyNodeLocal).(bool)
	return nodeLocal && ok
}

func pauseFrom(ctx context.Context) (bool, bool) {
	pause, ok := ctx.Value(ctxKeyPause).(*atomic.Bool)
	if !ok {
		return false, false
	}
	p := pause.Load()
	return p, true
}
