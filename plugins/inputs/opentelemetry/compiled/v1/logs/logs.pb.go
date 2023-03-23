// Copyright 2020, OpenTelemetry Authors
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
// source: opentelemetry/proto/logs/v1/logs.proto

package logs

import (
	common "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/opentelemetry/compiled/v1/common"
	resource "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/opentelemetry/compiled/v1/resource"
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

// Possible values for LogRecord.SeverityNumber.
type SeverityNumber int32

const (
	// UNSPECIFIED is the default SeverityNumber, it MUST NOT be used.
	SeverityNumber_SEVERITY_NUMBER_UNSPECIFIED SeverityNumber = 0
	SeverityNumber_SEVERITY_NUMBER_TRACE       SeverityNumber = 1
	SeverityNumber_SEVERITY_NUMBER_TRACE2      SeverityNumber = 2
	SeverityNumber_SEVERITY_NUMBER_TRACE3      SeverityNumber = 3
	SeverityNumber_SEVERITY_NUMBER_TRACE4      SeverityNumber = 4
	SeverityNumber_SEVERITY_NUMBER_DEBUG       SeverityNumber = 5
	SeverityNumber_SEVERITY_NUMBER_DEBUG2      SeverityNumber = 6
	SeverityNumber_SEVERITY_NUMBER_DEBUG3      SeverityNumber = 7
	SeverityNumber_SEVERITY_NUMBER_DEBUG4      SeverityNumber = 8
	SeverityNumber_SEVERITY_NUMBER_INFO        SeverityNumber = 9
	SeverityNumber_SEVERITY_NUMBER_INFO2       SeverityNumber = 10
	SeverityNumber_SEVERITY_NUMBER_INFO3       SeverityNumber = 11
	SeverityNumber_SEVERITY_NUMBER_INFO4       SeverityNumber = 12
	SeverityNumber_SEVERITY_NUMBER_WARN        SeverityNumber = 13
	SeverityNumber_SEVERITY_NUMBER_WARN2       SeverityNumber = 14
	SeverityNumber_SEVERITY_NUMBER_WARN3       SeverityNumber = 15
	SeverityNumber_SEVERITY_NUMBER_WARN4       SeverityNumber = 16
	SeverityNumber_SEVERITY_NUMBER_ERROR       SeverityNumber = 17
	SeverityNumber_SEVERITY_NUMBER_ERROR2      SeverityNumber = 18
	SeverityNumber_SEVERITY_NUMBER_ERROR3      SeverityNumber = 19
	SeverityNumber_SEVERITY_NUMBER_ERROR4      SeverityNumber = 20
	SeverityNumber_SEVERITY_NUMBER_FATAL       SeverityNumber = 21
	SeverityNumber_SEVERITY_NUMBER_FATAL2      SeverityNumber = 22
	SeverityNumber_SEVERITY_NUMBER_FATAL3      SeverityNumber = 23
	SeverityNumber_SEVERITY_NUMBER_FATAL4      SeverityNumber = 24
)

// Enum value maps for SeverityNumber.
var (
	SeverityNumber_name = map[int32]string{
		0:  "SEVERITY_NUMBER_UNSPECIFIED",
		1:  "SEVERITY_NUMBER_TRACE",
		2:  "SEVERITY_NUMBER_TRACE2",
		3:  "SEVERITY_NUMBER_TRACE3",
		4:  "SEVERITY_NUMBER_TRACE4",
		5:  "SEVERITY_NUMBER_DEBUG",
		6:  "SEVERITY_NUMBER_DEBUG2",
		7:  "SEVERITY_NUMBER_DEBUG3",
		8:  "SEVERITY_NUMBER_DEBUG4",
		9:  "SEVERITY_NUMBER_INFO",
		10: "SEVERITY_NUMBER_INFO2",
		11: "SEVERITY_NUMBER_INFO3",
		12: "SEVERITY_NUMBER_INFO4",
		13: "SEVERITY_NUMBER_WARN",
		14: "SEVERITY_NUMBER_WARN2",
		15: "SEVERITY_NUMBER_WARN3",
		16: "SEVERITY_NUMBER_WARN4",
		17: "SEVERITY_NUMBER_ERROR",
		18: "SEVERITY_NUMBER_ERROR2",
		19: "SEVERITY_NUMBER_ERROR3",
		20: "SEVERITY_NUMBER_ERROR4",
		21: "SEVERITY_NUMBER_FATAL",
		22: "SEVERITY_NUMBER_FATAL2",
		23: "SEVERITY_NUMBER_FATAL3",
		24: "SEVERITY_NUMBER_FATAL4",
	}
	SeverityNumber_value = map[string]int32{
		"SEVERITY_NUMBER_UNSPECIFIED": 0,
		"SEVERITY_NUMBER_TRACE":       1,
		"SEVERITY_NUMBER_TRACE2":      2,
		"SEVERITY_NUMBER_TRACE3":      3,
		"SEVERITY_NUMBER_TRACE4":      4,
		"SEVERITY_NUMBER_DEBUG":       5,
		"SEVERITY_NUMBER_DEBUG2":      6,
		"SEVERITY_NUMBER_DEBUG3":      7,
		"SEVERITY_NUMBER_DEBUG4":      8,
		"SEVERITY_NUMBER_INFO":        9,
		"SEVERITY_NUMBER_INFO2":       10,
		"SEVERITY_NUMBER_INFO3":       11,
		"SEVERITY_NUMBER_INFO4":       12,
		"SEVERITY_NUMBER_WARN":        13,
		"SEVERITY_NUMBER_WARN2":       14,
		"SEVERITY_NUMBER_WARN3":       15,
		"SEVERITY_NUMBER_WARN4":       16,
		"SEVERITY_NUMBER_ERROR":       17,
		"SEVERITY_NUMBER_ERROR2":      18,
		"SEVERITY_NUMBER_ERROR3":      19,
		"SEVERITY_NUMBER_ERROR4":      20,
		"SEVERITY_NUMBER_FATAL":       21,
		"SEVERITY_NUMBER_FATAL2":      22,
		"SEVERITY_NUMBER_FATAL3":      23,
		"SEVERITY_NUMBER_FATAL4":      24,
	}
)

func (x SeverityNumber) Enum() *SeverityNumber {
	p := new(SeverityNumber)
	*p = x
	return p
}

func (x SeverityNumber) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SeverityNumber) Descriptor() protoreflect.EnumDescriptor {
	return file_opentelemetry_proto_logs_v1_logs_proto_enumTypes[0].Descriptor()
}

