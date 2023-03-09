// Copyright 2019, OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: opentelemetry/proto/collector/trace/v1/trace_service.proto

package trace

import (
	trace "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/opentelemetry/compiled/v1/trace"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ExportTraceServiceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// An array of ResourceSpans.
	// For data coming from a single resource this array will typically contain
	// one element. Intermediary nodes (such as OpenTelemetry Collector) that
	// receive data from multiple origins typically batch the data before
	// forwarding further and in that case this array will contain multiple
	// elements.
	ResourceSpans []*trace.ResourceSpans `protobuf:"bytes,1,rep,name=resource_spans,json=resourceSpans,proto3" json:"resource_spans,omitempty"`
}

func (x *ExportTraceServiceRequest) Reset() {
	*x = ExportTraceServiceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_opentelemetry_proto_collector_trace_v1_trace_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExportTraceServiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExportTraceServiceRequest) ProtoMessage() {}

func (x *ExportTraceServiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_opentelemetry_proto_collector_trace_v1_trace_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExportTraceServiceRequest.ProtoReflect.Descriptor instead.
func (*ExportTraceServiceRequest) Descriptor() ([]byte, []int) {
	return file_opentelemetry_proto_collector_trace_v1_trace_service_proto_rawDescGZIP(), []int{0}
}

func (x *ExportTraceServiceRequest) GetResourceSpans() []*trace.ResourceSpans {
	if x != nil {
		return x.ResourceSpans
	}
	return nil
}

type ExportTraceServiceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The details of a partially successful export request.
	//
	// If the request is only partially accepted
	// (i.e. when the server accepts only parts of the data and rejects the rest)
	// the server MUST initialize the `partial_success` field and MUST
	// set the `rejected_<signal>` with the number of items it rejected.
	//
	// Servers MAY also make use of the `partial_success` field to convey
	// warnings/suggestions to senders even when the request was fully accepted.
	// In such cases, the `rejected_<signal>` MUST have a value of `0` and
	// the `error_message` MUST be non-empty.
	//
	// A `partial_success` message with an empty value (rejected_<signal> = 0 and
	// `error_message` = "") is equivalent to it not being set/present. Senders
	// SHOULD interpret it the same way as in the full success case.
	PartialSuccess *ExportTracePartialSuccess `protobuf:"bytes,1,opt,name=partial_success,json=partialSuccess,proto3" json:"partial_success,omitempty"`
}

func (x *ExportTraceServiceResponse) Reset() {
	*x = ExportTraceServiceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_opentelemetry_proto_collector_trace_v1_trace_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExportTraceServiceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExportTraceServiceResponse) ProtoMessage() {}

