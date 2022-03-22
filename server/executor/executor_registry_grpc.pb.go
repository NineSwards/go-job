// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: server/proto/executor_registry.proto

package server

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// JobExecutorRegisterClient is the client API for JobExecutorRegister service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type JobExecutorRegisterClient interface {
	ExecutorRegistry(ctx context.Context, in *RegistryRequest, opts ...grpc.CallOption) (*RegistryReply, error)
	ExecutorUnRegistry(ctx context.Context, in *RegistryRequest, opts ...grpc.CallOption) (*RegistryReply, error)
}

type jobExecutorRegisterClient struct {
	cc grpc.ClientConnInterface
}

func NewJobExecutorRegisterClient(cc grpc.ClientConnInterface) JobExecutorRegisterClient {
	return &jobExecutorRegisterClient{cc}
}

func (c *jobExecutorRegisterClient) ExecutorRegistry(ctx context.Context, in *RegistryRequest, opts ...grpc.CallOption) (*RegistryReply, error) {
	out := new(RegistryReply)
	err := c.cc.Invoke(ctx, "/proto.registry.JobExecutorRegister/ExecutorRegistry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobExecutorRegisterClient) ExecutorUnRegistry(ctx context.Context, in *RegistryRequest, opts ...grpc.CallOption) (*RegistryReply, error) {
	out := new(RegistryReply)
	err := c.cc.Invoke(ctx, "/proto.registry.JobExecutorRegister/ExecutorUnRegistry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JobExecutorRegisterServer is the server API for JobExecutorRegister service.
// All implementations must embed UnimplementedJobExecutorRegisterServer
// for forward compatibility
type JobExecutorRegisterServer interface {
	ExecutorRegistry(context.Context, *RegistryRequest) (*RegistryReply, error)
	ExecutorUnRegistry(context.Context, *RegistryRequest) (*RegistryReply, error)
	mustEmbedUnimplementedJobExecutorRegisterServer()
}

// UnimplementedJobExecutorRegisterServer must be embedded to have forward compatible implementations.
type UnimplementedJobExecutorRegisterServer struct {
}

func (UnimplementedJobExecutorRegisterServer) ExecutorRegistry(context.Context, *RegistryRequest) (*RegistryReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExecutorRegistry not implemented")
}
func (UnimplementedJobExecutorRegisterServer) ExecutorUnRegistry(context.Context, *RegistryRequest) (*RegistryReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExecutorUnRegistry not implemented")
}
func (UnimplementedJobExecutorRegisterServer) mustEmbedUnimplementedJobExecutorRegisterServer() {}

// UnsafeJobExecutorRegisterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to JobExecutorRegisterServer will
// result in compilation errors.
type UnsafeJobExecutorRegisterServer interface {
	mustEmbedUnimplementedJobExecutorRegisterServer()
}

func RegisterJobExecutorRegisterServer(s grpc.ServiceRegistrar, srv JobExecutorRegisterServer) {
	s.RegisterService(&JobExecutorRegister_ServiceDesc, srv)
}

func _JobExecutorRegister_ExecutorRegistry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegistryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobExecutorRegisterServer).ExecutorRegistry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.registry.JobExecutorRegister/ExecutorRegistry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobExecutorRegisterServer).ExecutorRegistry(ctx, req.(*RegistryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobExecutorRegister_ExecutorUnRegistry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegistryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobExecutorRegisterServer).ExecutorUnRegistry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.registry.JobExecutorRegister/ExecutorUnRegistry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobExecutorRegisterServer).ExecutorUnRegistry(ctx, req.(*RegistryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// JobExecutorRegister_ServiceDesc is the grpc.ServiceDesc for JobExecutorRegister service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var JobExecutorRegister_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.registry.JobExecutorRegister",
	HandlerType: (*JobExecutorRegisterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ExecutorRegistry",
			Handler:    _JobExecutorRegister_ExecutorRegistry_Handler,
		},
		{
			MethodName: "ExecutorUnRegistry",
			Handler:    _JobExecutorRegister_ExecutorUnRegistry_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server/proto/executor_registry.proto",
}