func (SeverityNumber) Type() protoreflect.EnumType {
	return &file_opentelemetry_proto_logs_v1_logs_proto_enumTypes[0]
}

func (x SeverityNumber) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SeverityNumber.Descriptor instead.
func (SeverityNumber) EnumDescriptor() ([]byte, []int) {
	return file_opentelemetry_proto_logs_v1_logs_proto_rawDescGZIP(), []int{0}
}

// Masks for LogRecord.flags field.
type LogRecordFlags int32

const (
	LogRecordFlags_LOG_RECORD_FLAG_UNSPECIFIED      LogRecordFlags = 0
	LogRecordFlags_LOG_RECORD_FLAG_TRACE_FLAGS_MASK LogRecordFlags = 255
)

// Enum value maps for LogRecordFlags.
var (
	LogRecordFlags_name = map[int32]string{
		0:   "LOG_RECORD_FLAG_UNSPECIFIED",
		255: "LOG_RECORD_FLAG_TRACE_FLAGS_MASK",
	}
	LogRecordFlags_value = map[string]int32{
		"LOG_RECORD_FLAG_UNSPECIFIED":      0,
		"LOG_RECORD_FLAG_TRACE_FLAGS_MASK": 255,
	}
)

func (x LogRecordFlags) Enum() *LogRecordFlags {
	p := new(LogRecordFlags)
	*p = x
	return p
}

func (x LogRecordFlags) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LogRecordFlags) Descriptor() protoreflect.EnumDescriptor {
	return file_opentelemetry_proto_logs_v1_logs_proto_enumTypes[1].Descriptor()
}

func (LogRecordFlags) Type() protoreflect.EnumType {
	return &file_opentelemetry_proto_logs_v1_logs_proto_enumTypes[1]
}

func (x LogRecordFlags) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LogRecordFlags.Descriptor instead.
func (LogRecordFlags) EnumDescriptor() ([]byte, []int) {
	return file_opentelemetry_proto_logs_v1_logs_proto_rawDescGZIP(), []int{1}
}

// LogsData represents the logs data that can be stored in a persistent storage,
// OR can be embedded by other protocols that transfer OTLP logs data but do not
// implement the OTLP protocol.
//
// The main difference between this message and collector protocol is that
// in this message there will not be any "control" or "metadata" specific to
// OTLP protocol.
//
// When new fields are added into this message, the OTLP request MUST be updated
// as well.
type LogsData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// An array of ResourceLogs.
	// For data coming from a single resource this array will typically contain
	// one element. Intermediary nodes that receive data from multiple origins
	// typically batch the data before forwarding further and in that case this
	// array will contain multiple elements.
	ResourceLogs []*ResourceLogs `protobuf:"bytes,1,rep,name=resource_logs,json=resourceLogs,proto3" json:"resource_logs,omitempty"`
}

func (x *LogsData) Reset() {
	*x = LogsData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_opentelemetry_proto_logs_v1_logs_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogsData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogsData) ProtoMessage() {}

