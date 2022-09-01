//
// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// The ASF licenses this file to You under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance with
// the License.  You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: language-agent/JVMMetric.proto

package v3

import (
	v3 "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/skywalking/compiled/common/v3"
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

type PoolType int32

const (
	PoolType_CODE_CACHE_USAGE PoolType = 0
	PoolType_NEWGEN_USAGE     PoolType = 1
	PoolType_OLDGEN_USAGE     PoolType = 2
	PoolType_SURVIVOR_USAGE   PoolType = 3
	PoolType_PERMGEN_USAGE    PoolType = 4
	PoolType_METASPACE_USAGE  PoolType = 5
)

// Enum value maps for PoolType.
var (
	PoolType_name = map[int32]string{
		0: "CODE_CACHE_USAGE",
		1: "NEWGEN_USAGE",
		2: "OLDGEN_USAGE",
		3: "SURVIVOR_USAGE",
		4: "PERMGEN_USAGE",
		5: "METASPACE_USAGE",
	}
	PoolType_value = map[string]int32{
		"CODE_CACHE_USAGE": 0,
		"NEWGEN_USAGE":     1,
		"OLDGEN_USAGE":     2,
		"SURVIVOR_USAGE":   3,
		"PERMGEN_USAGE":    4,
		"METASPACE_USAGE":  5,
	}
)

func (x PoolType) Enum() *PoolType {
	p := new(PoolType)
	*p = x
	return p
}

func (x PoolType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PoolType) Descriptor() protoreflect.EnumDescriptor {
	return file_language_agent_JVMMetric_proto_enumTypes[0].Descriptor()
}

func (PoolType) Type() protoreflect.EnumType {
	return &file_language_agent_JVMMetric_proto_enumTypes[0]
}

func (x PoolType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PoolType.Descriptor instead.
func (PoolType) EnumDescriptor() ([]byte, []int) {
	return file_language_agent_JVMMetric_proto_rawDescGZIP(), []int{0}
}

type GCPhase int32

const (
	GCPhase_NEW    GCPhase = 0
	GCPhase_OLD    GCPhase = 1
	GCPhase_NORMAL GCPhase = 2 // The type of GC doesn't have new and old phases, like Z Garbage
)

// Enum value maps for GCPhase.
var (
	GCPhase_name = map[int32]string{
		0: "NEW",
		1: "OLD",
		2: "NORMAL",
	}
	GCPhase_value = map[string]int32{
		"NEW":    0,
		"OLD":    1,
		"NORMAL": 2,
	}
)

func (x GCPhase) Enum() *GCPhase {
	p := new(GCPhase)
	*p = x
	return p
}

func (x GCPhase) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GCPhase) Descriptor() protoreflect.EnumDescriptor {
	return file_language_agent_JVMMetric_proto_enumTypes[1].Descriptor()
}

func (GCPhase) Type() protoreflect.EnumType {
	return &file_language_agent_JVMMetric_proto_enumTypes[1]
}

func (x GCPhase) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GCPhase.Descriptor instead.
func (GCPhase) EnumDescriptor() ([]byte, []int) {
	return file_language_agent_JVMMetric_proto_rawDescGZIP(), []int{1}
}

type JVMMetricCollection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metrics         []*JVMMetric `protobuf:"bytes,1,rep,name=metrics,proto3" json:"metrics,omitempty"`
	Service         string       `protobuf:"bytes,2,opt,name=service,proto3" json:"service,omitempty"`
	ServiceInstance string       `protobuf:"bytes,3,opt,name=serviceInstance,proto3" json:"serviceInstance,omitempty"`
}

func (x *JVMMetricCollection) Reset() {
	*x = JVMMetricCollection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_language_agent_JVMMetric_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JVMMetricCollection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JVMMetricCollection) ProtoMessage() {}

