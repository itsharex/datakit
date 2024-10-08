// Unless explicitly stated otherwise all files in this repository are licensed
// under the MIT License.
// This product includes software developed at Guance Cloud (https://www.guance.com/).
// Copyright 2021-present Guance, Inc.

package ddtrace

import (
	"errors"
	"fmt"
	"io"
	"mime"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/GuanceCloud/cliutils/point"
	jsoniter "github.com/json-iterator/go"
	"github.com/tinylib/msgp/msgp"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/internal/bufpool"
	itrace "gitlab.jiagouyun.com/cloudcare-tools/datakit/internal/trace"
)

const (
	// headerRatesPayloadVersion contains the version of sampling rates.
	// If both agent and client have the same version, the agent won't return rates in API response.
	headerRatesPayloadVersion = "Datadog-Rates-Payload-Version"
)

const (
	// KeySamplingPriority is the key of the sampling priority value in the metrics map of the root span.
	keyPriority = "_sampling_priority_v1"
)

var jsonIterator = jsoniter.ConfigFastest

func httpStatusRespFunc(resp http.ResponseWriter, req *http.Request, err error) {
	if err != nil {
		log.Error(err.Error())
		resp.WriteHeader(http.StatusBadRequest)

		return
	}

	switch req.URL.Path {
	case v1, v2, v3:
		io.WriteString(resp, "OK\n") // nolint: errcheck,gosec
	default:
		resp.Header().Set("Content-Type", "application/json")
		resp.Header().Set(headerRatesPayloadVersion, req.Header.Get(headerRatesPayloadVersion))
		resp.WriteHeader(http.StatusOK)
		resp.Write([]byte(`{}`)) // nolint: errcheck,gosec
	}
}

func handleDDTraces(resp http.ResponseWriter, req *http.Request) {
	if req.Header.Get("Content-Length") == "0" || req.Header.Get("X-Datadog-Trace-Count") == "0" {
		log.Debug("empty request body")
		httpStatusRespFunc(resp, req, nil)

		return
	}

	pbuf := bufpool.GetBuffer()
	defer bufpool.PutBuffer(pbuf)

	_, err := io.Copy(pbuf, req.Body)
	if err != nil {
		log.Error(err.Error())
		resp.WriteHeader(http.StatusBadRequest)

		return
	}

	param := &itrace.TraceParameters{
		URLPath: req.URL.Path,
		Media:   req.Header.Get("Content-Type"),
		Body:    pbuf,
	}
	log.Debugf("param body len=%d", param.Body.Len())
	if err = parseDDTraces(param); err != nil {
		if errors.Is(err, msgp.ErrShortBytes) {
			log.Warn(err.Error())
		} else {
			log.Errorf("### parse ddtrace failed: %s", err.Error())
		}
		resp.WriteHeader(http.StatusBadRequest)

		return
	}

	httpStatusRespFunc(resp, req, nil)
}

// TODO:.
func handleDDStats(resp http.ResponseWriter, req *http.Request) {
	log.Infof("### %s unsupported yet", req.URL.Path)
	resp.WriteHeader(http.StatusNotFound)
}

func (ipt *Input) handleDDProxy(resp http.ResponseWriter, req *http.Request) {
	bts, err := io.ReadAll(req.Body)
	defer req.Body.Close() //nolint
	if err != nil {
		log.Warnf("read body err=%v", err)
		resp.WriteHeader(http.StatusBadRequest)
		return
	}

	if ipt.om != nil {
		ipt.om.parseTelemetryRequest(req.Header, bts)
	}

	resp.WriteHeader(http.StatusOK)
}

func parseDDTraces(param *itrace.TraceParameters) error {
	dktraces, err := decodeDDTraces(param)
	if err != nil {
		return err
	}

	if len(dktraces) != 0 && afterGatherRun != nil {
		afterGatherRun.Run(inputName, dktraces)
	}

	return nil
}

func decodeDDTraces(param *itrace.TraceParameters) (itrace.DatakitTraces, error) {
	var (
		err      error
		dktraces itrace.DatakitTraces
	)
	traces := ddtracePool.Get().(DDTraces)
	defer func() {
		traces.reset()
		ddtracePool.Put(traces) //nolint
	}()

	switch param.URLPath {
	case v1:
		var spans DDTrace
		if err := jsonIterator.Unmarshal(param.Body.Bytes(), &spans); err != nil {
			return nil, err
		}
		traces = mergeSpans(spans)
	case v5:
		if err = traces.UnmarshalMsgDictionary(param.Body.Bytes()); err == nil {
			traces = mergeTraces(traces)
		}
	default:
		if err = decodeRequest(param, &traces); err != nil {
			// traces = mergeTraces(traces)
			return nil, err
		}
	}

	if len(traces) != 0 {
		for _, trace := range traces {
			if len(trace) == 0 {
				log.Debug("### empty trace in traces")
				continue
			}
			if dktrace := ddtraceToDkTrace(trace); len(dktrace) != 0 {
				dktraces = append(dktraces, dktrace)
			}
		}
	}
	return dktraces, err
}

