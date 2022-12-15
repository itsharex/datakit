// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v3

import (
	context "context"
	v3 "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/skywalking/compiled/v8.3.0/common/v3"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// ManagementServiceClient is the client API for ManagementService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ManagementServiceClient interface {
	// Report custom properties of a service instance.
	ReportInstanceProperties(ctx context.Context, in *InstanceProperties, opts ...grpc.CallOption) (*v3.Commands, error)
	// Keep the instance alive in the backend analysis.
	// Only recommend to do separate keepAlive report when no trace and metrics
	// needs to be reported. Otherwise, it is duplicated.
	KeepAlive(ctx context.Context, in *InstancePingPkg, opts ...grpc.CallOption) (*v3.Commands, error)
}

type managementServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewManagementServiceClient(cc grpc.ClientConnInterface) ManagementServiceClient {
	return &managementServiceClient{cc}
}

func (c *managementServiceClient) ReportInstanceProperties(ctx context.Context, in *InstanceProperties, opts ...grpc.CallOption) (*v3.Commands, error) {
	out := new(v3.Commands)
	err := c.cc.Invoke(ctx, "/ManagementService/reportInstanceProperties", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managementServiceClient) KeepAlive(ctx context.Context, in *InstancePingPkg, opts ...grpc.CallOption) (*v3.Commands, error) {
	out := new(v3.Commands)
	err := c.cc.Invoke(ctx, "/ManagementService/keepAlive", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ManagementServiceServer is the server API for ManagementService service.
// All implementations must embed UnimplementedManagementServiceServer
// for forward compatibility
type ManagementServiceServer interface {
	// Report custom properties of a service instance.
	ReportInstanceProperties(context.Context, *InstanceProperties) (*v3.Commands, error)
	// Keep the instance alive in the backend analysis.
	// Only recommend to do separate keepAlive report when no trace and metrics
	// needs to be reported. Otherwise, it is duplicated.
	KeepAlive(context.Context, *InstancePingPkg) (*v3.Commands, error)
	mustEmbedUnimplementedManagementServiceServer()
}

// UnimplementedManagementServiceServer must be embedded to have forward compatible implementations.
type UnimplementedManagementServiceServer struct {
}

func (UnimplementedManagementServiceServer) ReportInstanceProperties(context.Context, *InstanceProperties) (*v3.Commands, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReportInstanceProperties not implemented")
}
func (UnimplementedManagementServiceServer) KeepAlive(context.Context, *InstancePingPkg) (*v3.Commands, error) {
	return nil, status.Errorf(codes.Unimplemented, "method KeepAlive not implemented")
}
func (UnimplementedManagementServiceServer) mustEmbedUnimplementedManagementServiceServer() {}

// UnsafeManagementServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ManagementServiceServer will
// result in compilation errors.
type UnsafeManagementServiceServer interface {
	mustEmbedUnimplementedManagementServiceServer()
}

func RegisterManagementServiceServer(s grpc.ServiceRegistrar, srv ManagementServiceServer) {
	s.RegisterService(&_ManagementService_serviceDesc, srv)
}

func _ManagementService_ReportInstanceProperties_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InstanceProperties)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagementServiceServer).ReportInstanceProperties(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ManagementService/reportInstanceProperties",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagementServiceServer).ReportInstanceProperties(ctx, req.(*InstanceProperties))
	}
	return interceptor(ctx, in, info, handler)
}

func _ManagementService_KeepAlive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InstancePingPkg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagementServiceServer).KeepAlive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ManagementService/keepAlive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagementServiceServer).KeepAlive(ctx, req.(*InstancePingPkg))
	}
	return interceptor(ctx, in, info, handler)
}

var _ManagementService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ManagementService",
	HandlerType: (*ManagementServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "reportInstanceProperties",
			Handler:    _ManagementService_ReportInstanceProperties_Handler,
		},
		{
			MethodName: "keepAlive",
			Handler:    _ManagementService_KeepAlive_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "management/ManagementOld.proto",
}
