// Unless explicitly stated otherwise all files in this repository are licensed
// under the MIT License.
// This product includes software developed at Guance Cloud (https://www.guance.com/).
// Copyright 2021-present Guance, Inc.

// Package logstreaming handle remote logging data.
package logstreaming

import (
	"bufio"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	lp "gitlab.jiagouyun.com/cloudcare-tools/cliutils/lineproto"
	"gitlab.jiagouyun.com/cloudcare-tools/cliutils/logger"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/config"
	dhttp "gitlab.jiagouyun.com/cloudcare-tools/datakit/http"
	ihttp "gitlab.jiagouyun.com/cloudcare-tools/datakit/internal/http"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/io"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/pipeline"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/pipeline/scriptstore"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs"
)

const (
	inputName        = "logstreaming"
	defaultPercision = "s"

	sampleCfg = `
[inputs.logstreaming]
  ignore_url_tags = true
`
)

var l = logger.DefaultSLogger(inputName)

type Input struct {
	IgnoreURLTags bool `yaml:"ignore_url_tags"`
}

func (*Input) Catalog() string { return "log" }

func (*Input) SampleConfig() string { return sampleCfg }

func (*Input) AvailableArchs() []string { return datakit.AllArch }

func (*Input) SampleMeasurement() []inputs.Measurement {
	return []inputs.Measurement{&logstreamingMeasurement{}}
}

func (*Input) Run() {
	l.Info("register logstreaming router")
}

func (i *Input) RegHTTPHandler() {
	l = logger.SLogger(inputName)
	dhttp.RegHTTPHandler("POST", "/v1/write/logstreaming", ihttp.ProtectedHandlerFunc(i.handleLogstreaming, l))
}

type Result struct {
	Status   string `json:"status"`
	ErrorMsg string `json:"error_msg,omitempty"`
}

const (
	statusSuccess = "success"
	statusFail    = "fail"
)

func (i *Input) handleLogstreaming(resp http.ResponseWriter, req *http.Request) {
	var (
		urlstr       = req.URL.String()
		queries      = req.URL.Query()
		typee        = queries.Get("type")
		source       = completeSource(queries.Get("source"))
		precision    = completePrecision(queries.Get("precision"))
		pipelinePath = queries.Get("pipeline")

		// TODO
		// 每一条 request 都要解析和创建一个 tags，影响性能
		// 可以将其缓存起来，以 url 的 md5 值为 key
		extraTags = config.ParseGlobalTags(queries.Get("tags"))

		result = Result{Status: statusSuccess}
		err    error
	)

	if !i.IgnoreURLTags {
		extraTags["ip_or_hostname"] = req.URL.Hostname()
		extraTags["service"] = completeService(source, queries.Get("service"))
	}

	l.Debugf("req url %s", urlstr)

	if req.Body == nil {
		l.Debugf("body is nil")
		return
	}

	switch typee {
	case "influxdb":
		body, _err := ioutil.ReadAll(req.Body)
		if _err != nil {
			l.Errorf("url %s failed to read body: %s", urlstr, _err)
			err = _err
			break
		}

		pts, _err := lp.ParsePoints(body, &lp.Option{
			Time:      time.Now(),
			ExtraTags: extraTags,
			Strict:    true,
			Precision: precision,
		})
		if _err != nil {
			l.Errorf("url %s handler err: %s", urlstr, _err)
			err = _err
			break
		}

		if len(pts) == 0 {
			l.Debugf("len(points) is zero, skip")
			break
		}
		pts1 := io.WrapPoint(pts)
		pts1 = plRunPt(pts1)
		err = io.Feed(inputName, datakit.Logging, pts1, nil)
	default:
		scanner := bufio.NewScanner(req.Body)
		pending := []string{}
		for scanner.Scan() {
			pending = append(pending, scanner.Text())
		}
		pts := plRunCnt(source, pipelinePath, pending, extraTags)
		err = io.Feed(source, datakit.Logging, pts, nil)
	}

	if err != nil {
		result.Status = statusFail
		result.ErrorMsg = err.Error()
		resp.WriteHeader(http.StatusBadRequest)
	}
	resultBody, err := json.Marshal(result)
	if err != nil {
		l.Errorf("marshal reuslt err: %s", err)
	}

	if _, err := resp.Write(resultBody); err != nil {
		l.Warnf("Write: %s, ignored", err.Error())
	}
}

func completeSource(source string) string {
	if source != "" {
		return source
	}
	return "default"
}

func completeService(defaultService, service string) string {
	if service != "" {
		return service
	}
	return defaultService
}

func completePrecision(precision string) string {
	if precision == "" {
		return defaultPercision
	}
	return precision
}

func init() { //nolint:gochecknoinits
	inputs.Add(inputName, func() inputs.Input {
		return &Input{}
	})
}

type logstreamingMeasurement struct{}

func (*logstreamingMeasurement) LineProto() (*io.Point, error) { return nil, nil }

//nolint:lll
func (*logstreamingMeasurement) Info() *inputs.MeasurementInfo {
	return &inputs.MeasurementInfo{
		Type: "logging",
		Name: "logstreaming 日志接收",
		Desc: "非行协议数据格式时，使用 URL 中的 `source` 参数，如果该值为空，则默认为 `default`",
		Tags: map[string]interface{}{
			"service":        inputs.NewTagInfo("service 名称，对应 URL 中的 `service` 参数"),
			"ip_or_hostname": inputs.NewTagInfo("request IP or hostname"),
		},
		Fields: map[string]interface{}{
			"message": &inputs.FieldInfo{DataType: inputs.String, Unit: inputs.UnknownUnit, Desc: "日志正文，默认存在，可以使用 pipeline 删除此字段"},
		},
	}
}

func plCallback(ret *pipeline.Result) (*pipeline.Result, error) {
	return pipeline.ResultUtilsLoggingProcessor(ret, false, nil), nil
}

func plRunCnt(source, scriptName string, cnt []string, tags map[string]string) []*io.Point {
	ret := []*io.Point{}

	ts := time.Now().Add(-time.Duration(int(time.Nanosecond) * len(cnt)))
	for i, cnt := range cnt {
		pt, err := io.MakePoint(source, tags,
			map[string]interface{}{pipeline.PipelineMessageField: cnt},
			ts.Add(time.Duration(int(time.Nanosecond)*i)),
		)
		if err != nil {
			l.Error(err)
			continue
		}
		drop := false
		if script, ok := scriptstore.QueryScript(datakit.Logging, scriptName); ok {
			if ptRet, dropRet, err := pipeline.RunScript(pt, script, plCallback); err != nil {
				l.Error(err)
			} else {
				pt = ptRet
				drop = dropRet
			}
		}
		if !drop {
			ret = append(ret, pt)
		}
	}
	return ret
}

func plRunPt(pts []*io.Point) []*io.Point {
	ret := []*io.Point{}
	for _, pt := range pts {
		scriptName := pt.Name() + ".p"
		drop := false
		if script, ok := scriptstore.QueryScript(datakit.Logging, scriptName); ok {
			if ptRet, dropRet, err := pipeline.RunScript(pt, script, plCallback); err != nil {
				l.Error(err)
			} else {
				pt = ptRet
				drop = dropRet
			}
		}
		if !drop {
			ret = append(ret, pt)
		}
	}
	return ret
}
