// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.14.0
// source: service/service.proto

package event

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

// RpcCommandClient is the client API for RpcCommand service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RpcCommandClient interface {
	//设备端响应服务调用
	DeviceCommandReply(ctx context.Context, in *DeviceCommandReplyRequest, opts ...grpc.CallOption) (*common.CommonResponse, error)
}

type rpcCommandClient struct {
	cc grpc.ClientConnInterface
}

func NewRpcCommandClient(cc grpc.ClientConnInterface) RpcCommandClient {
	return &rpcCommandClient{cc}
}

func (c *rpcCommandClient) DeviceCommandReply(ctx context.Context, in *DeviceCommandReplyRequest, opts ...grpc.CallOption) (*common.CommonResponse, error) {
	out := new(common.CommonResponse)
	err := c.cc.Invoke(ctx, "/event.RpcCommand/DeviceCommandReply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RpcCommandServer is the server API for RpcCommand service.
// All implementations must embed UnimplementedRpcCommandServer
// for forward compatibility
type RpcCommandServer interface {
	//设备端响应服务调用
	DeviceCommandReply(context.Context, *DeviceCommandReplyRequest) (*common.CommonResponse, error)
	mustEmbedUnimplementedRpcCommandServer()
}

// UnimplementedRpcCommandServer must be embedded to have forward compatible implementations.
type UnimplementedRpcCommandServer struct {
}

func (UnimplementedRpcCommandServer) DeviceCommandReply(context.Context, *DeviceCommandReplyRequest) (*common.CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeviceCommandReply not implemented")
}
func (UnimplementedRpcCommandServer) mustEmbedUnimplementedRpcCommandServer() {}

// UnsafeRpcCommandServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RpcCommandServer will
// result in compilation errors.
type UnsafeRpcCommandServer interface {
	mustEmbedUnimplementedRpcCommandServer()
}

func RegisterRpcCommandServer(s grpc.ServiceRegistrar, srv RpcCommandServer) {
	s.RegisterService(&RpcCommand_ServiceDesc, srv)
}

func _RpcCommand_DeviceCommandReply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeviceCommandReplyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcCommandServer).DeviceCommandReply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/event.RpcCommand/DeviceCommandReply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcCommandServer).DeviceCommandReply(ctx, req.(*DeviceCommandReplyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RpcCommand_ServiceDesc is the grpc.ServiceDesc for RpcCommand service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RpcCommand_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "event.RpcCommand",
	HandlerType: (*RpcCommandServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DeviceCommandReply",
			Handler:    _RpcCommand_DeviceCommandReply_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service/service.proto",
}
