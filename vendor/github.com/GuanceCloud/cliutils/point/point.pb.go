// Unless explicitly stated otherwise all files in this repository are licensed
// under the MIT License.
// This product includes software developed at Guance Cloud (https://www.guance.com/).
// Copyright 2021-present Guance, Inc.

// Definition of point in protobuf
// Generate: see pb.sh

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.3
// source: point.proto

package point

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type KeyType int32

const (
	KeyType_X   KeyType = 0 // unknown
	KeyType_I   KeyType = 1
	KeyType_U   KeyType = 2
	KeyType_F   KeyType = 3
	KeyType_B   KeyType = 4
	KeyType_D   KeyType = 5
	KeyType_NIL KeyType = 6
	KeyType_S   KeyType = 7
	KeyType_A   KeyType = 8
)

// Enum value maps for KeyType.
var (
	KeyType_name = map[int32]string{
		0: "X",
		1: "I",
		2: "U",
		3: "F",
		4: "B",
		5: "D",
		6: "NIL",
		7: "S",
		8: "A",
	}
	KeyType_value = map[string]int32{
		"X":   0,
		"I":   1,
		"U":   2,
		"F":   3,
		"B":   4,
		"D":   5,
		"NIL": 6,
		"S":   7,
		"A":   8,
	}
)

func (x KeyType) Enum() *KeyType {
	p := new(KeyType)
	*p = x
	return p
}

