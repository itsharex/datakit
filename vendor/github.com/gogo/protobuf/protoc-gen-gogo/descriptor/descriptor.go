// Go support for Protocol Buffers - Google's data interchange format
//
// Copyright 2016 The Go Authors.  All rights reserved.
// https://github.com/golang/protobuf
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
//
//     * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
//     * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
//     * Neither the name of Google Inc. nor the names of its
// contributors may be used to endorse or promote products derived from
// this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

// Package descriptor provides functions for obtaining protocol buffer
// descriptors for generated Go types.
//
// These functions cannot go in package proto because they depend on the
// generated protobuf descriptor messages, which themselves depend on proto.
package descriptor

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io/ioutil"

	"github.com/gogo/protobuf/proto"
)

// extractFile extracts a FileDescriptorProto from a gzip'd buffer.
func extractFile(gz []byte) (*FileDescriptorProto, error) {
	r, err := gzip.NewReader(bytes.NewReader(gz))
	if err != nil {
		return nil, fmt.Errorf("failed to open gzip reader: %v", err)
	}
	defer r.Close()

	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("failed to uncompress descriptor: %v", err)
	}

	fd := new(FileDescriptorProto)
	if err := proto.Unmarshal(b, fd); err != nil {
		return nil, fmt.Errorf("malformed FileDescriptorProto: %v", err)
	}

	return fd, nil
}

// Message is a proto.Message with a method to return its descriptor.
//
// Message types generated by the protocol compiler always satisfy
// the Message interface.
type Message interface {
	proto.Message
	Descriptor() ([]byte, []int)
}

// ForMessage returns a FileDescriptorProto and a DescriptorProto from within it
// describing the given message.
func ForMessage(msg Message) (fd *FileDescriptorProto, md *DescriptorProto) {
	gz, path := msg.Descriptor()
	fd, err := extractFile(gz)
	if err != nil {
		panic(fmt.Sprintf("invalid FileDescriptorProto for %T: %v", msg, err))
	}

	md = fd.MessageType[path[0]]
	for _, i := range path[1:] {
		md = md.NestedType[i]
	}
	return fd, md
}

// Is this field a scalar numeric type?
func (field *FieldDescriptorProto) IsScalar() bool {
	if field.Type == nil {
		return false
	}
	switch *field.Type {
	case FieldDescriptorProto_TYPE_DOUBLE,
		FieldDescriptorProto_TYPE_FLOAT,
		FieldDescriptorProto_TYPE_INT64,
		FieldDescriptorProto_TYPE_UINT64,
		FieldDescriptorProto_TYPE_INT32,
		FieldDescriptorProto_TYPE_FIXED64,
		FieldDescriptorProto_TYPE_FIXED32,
		FieldDescriptorProto_TYPE_BOOL,
		FieldDescriptorProto_TYPE_UINT32,
		FieldDescriptorProto_TYPE_ENUM,
		FieldDescriptorProto_TYPE_SFIXED32,
		FieldDescriptorProto_TYPE_SFIXED64,
		FieldDescriptorProto_TYPE_SINT32,
		FieldDescriptorProto_TYPE_SINT64:
		return true
	default:
		return false
	}
}
