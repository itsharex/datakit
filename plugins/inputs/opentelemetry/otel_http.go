// Unless explicitly stated otherwise all files in this repository are licensed
// under the MIT License.
// This product includes software developed at Guance Cloud (https://www.guance.com/).
// Copyright 2021-present Guance, Inc.

package opentelemetry

import (
	"net/http"

	"gitlab.jiagouyun.com/cloudcare-tools/datakit"
	itrace "gitlab.jiagouyun.com/cloudcare-tools/datakit/internal/trace"
	dkio "gitlab.jiagouyun.com/cloudcare-tools/datakit/io"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/io/point"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/opentelemetry/compiled/v1/collector/metrics"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/opentelemetry/compiled/v1/collector/trace"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func httpStatusRespFunc(resp http.ResponseWriter, req *http.Request, err error) {
	if err != nil {
		log.Error(err.Error())
		resp.WriteHeader(http.StatusBadRequest)

		return
	}

	buf, err := proto.Marshal(&trace.ExportTraceServiceResponse{})
	if err != nil {
		log.Error(err.Error())
		resp.WriteHeader(http.StatusInternalServerError)

		return
	}

	resp.WriteHeader(http.StatusOK)
	resp.Write(buf)
}

func handleOTELTrace(resp http.ResponseWriter, req *http.Request) {
	media, _, buf, err := itrace.ParseTracerRequest(req)
	if err != nil {
		log.Error(err.Error())
		resp.WriteHeader(http.StatusBadRequest)

		return
	}

	tsreq := &trace.ExportTraceServiceRequest{}
	switch media {
	case "application/x-protobuf":
		err = proto.Unmarshal(buf, tsreq)
	case "application/json":
		err = protojson.Unmarshal(buf, tsreq)
	default:
		log.Error("unrecognized Content-Type")
		resp.WriteHeader(http.StatusBadRequest)

		return
	}

	if afterGatherRun != nil {
		if dktraces := parseResourceSpans(tsreq.ResourceSpans); len(dktraces) != 0 {
			afterGatherRun.Run(inputName, dktraces, false)
		}
	}
}

func handleOTElMetrics(resp http.ResponseWriter, req *http.Request) {
	media, _, buf, err := itrace.ParseTracerRequest(req)
	if err != nil {
		log.Error(err.Error())
		resp.WriteHeader(http.StatusBadRequest)

		return
	}

	msreq := &metrics.ExportMetricsServiceRequest{}
	switch media {
	case "application/x-protobuf":
		err = proto.Unmarshal(buf, msreq)
	case "application/json":
		err = protojson.Unmarshal(buf, msreq)
	default:
		log.Error("unrecognized Content-Type")
		resp.WriteHeader(http.StatusBadRequest)

		return
	}

	omcs := parseResourceMetrics(msreq.ResourceMetrics)
	var points []*point.Point
	for i := range omcs {
		if pts := omcs[i].getPoints(); len(pts) != 0 {
			points = append(points, pts...)
		}
	}
	if len(points) != 0 {
		if err = dkio.Feed(inputName, datakit.Metric, points, &dkio.Option{HighFreq: true}); err != nil {
			log.Error(err.Error())
		}
	}
}
