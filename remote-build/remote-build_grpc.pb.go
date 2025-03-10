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

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: remote-build/remote-build.proto

package remote_build

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ServerWorker_HelloWorker_FullMethodName = "/remote_build.Server_worker/HelloWorker"
)

// ServerWorkerClient is the client API for ServerWorker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Server to Worker
type ServerWorkerClient interface {
	// Sends a greeting
	HelloWorker(ctx context.Context, in *WorkRequest, opts ...grpc.CallOption) (*WorkResponce, error)
}

type serverWorkerClient struct {
	cc grpc.ClientConnInterface
}

func NewServerWorkerClient(cc grpc.ClientConnInterface) ServerWorkerClient {
	return &serverWorkerClient{cc}
}

func (c *serverWorkerClient) HelloWorker(ctx context.Context, in *WorkRequest, opts ...grpc.CallOption) (*WorkResponce, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(WorkResponce)
	err := c.cc.Invoke(ctx, ServerWorker_HelloWorker_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServerWorkerServer is the server API for ServerWorker service.
// All implementations must embed UnimplementedServerWorkerServer
// for forward compatibility.
//
// Server to Worker
type ServerWorkerServer interface {
	// Sends a greeting
	HelloWorker(context.Context, *WorkRequest) (*WorkResponce, error)
	mustEmbedUnimplementedServerWorkerServer()
}

// UnimplementedServerWorkerServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedServerWorkerServer struct{}

func (UnimplementedServerWorkerServer) HelloWorker(context.Context, *WorkRequest) (*WorkResponce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HelloWorker not implemented")
}
func (UnimplementedServerWorkerServer) mustEmbedUnimplementedServerWorkerServer() {}
func (UnimplementedServerWorkerServer) testEmbeddedByValue()                      {}

// UnsafeServerWorkerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServerWorkerServer will
// result in compilation errors.
type UnsafeServerWorkerServer interface {
	mustEmbedUnimplementedServerWorkerServer()
}

func RegisterServerWorkerServer(s grpc.ServiceRegistrar, srv ServerWorkerServer) {
	// If the following call pancis, it indicates UnimplementedServerWorkerServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ServerWorker_ServiceDesc, srv)
}

func _ServerWorker_HelloWorker_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WorkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerWorkerServer).HelloWorker(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServerWorker_HelloWorker_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerWorkerServer).HelloWorker(ctx, req.(*WorkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ServerWorker_ServiceDesc is the grpc.ServiceDesc for ServerWorker service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServerWorker_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "remote_build.Server_worker",
	HandlerType: (*ServerWorkerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HelloWorker",
			Handler:    _ServerWorker_HelloWorker_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "remote-build/remote-build.proto",
}

const (
	ClientServer_HelloServer_FullMethodName = "/remote_build.Client_server/HelloServer"
)

// ClientServerClient is the client API for ClientServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Clinet to Server
type ClientServerClient interface {
	HelloServer(ctx context.Context, in *BuildRequest, opts ...grpc.CallOption) (*BuildResponse, error)
}

type clientServerClient struct {
	cc grpc.ClientConnInterface
}

func NewClientServerClient(cc grpc.ClientConnInterface) ClientServerClient {
	return &clientServerClient{cc}
}

func (c *clientServerClient) HelloServer(ctx context.Context, in *BuildRequest, opts ...grpc.CallOption) (*BuildResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BuildResponse)
	err := c.cc.Invoke(ctx, ClientServer_HelloServer_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClientServerServer is the server API for ClientServer service.
// All implementations must embed UnimplementedClientServerServer
// for forward compatibility.
//
// Clinet to Server
type ClientServerServer interface {
	HelloServer(context.Context, *BuildRequest) (*BuildResponse, error)
	mustEmbedUnimplementedClientServerServer()
}

// UnimplementedClientServerServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedClientServerServer struct{}

func (UnimplementedClientServerServer) HelloServer(context.Context, *BuildRequest) (*BuildResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HelloServer not implemented")
}
func (UnimplementedClientServerServer) mustEmbedUnimplementedClientServerServer() {}
func (UnimplementedClientServerServer) testEmbeddedByValue()                      {}

// UnsafeClientServerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClientServerServer will
// result in compilation errors.
type UnsafeClientServerServer interface {
	mustEmbedUnimplementedClientServerServer()
}

func RegisterClientServerServer(s grpc.ServiceRegistrar, srv ClientServerServer) {
	// If the following call pancis, it indicates UnimplementedClientServerServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ClientServer_ServiceDesc, srv)
}

func _ClientServer_HelloServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BuildRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServerServer).HelloServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClientServer_HelloServer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServerServer).HelloServer(ctx, req.(*BuildRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ClientServer_ServiceDesc is the grpc.ServiceDesc for ClientServer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ClientServer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "remote_build.Client_server",
	HandlerType: (*ClientServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HelloServer",
			Handler:    _ClientServer_HelloServer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "remote-build/remote-build.proto",
}