func decodeRequest(param *itrace.TraceParameters, out *DDTraces) error {
	mediaType, _, err := mime.ParseMediaType(param.Media)
	if err != nil {
		log.Debug(err.Error())
	}
	switch mediaType {
	case "application/msgpack":
		_, err = out.UnmarshalMsg(param.Body.Bytes())
	case "application/json", "text/json", "":
		return jsonIterator.Unmarshal(param.Body.Bytes(), out)
	default:
		// do our best
		if err1 := jsonIterator.Unmarshal(param.Body.Bytes(), out); err1 != nil {
			if _, err2 := out.UnmarshalMsg(param.Body.Bytes()); err2 != nil {
				err = fmt.Errorf("### could not decode JSON (err:%s), nor Msgpack (err:%s)", err1.Error(), err2.Error()) // nolint:errorlint
			}
		}
	}

	return err
}

func mergeSpans(trace DDTrace) DDTraces {
	var (
		traces = DDTraces{}
		byID   = make(map[uint64]DDTrace)
	)
	for _, span := range trace {
		byID[span.TraceID] = append(byID[span.TraceID], span)
	}
	for _, trace := range byID {
		traces = append(traces, trace)
	}

	return traces
}

func mergeTraces(traces DDTraces) DDTraces {
	var (
		merged DDTraces
		byID   = make(map[uint64]DDTrace)
	)
	for i := range traces {
		if len(traces[i]) != 0 {
			byID[traces[i][0].TraceID] = append(byID[traces[i][0].TraceID], traces[i]...)
		}
	}
	for _, trace := range byID {
		merged = append(merged, trace)
	}

	return merged
}

var (
	ignoreTraceIDFromTag = false
	traceOpts            = []point.Option{}
)

func ddtraceToDkTrace(trace DDTrace) itrace.DatakitTrace {
	var (
		dktrace            itrace.DatakitTrace
		parentIDs, spanIDs = gatherSpansInfo(trace)
	)

	for _, span := range trace {
		if span == nil {
			continue
		}

		var spanKV point.KVs
		priority, ok := span.Metrics[keyPriority]
		if ok {
			if priority == -1 || priority == -3 {
				log.Debugf("drop this traceID=%s service=%s", span.TraceID, span.Service)
				return []*itrace.DkSpan{} // 此处应该返回空的数组。
			}

			if p, ok := itrace.DDPriorityMap[int(priority)]; ok {
				// 在采样的结果放到行协议中，如果 DK 有配置采样，则需要该值进行过滤。
				spanKV = spanKV.Add(itrace.SampleRateKey, p, true, true)
			}
		}

		spanKV = spanKV.Add(itrace.FieldTraceID, strconv.FormatUint(span.TraceID, traceBase), false, false).
			Add(itrace.FieldParentID, strconv.FormatUint(span.ParentID, spanBase), false, false).
			Add(itrace.FieldSpanid, strconv.FormatUint(span.SpanID, spanBase), false, false).
			AddTag(itrace.TagService, span.Service).
			Add(itrace.FieldResource, strings.ReplaceAll(span.Resource, "\n", " "), false, false).
			AddTag(itrace.TagOperation, span.Name).
			AddTag(itrace.TagSource, inputName).
			Add(itrace.TagSpanType,
				itrace.FindSpanTypeInMultiServersIntSpanID(span.SpanID, span.ParentID, span.Service, spanIDs, parentIDs), true, false).
			AddTag(itrace.TagSourceType, itrace.GetSpanSourceType(span.Type)).
			Add(itrace.FieldStart, span.Start/int64(time.Microsecond), false, false).
			Add(itrace.FieldDuration, span.Duration/int64(time.Microsecond), false, false)

		// runtime_id 作为链路和 profiling 关联的字段，由于历史问题，需要增加一个冗余字段。
		runTimeIDKey := "runtime-id"
		if v, ok := span.Meta[runTimeIDKey]; ok {
			spanKV = spanKV.AddTag("runtime_id", v).AddTag(runTimeIDKey, v)
			delete(span.Meta, runTimeIDKey)
		}
		if v, ok := span.Meta["trace_128_bit_id"]; !ignoreTraceIDFromTag && ok {
			spanKV = spanKV.Add(itrace.FieldTraceID, v, false, true)
		}

		for k, v := range tags {
			spanKV = spanKV.AddTag(k, v)
		}

		for k, v := range span.Meta {
			ddTagsLock.RLock()
			if replace, ok := ddTags[k]; ok {
				if len(v) > 1024 {
					spanKV = spanKV.Add(replace, v, false, true)
				} else {
					spanKV = spanKV.MustAddTag(replace, v)
				}
				// 从 message 中删除 key.
				delete(span.Meta, k)
			}
			ddTagsLock.RUnlock()
		}

		if span.Error != 0 {
			spanKV = spanKV.AddTag(itrace.TagSpanStatus, itrace.StatusErr)
		} else {
			spanKV = spanKV.AddTag(itrace.TagSpanStatus, itrace.StatusOk)
		}

		if !delMessage {
			if buf, err := jsonIterator.Marshal(span); err != nil {
				log.Warn(err.Error())
			} else {
				spanKV = spanKV.Add(itrace.FieldMessage, string(buf), false, false)
			}
		}

		t := time.Unix(span.Start/1e9, span.Start%1e9)
		pt := point.NewPointV2(inputName, spanKV, append(traceOpts, point.WithTime(t))...)
		dktrace = append(dktrace, &itrace.DkSpan{Point: pt})
	}

	return dktrace
}

func gatherSpansInfo(trace DDTrace) (parentIDs map[uint64]bool, spanIDs map[uint64]string) {
	parentIDs = make(map[uint64]bool)
	spanIDs = make(map[uint64]string)
	for _, span := range trace {
		if span == nil {
			continue
		}
		spanIDs[span.SpanID] = span.Service
		parentIDs[span.ParentID] = true
	}

	return
}
