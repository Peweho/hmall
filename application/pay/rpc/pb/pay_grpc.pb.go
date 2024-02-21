// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.3
// source: pay.proto

package pb

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

// PayClient is the client API for Pay service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PayClient interface {
	UpdatePayOrder(ctx context.Context, in *UpdatePayOrderReq, opts ...grpc.CallOption) (*UpdatePayOrderResp, error)
}

type payClient struct {
	cc grpc.ClientConnInterface
}

func NewPayClient(cc grpc.ClientConnInterface) PayClient {
	return &payClient{cc}
}

func (c *payClient) UpdatePayOrder(ctx context.Context, in *UpdatePayOrderReq, opts ...grpc.CallOption) (*UpdatePayOrderResp, error) {
	out := new(UpdatePayOrderResp)
	err := c.cc.Invoke(ctx, "/service.Pay/UpdatePayOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PayServer is the server API for Pay service.
// All implementations must embed UnimplementedPayServer
// for forward compatibility
type PayServer interface {
	UpdatePayOrder(context.Context, *UpdatePayOrderReq) (*UpdatePayOrderResp, error)
	mustEmbedUnimplementedPayServer()
}

// UnimplementedPayServer must be embedded to have forward compatible implementations.
type UnimplementedPayServer struct {
}

func (UnimplementedPayServer) UpdatePayOrder(context.Context, *UpdatePayOrderReq) (*UpdatePayOrderResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePayOrder not implemented")
}
func (UnimplementedPayServer) mustEmbedUnimplementedPayServer() {}

// UnsafePayServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PayServer will
// result in compilation errors.
type UnsafePayServer interface {
	mustEmbedUnimplementedPayServer()
}

func RegisterPayServer(s grpc.ServiceRegistrar, srv PayServer) {
	s.RegisterService(&Pay_ServiceDesc, srv)
}

func _Pay_UpdatePayOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePayOrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayServer).UpdatePayOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.Pay/UpdatePayOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayServer).UpdatePayOrder(ctx, req.(*UpdatePayOrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Pay_ServiceDesc is the grpc.ServiceDesc for Pay service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Pay_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.Pay",
	HandlerType: (*PayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdatePayOrder",
			Handler:    _Pay_UpdatePayOrder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pay.proto",
}