func (x *JVMMetricCollection) ProtoReflect() protoreflect.Message {
	mi := &file_language_agent_JVMMetric_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JVMMetricCollection.ProtoReflect.Descriptor instead.
func (*JVMMetricCollection) Descriptor() ([]byte, []int) {
	return file_language_agent_JVMMetric_proto_rawDescGZIP(), []int{0}
}

func (x *JVMMetricCollection) GetMetrics() []*JVMMetric {
	if x != nil {
		return x.Metrics
	}
	return nil
}

func (x *JVMMetricCollection) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *JVMMetricCollection) GetServiceInstance() string {
	if x != nil {
		return x.ServiceInstance
	}
	return ""
}

type JVMMetric struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time       int64         `protobuf:"varint,1,opt,name=time,proto3" json:"time,omitempty"`
	Cpu        *v3.CPU       `protobuf:"bytes,2,opt,name=cpu,proto3" json:"cpu,omitempty"`
	Memory     []*Memory     `protobuf:"bytes,3,rep,name=memory,proto3" json:"memory,omitempty"`
	MemoryPool []*MemoryPool `protobuf:"bytes,4,rep,name=memoryPool,proto3" json:"memoryPool,omitempty"`
	Gc         []*GC         `protobuf:"bytes,5,rep,name=gc,proto3" json:"gc,omitempty"`
	Thread     *Thread       `protobuf:"bytes,6,opt,name=thread,proto3" json:"thread,omitempty"`
	Clazz      *Class        `protobuf:"bytes,7,opt,name=clazz,proto3" json:"clazz,omitempty"`
}

func (x *JVMMetric) Reset() {
	*x = JVMMetric{}
	if protoimpl.UnsafeEnabled {
		mi := &file_language_agent_JVMMetric_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JVMMetric) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JVMMetric) ProtoMessage() {}

func (x *JVMMetric) ProtoReflect() protoreflect.Message {
	mi := &file_language_agent_JVMMetric_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JVMMetric.ProtoReflect.Descriptor instead.
func (*JVMMetric) Descriptor() ([]byte, []int) {
	return file_language_agent_JVMMetric_proto_rawDescGZIP(), []int{1}
}

func (x *JVMMetric) GetTime() int64 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *JVMMetric) GetCpu() *v3.CPU {
	if x != nil {
		return x.Cpu
	}
	return nil
}

func (x *JVMMetric) GetMemory() []*Memory {
	if x != nil {
		return x.Memory
	}
	return nil
}

func (x *JVMMetric) GetMemoryPool() []*MemoryPool {
	if x != nil {
		return x.MemoryPool
	}
	return nil
}

func (x *JVMMetric) GetGc() []*GC {
	if x != nil {
		return x.Gc
	}
	return nil
}

func (x *JVMMetric) GetThread() *Thread {
	if x != nil {
		return x.Thread
	}
	return nil
}

func (x *JVMMetric) GetClazz() *Class {
	if x != nil {
		return x.Clazz
	}
	return nil
}

type Memory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsHeap    bool  `protobuf:"varint,1,opt,name=isHeap,proto3" json:"isHeap,omitempty"`
	Init      int64 `protobuf:"varint,2,opt,name=init,proto3" json:"init,omitempty"`
	Max       int64 `protobuf:"varint,3,opt,name=max,proto3" json:"max,omitempty"`
	Used      int64 `protobuf:"varint,4,opt,name=used,proto3" json:"used,omitempty"`
	Committed int64 `protobuf:"varint,5,opt,name=committed,proto3" json:"committed,omitempty"`
}

func (x *Memory) Reset() {
	*x = Memory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_language_agent_JVMMetric_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Memory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Memory) ProtoMessage() {}

