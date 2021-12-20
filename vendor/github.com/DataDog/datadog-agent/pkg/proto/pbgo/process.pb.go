// Code generated by protoc-gen-go. DO NOT EDIT.
// source: datadog/process/process.proto

package pbgo

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// ProcessStatRequest is the request to get process stats.
type ProcessStatRequest struct {
	Pids                 []int32  `protobuf:"varint,1,rep,packed,name=pids,proto3" json:"pids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProcessStatRequest) Reset()         { *m = ProcessStatRequest{} }
func (m *ProcessStatRequest) String() string { return proto.CompactTextString(m) }
func (*ProcessStatRequest) ProtoMessage()    {}
func (*ProcessStatRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_07928c410743578c, []int{0}
}

func (m *ProcessStatRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProcessStatRequest.Unmarshal(m, b)
}
func (m *ProcessStatRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProcessStatRequest.Marshal(b, m, deterministic)
}
func (m *ProcessStatRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProcessStatRequest.Merge(m, src)
}
func (m *ProcessStatRequest) XXX_Size() int {
	return xxx_messageInfo_ProcessStatRequest.Size(m)
}
func (m *ProcessStatRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ProcessStatRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ProcessStatRequest proto.InternalMessageInfo

func (m *ProcessStatRequest) GetPids() []int32 {
	if m != nil {
		return m.Pids
	}
	return nil
}

func init() {
	proto.RegisterType((*ProcessStatRequest)(nil), "datadog.process.ProcessStatRequest")
}

func init() { proto.RegisterFile("datadog/process/process.proto", fileDescriptor_07928c410743578c) }

var fileDescriptor_07928c410743578c = []byte{
	// 110 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4d, 0x49, 0x2c, 0x49,
	0x4c, 0xc9, 0x4f, 0xd7, 0x2f, 0x28, 0xca, 0x4f, 0x4e, 0x2d, 0x2e, 0x86, 0xd1, 0x7a, 0x05, 0x45,
	0xf9, 0x25, 0xf9, 0x42, 0xfc, 0x50, 0x69, 0x3d, 0xa8, 0xb0, 0x92, 0x06, 0x97, 0x50, 0x00, 0x84,
	0x19, 0x5c, 0x92, 0x58, 0x12, 0x94, 0x5a, 0x58, 0x9a, 0x5a, 0x5c, 0x22, 0x24, 0xc4, 0xc5, 0x52,
	0x90, 0x99, 0x52, 0x2c, 0xc1, 0xa8, 0xc0, 0xac, 0xc1, 0x1a, 0x04, 0x66, 0x3b, 0x09, 0x44, 0xf1,
	0x15, 0x64, 0x83, 0xcd, 0x2d, 0xc9, 0xd7, 0x2f, 0x48, 0x4a, 0xcf, 0x4f, 0x62, 0x03, 0xb3, 0x8d,
	0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x3c, 0xb7, 0xfb, 0xee, 0x74, 0x00, 0x00, 0x00,
}
