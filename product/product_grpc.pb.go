// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.14.0
// source: product/product.proto

package product

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

// RpcProductClient is the client API for RpcProduct service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RpcProductClient interface {
	//查询所有产品
	QueryProductList(ctx context.Context, in *QueryProductListRequest, opts ...grpc.CallOption) (*QueryProductListResponse, error)
	//查询单个产品
	QueryProduct(ctx context.Context, in *QueryProductRequest, opts ...grpc.CallOption) (*QueryProductResponse, error)
}

type rpcProductClient struct {
	cc grpc.ClientConnInterface
}

func NewRpcProductClient(cc grpc.ClientConnInterface) RpcProductClient {
	return &rpcProductClient{cc}
}

func (c *rpcProductClient) QueryProductList(ctx context.Context, in *QueryProductListRequest, opts ...grpc.CallOption) (*QueryProductListResponse, error) {
	out := new(QueryProductListResponse)
	err := c.cc.Invoke(ctx, "/product.RpcProduct/QueryProductList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcProductClient) QueryProduct(ctx context.Context, in *QueryProductRequest, opts ...grpc.CallOption) (*QueryProductResponse, error) {
	out := new(QueryProductResponse)
	err := c.cc.Invoke(ctx, "/product.RpcProduct/QueryProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RpcProductServer is the server API for RpcProduct service.
// All implementations must embed UnimplementedRpcProductServer
// for forward compatibility
type RpcProductServer interface {
	//查询所有产品
	QueryProductList(context.Context, *QueryProductListRequest) (*QueryProductListResponse, error)
	//查询单个产品
	QueryProduct(context.Context, *QueryProductRequest) (*QueryProductResponse, error)
	mustEmbedUnimplementedRpcProductServer()
}

// UnimplementedRpcProductServer must be embedded to have forward compatible implementations.
type UnimplementedRpcProductServer struct {
}

func (UnimplementedRpcProductServer) QueryProductList(context.Context, *QueryProductListRequest) (*QueryProductListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryProductList not implemented")
}
func (UnimplementedRpcProductServer) QueryProduct(context.Context, *QueryProductRequest) (*QueryProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryProduct not implemented")
}
func (UnimplementedRpcProductServer) mustEmbedUnimplementedRpcProductServer() {}

// UnsafeRpcProductServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RpcProductServer will
// result in compilation errors.
type UnsafeRpcProductServer interface {
	mustEmbedUnimplementedRpcProductServer()
}

func RegisterRpcProductServer(s grpc.ServiceRegistrar, srv RpcProductServer) {
	s.RegisterService(&RpcProduct_ServiceDesc, srv)
}

func _RpcProduct_QueryProductList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryProductListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcProductServer).QueryProductList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.RpcProduct/QueryProductList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcProductServer).QueryProductList(ctx, req.(*QueryProductListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RpcProduct_QueryProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcProductServer).QueryProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.RpcProduct/QueryProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcProductServer).QueryProduct(ctx, req.(*QueryProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RpcProduct_ServiceDesc is the grpc.ServiceDesc for RpcProduct service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RpcProduct_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "product.RpcProduct",
	HandlerType: (*RpcProductServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "QueryProductList",
			Handler:    _RpcProduct_QueryProductList_Handler,
		},
		{
			MethodName: "QueryProduct",
			Handler:    _RpcProduct_QueryProduct_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "product/product.proto",
}