func (x *LogsData) ProtoReflect() protoreflect.Message {
	mi := &file_opentelemetry_proto_logs_v1_logs_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogsData.ProtoReflect.Descriptor instead.
func (*LogsData) Descriptor() ([]byte, []int) {
	return file_opentelemetry_proto_logs_v1_logs_proto_rawDescGZIP(), []int{0}
}

func (x *LogsData) GetResourceLogs() []*ResourceLogs {
	if x != nil {
		return x.ResourceLogs
	}
	return nil
}

// A collection of ScopeLogs from a Resource.
type ResourceLogs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The resource for the logs in this message.
	// If this field is not set then resource info is unknown.
	Resource *resource.Resource `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
	// A list of ScopeLogs that originate from a resource.
	ScopeLogs []*ScopeLogs `protobuf:"bytes,2,rep,name=scope_logs,json=scopeLogs,proto3" json:"scope_logs,omitempty"`
	// This schema_url applies to the data in the "resource" field. It does not
	// apply to the data in the "scope_logs" field which have their own schema_url
	// field.
	SchemaUrl string `protobuf:"bytes,3,opt,name=schema_url,json=schemaUrl,proto3" json:"schema_url,omitempty"`
}

func (x *ResourceLogs) Reset() {
	*x = ResourceLogs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_opentelemetry_proto_logs_v1_logs_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceLogs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceLogs) ProtoMessage() {}

func (x *ResourceLogs) ProtoReflect() protoreflect.Message {
	mi := &file_opentelemetry_proto_logs_v1_logs_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceLogs.ProtoReflect.Descriptor instead.
func (*ResourceLogs) Descriptor() ([]byte, []int) {
	return file_opentelemetry_proto_logs_v1_logs_proto_rawDescGZIP(), []int{1}
}

func (x *ResourceLogs) GetResource() *resource.Resource {
	if x != nil {
		return x.Resource
	}
	return nil
}

func (x *ResourceLogs) GetScopeLogs() []*ScopeLogs {
	if x != nil {
		return x.ScopeLogs
	}
	return nil
}

func (x *ResourceLogs) GetSchemaUrl() string {
	if x != nil {
		return x.SchemaUrl
	}
	return ""
}

// A collection of Logs produced by a Scope.
type ScopeLogs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The instrumentation scope information for the logs in this message.
	// Semantically when InstrumentationScope isn't set, it is equivalent with
	// an empty instrumentation scope name (unknown).
	Scope *common.InstrumentationScope `protobuf:"bytes,1,opt,name=scope,proto3" json:"scope,omitempty"`
	// A list of log records.
	LogRecords []*LogRecord `protobuf:"bytes,2,rep,name=log_records,json=logRecords,proto3" json:"log_records,omitempty"`
	// This schema_url applies to all logs in the "logs" field.
	SchemaUrl string `protobuf:"bytes,3,opt,name=schema_url,json=schemaUrl,proto3" json:"schema_url,omitempty"`
}

func (x *ScopeLogs) Reset() {
	*x = ScopeLogs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_opentelemetry_proto_logs_v1_logs_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScopeLogs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScopeLogs) ProtoMessage() {}

func (x *ScopeLogs) ProtoReflect() protoreflect.Message {
	mi := &file_opentelemetry_proto_logs_v1_logs_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScopeLogs.ProtoReflect.Descriptor instead.
func (*ScopeLogs) Descriptor() ([]byte, []int) {
	return file_opentelemetry_proto_logs_v1_logs_proto_rawDescGZIP(), []int{2}
}

func (x *ScopeLogs) GetScope() *common.InstrumentationScope {
	if x != nil {
		return x.Scope
	}
	return nil
}

func (x *ScopeLogs) GetLogRecords() []*LogRecord {
	if x != nil {
		return x.LogRecords
	}
	return nil
}

func (x *ScopeLogs) GetSchemaUrl() string {
	if x != nil {
		return x.SchemaUrl
	}
	return ""
}

// A log record according to OpenTelemetry Log Data Model:
// https://github.com/open-telemetry/oteps/blob/main/text/logs/0097-log-data-model.md
type LogRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// time_unix_nano is the time when the event occurred.
	// Value is UNIX Epoch time in nanoseconds since 00:00:00 UTC on 1 January
	// 1970. Value of 0 indicates unknown or missing timestamp.
	TimeUnixNano uint64 `protobuf:"fixed64,1,opt,name=time_unix_nano,json=timeUnixNano,proto3" json:"time_unix_nano,omitempty"`
	// Time when the event was observed by the collection system.
	// For events that originate in OpenTelemetry (e.g. using OpenTelemetry
	// Logging SDK) this timestamp is typically set at the generation time and is
	// equal to Timestamp. For events originating externally and collected by
	// OpenTelemetry (e.g. using Collector) this is the time when OpenTelemetry's
	// code observed the event measured by the clock of the OpenTelemetry code.
	// This field MUST be set once the event is observed by OpenTelemetry.
	//
	// For converting OpenTelemetry log data to formats that support only one
	// timestamp or when receiving OpenTelemetry log data by recipients that
	// support only one timestamp internally the following logic is recommended:
	//   - Use time_unix_nano if it is present, otherwise use
	//   observed_time_unix_nano.
	//
	// Value is UNIX Epoch time in nanoseconds since 00:00:00 UTC on 1 January
	// 1970. Value of 0 indicates unknown or missing timestamp.
	ObservedTimeUnixNano uint64 `protobuf:"fixed64,11,opt,name=observed_time_unix_nano,json=observedTimeUnixNano,proto3" json:"observed_time_unix_nano,omitempty"`
	// Numerical value of the severity, normalized to values described in Log Data
	// Model. [Optional].
	SeverityNumber SeverityNumber `protobuf:"varint,2,opt,name=severity_number,json=severityNumber,proto3,enum=opentelemetry.proto.logs.v1.SeverityNumber" json:"severity_number,omitempty"`
	// The severity text (also known as log level). The original string
	// representation as it is known at the source. [Optional].
	SeverityText string `protobuf:"bytes,3,opt,name=severity_text,json=severityText,proto3" json:"severity_text,omitempty"`
	// A value containing the body of the log record. Can be for example a
	// human-readable string message (including multi-line) describing the event
	// in a free form or it can be a structured data composed of arrays and maps
	// of other values. [Optional].
	Body *common.AnyValue `protobuf:"bytes,5,opt,name=body,proto3" json:"body,omitempty"`
	// Additional attributes that describe the specific event occurrence.
	// [Optional]. Attribute keys MUST be unique (it is not allowed to have more
	// than one attribute with the same key).
	Attributes             []*common.KeyValue `protobuf:"bytes,6,rep,name=attributes,proto3" json:"attributes,omitempty"`
	DroppedAttributesCount uint32             `protobuf:"varint,7,opt,name=dropped_attributes_count,json=droppedAttributesCount,proto3" json:"dropped_attributes_count,omitempty"`
	// Flags, a bit field. 8 least significant bits are the trace flags as
	// defined in W3C Trace Context specification. 24 most significant bits are
	// reserved and must be set to 0. Readers must not assume that 24 most
	// significant bits will be zero and must correctly mask the bits when reading
	// 8-bit trace flag (use flags & TRACE_FLAGS_MASK). [Optional].
	Flags uint32 `protobuf:"fixed32,8,opt,name=flags,proto3" json:"flags,omitempty"`
	// A unique identifier for a trace. All logs from the same trace share
	// the same `trace_id`. The ID is a 16-byte array. An ID with all zeroes OR
	// of length other than 16 bytes is considered invalid (empty string in
	// OTLP/JSON is zero-length and thus is also invalid).
	//
	// This field is optional.
	//
	// The receivers SHOULD assume that the log record is not associated with a
	// trace if any of the following is true:
	//   - the field is not present,
	//   - the field contains an invalid value.
	TraceId []byte `protobuf:"bytes,9,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
	// A unique identifier for a span within a trace, assigned when the span
	// is created. The ID is an 8-byte array. An ID with all zeroes OR of length
	// other than 8 bytes is considered invalid (empty string in OTLP/JSON
	// is zero-length and thus is also invalid).
	//
	// This field is optional. If the sender specifies a valid span_id then it
	// SHOULD also specify a valid trace_id.
	//
	// The receivers SHOULD assume that the log record is not associated with a
	// span if any of the following is true:
	//   - the field is not present,
	//   - the field contains an invalid value.
	SpanId []byte `protobuf:"bytes,10,opt,name=span_id,json=spanId,proto3" json:"span_id,omitempty"`
}