func (x KeyType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (KeyType) Descriptor() protoreflect.EnumDescriptor {
	return file_point_proto_enumTypes[0].Descriptor()
}

func (KeyType) Type() protoreflect.EnumType {
	return &file_point_proto_enumTypes[0]
}

func (x KeyType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use KeyType.Descriptor instead.
func (KeyType) EnumDescriptor() ([]byte, []int) {
	return file_point_proto_rawDescGZIP(), []int{0}
}

type MetricType int32

const (
	MetricType_UNSPECIFIED MetricType = 0
	MetricType_COUNT       MetricType = 1
	MetricType_RATE        MetricType = 2
	MetricType_GAUGE       MetricType = 3
)

// Enum value maps for MetricType.
var (
	MetricType_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "COUNT",
		2: "RATE",
		3: "GAUGE",
	}
	MetricType_value = map[string]int32{
		"UNSPECIFIED": 0,
		"COUNT":       1,
		"RATE":        2,
		"GAUGE":       3,
	}
)

func (x MetricType) Enum() *MetricType {
	p := new(MetricType)
	*p = x
	return p
}

func (x MetricType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MetricType) Descriptor() protoreflect.EnumDescriptor {
	return file_point_proto_enumTypes[1].Descriptor()
}

func (MetricType) Type() protoreflect.EnumType {
	return &file_point_proto_enumTypes[1]
}

func (x MetricType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MetricType.Descriptor instead.
func (MetricType) EnumDescriptor() ([]byte, []int) {
	return file_point_proto_rawDescGZIP(), []int{1}
}

// Debug used to attached some debug info for the point, these debug info
// will encoded into payload, storage can take optional handle on these debug
// info.
type Debug struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info string `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *Debug) Reset() {
	*x = Debug{}
	if protoimpl.UnsafeEnabled {
		mi := &file_point_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Debug) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Debug) ProtoMessage() {}

func (x *Debug) ProtoReflect() protoreflect.Message {
	mi := &file_point_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Debug.ProtoReflect.Descriptor instead.
func (*Debug) Descriptor() ([]byte, []int) {
	return file_point_proto_rawDescGZIP(), []int{0}
}

func (x *Debug) GetInfo() string {
	if x != nil {
		return x.Info
	}
	return ""
}

// example of pb.Any
type AnyDemo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Demo string `protobuf:"bytes,1,opt,name=demo,proto3" json:"demo,omitempty"`
}

func (x *AnyDemo) Reset() {
	*x = AnyDemo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_point_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnyDemo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnyDemo) ProtoMessage() {}

func (x *AnyDemo) ProtoReflect() protoreflect.Message {
	mi := &file_point_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnyDemo.ProtoReflect.Descriptor instead.
func (*AnyDemo) Descriptor() ([]byte, []int) {
	return file_point_proto_rawDescGZIP(), []int{1}
}

func (x *AnyDemo) GetDemo() string {
	if x != nil {
		return x.Demo
	}
	return ""
}

type BasicTypes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to X:
	//	*BasicTypes_I
	//	*BasicTypes_U
	//	*BasicTypes_F
	//	*BasicTypes_B
	//	*BasicTypes_D
	//	*BasicTypes_S
	X isBasicTypes_X `protobuf_oneof:"x"`
}

func (x *BasicTypes) Reset() {
	*x = BasicTypes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_point_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BasicTypes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BasicTypes) ProtoMessage() {}

func (x *BasicTypes) ProtoReflect() protoreflect.Message {
	mi := &file_point_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BasicTypes.ProtoReflect.Descriptor instead.
func (*BasicTypes) Descriptor() ([]byte, []int) {
	return file_point_proto_rawDescGZIP(), []int{2}
}

func (m *BasicTypes) GetX() isBasicTypes_X {
	if m != nil {
		return m.X
	}
	return nil
}

func (x *BasicTypes) GetI() int64 {
	if x, ok := x.GetX().(*BasicTypes_I); ok {
		return x.I
	}
	return 0
}

func (x *BasicTypes) GetU() uint64 {
	if x, ok := x.GetX().(*BasicTypes_U); ok {
		return x.U
	}
	return 0
}

func (x *BasicTypes) GetF() float64 {
	if x, ok := x.GetX().(*BasicTypes_F); ok {
		return x.F
	}
	return 0
}

func (x *BasicTypes) GetB() bool {
	if x, ok := x.GetX().(*BasicTypes_B); ok {
		return x.B
	}
	return false
}

func (x *BasicTypes) GetD() []byte {
	if x, ok := x.GetX().(*BasicTypes_D); ok {
		return x.D
	}
	return nil
}

func (x *BasicTypes) GetS() string {
	if x, ok := x.GetX().(*BasicTypes_S); ok {
		return x.S
	}
	return ""
}

type isBasicTypes_X interface {
	isBasicTypes_X()
}

type BasicTypes_I struct {
	I int64 `protobuf:"varint,1,opt,name=i,proto3,oneof"` // signed int
}

type BasicTypes_U struct {
	U uint64 `protobuf:"varint,2,opt,name=u,proto3,oneof"` // unsigned int
}

type BasicTypes_F struct {
	F float64 `protobuf:"fixed64,3,opt,name=f,proto3,oneof"` // float64
}

type BasicTypes_B struct {
	B bool `protobuf:"varint,4,opt,name=b,proto3,oneof"` // bool
}

type BasicTypes_D struct {
	D []byte `protobuf:"bytes,5,opt,name=d,proto3,oneof"` // bytes, for binary data
}

type BasicTypes_S struct {
	S string `protobuf:"bytes,6,opt,name=s,proto3,oneof"` // string, for string data
}

func (*BasicTypes_I) isBasicTypes_X() {}

func (*BasicTypes_U) isBasicTypes_X() {}

func (*BasicTypes_F) isBasicTypes_X() {}

func (*BasicTypes_B) isBasicTypes_X() {}

func (*BasicTypes_D) isBasicTypes_X() {}

func (*BasicTypes_S) isBasicTypes_X() {}

type Array struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Arr []*BasicTypes `protobuf:"bytes,1,rep,name=arr,proto3" json:"arr,omitempty"`
}

func (x *Array) Reset() {
	*x = Array{}
	if protoimpl.UnsafeEnabled {
		mi := &file_point_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Array) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Array) ProtoMessage() {}

func (x *Array) ProtoReflect() protoreflect.Message {
	mi := &file_point_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Array.ProtoReflect.Descriptor instead.
func (*Array) Descriptor() ([]byte, []int) {
	return file_point_proto_rawDescGZIP(), []int{3}
}

func (x *Array) GetArr() []*BasicTypes {
	if x != nil {
		return x.Arr
	}
	return nil
}

type Map struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Map map[string]*BasicTypes `protobuf:"bytes,1,rep,name=map,proto3" json:"map,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Map) Reset() {
	*x = Map{}
	if protoimpl.UnsafeEnabled {
		mi := &file_point_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Map) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Map) ProtoMessage() {}

func (x *Map) ProtoReflect() protoreflect.Message {
	mi := &file_point_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Map.ProtoReflect.Descriptor instead.
func (*Map) Descriptor() ([]byte, []int) {
	return file_point_proto_rawDescGZIP(), []int{4}
}

func (x *Map) GetMap() map[string]*BasicTypes {
	if x != nil {
		return x.Map
	}
	return nil
}

type Field struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"` // field name
	// See https://developers.google.com/protocol-buffers/docs/proto3#json
	//
	// Types that are assignable to Val:
	//	*Field_I
	//	*Field_U
	//	*Field_F
	//	*Field_B
	//	*Field_D
	//	*Field_S
	//	*Field_A
	Val   isField_Val `protobuf_oneof:"val"`
	IsTag bool        `protobuf:"varint,8,opt,name=is_tag,proto3" json:"is_tag,omitempty"` // set field as a tag or not
	Type  MetricType  `protobuf:"varint,9,opt,name=type,proto3,enum=point.MetricType" json:"type,omitempty"`
	// field unit name
	Unit string `protobuf:"bytes,10,opt,name=unit,proto3" json:"unit,omitempty"` // metric unit, such as bytes(B), duration(ms/us) and so on.
}

func (x *Field) Reset() {
	*x = Field{}
	if protoimpl.UnsafeEnabled {
		mi := &file_point_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Field) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Field) ProtoMessage() {}

func (x *Field) ProtoReflect() protoreflect.Message {
	mi := &file_point_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Field.ProtoReflect.Descriptor instead.
func (*Field) Descriptor() ([]byte, []int) {
	return file_point_proto_rawDescGZIP(), []int{5}
}

func (x *Field) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (m *Field) GetVal() isField_Val {
	if m != nil {
		return m.Val
	}
	return nil
}

func (x *Field) GetI() int64 {
	if x, ok := x.GetVal().(*Field_I); ok {
		return x.I
	}
	return 0
}

func (x *Field) GetU() uint64 {
	if x, ok := x.GetVal().(*Field_U); ok {
		return x.U
	}
	return 0
}

func (x *Field) GetF() float64 {
	if x, ok := x.GetVal().(*Field_F); ok {
		return x.F
	}
	return 0
}

func (x *Field) GetB() bool {
	if x, ok := x.GetVal().(*Field_B); ok {
		return x.B
	}
	return false
}

func (x *Field) GetD() []byte {
	if x, ok := x.GetVal().(*Field_D); ok {
		return x.D
	}
	return nil
}

func (x *Field) GetS() string {
	if x, ok := x.GetVal().(*Field_S); ok {
		return x.S
	}
	return ""
}

func (x *Field) GetA() *anypb.Any {
	if x, ok := x.GetVal().(*Field_A); ok {
		return x.A
	}
	return nil
}

func (x *Field) GetIsTag() bool {
	if x != nil {
		return x.IsTag
	}
	return false
}

func (x *Field) GetType() MetricType {
	if x != nil {
		return x.Type
	}
	return MetricType_UNSPECIFIED
}

func (x *Field) GetUnit() string {
	if x != nil {
		return x.Unit
	}
	return ""
}

type isField_Val interface {
	isField_Val()
}

type Field_I struct {
	I int64 `protobuf:"varint,2,opt,name=i,proto3,oneof"` // signed int
}

type Field_U struct {
	U uint64 `protobuf:"varint,3,opt,name=u,proto3,oneof"` // unsigned int
}

type Field_F struct {
	F float64 `protobuf:"fixed64,4,opt,name=f,proto3,oneof"` // float64
}

type Field_B struct {
	B bool `protobuf:"varint,5,opt,name=b,proto3,oneof"` // bool
}

type Field_D struct {
	D []byte `protobuf:"bytes,6,opt,name=d,proto3,oneof"` // bytes, for binary data
}

type Field_S struct {
	S string `protobuf:"bytes,11,opt,name=s,proto3,oneof"` // string, for string data
}

type Field_A struct {
	// XXX: not used
	A *anypb.Any `protobuf:"bytes,7,opt,name=a,proto3,oneof"` // any data
}

func (*Field_I) isField_Val() {}

func (*Field_U) isField_Val() {}

func (*Field_F) isField_Val() {}

func (*Field_B) isField_Val() {}

func (*Field_D) isField_Val() {}

func (*Field_S) isField_Val() {}

func (*Field_A) isField_Val() {}

// Warn used to attach some warning message during building the point.
type Warn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,json=message,proto3" json:"msg,omitempty"`
}

func (x *Warn) Reset() {
	*x = Warn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_point_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Warn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Warn) ProtoMessage() {}

func (x *Warn) ProtoReflect() protoreflect.Message {
	mi := &file_point_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Warn.ProtoReflect.Descriptor instead.
func (*Warn) Descriptor() ([]byte, []int) {
	return file_point_proto_rawDescGZIP(), []int{6}
}

func (x *Warn) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Warn) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type PBPoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Fields []*Field `protobuf:"bytes,2,rep,name=fields,proto3" json:"fields,omitempty"`
	Time   int64    `protobuf:"varint,3,opt,name=time,proto3" json:"time,omitempty"`
	// Auxiliary fields for the point, they should not
	// write to the final storage on production.
	Warns  []*Warn  `protobuf:"bytes,4,rep,name=warns,proto3" json:"warns,omitempty"`
	Debugs []*Debug `protobuf:"bytes,5,rep,name=debugs,proto3" json:"debugs,omitempty"`
}

