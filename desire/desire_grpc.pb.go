// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.14.0
// source: desire/desire.proto

package desire

import (
	context "context"
	common "github.com/winc-link/iot-cloud-proto/common"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// RpcDesireClient is the client API for RpcDesire service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RpcDesireClient interface {
	// 设备期望属性获取
	DeviceDesireGet(ctx context.Context, in *DeviceDesireGetRequest, opts ...grpc.CallOption) (*common.CommonResponse, error)
	// 设备期望属性删除
	DeviceDesireDel(ctx context.Context, in *DeviceDesireGetRequest, opts ...grpc.CallOption) (*common.CommonResponse, error)
}

type rpcDesireClient struct {
	cc grpc.ClientConnInterface
}

func NewRpcDesireClient(cc grpc.ClientConnInterface) RpcDesireClient {
	return &rpcDesireClient{cc}
}

func (c *rpcDesireClient) DeviceDesireGet(ctx context.Context, in *DeviceDesireGetRequest, opts ...grpc.CallOption) (*common.CommonResponse, error) {
	out := new(common.CommonResponse)
	err := c.cc.Invoke(ctx, "/desire.RpcDesire/DeviceDesireGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcDesireClient) DeviceDesireDel(ctx context.Context, in *DeviceDesireGetRequest, opts ...grpc.CallOption) (*common.CommonResponse, error) {
	out := new(common.CommonResponse)
	err := c.cc.Invoke(ctx, "/desire.RpcDesire/DeviceDesireDel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RpcDesireServer is the server API for RpcDesire service.
// All implementations must embed UnimplementedRpcDesireServer
// for forward compatibility
type RpcDesireServer interface {
	// 设备期望属性获取
	DeviceDesireGet(context.Context, *DeviceDesireGetRequest) (*common.CommonResponse, error)
	// 设备期望属性删除
	DeviceDesireDel(context.Context, *DeviceDesireGetRequest) (*common.CommonResponse, error)
	mustEmbedUnimplementedRpcDesireServer()
}

// UnimplementedRpcDesireServer must be embedded to have forward compatible implementations.
type UnimplementedRpcDesireServer struct {
}

func (UnimplementedRpcDesireServer) DeviceDesireGet(context.Context, *DeviceDesireGetRequest) (*common.CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeviceDesireGet not implemented")
}
func (UnimplementedRpcDesireServer) DeviceDesireDel(context.Context, *DeviceDesireGetRequest) (*common.CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeviceDesireDel not implemented")
}
func (UnimplementedRpcDesireServer) mustEmbedUnimplementedRpcDesireServer() {}

// UnsafeRpcDesireServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RpcDesireServer will
// result in compilation errors.
type UnsafeRpcDesireServer interface {
	mustEmbedUnimplementedRpcDesireServer()
}

func RegisterRpcDesireServer(s grpc.ServiceRegistrar, srv RpcDesireServer) {
	s.RegisterService(&RpcDesire_ServiceDesc, srv)
}

func _RpcDesire_DeviceDesireGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeviceDesireGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcDesireServer).DeviceDesireGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/desire.RpcDesire/DeviceDesireGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcDesireServer).DeviceDesireGet(ctx, req.(*DeviceDesireGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RpcDesire_DeviceDesireDel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeviceDesireGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcDesireServer).DeviceDesireDel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/desire.RpcDesire/DeviceDesireDel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcDesireServer).DeviceDesireDel(ctx, req.(*DeviceDesireGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RpcDesire_ServiceDesc is the grpc.ServiceDesc for RpcDesire service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RpcDesire_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "desire.RpcDesire",
	HandlerType: (*RpcDesireServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DeviceDesireGet",
			Handler:    _RpcDesire_DeviceDesireGet_Handler,
		},
		{
			MethodName: "DeviceDesireDel",
			Handler:    _RpcDesire_DeviceDesireDel_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "desire/desire.proto",
}
