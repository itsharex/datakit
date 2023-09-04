// Unless explicitly stated otherwise all files in this repository are licensed
// under the MIT License.
// This product includes software developed at Guance Cloud (https://www.guance.com/).
// Copyright 2021-present Guance, Inc.

// Package prom used to parsing promemetheuse exportor metrics.
package prom

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"time"

	"github.com/GuanceCloud/cliutils/logger"
	"github.com/GuanceCloud/cliutils/point"
	dto "github.com/prometheus/client_model/go"
	"github.com/prometheus/common/expfmt"

	"gitlab.jiagouyun.com/cloudcare-tools/datakit/internal/datakit"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/internal/httpcli"
	dnet "gitlab.jiagouyun.com/cloudcare-tools/datakit/internal/net"
)

type Rule struct {
	Pattern string `toml:"pattern" json:"pattern"`
	Prefix  string `toml:"prefix" json:"prefix"`
	Name    string `toml:"name" json:"name"`
}

type RenameTags struct {
	OverwriteExistTags bool              `toml:"overwrite_exist_tags" json:"overwrite_exist_tags"`
	Mapping            map[string]string `toml:"mapping" json:"mapping"`
}

type AsLogging struct {
	Enable  bool   `toml:"enable" json:"enable"`
	Service string `toml:"service" json:"service"`
}

type IgnoreTagKeyValMatch map[string][]*regexp.Regexp

func (opt *option) GetSource(defaultSource ...string) string {
	if opt.source != "" {
		return opt.source
	}
	if len(defaultSource) > 0 {
		return defaultSource[0]
	}
	return "prom" //nolint:goconst
}

const (
	httpTimeout               = time.Second * 3
	defaultInsecureSkipVerify = false
)

type Prom struct {
	opt      *option
	client   *http.Client
	parser   expfmt.TextParser
	InfoTags map[string]string
	ptCount  int
}

func NewProm(promOpts ...PromOption) (*Prom, error) {
	opt := option{}
	for idx := range promOpts {
		if promOpts[idx] != nil {
			promOpts[idx](&opt)
		}
	}

	if opt.l == nil {
		opt.l = logger.DefaultSLogger("prom")
	}

	if opt.timeout < httpTimeout {
		opt.timeout = httpTimeout
	}

	p := Prom{opt: &opt, InfoTags: make(map[string]string)}

	var f expfmt.BatchCallback = func(mf map[string]*dto.MetricFamily) error {
		pts, err := p.MetricFamilies2points(mf, "")
		if err != nil {
			return err
		}

		return p.opt.batchCallback(pts)
	}
	if opt.streamSize > 0 && opt.batchCallback != nil {
		parse := expfmt.NewTextParser(expfmt.WithBatchCallback(opt.streamSize, f))
		p.parser = *parse
	}

	cliopts := httpcli.NewOptions()
	cliopts.DialTimeout = opt.timeout
	cliopts.DialKeepAlive = opt.keepAlive

	if opt.tlsOpen {
		caCerts := []string{}
		insecureSkipVerify := defaultInsecureSkipVerify
		if len(opt.cacertFile) != 0 {
			caCerts = append(caCerts, opt.cacertFile)
		} else {
			insecureSkipVerify = true
		}
		tc := &dnet.TLSClientConfig{
			CaCerts:            caCerts,
			Cert:               opt.certFile,
			CertKey:            opt.keyFile,
			InsecureSkipVerify: insecureSkipVerify,
		}

		tlsConfig, err := tc.TLSConfig()
		if err != nil {
			return nil, err
		}
		cliopts.TLSClientConfig = tlsConfig
	}

	if p.opt.udsPath != "" {
		cliopts.DialContext = func(_ context.Context, _, _ string) (net.Conn, error) {
			return net.Dial("unix", p.opt.udsPath)
		}
	}

	p.SetClient(httpcli.Cli(cliopts))
	return &p, nil
}

func (p *Prom) Option() *option {
	return p.opt
}

