// Unless explicitly stated otherwise all files in this repository are licensed
// under the MIT License.
// This product includes software developed at Guance Cloud (https://www.guance.com/).
// Copyright 2021-present Guance, Inc.

package rum

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	uhttp "github.com/GuanceCloud/cliutils/network/http"
	plmanager "github.com/GuanceCloud/cliutils/pipeline/manager"
	"github.com/GuanceCloud/cliutils/point"
	"github.com/gin-gonic/gin/binding"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/internal/config"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/internal/httpapi"
	dkio "gitlab.jiagouyun.com/cloudcare-tools/datakit/internal/io"
)

func httpStatusRespFunc(resp http.ResponseWriter, req *http.Request, err error) {
	if err != nil {
		httpErr(resp, err)
	} else {
		httpOK(resp, nil)
	}
}

func (ipt *Input) parseCallback(p *point.Point) (*point.Point, error) {
	name := p.Name()

	appid, ok := p.Get(rumMetricAppID).(string)
	if !ok {
		return nil, fmt.Errorf("invalid key %q", rumMetricAppID)
	}

	if !contains(appid, config.Cfg.HTTPAPI.RUMAppIDWhiteList) {
		return nil, httpapi.ErrRUMAppIDNotInWhiteList
	}

	if _, ok := ipt.measurementMap[name]; !ok {
		return nil, uhttp.Errorf(httpapi.ErrUnknownRUMMeasurement, "unknown RUM measurement: %s", name)
	}

	if name == Error {
		// handle sourcemap
		sdkName, ok := p.Get("sdk_name").(string)
		if !ok {
			return nil, fmt.Errorf("invalid key %q", "sdk_name")
		}

		status := &sourceMapStatus{
			appid:   appid,
			sdkName: sdkName,
			status:  StatusUnknown,
			remark:  "",
		}
		px, err := ipt.parseSourcemap(p, sdkName, status)
		if err != nil {
			log.Errorf("handle source map failed: %s", err.Error())
			if status.status != StatusZipNotFound && status.status != StatusToolNotFound {
				status.status = StatusError
			}
			status.remark = err.Error()
			// Do nothing, return original point.
			sourceMapCount.WithLabelValues(status.appid, status.sdkName, status.status, utf8SubStr(status.remark, 8)).Inc()
			return p, nil
		}
		return px, nil
	} else if name == Resource {
		// handle resource provider
		px, err := ipt.handleProvider(p)
		if err != nil {
			log.Errorf("unable to new point: %s", err)
			// If err prompt, we return the original point.
			return p, nil
		}
		return px, nil
	}
	return p, nil
}