func (x *Memory) ProtoReflect() protoreflect.Message {
	mi := &file_language_agent_JVMMetric_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Memory.ProtoReflect.Descriptor instead.
func (*Memory) Descriptor() ([]byte, []int) {
	return file_language_agent_JVMMetric_proto_rawDescGZIP(), []int{2}
}

func (x *Memory) GetIsHeap() bool {
	if x != nil {
		return x.IsHeap
	}
	return false
}

func (x *Memory) GetInit() int64 {
	if x != nil {
		return x.Init
	}
	return 0
}

func (x *Memory) GetMax() int64 {
	if x != nil {
		return x.Max
	}
	return 0
}

func (x *Memory) GetUsed() int64 {
	if x != nil {
		return x.Used
	}
	return 0
}

func (x *Memory) GetCommitted() int64 {
	if x != nil {
		return x.Committed
	}
	return 0
}

type MemoryPool struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type      PoolType `protobuf:"varint,1,opt,name=type,proto3,enum=skywalking.v3.PoolType" json:"type,omitempty"`
	Init      int64    `protobuf:"varint,2,opt,name=init,proto3" json:"init,omitempty"`
	Max       int64    `protobuf:"varint,3,opt,name=max,proto3" json:"max,omitempty"`
	Used      int64    `protobuf:"varint,4,opt,name=used,proto3" json:"used,omitempty"`
	Committed int64    `protobuf:"varint,5,opt,name=committed,proto3" json:"committed,omitempty"`
}

func (x *MemoryPool) Reset() {
	*x = MemoryPool{}
	if protoimpl.UnsafeEnabled {
		mi := &file_language_agent_JVMMetric_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemoryPool) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemoryPool) ProtoMessage() {}

func (x *MemoryPool) ProtoReflect() protoreflect.Message {
	mi := &file_language_agent_JVMMetric_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemoryPool.ProtoReflect.Descriptor instead.
func (*MemoryPool) Descriptor() ([]byte, []int) {
	return file_language_agent_JVMMetric_proto_rawDescGZIP(), []int{3}
}

func (x *MemoryPool) GetType() PoolType {
	if x != nil {
		return x.Type
	}
	return PoolType_CODE_CACHE_USAGE
}

func (x *MemoryPool) GetInit() int64 {
	if x != nil {
		return x.Init
	}
	return 0
}

func (x *MemoryPool) GetMax() int64 {
	if x != nil {
		return x.Max
	}
	return 0
}

func (x *MemoryPool) GetUsed() int64 {
	if x != nil {
		return x.Used
	}
	return 0
}

func (x *MemoryPool) GetCommitted() int64 {
	if x != nil {
		return x.Committed
	}
	return 0
}

type GC struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Phase GCPhase `protobuf:"varint,1,opt,name=phase,proto3,enum=skywalking.v3.GCPhase" json:"phase,omitempty"`
	Count int64   `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	Time  int64   `protobuf:"varint,3,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *GC) Reset() {
	*x = GC{}
	if protoimpl.UnsafeEnabled {
		mi := &file_language_agent_JVMMetric_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GC) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GC) ProtoMessage() {}

func (x *GC) ProtoReflect() protoreflect.Message {
	mi := &file_language_agent_JVMMetric_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GC.ProtoReflect.Descriptor instead.
func (*GC) Descriptor() ([]byte, []int) {
	return file_language_agent_JVMMetric_proto_rawDescGZIP(), []int{4}
}

func (x *GC) GetPhase() GCPhase {
	if x != nil {
		return x.Phase
	}
	return GCPhase_NEW
}

func (x *GC) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *GC) GetTime() int64 {
	if x != nil {
		return x.Time
	}
	return 0
}

// See:
// https://docs.oracle.com/javase/8/docs/api/java/lang/management/ThreadMXBean.html
type Thread struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LiveCount                    int64 `protobuf:"varint,1,opt,name=liveCount,proto3" json:"liveCount,omitempty"`
	DaemonCount                  int64 `protobuf:"varint,2,opt,name=daemonCount,proto3" json:"daemonCount,omitempty"`
	PeakCount                    int64 `protobuf:"varint,3,opt,name=peakCount,proto3" json:"peakCount,omitempty"`
	RunnableStateThreadCount     int64 `protobuf:"varint,4,opt,name=runnableStateThreadCount,proto3" json:"runnableStateThreadCount,omitempty"`
	BlockedStateThreadCount      int64 `protobuf:"varint,5,opt,name=blockedStateThreadCount,proto3" json:"blockedStateThreadCount,omitempty"`
	WaitingStateThreadCount      int64 `protobuf:"varint,6,opt,name=waitingStateThreadCount,proto3" json:"waitingStateThreadCount,omitempty"`
	TimedWaitingStateThreadCount int64 `protobuf:"varint,7,opt,name=timedWaitingStateThreadCount,proto3" json:"timedWaitingStateThreadCount,omitempty"`
}