func (x *ExportTraceServiceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_opentelemetry_proto_collector_trace_v1_trace_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExportTraceServiceResponse.ProtoReflect.Descriptor instead.
func (*ExportTraceServiceResponse) Descriptor() ([]byte, []int) {
	return file_opentelemetry_proto_collector_trace_v1_trace_service_proto_rawDescGZIP(), []int{1}
}

func (x *ExportTraceServiceResponse) GetPartialSuccess() *ExportTracePartialSuccess {
	if x != nil {
		return x.PartialSuccess
	}
	return nil
}

type ExportTracePartialSuccess struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The number of rejected spans.
	//
	// A `rejected_<signal>` field holding a `0` value indicates that the
	// request was fully accepted.
	RejectedSpans int64 `protobuf:"varint,1,opt,name=rejected_spans,json=rejectedSpans,proto3" json:"rejected_spans,omitempty"`
	// A developer-facing human-readable message in English. It should be used
	// either to explain why the server rejected parts of the data during a
	// partial success or to convey warnings/suggestions during a full success.
	// The message should offer guidance on how users can address such issues.
	//
	// error_message is an optional field. An error_message with an empty value
	// is equivalent to it not being set.
	ErrorMessage string `protobuf:"bytes,2,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
}

func (x *ExportTracePartialSuccess) Reset() {
	*x = ExportTracePartialSuccess{}
	if protoimpl.UnsafeEnabled {
		mi := &file_opentelemetry_proto_collector_trace_v1_trace_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExportTracePartialSuccess) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExportTracePartialSuccess) ProtoMessage() {}

func (x *ExportTracePartialSuccess) ProtoReflect() protoreflect.Message {
	mi := &file_opentelemetry_proto_collector_trace_v1_trace_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExportTracePartialSuccess.ProtoReflect.Descriptor instead.
func (*ExportTracePartialSuccess) Descriptor() ([]byte, []int) {
	return file_opentelemetry_proto_collector_trace_v1_trace_service_proto_rawDescGZIP(), []int{2}
}

func (x *ExportTracePartialSuccess) GetRejectedSpans() int64 {
	if x != nil {
		return x.RejectedSpans
	}
	return 0
}

func (x *ExportTracePartialSuccess) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

var File_opentelemetry_proto_collector_trace_v1_trace_service_proto protoreflect.FileDescriptor

var file_opentelemetry_proto_collector_trace_v1_trace_service_proto_rawDesc = []byte{
	0x0a, 0x3a, 0x6f, 0x70, 0x65, 0x6e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2f,
	0x74, 0x72, 0x61, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x72, 0x61, 0x63, 0x65, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x26, 0x6f, 0x70,
	0x65, 0x6e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x74, 0x72, 0x61, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x1a, 0x28, 0x6f, 0x70, 0x65, 0x6e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65,
	0x74, 0x72, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x72, 0x61, 0x63, 0x65, 0x2f,
	0x76, 0x31, 0x2f, 0x74, 0x72, 0x61, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6f,
	0x0a, 0x19, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x72, 0x61, 0x63, 0x65, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x52, 0x0a, 0x0e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x73, 0x70, 0x61, 0x6e, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65,
	0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x70, 0x61, 0x6e, 0x73,
	0x52, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x70, 0x61, 0x6e, 0x73, 0x22,
	0x88, 0x01, 0x0a, 0x1a, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x72, 0x61, 0x63, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6a,
	0x0a, 0x0f, 0x70, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x41, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x74, 0x65,
	0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x72, 0x61, 0x63, 0x65, 0x50, 0x61, 0x72, 0x74,
	0x69, 0x61, 0x6c, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x0e, 0x70, 0x61, 0x72, 0x74,
	0x69, 0x61, 0x6c, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x67, 0x0a, 0x19, 0x45, 0x78,
	0x70, 0x6f, 0x72, 0x74, 0x54, 0x72, 0x61, 0x63, 0x65, 0x50, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c,
	0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65, 0x6a, 0x65, 0x63,
	0x74, 0x65, 0x64, 0x5f, 0x73, 0x70, 0x61, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0d, 0x72, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x65, 0x64, 0x53, 0x70, 0x61, 0x6e, 0x73, 0x12, 0x23,
	0x0a, 0x0d, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x32, 0xa2, 0x01, 0x0a, 0x0c, 0x54, 0x72, 0x61, 0x63, 0x65, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x91, 0x01, 0x0a, 0x06, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x12,
	0x41, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e,
	0x74, 0x72, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x54,
	0x72, 0x61, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x42, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74,
	0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x70, 0x6f,
	0x72, 0x74, 0x54, 0x72, 0x61, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0xd0, 0x01, 0x0a, 0x29, 0x69, 0x6f, 0x2e,
	0x6f, 0x70, 0x65, 0x6e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x74, 0x72,
	0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x11, 0x54, 0x72, 0x61, 0x63, 0x65, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x65, 0x67, 0x69, 0x74,
	0x6c, 0x61, 0x62, 0x2e, 0x6a, 0x69, 0x61, 0x67, 0x6f, 0x75, 0x79, 0x75, 0x6e, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x63, 0x61, 0x72, 0x65, 0x2d, 0x74, 0x6f, 0x6f, 0x6c,
	0x73, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x6b, 0x69, 0x74, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x73, 0x2f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x74, 0x65, 0x6c,
	0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x64, 0x2f,
	0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2f, 0x74, 0x72, 0x61,
	0x63, 0x65, 0xaa, 0x02, 0x26, 0x4f, 0x70, 0x65, 0x6e, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74,
	0x72, 0x79, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x65, 0x2e, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_opentelemetry_proto_collector_trace_v1_trace_service_proto_rawDescOnce sync.Once
	file_opentelemetry_proto_collector_trace_v1_trace_service_proto_rawDescData = file_opentelemetry_proto_collector_trace_v1_trace_service_proto_rawDesc
)

func file_opentelemetry_proto_collector_trace_v1_trace_service_proto_rawDescGZIP() []byte {
	file_opentelemetry_proto_collector_trace_v1_trace_service_proto_rawDescOnce.Do(func() {
		file_opentelemetry_proto_collector_trace_v1_trace_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_opentelemetry_proto_collector_trace_v1_trace_service_proto_rawDescData)
	})
	return file_opentelemetry_proto_collector_trace_v1_trace_service_proto_rawDescData
}

var file_opentelemetry_proto_collector_trace_v1_trace_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_opentelemetry_proto_collector_trace_v1_trace_service_proto_goTypes = []interface{}{
	(*ExportTraceServiceRequest)(nil),  // 0: opentelemetry.proto.collector.trace.v1.ExportTraceServiceRequest
	(*ExportTraceServiceResponse)(nil), // 1: opentelemetry.proto.collector.trace.v1.ExportTraceServiceResponse
	(*ExportTracePartialSuccess)(nil),  // 2: opentelemetry.proto.collector.trace.v1.ExportTracePartialSuccess
	(*trace.ResourceSpans)(nil),        // 3: opentelemetry.proto.trace.v1.ResourceSpans
}
var file_opentelemetry_proto_collector_trace_v1_trace_service_proto_depIdxs = []int32{
	3, // 0: opentelemetry.proto.collector.trace.v1.ExportTraceServiceRequest.resource_spans:type_name -> opentelemetry.proto.trace.v1.ResourceSpans
	2, // 1: opentelemetry.proto.collector.trace.v1.ExportTraceServiceResponse.partial_success:type_name -> opentelemetry.proto.collector.trace.v1.ExportTracePartialSuccess
	0, // 2: opentelemetry.proto.collector.trace.v1.TraceService.Export:input_type -> opentelemetry.proto.collector.trace.v1.ExportTraceServiceRequest
	1, // 3: opentelemetry.proto.collector.trace.v1.TraceService.Export:output_type -> opentelemetry.proto.collector.trace.v1.ExportTraceServiceResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_opentelemetry_proto_collector_trace_v1_trace_service_proto_init() }
func file_opentelemetry_proto_collector_trace_v1_trace_service_proto_init() {
	if File_opentelemetry_proto_collector_trace_v1_trace_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_opentelemetry_proto_collector_trace_v1_trace_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExportTraceServiceRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_opentelemetry_proto_collector_trace_v1_trace_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExportTraceServiceResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_opentelemetry_proto_collector_trace_v1_trace_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExportTracePartialSuccess); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_opentelemetry_proto_collector_trace_v1_trace_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_opentelemetry_proto_collector_trace_v1_trace_service_proto_goTypes,
		DependencyIndexes: file_opentelemetry_proto_collector_trace_v1_trace_service_proto_depIdxs,
		MessageInfos:      file_opentelemetry_proto_collector_trace_v1_trace_service_proto_msgTypes,
	}.Build()
	File_opentelemetry_proto_collector_trace_v1_trace_service_proto = out.File
	file_opentelemetry_proto_collector_trace_v1_trace_service_proto_rawDesc = nil
	file_opentelemetry_proto_collector_trace_v1_trace_service_proto_goTypes = nil
	file_opentelemetry_proto_collector_trace_v1_trace_service_proto_depIdxs = nil
}