func (x *LogRecord) Reset() {
	*x = LogRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_opentelemetry_proto_logs_v1_logs_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogRecord) ProtoMessage() {}

func (x *LogRecord) ProtoReflect() protoreflect.Message {
	mi := &file_opentelemetry_proto_logs_v1_logs_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogRecord.ProtoReflect.Descriptor instead.
func (*LogRecord) Descriptor() ([]byte, []int) {
	return file_opentelemetry_proto_logs_v1_logs_proto_rawDescGZIP(), []int{3}
}

func (x *LogRecord) GetTimeUnixNano() uint64 {
	if x != nil {
		return x.TimeUnixNano
	}
	return 0
}

func (x *LogRecord) GetObservedTimeUnixNano() uint64 {
	if x != nil {
		return x.ObservedTimeUnixNano
	}
	return 0
}

func (x *LogRecord) GetSeverityNumber() SeverityNumber {
	if x != nil {
		return x.SeverityNumber
	}
	return SeverityNumber_SEVERITY_NUMBER_UNSPECIFIED
}

func (x *LogRecord) GetSeverityText() string {
	if x != nil {
		return x.SeverityText
	}
	return ""
}

func (x *LogRecord) GetBody() *common.AnyValue {
	if x != nil {
		return x.Body
	}
	return nil
}

func (x *LogRecord) GetAttributes() []*common.KeyValue {
	if x != nil {
		return x.Attributes
	}
	return nil
}

func (x *LogRecord) GetDroppedAttributesCount() uint32 {
	if x != nil {
		return x.DroppedAttributesCount
	}
	return 0
}

func (x *LogRecord) GetFlags() uint32 {
	if x != nil {
		return x.Flags
	}
	return 0
}

func (x *LogRecord) GetTraceId() []byte {
	if x != nil {
		return x.TraceId
	}
	return nil
}

func (x *LogRecord) GetSpanId() []byte {
	if x != nil {
		return x.SpanId
	}
	return nil
}

var File_opentelemetry_proto_logs_v1_logs_proto protoreflect.FileDescriptor

var file_opentelemetry_proto_logs_v1_logs_proto_rawDesc = []byte{
	0x0a, 0x26, 0x6f, 0x70, 0x65, 0x6e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6c, 0x6f, 0x67, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f,
	0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x6f, 0x70, 0x65, 0x6e, 0x74, 0x65,
	0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6c, 0x6f,
	0x67, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x2a, 0x6f, 0x70, 0x65, 0x6e, 0x74, 0x65, 0x6c, 0x65, 0x6d,
	0x65, 0x74, 0x72, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f,
	0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x5a, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x73, 0x44, 0x61, 0x74, 0x61, 0x12, 0x4e, 0x0a,
	0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6c, 0x6f, 0x67, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x74, 0x65, 0x6c, 0x65, 0x6d,
	0x65, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6c, 0x6f, 0x67, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4c, 0x6f, 0x67, 0x73, 0x52,
	0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4c, 0x6f, 0x67, 0x73, 0x22, 0xc3, 0x01,
	0x0a, 0x0c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4c, 0x6f, 0x67, 0x73, 0x12, 0x45,
	0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x29, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x45, 0x0a, 0x0a, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x5f, 0x6c,
	0x6f, 0x67, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x6f, 0x70, 0x65, 0x6e,
	0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x6c, 0x6f, 0x67, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x4c, 0x6f, 0x67,
	0x73, 0x52, 0x09, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x4c, 0x6f, 0x67, 0x73, 0x12, 0x1d, 0x0a, 0x0a,
	0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x55, 0x72, 0x6c, 0x4a, 0x06, 0x08, 0xe8, 0x07,
	0x10, 0xe9, 0x07, 0x22, 0xbe, 0x01, 0x0a, 0x09, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x4c, 0x6f, 0x67,
	0x73, 0x12, 0x49, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x33, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x63, 0x6f, 0x70, 0x65, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x47, 0x0a, 0x0b,
	0x6c, 0x6f, 0x67, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x26, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6c, 0x6f, 0x67, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x6f, 0x67, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x0a, 0x6c, 0x6f, 0x67, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x5f,
	0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x55, 0x72, 0x6c, 0x22, 0xf3, 0x03, 0x0a, 0x09, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x75, 0x6e, 0x69, 0x78, 0x5f,
	0x6e, 0x61, 0x6e, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x06, 0x52, 0x0c, 0x74, 0x69, 0x6d, 0x65,
	0x55, 0x6e, 0x69, 0x78, 0x4e, 0x61, 0x6e, 0x6f, 0x12, 0x35, 0x0a, 0x17, 0x6f, 0x62, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x75, 0x6e, 0x69, 0x78, 0x5f, 0x6e,
	0x61, 0x6e, 0x6f, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x06, 0x52, 0x14, 0x6f, 0x62, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x55, 0x6e, 0x69, 0x78, 0x4e, 0x61, 0x6e, 0x6f, 0x12,
	0x54, 0x0a, 0x0f, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2b, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x74,
	0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6c,
	0x6f, 0x67, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x0e, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74,
	0x79, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x65,
	0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x54, 0x65, 0x78, 0x74, 0x12, 0x3b, 0x0a, 0x04, 0x62, 0x6f,
	0x64, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x74,
	0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6e, 0x79, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x47, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x6f, 0x70,
	0x65, 0x6e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4b, 0x65, 0x79, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73,
	0x12, 0x38, 0x0a, 0x18, 0x64, 0x72, 0x6f, 0x70, 0x70, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x16, 0x64, 0x72, 0x6f, 0x70, 0x70, 0x65, 0x64, 0x41, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x6c,
	0x61, 0x67, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x07, 0x52, 0x05, 0x66, 0x6c, 0x61, 0x67, 0x73,
	0x12, 0x19, 0x0a, 0x08, 0x74, 0x72, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x07, 0x74, 0x72, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x73,
	0x70, 0x61, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x73, 0x70,
	0x61, 0x6e, 0x49, 0x64, 0x4a, 0x04, 0x08, 0x04, 0x10, 0x05, 0x2a, 0xc3, 0x05, 0x0a, 0x0e, 0x53,
	0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1f, 0x0a,
	0x1b, 0x53, 0x45, 0x56, 0x45, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x4e, 0x55, 0x4d, 0x42, 0x45, 0x52,
	0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x19,
	0x0a, 0x15, 0x53, 0x45, 0x56, 0x45, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x4e, 0x55, 0x4d, 0x42, 0x45,
	0x52, 0x5f, 0x54, 0x52, 0x41, 0x43, 0x45, 0x10, 0x01, 0x12, 0x1a, 0x0a, 0x16, 0x53, 0x45, 0x56,
	0x45, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x4e, 0x55, 0x4d, 0x42, 0x45, 0x52, 0x5f, 0x54, 0x52, 0x41,
	0x43, 0x45, 0x32, 0x10, 0x02, 0x12, 0x1a, 0x0a, 0x16, 0x53, 0x45, 0x56, 0x45, 0x52, 0x49, 0x54,
	0x59, 0x5f, 0x4e, 0x55, 0x4d, 0x42, 0x45, 0x52, 0x5f, 0x54, 0x52, 0x41, 0x43, 0x45, 0x33, 0x10,
	0x03, 0x12, 0x1a, 0x0a, 0x16, 0x53, 0x45, 0x56, 0x45, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x4e, 0x55,
	0x4d, 0x42, 0x45, 0x52, 0x5f, 0x54, 0x52, 0x41, 0x43, 0x45, 0x34, 0x10, 0x04, 0x12, 0x19, 0x0a,
	0x15, 0x53, 0x45, 0x56, 0x45, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x4e, 0x55, 0x4d, 0x42, 0x45, 0x52,
	0x5f, 0x44, 0x45, 0x42, 0x55, 0x47, 0x10, 0x05, 0x12, 0x1a, 0x0a, 0x16, 0x53, 0x45, 0x56, 0x45,
	0x52, 0x49, 0x54, 0x59, 0x5f, 0x4e, 0x55, 0x4d, 0x42, 0x45, 0x52, 0x5f, 0x44, 0x45, 0x42, 0x55,
	0x47, 0x32, 0x10, 0x06, 0x12, 0x1a, 0x0a, 0x16, 0x53, 0x45, 0x56, 0x45, 0x52, 0x49, 0x54, 0x59,
	0x5f, 0x4e, 0x55, 0x4d, 0x42, 0x45, 0x52, 0x5f, 0x44, 0x45, 0x42, 0x55, 0x47, 0x33, 0x10, 0x07,
	0x12, 0x1a, 0x0a, 0x16, 0x53, 0x45, 0x56, 0x45, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x4e, 0x55, 0x4d,
	0x42, 0x45, 0x52, 0x5f, 0x44, 0x45, 0x42, 0x55, 0x47, 0x34, 0x10, 0x08, 0x12, 0x18, 0x0a, 0x14,
	0x53, 0x45, 0x56, 0x45, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x4e, 0x55, 0x4d, 0x42, 0x45, 0x52, 0x5f,
	0x49, 0x4e, 0x46, 0x4f, 0x10, 0x09, 0x12, 0x19, 0x0a, 0x15, 0x53, 0x45, 0x56, 0x45, 0x52, 0x49,
	0x54, 0x59, 0x5f, 0x4e, 0x55, 0x4d, 0x42, 0x45, 0x52, 0x5f, 0x49, 0x4e, 0x46, 0x4f, 0x32, 0x10,
	0x0a, 0x12, 0x19, 0x0a, 0x15, 0x53, 0x45, 0x56, 0x45, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x4e, 0x55,
	0x4d, 0x42, 0x45, 0x52, 0x5f, 0x49, 0x4e, 0x46, 0x4f, 0x33, 0x10, 0x0b, 0x12, 0x19, 0x0a, 0x15,
	0x53, 0x45, 0x56, 0x45, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x4e, 0x55, 0x4d, 0x42, 0x45, 0x52, 0x5f,
	0x49, 0x4e, 0x46, 0x4f, 0x34, 0x10, 0x0c, 0x12, 0x18, 0x0a, 0x14, 0x53, 0x45, 0x56, 0x45, 0x52,
	0x49, 0x54, 0x59, 0x5f, 0x4e, 0x55, 0x4d, 0x42, 0x45, 0x52, 0x5f, 0x57, 0x41, 0x52, 0x4e, 0x10,
	0x0d, 0x12, 0x19, 0x0a, 0x15, 0x53, 0x45, 0x56, 0x45, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x4e, 0x55,
	0x4d, 0x42, 0x45, 0x52, 0x5f, 0x57, 0x41, 0x52, 0x4e, 0x32, 0x10, 0x0e, 0x12, 0x19, 0x0a, 0x15,
	0x53, 0x45, 0x56, 0x45, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x4e, 0x55, 0x4d, 0x42, 0x45, 0x52, 0x5f,
	0x57, 0x41, 0x52, 0x4e, 0x33, 0x10, 0x0f, 0x12, 0x19, 0x0a, 0x15, 0x53, 0x45, 0x56, 0x45, 0x52,
	0x49, 0x54, 0x59, 0x5f, 0x4e, 0x55, 0x4d, 0x42, 0x45, 0x52, 0x5f, 0x57, 0x41, 0x52, 0x4e, 0x34,
	0x10, 0x10, 0x12, 0x19, 0x0a, 0x15, 0x53, 0x45, 0x56, 0x45, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x4e,
	0x55, 0x4d, 0x42, 0x45, 0x52, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x11, 0x12, 0x1a, 0x0a,
	0x16, 0x53, 0x45, 0x56, 0x45, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x4e, 0x55, 0x4d, 0x42, 0x45, 0x52,
	0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x32, 0x10, 0x12, 0x12, 0x1a, 0x0a, 0x16, 0x53, 0x45, 0x56,
	0x45, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x4e, 0x55, 0x4d, 0x42, 0x45, 0x52, 0x5f, 0x45, 0x52, 0x52,
	0x4f, 0x52, 0x33, 0x10, 0x13, 0x12, 0x1a, 0x0a, 0x16, 0x53, 0x45, 0x56, 0x45, 0x52, 0x49, 0x54,
	0x59, 0x5f, 0x4e, 0x55, 0x4d, 0x42, 0x45, 0x52, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x34, 0x10,
	0x14, 0x12, 0x19, 0x0a, 0x15, 0x53, 0x45, 0x56, 0x45, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x4e, 0x55,
	0x4d, 0x42, 0x45, 0x52, 0x5f, 0x46, 0x41, 0x54, 0x41, 0x4c, 0x10, 0x15, 0x12, 0x1a, 0x0a, 0x16,
	0x53, 0x45, 0x56, 0x45, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x4e, 0x55, 0x4d, 0x42, 0x45, 0x52, 0x5f,
	0x46, 0x41, 0x54, 0x41, 0x4c, 0x32, 0x10, 0x16, 0x12, 0x1a, 0x0a, 0x16, 0x53, 0x45, 0x56, 0x45,
	0x52, 0x49, 0x54, 0x59, 0x5f, 0x4e, 0x55, 0x4d, 0x42, 0x45, 0x52, 0x5f, 0x46, 0x41, 0x54, 0x41,
	0x4c, 0x33, 0x10, 0x17, 0x12, 0x1a, 0x0a, 0x16, 0x53, 0x45, 0x56, 0x45, 0x52, 0x49, 0x54, 0x59,
	0x5f, 0x4e, 0x55, 0x4d, 0x42, 0x45, 0x52, 0x5f, 0x46, 0x41, 0x54, 0x41, 0x4c, 0x34, 0x10, 0x18,
	0x2a, 0x58, 0x0a, 0x0e, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x46, 0x6c, 0x61,
	0x67, 0x73, 0x12, 0x1f, 0x0a, 0x1b, 0x4c, 0x4f, 0x47, 0x5f, 0x52, 0x45, 0x43, 0x4f, 0x52, 0x44,
	0x5f, 0x46, 0x4c, 0x41, 0x47, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x25, 0x0a, 0x20, 0x4c, 0x4f, 0x47, 0x5f, 0x52, 0x45, 0x43, 0x4f, 0x52,
	0x44, 0x5f, 0x46, 0x4c, 0x41, 0x47, 0x5f, 0x54, 0x52, 0x41, 0x43, 0x45, 0x5f, 0x46, 0x4c, 0x41,
	0x47, 0x53, 0x5f, 0x4d, 0x41, 0x53, 0x4b, 0x10, 0xff, 0x01, 0x42, 0xa7, 0x01, 0x0a, 0x1e, 0x69,
	0x6f, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6c, 0x6f, 0x67, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x09, 0x4c,
	0x6f, 0x67, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x5a, 0x67, 0x69, 0x74, 0x6c,
	0x61, 0x62, 0x2e, 0x6a, 0x69, 0x61, 0x67, 0x6f, 0x75, 0x79, 0x75, 0x6e, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x63, 0x61, 0x72, 0x65, 0x2d, 0x74, 0x6f, 0x6f, 0x6c, 0x73,
	0x2f, 0x64, 0x61, 0x74, 0x61, 0x6b, 0x69, 0x74, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73,
	0x2f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x74, 0x65, 0x6c, 0x65,
	0x6d, 0x65, 0x74, 0x72, 0x79, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x64, 0x2f, 0x76,
	0x31, 0x2f, 0x6c, 0x6f, 0x67, 0x73, 0xaa, 0x02, 0x1b, 0x4f, 0x70, 0x65, 0x6e, 0x54, 0x65, 0x6c,
	0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x6f, 0x67,
	0x73, 0x2e, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_opentelemetry_proto_logs_v1_logs_proto_rawDescOnce sync.Once
	file_opentelemetry_proto_logs_v1_logs_proto_rawDescData = file_opentelemetry_proto_logs_v1_logs_proto_rawDesc
)

func file_opentelemetry_proto_logs_v1_logs_proto_rawDescGZIP() []byte {
	file_opentelemetry_proto_logs_v1_logs_proto_rawDescOnce.Do(func() {
		file_opentelemetry_proto_logs_v1_logs_proto_rawDescData = protoimpl.X.CompressGZIP(file_opentelemetry_proto_logs_v1_logs_proto_rawDescData)
	})
	return file_opentelemetry_proto_logs_v1_logs_proto_rawDescData
}

var file_opentelemetry_proto_logs_v1_logs_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_opentelemetry_proto_logs_v1_logs_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_opentelemetry_proto_logs_v1_logs_proto_goTypes = []interface{}{
	(SeverityNumber)(0),                 // 0: opentelemetry.proto.logs.v1.SeverityNumber
	(LogRecordFlags)(0),                 // 1: opentelemetry.proto.logs.v1.LogRecordFlags
	(*LogsData)(nil),                    // 2: opentelemetry.proto.logs.v1.LogsData
	(*ResourceLogs)(nil),                // 3: opentelemetry.proto.logs.v1.ResourceLogs
	(*ScopeLogs)(nil),                   // 4: opentelemetry.proto.logs.v1.ScopeLogs
	(*LogRecord)(nil),                   // 5: opentelemetry.proto.logs.v1.LogRecord
	(*resource.Resource)(nil),           // 6: opentelemetry.proto.resource.v1.Resource
	(*common.InstrumentationScope)(nil), // 7: opentelemetry.proto.common.v1.InstrumentationScope
	(*common.AnyValue)(nil),             // 8: opentelemetry.proto.common.v1.AnyValue
	(*common.KeyValue)(nil),             // 9: opentelemetry.proto.common.v1.KeyValue
}
var file_opentelemetry_proto_logs_v1_logs_proto_depIdxs = []int32{
	3, // 0: opentelemetry.proto.logs.v1.LogsData.resource_logs:type_name -> opentelemetry.proto.logs.v1.ResourceLogs
	6, // 1: opentelemetry.proto.logs.v1.ResourceLogs.resource:type_name -> opentelemetry.proto.resource.v1.Resource
	4, // 2: opentelemetry.proto.logs.v1.ResourceLogs.scope_logs:type_name -> opentelemetry.proto.logs.v1.ScopeLogs
	7, // 3: opentelemetry.proto.logs.v1.ScopeLogs.scope:type_name -> opentelemetry.proto.common.v1.InstrumentationScope
	5, // 4: opentelemetry.proto.logs.v1.ScopeLogs.log_records:type_name -> opentelemetry.proto.logs.v1.LogRecord
	0, // 5: opentelemetry.proto.logs.v1.LogRecord.severity_number:type_name -> opentelemetry.proto.logs.v1.SeverityNumber
	8, // 6: opentelemetry.proto.logs.v1.LogRecord.body:type_name -> opentelemetry.proto.common.v1.AnyValue
	9, // 7: opentelemetry.proto.logs.v1.LogRecord.attributes:type_name -> opentelemetry.proto.common.v1.KeyValue
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_opentelemetry_proto_logs_v1_logs_proto_init() }
func file_opentelemetry_proto_logs_v1_logs_proto_init() {
	if File_opentelemetry_proto_logs_v1_logs_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_opentelemetry_proto_logs_v1_logs_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogsData); i {
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
		file_opentelemetry_proto_logs_v1_logs_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResourceLogs); i {
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
		file_opentelemetry_proto_logs_v1_logs_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScopeLogs); i {
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
		file_opentelemetry_proto_logs_v1_logs_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogRecord); i {
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
			RawDescriptor: file_opentelemetry_proto_logs_v1_logs_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_opentelemetry_proto_logs_v1_logs_proto_goTypes,
		DependencyIndexes: file_opentelemetry_proto_logs_v1_logs_proto_depIdxs,
		EnumInfos:         file_opentelemetry_proto_logs_v1_logs_proto_enumTypes,
		MessageInfos:      file_opentelemetry_proto_logs_v1_logs_proto_msgTypes,
	}.Build()
	File_opentelemetry_proto_logs_v1_logs_proto = out.File
	file_opentelemetry_proto_logs_v1_logs_proto_rawDesc = nil
	file_opentelemetry_proto_logs_v1_logs_proto_goTypes = nil
	file_opentelemetry_proto_logs_v1_logs_proto_depIdxs = nil
}