func (x *Thread) Reset() {
	*x = Thread{}
	if protoimpl.UnsafeEnabled {
		mi := &file_language_agent_JVMMetric_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Thread) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Thread) ProtoMessage() {}

func (x *Thread) ProtoReflect() protoreflect.Message {
	mi := &file_language_agent_JVMMetric_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Thread.ProtoReflect.Descriptor instead.
func (*Thread) Descriptor() ([]byte, []int) {
	return file_language_agent_JVMMetric_proto_rawDescGZIP(), []int{5}
}

func (x *Thread) GetLiveCount() int64 {
	if x != nil {
		return x.LiveCount
	}
	return 0
}

func (x *Thread) GetDaemonCount() int64 {
	if x != nil {
		return x.DaemonCount
	}
	return 0
}

func (x *Thread) GetPeakCount() int64 {
	if x != nil {
		return x.PeakCount
	}
	return 0
}

func (x *Thread) GetRunnableStateThreadCount() int64 {
	if x != nil {
		return x.RunnableStateThreadCount
	}
	return 0
}

func (x *Thread) GetBlockedStateThreadCount() int64 {
	if x != nil {
		return x.BlockedStateThreadCount
	}
	return 0
}

func (x *Thread) GetWaitingStateThreadCount() int64 {
	if x != nil {
		return x.WaitingStateThreadCount
	}
	return 0
}

func (x *Thread) GetTimedWaitingStateThreadCount() int64 {
	if x != nil {
		return x.TimedWaitingStateThreadCount
	}
	return 0
}

// See:
// https://docs.oracle.com/javase/8/docs/api/java/lang/management/ClassLoadingMXBean.html
type Class struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LoadedClassCount        int64 `protobuf:"varint,1,opt,name=loadedClassCount,proto3" json:"loadedClassCount,omitempty"`
	TotalUnloadedClassCount int64 `protobuf:"varint,2,opt,name=totalUnloadedClassCount,proto3" json:"totalUnloadedClassCount,omitempty"`
	TotalLoadedClassCount   int64 `protobuf:"varint,3,opt,name=totalLoadedClassCount,proto3" json:"totalLoadedClassCount,omitempty"`
}

func (x *Class) Reset() {
	*x = Class{}
	if protoimpl.UnsafeEnabled {
		mi := &file_language_agent_JVMMetric_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Class) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Class) ProtoMessage() {}