func (x *PBPoint) Reset() {
	*x = PBPoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_point_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PBPoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PBPoint) ProtoMessage() {}

func (x *PBPoint) ProtoReflect() protoreflect.Message {
	mi := &file_point_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PBPoint.ProtoReflect.Descriptor instead.
func (*PBPoint) Descriptor() ([]byte, []int) {
	return file_point_proto_rawDescGZIP(), []int{7}
}

func (x *PBPoint) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PBPoint) GetFields() []*Field {
	if x != nil {
		return x.Fields
	}
	return nil
}

func (x *PBPoint) GetTime() int64 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *PBPoint) GetWarns() []*Warn {
	if x != nil {
		return x.Warns
	}
	return nil
}

func (x *PBPoint) GetDebugs() []*Debug {
	if x != nil {
		return x.Debugs
	}
	return nil
}

// batch of pbpoint.
type PBPoints struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Arr []*PBPoint `protobuf:"bytes,1,rep,name=arr,proto3" json:"arr,omitempty"`
}

func (x *PBPoints) Reset() {
	*x = PBPoints{}
	if protoimpl.UnsafeEnabled {
		mi := &file_point_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PBPoints) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PBPoints) ProtoMessage() {}

func (x *PBPoints) ProtoReflect() protoreflect.Message {
	mi := &file_point_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PBPoints.ProtoReflect.Descriptor instead.
func (*PBPoints) Descriptor() ([]byte, []int) {
	return file_point_proto_rawDescGZIP(), []int{8}
}

func (x *PBPoints) GetArr() []*PBPoint {
	if x != nil {
		return x.Arr
	}
	return nil
}

var File_point_proto protoreflect.FileDescriptor

var file_point_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x1b, 0x0a, 0x05, 0x44, 0x65, 0x62, 0x75, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x1d, 0x0a, 0x07,
	0x41, 0x6e, 0x79, 0x44, 0x65, 0x6d, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x6d, 0x6f, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x65, 0x6d, 0x6f, 0x22, 0x71, 0x0a, 0x0a, 0x42,
	0x61, 0x73, 0x69, 0x63, 0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x0e, 0x0a, 0x01, 0x69, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x01, 0x69, 0x12, 0x0e, 0x0a, 0x01, 0x75, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x04, 0x48, 0x00, 0x52, 0x01, 0x75, 0x12, 0x0e, 0x0a, 0x01, 0x66, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x01, 0x48, 0x00, 0x52, 0x01, 0x66, 0x12, 0x0e, 0x0a, 0x01, 0x62, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x01, 0x62, 0x12, 0x0e, 0x0a, 0x01, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x01, 0x64, 0x12, 0x0e, 0x0a, 0x01, 0x73, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x01, 0x73, 0x42, 0x03, 0x0a, 0x01, 0x78, 0x22, 0x2c,
	0x0a, 0x05, 0x41, 0x72, 0x72, 0x61, 0x79, 0x12, 0x23, 0x0a, 0x03, 0x61, 0x72, 0x72, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x42, 0x61, 0x73,
	0x69, 0x63, 0x54, 0x79, 0x70, 0x65, 0x73, 0x52, 0x03, 0x61, 0x72, 0x72, 0x22, 0x77, 0x0a, 0x03,
	0x4d, 0x61, 0x70, 0x12, 0x25, 0x0a, 0x03, 0x6d, 0x61, 0x70, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x4d, 0x61, 0x70, 0x2e, 0x4d, 0x61, 0x70,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x03, 0x6d, 0x61, 0x70, 0x1a, 0x49, 0x0a, 0x08, 0x4d, 0x61,
	0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x27, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e,
	0x42, 0x61, 0x73, 0x69, 0x63, 0x54, 0x79, 0x70, 0x65, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xf9, 0x01, 0x0a, 0x05, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x0e, 0x0a, 0x01, 0x69, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x01,
	0x69, 0x12, 0x0e, 0x0a, 0x01, 0x75, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x48, 0x00, 0x52, 0x01,
	0x75, 0x12, 0x0e, 0x0a, 0x01, 0x66, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x48, 0x00, 0x52, 0x01,
	0x66, 0x12, 0x0e, 0x0a, 0x01, 0x62, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x01,
	0x62, 0x12, 0x0e, 0x0a, 0x01, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x01,
	0x64, 0x12, 0x0e, 0x0a, 0x01, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x01,
	0x73, 0x12, 0x24, 0x0a, 0x01, 0x61, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41,
	0x6e, 0x79, 0x48, 0x00, 0x52, 0x01, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x73, 0x5f, 0x74, 0x61,
	0x67, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73, 0x5f, 0x74, 0x61, 0x67, 0x12,
	0x25, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x42, 0x05, 0x0a, 0x03, 0x76, 0x61,
	0x6c, 0x22, 0x30, 0x0a, 0x04, 0x57, 0x61, 0x72, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a,
	0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0xa0, 0x01, 0x0a, 0x07, 0x50, 0x42, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x21, 0x0a,
	0x05, 0x77, 0x61, 0x72, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x57, 0x61, 0x72, 0x6e, 0x52, 0x05, 0x77, 0x61, 0x72, 0x6e, 0x73,
	0x12, 0x24, 0x0a, 0x06, 0x64, 0x65, 0x62, 0x75, 0x67, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x44, 0x65, 0x62, 0x75, 0x67, 0x52, 0x06,
	0x64, 0x65, 0x62, 0x75, 0x67, 0x73, 0x22, 0x2c, 0x0a, 0x08, 0x50, 0x42, 0x50, 0x6f, 0x69, 0x6e,
	0x74, 0x73, 0x12, 0x20, 0x0a, 0x03, 0x61, 0x72, 0x72, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0e, 0x2e, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x50, 0x42, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52,
	0x03, 0x61, 0x72, 0x72, 0x2a, 0x4a, 0x0a, 0x07, 0x4b, 0x65, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x05, 0x0a, 0x01, 0x58, 0x10, 0x00, 0x12, 0x05, 0x0a, 0x01, 0x49, 0x10, 0x01, 0x12, 0x05, 0x0a,
	0x01, 0x55, 0x10, 0x02, 0x12, 0x05, 0x0a, 0x01, 0x46, 0x10, 0x03, 0x12, 0x05, 0x0a, 0x01, 0x42,
	0x10, 0x04, 0x12, 0x05, 0x0a, 0x01, 0x44, 0x10, 0x05, 0x12, 0x07, 0x0a, 0x03, 0x4e, 0x49, 0x4c,
	0x10, 0x06, 0x12, 0x05, 0x0a, 0x01, 0x53, 0x10, 0x07, 0x12, 0x05, 0x0a, 0x01, 0x41, 0x10, 0x08,
	0x2a, 0x3d, 0x0a, 0x0a, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0f,
	0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x09, 0x0a, 0x05, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x52, 0x41,
	0x54, 0x45, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x47, 0x41, 0x55, 0x47, 0x45, 0x10, 0x03, 0x42,
	0x12, 0x42, 0x05, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x50, 0x01, 0x5a, 0x07, 0x2f, 0x3b, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_point_proto_rawDescOnce sync.Once
	file_point_proto_rawDescData = file_point_proto_rawDesc
)

