// Unless explicitly stated otherwise all files in this repository are licensed
// under the MIT License.
// This product includes software developed at Guance Cloud (https://www.guance.com/).
// Copyright 2021-present Guance, Inc.

package sqlserver

import (
	"fmt"
	"net/netip"
	"sync"
	T "testing"
	"time"

	"github.com/BurntSushi/toml"
	dt "github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/stretchr/testify/assert"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/internal/testutils"
	tu "gitlab.jiagouyun.com/cloudcare-tools/datakit/internal/testutils"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/io"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/io/point"
)

type caseSpec struct {
	t *T.T

	name        string
	repo        string
	repoTag     string
	envs        []string
	servicePort string

	ipt    *Input
	feeder *io.MockedFeeder

	pool     *dt.Pool
	resource *dt.Resource

	cr *testutils.CaseResult
}

func (cs *caseSpec) checkPoint(pts []*point.Point) error {
	for _, pt := range pts {
		switch pt.Name() {
		case "sqlserver_performance":
			//cs.t.Logf("get %s", pt.String())

			// TODO: check pt according to Performance

		default: // TODO: check other measurement
		}

		// check if tag appended
		if len(cs.ipt.Tags) != 0 {
			tags := pt.Tags()
			for k, expect := range cs.ipt.Tags {
				if got, ok := tags[k]; !ok {
					return fmt.Errorf("tag %s not found", k)
				} else if got != expect {
					return fmt.Errorf("expect tag value %s, got %s", expect, got)
				}
			}
		}
	}

	// TODO: some other checking on @pts, such as `if some required measurements exist'...

	return nil
}

func (cs *caseSpec) run() error {
	// start remote sqlserver
	r := tu.GetRemote()
	dockerTCP := r.TCPURL()

	cs.t.Logf("get remote: %+#v, TCP: %s", r, dockerTCP)

	p, err := dt.NewPool(dockerTCP)
	if err != nil {
		return err
	}

	if err := p.RemoveContainerByName(cs.name); err != nil {
		return err
	}

	resource, err := p.RunWithOptions(&dt.RunOptions{
		// specify container image & tag
		Repository: cs.repo,
		Tag:        cs.repoTag,

		// port binding
		PortBindings: map[docker.Port][]docker.PortBinding{
			"1433/tcp": {{HostIP: "0.0.0.0", HostPort: cs.servicePort}},
		},

		// container name
		Name: cs.name,

		// container run-time envs
		Env: cs.envs,
	}, func(c *docker.HostConfig) {

		c.RestartPolicy = docker.RestartPolicy{Name: "no"}
	})
	if err != nil {
		return err
	}

	cs.pool = p
	cs.resource = resource

	cs.t.Logf("check service(%s:%s)...", r.Host, cs.servicePort)
	if  !r.PortOK(cs.servicePort, time.Minute) {
		return fmt.Errorf("service checking failed")
	}

	var wg sync.WaitGroup

	// start input
	cs.t.Logf("start input...")
	wg.Add(1)
	go func() {
		defer wg.Done()
		cs.ipt.Run()
	}()

	// wait data
	cs.t.Logf("wait points...")
	pts, err := cs.feeder.AnyPoints()
	if err != nil {
		return err
	}

	cs.t.Logf("get %d points", len(pts))
	if err:= cs.checkPoint(pts)); err!= nil {
		return err
	}

	cs.t.Logf("stop input...")
	cs.ipt.Terminate()

	cs.t.Logf("exit...")
	wg.Wait()

	return nil
}

func buildCases(t *T.T) ([]*caseSpec, error) {
	t.Helper()

	bases := []struct {
		name string
		conf string
	}{
		{
			name: "remote-sqlserver",

			conf: fmt.Sprintf(`
host = "%s"
user = "sa"
password = "Abc123abC$"`, tu.GetRemote().Host+":1433"),
		},

		{
			name: "remote-sqlserver-with-extra-tags",

			conf: fmt.Sprintf(`
host = "%s"
user = "sa"
password = "Abc123abC$"
[sqlserver.tags]
  tag1 = "some_value"
  tag2 = "some_other_value"`, tu.GetRemote().Host+":2433"),
		},
	}

	images := [][2]string{
		{"mcr.microsoft.com/mssql/server", "2017-latest"},
		{"mcr.microsoft.com/mssql/server", "2019-latest"},
		{"mcr.microsoft.com/mssql/server", "2022-latest"},
	}

	// TODO: add per-image configs
	perImageCfgs := []interface{}{}
	_ = perImageCfgs

	var cases []*caseSpec

	// compose cases
	for _, img := range images {
		for _, base := range bases {
			feeder := io.NewMockedFeeder()

			ipt := defaultInput()
			ipt.feeder = feeder

			_, err := toml.Decode(base.conf, ipt)
			assert.NoError(t, err)

			envs := []string{
				"ACCEPT_EULA=Y",
				fmt.Sprintf("SA_PASSWORD=%s", ipt.Password),
			}

			ipport, err := netip.ParseAddrPort(ipt.Host)
			assert.NoError(t, err, "parse %s failed: %s", ipt.Host, err)

			cases = append(cases, &caseSpec{
				t:      t,
				ipt:    ipt,
				name:   base.name,
				feeder: feeder,
				envs:   envs,

				repo:    img[0],
				repoTag: img[1],

				servicePort: fmt.Sprintf("%d", ipport.Port()),

				cr: &testutils.CaseResult{
					Name: t.Name(),
					Case: base.name,
					ExtraTags: map[string]string{
						"image":         img[0],
						"image_tag":     img[1],
						"remote_server": ipt.Host,
					},
				},
			})
		}
	}
	return cases, nil
}

func TestSQLServerInput(t *T.T) {
	start := time.Now()
	cases, err := buildCases(t)
	if err != nil {
		cs := testutils.CaseResult{
			Name:          t.Name(),
			Status:        testutils.CasePassed,
			FailedMessage: err.Error(),
			Cost:          time.Since(start),
		}

		_ = cs.Flush()
		return
	}

	t.Logf("testing %d cases...", len(cases))

	for _, tc := range cases {
		t.Run(tc.name, func(t *T.T) {
			caseStart := time.Now()

			t.Logf("testing %s...", tc.name)
			if err := tc.run(); err != nil {
				tc.cr.Status = testutils.CaseFailed
				tc.cr.FailedMessage = err.Error()
			} else {
				tc.cr.Status = testutils.CasePassed
			}

			tc.cr.Cost = time.Since(caseStart)

			assert.NoError(t, tc.cr.Flush())

			t.Cleanup(func() {
				// clean remote docker resources
				if tc.resource == nil {
					return
				}

				assert.NoError(t, tc.pool.Purge(tc.resource))
			})
		})
	}
}