func (x *Class) ProtoReflect() protoreflect.Message {
	mi := &file_language_agent_JVMMetric_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Class.ProtoReflect.Descriptor instead.
func (*Class) Descriptor() ([]byte, []int) {
	return file_language_agent_JVMMetric_proto_rawDescGZIP(), []int{6}
}

func (x *Class) GetLoadedClassCount() int64 {
	if x != nil {
		return x.LoadedClassCount
	}
	return 0
}

func (x *Class) GetTotalUnloadedClassCount() int64 {
	if x != nil {
		return x.TotalUnloadedClassCount
	}
	return 0
}

func (x *Class) GetTotalLoadedClassCount() int64 {
	if x != nil {
		return x.TotalLoadedClassCount
	}
	return 0
}

var File_language_agent_JVMMetric_proto protoreflect.FileDescriptor

var file_language_agent_JVMMetric_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2d, 0x61, 0x67, 0x65, 0x6e, 0x74,
	0x2f, 0x4a, 0x56, 0x4d, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0d, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x1a,
	0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8d, 0x01, 0x0a, 0x13, 0x4a, 0x56, 0x4d, 0x4d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x32, 0x0a, 0x07,
	0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e,
	0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x4a, 0x56,
	0x4d, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x52, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x22, 0xad, 0x02, 0x0a, 0x09, 0x4a, 0x56, 0x4d, 0x4d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x03, 0x63, 0x70, 0x75, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67,
	0x2e, 0x76, 0x33, 0x2e, 0x43, 0x50, 0x55, 0x52, 0x03, 0x63, 0x70, 0x75, 0x12, 0x2d, 0x0a, 0x06,
	0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73,
	0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x4d, 0x65, 0x6d,
	0x6f, 0x72, 0x79, 0x52, 0x06, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x12, 0x39, 0x0a, 0x0a, 0x6d,
	0x65, 0x6d, 0x6f, 0x72, 0x79, 0x50, 0x6f, 0x6f, 0x6c, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e,
	0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x50, 0x6f, 0x6f, 0x6c, 0x52, 0x0a, 0x6d, 0x65, 0x6d, 0x6f,
	0x72, 0x79, 0x50, 0x6f, 0x6f, 0x6c, 0x12, 0x21, 0x0a, 0x02, 0x67, 0x63, 0x18, 0x05, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x33, 0x2e, 0x47, 0x43, 0x52, 0x02, 0x67, 0x63, 0x12, 0x2d, 0x0a, 0x06, 0x74, 0x68, 0x72,
	0x65, 0x61, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x6b, 0x79, 0x77,
	0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64,
	0x52, 0x06, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x12, 0x2a, 0x0a, 0x05, 0x63, 0x6c, 0x61, 0x7a,
	0x7a, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c,
	0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x52, 0x05, 0x63,
	0x6c, 0x61, 0x7a, 0x7a, 0x22, 0x78, 0x0a, 0x06, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x12, 0x16,
	0x0a, 0x06, 0x69, 0x73, 0x48, 0x65, 0x61, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06,
	0x69, 0x73, 0x48, 0x65, 0x61, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x6e, 0x69, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x69, 0x6e, 0x69, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x61,
	0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6d, 0x61, 0x78, 0x12, 0x12, 0x0a, 0x04,
	0x75, 0x73, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x75, 0x73, 0x65, 0x64,
	0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x64, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x64, 0x22, 0x91,
	0x01, 0x0a, 0x0a, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x50, 0x6f, 0x6f, 0x6c, 0x12, 0x2b, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x73, 0x6b,
	0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x50, 0x6f, 0x6f, 0x6c,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x6e,
	0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x69, 0x6e, 0x69, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x6d, 0x61, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6d, 0x61, 0x78,
	0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04,
	0x75, 0x73, 0x65, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x74, 0x65,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x74,
	0x65, 0x64, 0x22, 0x5c, 0x0a, 0x02, 0x47, 0x43, 0x12, 0x2c, 0x0a, 0x05, 0x70, 0x68, 0x61, 0x73,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c,
	0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x47, 0x43, 0x50, 0x68, 0x61, 0x73, 0x65, 0x52,
	0x05, 0x70, 0x68, 0x61, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65,
	0x22, 0xda, 0x02, 0x0a, 0x06, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x6c,
	0x69, 0x76, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x6c, 0x69, 0x76, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x61, 0x65,
	0x6d, 0x6f, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b,
	0x64, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70,
	0x65, 0x61, 0x6b, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x70, 0x65, 0x61, 0x6b, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x3a, 0x0a, 0x18, 0x72, 0x75, 0x6e,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x18, 0x72, 0x75, 0x6e,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x38, 0x0a, 0x17, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x17, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x38, 0x0a, 0x17, 0x77, 0x61, 0x69, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x65, 0x54,
	0x68, 0x72, 0x65, 0x61, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x17, 0x77, 0x61, 0x69, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x65, 0x54, 0x68,
	0x72, 0x65, 0x61, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x42, 0x0a, 0x1c, 0x74, 0x69, 0x6d,
	0x65, 0x64, 0x57, 0x61, 0x69, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x65, 0x54, 0x68,
	0x72, 0x65, 0x61, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x1c, 0x74, 0x69, 0x6d, 0x65, 0x64, 0x57, 0x61, 0x69, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0xa3, 0x01,
	0x0a, 0x05, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x12, 0x2a, 0x0a, 0x10, 0x6c, 0x6f, 0x61, 0x64, 0x65,
	0x64, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x10, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x64, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x38, 0x0a, 0x17, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x55, 0x6e, 0x6c, 0x6f,
	0x61, 0x64, 0x65, 0x64, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x17, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x55, 0x6e, 0x6c, 0x6f, 0x61,
	0x64, 0x65, 0x64, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x34, 0x0a,
	0x15, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x4c, 0x6f, 0x61, 0x64, 0x65, 0x64, 0x43, 0x6c, 0x61, 0x73,
	0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x15, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x4c, 0x6f, 0x61, 0x64, 0x65, 0x64, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x2a, 0x80, 0x01, 0x0a, 0x08, 0x50, 0x6f, 0x6f, 0x6c, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x14, 0x0a, 0x10, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x43, 0x41, 0x43, 0x48, 0x45, 0x5f, 0x55,
	0x53, 0x41, 0x47, 0x45, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x4e, 0x45, 0x57, 0x47, 0x45, 0x4e,
	0x5f, 0x55, 0x53, 0x41, 0x47, 0x45, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x4f, 0x4c, 0x44, 0x47,
	0x45, 0x4e, 0x5f, 0x55, 0x53, 0x41, 0x47, 0x45, 0x10, 0x02, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x55,
	0x52, 0x56, 0x49, 0x56, 0x4f, 0x52, 0x5f, 0x55, 0x53, 0x41, 0x47, 0x45, 0x10, 0x03, 0x12, 0x11,
	0x0a, 0x0d, 0x50, 0x45, 0x52, 0x4d, 0x47, 0x45, 0x4e, 0x5f, 0x55, 0x53, 0x41, 0x47, 0x45, 0x10,
	0x04, 0x12, 0x13, 0x0a, 0x0f, 0x4d, 0x45, 0x54, 0x41, 0x53, 0x50, 0x41, 0x43, 0x45, 0x5f, 0x55,
	0x53, 0x41, 0x47, 0x45, 0x10, 0x05, 0x2a, 0x27, 0x0a, 0x07, 0x47, 0x43, 0x50, 0x68, 0x61, 0x73,
	0x65, 0x12, 0x07, 0x0a, 0x03, 0x4e, 0x45, 0x57, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x4f, 0x4c,
	0x44, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x4e, 0x4f, 0x52, 0x4d, 0x41, 0x4c, 0x10, 0x02, 0x32,
	0x62, 0x0a, 0x16, 0x4a, 0x56, 0x4d, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x52, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x48, 0x0a, 0x07, 0x63, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x12, 0x22, 0x2e, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e,
	0x67, 0x2e, 0x76, 0x33, 0x2e, 0x4a, 0x56, 0x4d, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x43, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x17, 0x2e, 0x73, 0x6b, 0x79, 0x77, 0x61,
	0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x73, 0x22, 0x00, 0x42, 0xba, 0x01, 0x0a, 0x33, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x70, 0x61, 0x63,
	0x68, 0x65, 0x2e, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70,
	0x6d, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61,
	0x67, 0x65, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x33, 0x50, 0x01, 0x5a, 0x61, 0x67,
	0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x6a, 0x69, 0x61, 0x67, 0x6f, 0x75, 0x79, 0x75, 0x6e, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x63, 0x61, 0x72, 0x65, 0x2d, 0x74, 0x6f,
	0x6f, 0x6c, 0x73, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x6b, 0x69, 0x74, 0x2f, 0x70, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x73, 0x2f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x2f, 0x73, 0x6b, 0x79, 0x77, 0x61,
	0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x64, 0x2f, 0x6c,
	0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x33,
	0xaa, 0x02, 0x1d, 0x53, 0x6b, 0x79, 0x57, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x4e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x56, 0x33,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_language_agent_JVMMetric_proto_rawDescOnce sync.Once
	file_language_agent_JVMMetric_proto_rawDescData = file_language_agent_JVMMetric_proto_rawDesc
)