func (ipt *Input) handleRUM(resp http.ResponseWriter, req *http.Request) {
	log.Debugf("### RUM request from %s", req.URL.String())

	opts := []point.Option{
		point.WithPrecision(point.PrecNS),
		point.WithTime(time.Now()),
	}

	var (
		query                   = req.URL.Query()
		version, pipelineSource string
	)
	if x := query.Get(httpapi.ArgVersion); x != "" {
		version = x
	}
	if x := query.Get(httpapi.ArgPipelineSource); x != "" {
		pipelineSource = x
	}
	if x := query.Get(httpapi.ArgPrecision); x != "" {
		opts = append(opts, point.WithPrecision(point.PrecStr(x)))
	}

	body, err := uhttp.ReadBody(req)
	if err != nil {
		log.Error(err.Error())
		httpErr(resp, err)

		return
	}

	if len(body) == 0 {
		log.Debug(httpapi.ErrEmptyBody.Err.Error())
		httpErr(resp, httpapi.ErrEmptyBody)

		return
	}

	var (
		pts       []*point.Point
		apiConfig = config.Cfg.HTTPAPI
		ct        = httpapi.GetPointEncoding(req.Header)
	)

	ipStatus := &ipLocationStatus{}
	defer func() {
		ClientRealIPCounter.WithLabelValues(ipStatus.appid, ipStatus.ipStatus, ipStatus.locateStatus).Inc()
	}()

	opts = append(opts, point.WithExtraTags(geoTags(getSrcIP(apiConfig, req, ipStatus), ipStatus)), point.WithCallback(ipt.parseCallback))

	if pts, err = httpapi.HandleWriteBody(body, ct, opts...); err != nil {
		log.Errorf("httpapi.HandleWriteBody: %s", err.Error())
		httpErr(resp, err)
		return
	}

	if len(pts) == 0 {
		log.Debug(httpapi.ErrNoPoints.Err.Error())
		httpErr(resp, httpapi.ErrNoPoints)

		return
	}

	ipStatus.appid, _ = pts[0].Get("app_id").(string)

	log.Debugf("### received %d(%s) points from %s, pipeline source: %v",
		len(pts), req.URL.Path, inputName, pipelineSource)

	feedOpt := &dkio.Option{Version: version}
	if pipelineSource != "" {
		feedOpt.PlOption = &plmanager.Option{
			ScriptMap: map[string]string{pipelineSource: pipelineSource + ".p"},
		}
	}
	if err := ipt.feeder.FeedV2(point.RUM, pts,
		dkio.WithInputName(inputName),
	); err != nil {
		log.Error(err.Error())
		httpErr(resp, err)
		return
	}

	if query.Get(httpapi.ArgEchoLineProto) != "" {
		var arr []string
		for _, x := range pts {
			arr = append(arr, x.LineProto())
		}
		httpOK(resp, strings.Join(arr, "\n"))
		return
	}

	httpOK(resp, nil)
}

func httpOK(w http.ResponseWriter, body interface{}) {
	if body == nil {
		if err := writeBody(w, httpapi.OK.HttpCode, binding.MIMEJSON, nil); err != nil {
			log.Error(err.Error())
		}

		return
	}

	var (
		bodyBytes   []byte
		contentType string
		err         error
	)
	switch x := body.(type) {
	case []byte:
		bodyBytes = x
	default:
		resp := &uhttp.BodyResp{
			HttpError: httpapi.OK,
			Content:   body,
		}
		contentType = `application/json`

		if bodyBytes, err = json.Marshal(resp); err != nil {
			log.Error(err.Error())
			jsonReturnf(uhttp.NewErr(err, http.StatusInternalServerError), w, "%s: %+#v", "json.Marshal() failed", resp)

			return
		}
	}

	if err := writeBody(w, httpapi.OK.HttpCode, contentType, bodyBytes); err != nil {
		log.Error(err.Error())
	}
}

func httpErr(w http.ResponseWriter, err error) {
	switch e := err.(type) { // nolint:errorlint
	case *uhttp.HttpError:
		jsonReturnf(e, w, "")
	case *uhttp.MsgError:
		if e.Args != nil {
			jsonReturnf(e.HttpError, w, e.Fmt, e.Args)
		}
	default:
		jsonReturnf(uhttp.NewErr(err, http.StatusInternalServerError), w, "")
	}
}

func writeBody(w http.ResponseWriter, statusCode int, contentType string, body []byte) error {
	w.WriteHeader(statusCode)
	if body != nil {
		w.Header().Set("Content-Type", contentType)
		n, err := w.Write(body)
		if n != len(body) || err != nil {
			return fmt.Errorf("unable to send http response, full body length(%d), send length(%d), err: %w", len(body), n, err)
		}
	}
	return nil
}

func jsonReturnf(httpErr *uhttp.HttpError, w http.ResponseWriter, format string, args ...interface{}) {
	resp := &uhttp.BodyResp{
		HttpError: httpErr,
	}

	if args != nil {
		resp.Message = fmt.Sprintf(format, args...)
	}

	buf, err := json.Marshal(resp)
	if err != nil {
		jsonReturnf(uhttp.NewErr(err, http.StatusInternalServerError), w, "%s: %+#v", "json.Marshal() failed", resp)

		return
	}

	if err := writeBody(w, httpErr.HttpCode, binding.MIMEJSON, buf); err != nil {
		log.Error(err.Error())
	}
}