func (p *Prom) SetClient(cli *http.Client) {
	p.client = cli
}

func (p *Prom) GetReq(url string) (*http.Request, error) {
	var (
		req *http.Request
		err error
	)

	if len(p.opt.auth) > 0 {
		if authType, ok := p.opt.auth["type"]; ok {
			if authFunc, ok := AuthMaps[authType]; ok {
				req, err = authFunc(p.opt.auth, url)
			} else {
				req, err = http.NewRequest("GET", url, nil)
			}
		}
	} else {
		req, err = http.NewRequest("GET", url, nil)
	}
	for k, v := range p.opt.httpHeaders {
		req.Header.Set(k, v)
	}
	return req, err
}

func (p *Prom) Request(url string) (*http.Response, error) {
	start := time.Now()
	defer func() {
		httpLatencyVec.WithLabelValues(p.opt.source).Observe(float64(time.Since(start)) / float64(time.Second))
	}()
	req, err := p.GetReq(url)
	if err != nil {
		return nil, err
	}

	r, err := p.client.Do(req)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// CollectFromHTTPV2 collect points.
func (p *Prom) CollectFromHTTPV2(u string) ([]*point.Point, error) {
	resp, err := p.Request(u)
	if err != nil {
		if p.opt.ignoreReqErr {
			return []*point.Point{}, nil
		} else {
			return nil, fmt.Errorf("collect from %s: %w", u, err)
		}
	}
	defer resp.Body.Close() //nolint:errcheck

	// A agent used to count bytes.
	wCounter := &writeCounter{}
	pts, err := p.text2Metrics(io.TeeReader(resp.Body, wCounter), u)
	if err != nil {
		return nil, err
	}
	defer func() {
		collectPointsTotalVec.WithLabelValues(p.opt.source).Observe(float64(p.ptCount))
		p.ptCount = 0
		httpGetBytesVec.WithLabelValues(p.opt.source).Observe(float64(wCounter.total))
	}()
	return pts, nil
}

// CollectFromFileV2 collect points.
func (p *Prom) CollectFromFileV2(filepath string) ([]*point.Point, error) {
	f, err := os.OpenFile(filepath, os.O_RDONLY, 0o600) //nolint:gosec
	if err != nil {
		return nil, err
	}
	defer f.Close() //nolint:errcheck,gosec
	return p.text2Metrics(f, "")
}

// WriteMetricText2File scrapes raw prometheus metric text from u
// then appends them directly to file p.opt.Output.
func (p *Prom) WriteMetricText2File(u string) error {
	fp := p.opt.output
	if !path.IsAbs(fp) {
		fp = filepath.Join(datakit.InstallDir, fp)
	}
	// Append to file if already exist.
	f, err := os.OpenFile(fp, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0o660) //nolint:gosec
	if err != nil {
		return err
	}
	defer f.Close() //nolint:errcheck,gosec

	uu, err := url.Parse(u)
	if err != nil {
		return fmt.Errorf("url parse error, %w", err)
	}
	// If url is configured as local path file, prom does not collect from it.
	if uu.Scheme != "http" && uu.Scheme != "https" {
		return fmt.Errorf("url is neither http nor https")
	}

	resp, err := p.client.Get(u)
	if err != nil {
		return err
	}
	defer resp.Body.Close() //nolint:errcheck

	if resp.ContentLength > p.opt.maxFileSize {
		return fmt.Errorf("content length is too large to handle, max: %d, got: %d", p.opt.maxFileSize, resp.ContentLength)
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if int64(len(data)) > p.opt.maxFileSize {
		return fmt.Errorf("content length is too large to handle, max: %d, got: %d", p.opt.maxFileSize, len(data))
	}
	if _, err := f.Write(data); err != nil {
		return err
	}
	return nil
}

// A agent used to count bytes.
type writeCounter struct {
	total uint64
}

// A agent used to count bytes.
func (wc *writeCounter) Write(p []byte) (int, error) {
	wc.total += uint64(len(p))
	return 0, nil
}
