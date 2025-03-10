// Copyright 2015 gRPC authors.
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
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: remote-build/remote-build.proto

package remote_build

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The request message containing the user's name.
type WorkRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WorkRequest) Reset() {
	*x = WorkRequest{}
	mi := &file_remote_build_remote_build_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WorkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkRequest) ProtoMessage() {}

func (x *WorkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_remote_build_remote_build_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkRequest.ProtoReflect.Descriptor instead.
func (*WorkRequest) Descriptor() ([]byte, []int) {
	return file_remote_build_remote_build_proto_rawDescGZIP(), []int{0}
}

func (x *WorkRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// The response message containing the greetings
type WorkResponce struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WorkResponce) Reset() {
	*x = WorkResponce{}
	mi := &file_remote_build_remote_build_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WorkResponce) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkResponce) ProtoMessage() {}

func (x *WorkResponce) ProtoReflect() protoreflect.Message {
	mi := &file_remote_build_remote_build_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkResponce.ProtoReflect.Descriptor instead.
func (*WorkResponce) Descriptor() ([]byte, []int) {
	return file_remote_build_remote_build_proto_rawDescGZIP(), []int{1}
}

func (x *WorkResponce) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type BuildRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BuildRequest) Reset() {
	*x = BuildRequest{}
	mi := &file_remote_build_remote_build_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BuildRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuildRequest) ProtoMessage() {}

func (x *BuildRequest) ProtoReflect() protoreflect.Message {
	mi := &file_remote_build_remote_build_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuildRequest.ProtoReflect.Descriptor instead.
func (*BuildRequest) Descriptor() ([]byte, []int) {
	return file_remote_build_remote_build_proto_rawDescGZIP(), []int{2}
}

func (x *BuildRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type BuildResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BuildResponse) Reset() {
	*x = BuildResponse{}
	mi := &file_remote_build_remote_build_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BuildResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuildResponse) ProtoMessage() {}

func (x *BuildResponse) ProtoReflect() protoreflect.Message {
	mi := &file_remote_build_remote_build_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuildResponse.ProtoReflect.Descriptor instead.
func (*BuildResponse) Descriptor() ([]byte, []int) {
	return file_remote_build_remote_build_proto_rawDescGZIP(), []int{3}
}

func (x *BuildResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_remote_build_remote_build_proto protoreflect.FileDescriptor

var file_remote_build_remote_build_proto_rawDesc = string([]byte{
	0x0a, 0x1f, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2d, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x72,
	0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2d, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0c, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x22,
	0x21, 0x0a, 0x0b, 0x57, 0x6f, 0x72, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x28, 0x0a, 0x0c, 0x57, 0x6f, 0x72, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x22, 0x0a, 0x0c,
	0x42, 0x75, 0x69, 0x6c, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x29, 0x0a, 0x0d, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x57, 0x0a, 0x0d, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x12, 0x46, 0x0a, 0x0b,
	0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x12, 0x19, 0x2e, 0x72, 0x65,
	0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x63, 0x65, 0x22, 0x00, 0x32, 0x59, 0x0a, 0x0d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x48, 0x0a, 0x0b, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x12, 0x1a, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x62, 0x75,
	0x69, 0x6c, 0x64, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1b, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e,
	0x42, 0x75, 0x69, 0x6c, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x1b, 0x5a, 0x19, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2d, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f,
	0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2d, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_remote_build_remote_build_proto_rawDescOnce sync.Once
	file_remote_build_remote_build_proto_rawDescData []byte
)

func file_remote_build_remote_build_proto_rawDescGZIP() []byte {
	file_remote_build_remote_build_proto_rawDescOnce.Do(func() {
		file_remote_build_remote_build_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_remote_build_remote_build_proto_rawDesc), len(file_remote_build_remote_build_proto_rawDesc)))
	})
	return file_remote_build_remote_build_proto_rawDescData
}

var file_remote_build_remote_build_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_remote_build_remote_build_proto_goTypes = []any{
	(*WorkRequest)(nil),   // 0: remote_build.WorkRequest
	(*WorkResponce)(nil),  // 1: remote_build.WorkResponce
	(*BuildRequest)(nil),  // 2: remote_build.BuildRequest
	(*BuildResponse)(nil), // 3: remote_build.BuildResponse
}
var file_remote_build_remote_build_proto_depIdxs = []int32{
	0, // 0: remote_build.Server_worker.HelloWorker:input_type -> remote_build.WorkRequest
	2, // 1: remote_build.Client_server.HelloServer:input_type -> remote_build.BuildRequest
	1, // 2: remote_build.Server_worker.HelloWorker:output_type -> remote_build.WorkResponce
	3, // 3: remote_build.Client_server.HelloServer:output_type -> remote_build.BuildResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_remote_build_remote_build_proto_init() }
func file_remote_build_remote_build_proto_init() {
	if File_remote_build_remote_build_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_remote_build_remote_build_proto_rawDesc), len(file_remote_build_remote_build_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_remote_build_remote_build_proto_goTypes,
		DependencyIndexes: file_remote_build_remote_build_proto_depIdxs,
		MessageInfos:      file_remote_build_remote_build_proto_msgTypes,
	}.Build()
	File_remote_build_remote_build_proto = out.File
	file_remote_build_remote_build_proto_goTypes = nil
	file_remote_build_remote_build_proto_depIdxs = nil
}
