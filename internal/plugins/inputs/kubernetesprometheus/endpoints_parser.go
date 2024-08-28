// Unless explicitly stated otherwise all files in this repository are licensed
// under the MIT License.
// This product includes software developed at Guance Cloud (https://www.guance.com/).
// Copyright 2021-present Guance, Inc.

package kubernetesprometheus

import (
	"fmt"
	"regexp"
	"strconv"

	corev1 "k8s.io/api/core/v1"
)

// EndpointsName                      = "__kubernetes_endpoints_name"
// EndpointsNamespace                 = "__kubernetes_endpoints_namespace"
// EndpointsLabel                     = "__kubernetes_endpoints_label_%s"
// EndpointsAnnotation                = "__kubernetes_endpoints_annotation_%s"
// EndpointsAddressNodeName           = "__kubernetes_endpoints_address_node_name"
// EndpointsAddressIP                 = "__kubernetes_endpoints_address_ip"
// EndpointsaddressTargetPodName      = "__kubernetes_endpoints_address_target_pod_name"
// EndpointsaddressTargetPodNamespace = "__kubernetes_endpoints_address_target_pod_namespace"
// EndpointsaddressPortNumber         = "__kubernetes_endpoints_port_%s_number"

var (
	EndpointsValueFroms = []struct {
		key keyMatcher
		fn  func(*corev1.Endpoints, []string) string
	}{
		{
			key: newKeyMatcher("__kubernetes_endpoints_name"),
			fn:  func(item *corev1.Endpoints, _ []string) string { return item.Name },
		},
		{
			key: newKeyMatcher("__kubernetes_endpoints_namespace"),
			fn:  func(item *corev1.Endpoints, _ []string) string { return item.Namespace },
		},
		{
			// e.g. __kubernetes_endpoints_label_app
			key: newKeyMatcherWithRegexp(regexp.MustCompile(`__kubernetes_endpoints_label_(.+)`)),
			fn: func(item *corev1.Endpoints, args []string) string {
				if len(args) != 1 {
					return ""
				}
				return item.Labels[args[0]]
			},
		},
		{
			// e.g. __kubernetes_endpoints_annotation_prometheus_io_scheme
			key: newKeyMatcherWithRegexp(regexp.MustCompile(`__kubernetes_endpoints_annotation_(.+)`)),
			fn: func(item *corev1.Endpoints, args []string) string {
				if len(args) != 1 {
					return ""
				}
				return item.Annotations[args[0]]
			},
		},
	}

	EndpointAddressValueFroms = []struct {
		key keyMatcher
		fn  func(*corev1.EndpointAddress, []string) string
	}{
		{
			key: newKeyMatcher("__kubernetes_endpoints_address_node_name"),
			fn: func(item *corev1.EndpointAddress, _ []string) string {
				if item.NodeName != nil {
					return *item.NodeName
				}
				return ""
			},
		},
		{
			key: newKeyMatcher("__kubernetes_endpoints_address_ip"),
			fn:  func(item *corev1.EndpointAddress, _ []string) string { return item.IP },
		},
		{
			key: newKeyMatcher("__kubernetes_endpoints_address_target_kind"),
			fn: func(item *corev1.EndpointAddress, _ []string) string {
				if item.TargetRef != nil {
					return item.TargetRef.Kind
				}
				return ""
			},
		},
		{
			key: newKeyMatcher("__kubernetes_endpoints_address_target_name"),
			fn: func(item *corev1.EndpointAddress, _ []string) string {
				if item.TargetRef != nil {
					return item.TargetRef.Name
				}
				return ""
			},
		},
		{
			key: newKeyMatcher("__kubernetes_endpoints_address_target_namespace"),
			fn: func(item *corev1.EndpointAddress, _ []string) string {
				if item.TargetRef != nil {
					return item.TargetRef.Namespace
				}
				return ""
			},
		},
		// deprecated, use __kubernetes_endpoints_address_target_name
		{
			key: newKeyMatcher("__kubernetes_endpoints_address_target_pod_name"),
			fn: func(item *corev1.EndpointAddress, _ []string) string {
				if item.TargetRef != nil {
					return item.TargetRef.Name
				}
				return ""
			},
		},
		// deprecated, use __kubernetes_endpoints_address_target_namespace
		{
			key: newKeyMatcher("__kubernetes_endpoints_address_target_pod_namespace"),
			fn: func(item *corev1.EndpointAddress, _ []string) string {
				if item.TargetRef != nil {
					return item.TargetRef.Namespace
				}
				return ""
			},
		},
	}

	EndpointPortValueFroms = []struct {
		key keyMatcher
		fn  func(*corev1.EndpointPort, []string) string
	}{
		{
			// e.g. __kubernetes_endpoints_port_metrics_number
			key: newKeyMatcherWithRegexp(regexp.MustCompile(`^__kubernetes_endpoints_port_(.*?)_number$`)),
			fn: func(item *corev1.EndpointPort, args []string) string {
				if len(args) != 1 {
					return ""
				}
				if item.Name == args[0] {
					return strconv.Itoa(int(item.Port))
				}
				return ""
			},
		},
	}
)

