// Unless explicitly stated otherwise all files in this repository are licensed
// under the MIT License.
// This product includes software developed at Guance Cloud (https://www.guance.com/).
// Copyright 2021-present Guance, Inc.

// Package dataway implement API request to dataway.
package dataway

import (
	"fmt"
	"runtime"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/GuanceCloud/cliutils/logger"
	"github.com/GuanceCloud/cliutils/point"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/internal/datakit"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/internal/git"
)

const (
	HeaderXGlobalTags = "X-Global-Tags"
	DefaultRetryCount = 4
	DefaultRetryDelay = time.Millisecond * 200
)

type IDataway interface {
	Write(...WriteOption) error
	Pull(what string) ([]byte, error)
}

var (
	dwAPIs = []string{
		point.MetricDeprecated.URL(),
		point.Metric.URL(),
		point.Network.URL(),
		point.KeyEvent.URL(),
		point.Object.URL(),
		point.CustomObject.URL(),
		point.Logging.URL(),
		point.Tracing.URL(),
		point.RUM.URL(),
		point.Security.URL(),
		point.Profiling.URL(),

		datakit.DatakitPull,
		datakit.LogFilter,
		datakit.SessionReplayUpload,
		datakit.HeartBeat,
		datakit.Election,
		datakit.ElectionHeartbeat,
		datakit.QueryRaw,
		datakit.Workspace,
		datakit.ListDataWay,
		datakit.ObjectLabel,
		datakit.LogUpload,
		datakit.PipelinePull,
		datakit.ProfilingUpload,
		datakit.TokenCheck,
	}

	AvailableDataways          = []string{}
	log                        = logger.DefaultSLogger("dataway")
	datawayListIntervalDefault = 60
)

type Dataway struct {
	URLs []string `toml:"urls"`

	DeprecatedHTTPTimeout string        `toml:"timeout,omitempty"`
	HTTPTimeout           time.Duration `toml:"timeout_v2"`
	MaxRetryCount         int           `toml:"max_retry_count"`
	RetryDelay            time.Duration `toml:"retry_delay"`

	HTTPProxy string `toml:"http_proxy"`

	Hostname string `toml:"-"`

	// Deprecated
	DeprecatedHost   string `toml:"host,omitempty"`
	DeprecatedScheme string `toml:"scheme,omitempty"`
	DeprecatedToken  string `toml:"token,omitempty"`
	DeprecatedURL    string `toml:"url,omitempty"`

	MaxIdleConnsPerHost int           `toml:"max_idle_conns_per_host,omitempty"`
	MaxIdleConns        int           `toml:"max_idle_conns"`
	IdleTimeout         time.Duration `toml:"idle_timeout"`

	Proxy bool `toml:"proxy,omitempty"`

	EnableHTTPTrace bool `toml:"enable_httptrace"`
	EnableSinker    bool `toml:"enable_sinker"`

	GlobalCustomerKeys []string `toml:"global_customer_keys"`

	eps        []*endPoint
	locker     sync.RWMutex
	dnsCachers []*dnsCacher

	globalTags                map[string]string
	globalTagsHTTPHeaderValue string

	// metrics
}

type dwopt func(*Dataway)

func ParseGlobalCustomerKeys(v string) (arr []string) {
	for _, elem := range strings.Split(v, ",") { // remove white space
		if x := strings.TrimSpace(elem); len(x) > 0 {
			arr = append(arr, x)
		}
	}
	return
}

func WithGlobalTags(maps ...map[string]string) dwopt {
	return func(dw *Dataway) {
		if dw.globalTags == nil {
			dw.globalTags = map[string]string{}
		}

		for _, tags := range maps {
			for k, v := range tags {
				dw.globalTags[k] = v
			}
		}

		log.Infof("dataway set globals: %+#v", dw.globalTags)
	}
}

func (dw *Dataway) Init(opts ...dwopt) error {
	log = logger.SLogger("dataway")

	for _, opt := range opts {
		if opt != nil {
			opt(dw)
		}
	}

	if err := dw.doInit(); err != nil {
		return err
	}

	return nil
}

