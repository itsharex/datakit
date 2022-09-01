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
// browser/BrowserPerfCompat.proto is a deprecated file.

package compat

import (
	v31 "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/skywalking/compiled/common/v3"
	v3 "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/skywalking/compiled/language/agent/v3"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_browser_BrowserPerfCompat_proto protoreflect.FileDescriptor

var file_browser_BrowserPerfCompat_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x62, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x72, 0x2f, 0x42, 0x72, 0x6f, 0x77, 0x73, 0x65,
	0x72, 0x50, 0x65, 0x72, 0x66, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x62, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x72, 0x2f,
	0x42, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x72, 0x50, 0x65, 0x72, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x32, 0xb3, 0x01, 0x0a, 0x12, 0x42, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x72, 0x50, 0x65, 0x72,
	0x66, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4c, 0x0a, 0x0f, 0x63, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x50, 0x65, 0x72, 0x66, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1e, 0x2e, 0x73, 0x6b,
	0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x42, 0x72, 0x6f, 0x77,
	0x73, 0x65, 0x72, 0x50, 0x65, 0x72, 0x66, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x17, 0x2e, 0x73, 0x6b,
	0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x43, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x73, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x10, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x4c, 0x6f, 0x67, 0x73, 0x12, 0x1e, 0x2e, 0x73, 0x6b, 0x79,
	0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x42, 0x72, 0x6f, 0x77, 0x73,
	0x65, 0x72, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x4c, 0x6f, 0x67, 0x1a, 0x17, 0x2e, 0x73, 0x6b, 0x79,
	0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x73, 0x22, 0x00, 0x28, 0x01, 0x42, 0xd2, 0x01, 0x0a, 0x3a, 0x6f, 0x72, 0x67, 0x2e,
	0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e,
	0x67, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x6c, 0x61,
	0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x33, 0x2e,
	0x63, 0x6f, 0x6d, 0x70, 0x61, 0x74, 0x50, 0x01, 0x5a, 0x68, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62,
	0x2e, 0x6a, 0x69, 0x61, 0x67, 0x6f, 0x75, 0x79, 0x75, 0x6e, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x63, 0x61, 0x72, 0x65, 0x2d, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2f, 0x64,
	0x61, 0x74, 0x61, 0x6b, 0x69, 0x74, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2f, 0x69,
	0x6e, 0x70, 0x75, 0x74, 0x73, 0x2f, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67,
	0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x64, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61,
	0x67, 0x65, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x33, 0x2f, 0x63, 0x6f, 0x6d, 0x70,
	0x61, 0x74, 0xb8, 0x01, 0x01, 0xaa, 0x02, 0x24, 0x53, 0x6b, 0x79, 0x57, 0x61, 0x6c, 0x6b, 0x69,
	0x6e, 0x67, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x2e, 0x56, 0x33, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x74, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var file_browser_BrowserPerfCompat_proto_goTypes = []interface{}{
	(*v3.BrowserPerfData)(nil), // 0: skywalking.v3.BrowserPerfData
	(*v3.BrowserErrorLog)(nil), // 1: skywalking.v3.BrowserErrorLog
	(*v31.Commands)(nil),       // 2: skywalking.v3.Commands
}
var file_browser_BrowserPerfCompat_proto_depIdxs = []int32{
	0, // 0: BrowserPerfService.collectPerfData:input_type -> skywalking.v3.BrowserPerfData
	1, // 1: BrowserPerfService.collectErrorLogs:input_type -> skywalking.v3.BrowserErrorLog
	2, // 2: BrowserPerfService.collectPerfData:output_type -> skywalking.v3.Commands
	2, // 3: BrowserPerfService.collectErrorLogs:output_type -> skywalking.v3.Commands
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_browser_BrowserPerfCompat_proto_init() }
func file_browser_BrowserPerfCompat_proto_init() {
	if File_browser_BrowserPerfCompat_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_browser_BrowserPerfCompat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_browser_BrowserPerfCompat_proto_goTypes,
		DependencyIndexes: file_browser_BrowserPerfCompat_proto_depIdxs,
	}.Build()
	File_browser_BrowserPerfCompat_proto = out.File
	file_browser_BrowserPerfCompat_proto_rawDesc = nil
	file_browser_BrowserPerfCompat_proto_goTypes = nil
	file_browser_BrowserPerfCompat_proto_depIdxs = nil
}