func file_point_proto_rawDescGZIP() []byte {
	file_point_proto_rawDescOnce.Do(func() {
		file_point_proto_rawDescData = protoimpl.X.CompressGZIP(file_point_proto_rawDescData)
	})
	return file_point_proto_rawDescData
}

var file_point_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_point_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_point_proto_goTypes = []interface{}{
	(KeyType)(0),       // 0: point.KeyType
	(MetricType)(0),    // 1: point.MetricType
	(*Debug)(nil),      // 2: point.Debug
	(*AnyDemo)(nil),    // 3: point.AnyDemo
	(*BasicTypes)(nil), // 4: point.BasicTypes
	(*Array)(nil),      // 5: point.Array
	(*Map)(nil),        // 6: point.Map
	(*Field)(nil),      // 7: point.Field
	(*Warn)(nil),       // 8: point.Warn
	(*PBPoint)(nil),    // 9: point.PBPoint
	(*PBPoints)(nil),   // 10: point.PBPoints
	nil,                // 11: point.Map.MapEntry
	(*anypb.Any)(nil),  // 12: google.protobuf.Any
}
var file_point_proto_depIdxs = []int32{
	4,  // 0: point.Array.arr:type_name -> point.BasicTypes
	11, // 1: point.Map.map:type_name -> point.Map.MapEntry
	12, // 2: point.Field.a:type_name -> google.protobuf.Any
	1,  // 3: point.Field.type:type_name -> point.MetricType
	7,  // 4: point.PBPoint.fields:type_name -> point.Field
	8,  // 5: point.PBPoint.warns:type_name -> point.Warn
	2,  // 6: point.PBPoint.debugs:type_name -> point.Debug
	9,  // 7: point.PBPoints.arr:type_name -> point.PBPoint
	4,  // 8: point.Map.MapEntry.value:type_name -> point.BasicTypes
	9,  // [9:9] is the sub-list for method output_type
	9,  // [9:9] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_point_proto_init() }
func file_point_proto_init() {
	if File_point_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_point_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Debug); i {
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
		file_point_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnyDemo); i {
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
		file_point_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BasicTypes); i {
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
		file_point_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Array); i {
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
		file_point_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Map); i {
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
		file_point_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Field); i {
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
		file_point_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Warn); i {
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
		file_point_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PBPoint); i {
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
		file_point_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PBPoints); i {
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
	file_point_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*BasicTypes_I)(nil),
		(*BasicTypes_U)(nil),
		(*BasicTypes_F)(nil),
		(*BasicTypes_B)(nil),
		(*BasicTypes_D)(nil),
		(*BasicTypes_S)(nil),
	}
	file_point_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*Field_I)(nil),
		(*Field_U)(nil),
		(*Field_F)(nil),
		(*Field_B)(nil),
		(*Field_D)(nil),
		(*Field_S)(nil),
		(*Field_A)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_point_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_point_proto_goTypes,
		DependencyIndexes: file_point_proto_depIdxs,
		EnumInfos:         file_point_proto_enumTypes,
		MessageInfos:      file_point_proto_msgTypes,
	}.Build()
	File_point_proto = out.File
	file_point_proto_rawDesc = nil
	file_point_proto_goTypes = nil
	file_point_proto_depIdxs = nil
}
