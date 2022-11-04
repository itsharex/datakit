// Unless explicitly stated otherwise all files in this repository are licensed
// under the MIT License.
// This product includes software developed at Guance Cloud (https://www.guance.com/).
// Copyright 2021-present Guance, Inc.

// Package skywalkingapi handle SkyWalking tracing metrics.
package skywalkingapi

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"gitlab.jiagouyun.com/cloudcare-tools/datakit/internal/storage"
	itrace "gitlab.jiagouyun.com/cloudcare-tools/datakit/internal/trace"
	commonv3 "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/skywalking/compiled/common/v3"
	agentv3 "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/skywalking/compiled/language/agent/v3"
	"google.golang.org/protobuf/proto"
)

func (api *SkyAPI) ProcessSegment(segment *agentv3.SegmentObject) {
	if api.localCache == nil || !api.localCache.Enabled() {
		api.parseSegmentObject(segment)
	} else {
		if buf, err := proto.Marshal(segment); err != nil {
			api.log.Error(err.Error())
		} else {
			if err = api.localCache.Put(storage.SKY_WALKING_GRPC_KEY, buf); err != nil {
				api.log.Error(err.Error())
			}
		}
	}
}

func (api *SkyAPI) parseSegmentObject(segment *agentv3.SegmentObject) {
	var dktrace itrace.DatakitTrace
	for _, span := range segment.Spans {
		if span == nil {
			continue
		}

		dkspan := &itrace.DatakitSpan{
			TraceID:    segment.TraceId,
			SpanID:     fmt.Sprintf("%s%d", segment.TraceSegmentId, span.SpanId),
			Service:    segment.Service,
			Resource:   span.OperationName,
			Operation:  span.OperationName,
			Source:     api.inputName,
			SourceType: itrace.SPAN_SOURCE_CUSTOMER,
			Start:      span.StartTime * int64(time.Millisecond),
			Duration:   (span.EndTime - span.StartTime) * int64(time.Millisecond),
		}

		if span.ParentSpanId < 0 {
			if len(span.Refs) > 0 {
				dkspan.ParentID = fmt.Sprintf("%s%d", span.Refs[0].ParentTraceSegmentId, span.Refs[0].ParentSpanId)
				if span.Refs[0].RefType == agentv3.RefType_CrossProcess && strings.Contains(span.Refs[0].ParentService, "_rum_") {
					dktrace = append(dktrace, &itrace.DatakitSpan{
						TraceID:    segment.TraceId,
						ParentID:   "0",
						SpanID:     dkspan.ParentID,
						Service:    span.Refs[0].ParentService,
						Resource:   span.Refs[0].ParentEndpoint,
						Operation:  span.Refs[0].ParentEndpoint,
						Source:     api.inputName,
						SpanType:   itrace.SPAN_TYPE_ENTRY,
						SourceType: itrace.SPAN_SOURCE_WEB,
						Start:      dkspan.Start - int64(time.Millisecond),
						Duration:   int64(time.Millisecond),
						Status:     itrace.STATUS_OK,
					})
					if endpoint := span.Refs[0].GetNetworkAddressUsedAtPeer(); endpoint != "" {
						dkspan.Tags = map[string]string{itrace.TAG_ENDPOINT: endpoint}
					}
				}
			} else {
				dkspan.ParentID = "0"
			}
		} else {
			if len(span.Refs) > 0 {
				dkspan.ParentID = fmt.Sprintf("%s%d", span.Refs[0].ParentTraceSegmentId, span.Refs[0].ParentSpanId)
			} else {
				dkspan.ParentID = fmt.Sprintf("%s%d", segment.TraceSegmentId, span.ParentSpanId)
			}
		}

		dkspan.Status = itrace.STATUS_OK
		if span.IsError {
			dkspan.Status = itrace.STATUS_ERR
		}

		switch span.SpanType {
		case agentv3.SpanType_Entry:
			dkspan.SpanType = itrace.SPAN_TYPE_ENTRY
		case agentv3.SpanType_Local:
			dkspan.SpanType = itrace.SPAN_TYPE_LOCAL
		case agentv3.SpanType_Exit:
			dkspan.SpanType = itrace.SPAN_TYPE_EXIT
		default:
			dkspan.SpanType = itrace.SPAN_TYPE_ENTRY
		}

		for i := range api.plugins {
			if value, ok := getTagValue(span.Tags, api.plugins[i]); ok {
				dkspan.Service = value
				dkspan.SpanType = itrace.SPAN_TYPE_ENTRY
				dkspan.SourceType = mapToSpanSourceType(span.SpanLayer)
				switch span.SpanLayer { // nolint: exhaustive
				case agentv3.SpanLayer_Database, agentv3.SpanLayer_Cache:
					if res, ok := getTagValue(span.Tags, "db.statement"); ok {
						dkspan.Resource = res
					}
				case agentv3.SpanLayer_MQ:
				case agentv3.SpanLayer_Http:
				case agentv3.SpanLayer_RPCFramework:
				case agentv3.SpanLayer_FAAS:
				case agentv3.SpanLayer_Unknown:
				}
			}
		}

		sourceTags := make(map[string]string)
		for _, tag := range span.Tags {
			sourceTags[tag.Key] = tag.Value
		}
		dkspan.Tags = itrace.MergeInToCustomerTags(api.customerKeys, api.tags, sourceTags)
		if span.Peer != "" {
			dkspan.Tags[itrace.TAG_ENDPOINT] = span.Peer
		}

		if buf, err := json.Marshal(span); err != nil {
			api.log.Warn(err.Error())
		} else {
			dkspan.Content = string(buf)
		}

		dktrace = append(dktrace, dkspan)
	}
	if len(dktrace) != 0 {
		dktrace[0].Metrics = make(map[string]interface{})
		dktrace[0].Metrics[itrace.FIELD_PRIORITY] = itrace.PRIORITY_AUTO_KEEP
	}

	if len(dktrace) != 0 && api.afterGatherRun != nil {
		api.afterGatherRun.Run(api.inputName, itrace.DatakitTraces{dktrace}, false)
	}
}

func getTagValue(tags []*commonv3.KeyStringValuePair, key string) (value string, ok bool) {
	for i := range tags {
		if key == tags[i].Key {
			if len(tags[i].Value) == 0 {
				return "", false
			} else {
				return tags[i].Value, true
			}
		}
	}

	return "", false
}

func mapToSpanSourceType(layer agentv3.SpanLayer) string {
	switch layer {
	case agentv3.SpanLayer_Database:
		return itrace.SPAN_SOURCE_DB
	case agentv3.SpanLayer_Cache:
		return itrace.SPAN_SOURCE_CACHE
	case agentv3.SpanLayer_RPCFramework:
		return itrace.SPAN_SOURCE_FRAMEWORK
	case agentv3.SpanLayer_Http:
		return itrace.SPAN_SOURCE_WEB
	case agentv3.SpanLayer_MQ:
		return itrace.SPAN_SOURCE_MSGQUE
	case agentv3.SpanLayer_FAAS:
		return itrace.SPAN_SOURCE_APP
	case agentv3.SpanLayer_Unknown:
		return itrace.SPAN_SOURCE_CUSTOMER
	default:
		return itrace.SPAN_SOURCE_CUSTOMER
	}
}