type endpointsParser struct{ item *corev1.Endpoints }

func newEndpointsParser(item *corev1.Endpoints) *endpointsParser { return &endpointsParser{item} }

func (p *endpointsParser) shouldScrape(scrape string) bool {
	if scrape == matchedScrape {
		return true
	}
	for _, v := range EndpointsValueFroms {
		matched, args := v.key.matches(scrape)
		if !matched {
			continue
		}
		if res := v.fn(p.item, args); res != "" {
			return res == matchedScrape
		}
	}
	return false
}

func (p *endpointsParser) parsePromConfig(ins *Instance) ([]*basePromConfig, error) {
	var configs []*basePromConfig

	for _, set := range p.item.Subsets {
		for addressIdx, address := range set.Addresses {
			// length 5
			oldElems := []string{ins.Scheme, ins.Address, ins.Port, ins.Path, ins.Measurement}
			newElems := deepCopySlice(oldElems)

			tagKeys := []string{}
			for k, v := range ins.Tags {
				tagKeys = append(tagKeys, k)
				newElems = append(newElems, v)
			}

			for idx, elem := range newElems {
				if matched, res := p.matchEndpoints(elem); matched && res != "" {
					newElems[idx] = res
					continue
				}
				if matched, res := p.matchAddress(&set.Addresses[addressIdx], elem); matched && res != "" {
					newElems[idx] = res
					continue
				}
				if matched, res := p.matchPort(set.Ports, elem); matched && res != "" {
					newElems[idx] = res
					continue
				}
				newElems[idx] = elem
			}

			u, err := buildURLWithParams(newElems[0], newElems[1], newElems[2], newElems[3], ins.Params)
			if err != nil {
				return nil, err
			}
			measurement := newElems[4]

			tags := map[string]string{}

			if len(tagKeys)+len(oldElems) != len(newElems) {
				return nil, fmt.Errorf("unexpected tags length %d-%d", len(tagKeys), len(newElems)-len(oldElems))
			}

			for idx, k := range tagKeys {
				tags[k] = newElems[idx+len(oldElems)]
			}

			for k, v := range tags {
				switch v {
				case MateInstanceTag:
					tags[k] = u.Host
				case MateHostTag:
					if host := splitHost(u.Host); host != "" {
						tags[k] = host
					}
				default:
					// nil
				}
			}

			nodeName := ""
			if address.NodeName != nil {
				nodeName = *address.NodeName
			}

			configs = append(configs, &basePromConfig{
				urlstr:      u.String(),
				measurement: measurement,
				tags:        tags,
				nodeName:    nodeName,
			})
		}
	}

	return configs, nil
}

func (p *endpointsParser) matchEndpoints(key string) (matched bool, res string) {
	for _, v := range EndpointsValueFroms {
		matched, args := v.key.matches(key)
		if !matched {
			continue
		}
		if res := v.fn(p.item, args); res != "" {
			return true, res
		}
	}
	return false, ""
}

func (p *endpointsParser) matchAddress(address *corev1.EndpointAddress, key string) (matched bool, res string) {
	for _, v := range EndpointAddressValueFroms {
		matched, args := v.key.matches(key)
		if !matched {
			continue
		}
		if res := v.fn(address, args); res != "" {
			return true, res
		}
	}
	return false, ""
}

func (p *endpointsParser) matchPort(ports []corev1.EndpointPort, key string) (matched bool, res string) {
	for idx := range ports {
		for _, v := range EndpointPortValueFroms {
			matched, args := v.key.matches(key)
			if !matched {
				continue
			}
			if res := v.fn(&ports[idx], args); res != "" {
				return true, res
			}
		}
	}
	return false, ""
}