func file_language_agent_JVMMetric_proto_rawDescGZIP() []byte {
	file_language_agent_JVMMetric_proto_rawDescOnce.Do(func() {
		file_language_agent_JVMMetric_proto_rawDescData = protoimpl.X.CompressGZIP(file_language_agent_JVMMetric_proto_rawDescData)
	})
	return file_language_agent_JVMMetric_proto_rawDescData
}

var file_language_agent_JVMMetric_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_language_agent_JVMMetric_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_language_agent_JVMMetric_proto_goTypes = []interface{}{
	(PoolType)(0),               // 0: skywalking.v3.PoolType
	(GCPhase)(0),                // 1: skywalking.v3.GCPhase
	(*JVMMetricCollection)(nil), // 2: skywalking.v3.JVMMetricCollection
	(*JVMMetric)(nil),           // 3: skywalking.v3.JVMMetric
	(*Memory)(nil),              // 4: skywalking.v3.Memory
	(*MemoryPool)(nil),          // 5: skywalking.v3.MemoryPool
	(*GC)(nil),                  // 6: skywalking.v3.GC
	(*Thread)(nil),              // 7: skywalking.v3.Thread
	(*Class)(nil),               // 8: skywalking.v3.Class
	(*v3.CPU)(nil),              // 9: skywalking.v3.CPU
	(*v3.Commands)(nil),         // 10: skywalking.v3.Commands
}
var file_language_agent_JVMMetric_proto_depIdxs = []int32{
	3,  // 0: skywalking.v3.JVMMetricCollection.metrics:type_name -> skywalking.v3.JVMMetric
	9,  // 1: skywalking.v3.JVMMetric.cpu:type_name -> skywalking.v3.CPU
	4,  // 2: skywalking.v3.JVMMetric.memory:type_name -> skywalking.v3.Memory
	5,  // 3: skywalking.v3.JVMMetric.memoryPool:type_name -> skywalking.v3.MemoryPool
	6,  // 4: skywalking.v3.JVMMetric.gc:type_name -> skywalking.v3.GC
	7,  // 5: skywalking.v3.JVMMetric.thread:type_name -> skywalking.v3.Thread
	8,  // 6: skywalking.v3.JVMMetric.clazz:type_name -> skywalking.v3.Class
	0,  // 7: skywalking.v3.MemoryPool.type:type_name -> skywalking.v3.PoolType
	1,  // 8: skywalking.v3.GC.phase:type_name -> skywalking.v3.GCPhase
	2,  // 9: skywalking.v3.JVMMetricReportService.collect:input_type -> skywalking.v3.JVMMetricCollection
	10, // 10: skywalking.v3.JVMMetricReportService.collect:output_type -> skywalking.v3.Commands
	10, // [10:11] is the sub-list for method output_type
	9,  // [9:10] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_language_agent_JVMMetric_proto_init() }
func file_language_agent_JVMMetric_proto_init() {
	if File_language_agent_JVMMetric_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_language_agent_JVMMetric_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JVMMetricCollection); i {
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
		file_language_agent_JVMMetric_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JVMMetric); i {
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
		file_language_agent_JVMMetric_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Memory); i {
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
		file_language_agent_JVMMetric_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemoryPool); i {
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
		file_language_agent_JVMMetric_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GC); i {
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
		file_language_agent_JVMMetric_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Thread); i {
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
		file_language_agent_JVMMetric_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Class); i {
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
			RawDescriptor: file_language_agent_JVMMetric_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_language_agent_JVMMetric_proto_goTypes,
		DependencyIndexes: file_language_agent_JVMMetric_proto_depIdxs,
		EnumInfos:         file_language_agent_JVMMetric_proto_enumTypes,
		MessageInfos:      file_language_agent_JVMMetric_proto_msgTypes,
	}.Build()
	File_language_agent_JVMMetric_proto = out.File
	file_language_agent_JVMMetric_proto_rawDesc = nil
	file_language_agent_JVMMetric_proto_goTypes = nil
	file_language_agent_JVMMetric_proto_depIdxs = nil
}