func (dw *Dataway) String() string {
	arr := []string{fmt.Sprintf("dataways: [%s]", strings.Join(dw.URLs, ","))}

	for _, x := range dw.eps {
		arr = append(arr, "---------------------------------")
		for k, v := range x.categoryURL {
			arr = append(arr, fmt.Sprintf("% 24s: %s", k, v))
		}
	}

	return strings.Join(arr, "\n")
}

func (dw *Dataway) ClientsCount() int {
	return len(dw.eps)
}

func (dw *Dataway) GetTokens() []string {
	var arr []string
	for _, ep := range dw.eps {
		if ep.token != "" {
			arr = append(arr, ep.token)
		}
	}

	return arr
}

// TagHeaderValue create X-Global-Tags header value in the
// form of key=val,key=val with ASC sorted.
func TagHeaderValue(tags map[string]string) string {
	var arr []string
	for k, v := range tags {
		arr = append(arr, fmt.Sprintf("%s=%s", k, v))
	}
	sort.Strings(arr)
	return strings.Join(arr, ",")
}

func (dw *Dataway) doInit() error {
	// 如果 env 已传入了 dataway 配置, 则不再追加老的 dataway 配置,
	// 避免俩边配置了同样的 dataway, 造成数据混乱
	if dw.DeprecatedURL != "" && len(dw.URLs) == 0 {
		dw.URLs = []string{dw.DeprecatedURL}
	}

	if len(dw.URLs) == 0 {
		return fmt.Errorf("dataway not set: urls is empty")
	}

	if dw.HTTPTimeout <= time.Duration(0) {
		dw.HTTPTimeout = time.Second * 30
	}

	if dw.MaxIdleConnsPerHost == 0 {
		dw.MaxIdleConnsPerHost = 64
	}

	log.Infof("set %d global tags to dataway", len(dw.globalTags))
	if len(dw.globalTags) > 0 && dw.EnableSinker {
		dw.globalTagsHTTPHeaderValue = TagHeaderValue(dw.globalTags)
	}

	for _, u := range dw.URLs {
		ep, err := newEndpoint(u,
			withProxy(dw.HTTPProxy),
			withAPIs(dwAPIs),
			withHTTPHeaders(map[string]string{
				HeaderXGlobalTags: dw.globalTagsHTTPHeaderValue,

				// DatakitUserAgent define HTTP User-Agent header.
				// user-agent format. See
				// 	 https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/User-Agent
				"User-Agent": fmt.Sprintf("datakit-%s-%s/%s/%s",
					runtime.GOOS, runtime.GOARCH, git.Version, datakit.DatakitHostName),
			}),
			withHTTPTimeout(dw.HTTPTimeout),
			withHTTPTrace(dw.EnableHTTPTrace),
			withMaxHTTPIdleConnectionPerHost(dw.MaxIdleConnsPerHost),
			withMaxHTTPConnections(dw.MaxIdleConns),
			withHTTPIdleTimeout(dw.IdleTimeout),
			withMaxRetryCount(dw.MaxRetryCount),
			withRetryDelay(dw.RetryDelay),
		)
		if err != nil {
			log.Errorf("init dataway url %s failed: %s", u, err.Error())
			return err
		}

		dw.eps = append(dw.eps, ep)

		dw.addDNSCache(ep.host)
	}

	return nil
}

func (dw *Dataway) GlobalTags() map[string]string {
	return dw.globalTags
}

func (dw *Dataway) CustomTagKeys() []string {
	return dw.GlobalCustomerKeys
}

func (dw *Dataway) GlobalTagsHTTPHeaderValue() string {
	return dw.globalTagsHTTPHeaderValue
}

func (dw *Dataway) addDNSCache(host string) {
	for _, v := range dw.dnsCachers {
		if v.GetDomain() == host {
			return // avoid repeat add same domain
		}
	}

	dnsCache := &dnsCacher{}
	dnsCache.initDNSCache(host, dw.initEndpoints)

	dw.dnsCachers = append(dw.dnsCachers, dnsCache)
}

func (dw *Dataway) initEndpoints() error {
	dw.locker.Lock()
	defer dw.locker.Unlock()

	for _, ep := range dw.eps {
		if err := ep.setupHTTP(); err != nil {
			return err
		}
	}

	return nil
}
