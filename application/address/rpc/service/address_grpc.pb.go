// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.3
// source: address.proto

package service

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

// AddressClient is the client API for Address service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AddressClient interface {
	FindAdressById(ctx context.Context, in *FindAdressByIdReq, opts ...grpc.CallOption) (*FindAdressByIdResp, error)
	GetUserDefaultAddress(ctx context.Context, in *GetUserDefaultAddressReq, opts ...grpc.CallOption) (*FindAdressByIdResp, error)
}

type addressClient struct {
	cc grpc.ClientConnInterface
}

func NewAddressClient(cc grpc.ClientConnInterface) AddressClient {
	return &addressClient{cc}
}

func (c *addressClient) FindAdressById(ctx context.Context, in *FindAdressByIdReq, opts ...grpc.CallOption) (*FindAdressByIdResp, error) {
	out := new(FindAdressByIdResp)
	err := c.cc.Invoke(ctx, "/service.Address/FindAdressById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addressClient) GetUserDefaultAddress(ctx context.Context, in *GetUserDefaultAddressReq, opts ...grpc.CallOption) (*FindAdressByIdResp, error) {
	out := new(FindAdressByIdResp)
	err := c.cc.Invoke(ctx, "/service.Address/GetUserDefaultAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AddressServer is the server API for Address service.
// All implementations must embed UnimplementedAddressServer
// for forward compatibility
type AddressServer interface {
	FindAdressById(context.Context, *FindAdressByIdReq) (*FindAdressByIdResp, error)
	GetUserDefaultAddress(context.Context, *GetUserDefaultAddressReq) (*FindAdressByIdResp, error)
	mustEmbedUnimplementedAddressServer()
}

// UnimplementedAddressServer must be embedded to have forward compatible implementations.
type UnimplementedAddressServer struct {
}

func (UnimplementedAddressServer) FindAdressById(context.Context, *FindAdressByIdReq) (*FindAdressByIdResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAdressById not implemented")
}
func (UnimplementedAddressServer) GetUserDefaultAddress(context.Context, *GetUserDefaultAddressReq) (*FindAdressByIdResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserDefaultAddress not implemented")
}
func (UnimplementedAddressServer) mustEmbedUnimplementedAddressServer() {}

// UnsafeAddressServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AddressServer will
// result in compilation errors.
type UnsafeAddressServer interface {
	mustEmbedUnimplementedAddressServer()
}

func RegisterAddressServer(s grpc.ServiceRegistrar, srv AddressServer) {
	s.RegisterService(&Address_ServiceDesc, srv)
}

func _Address_FindAdressById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAdressByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddressServer).FindAdressById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.Address/FindAdressById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddressServer).FindAdressById(ctx, req.(*FindAdressByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Address_GetUserDefaultAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserDefaultAddressReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddressServer).GetUserDefaultAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.Address/GetUserDefaultAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddressServer).GetUserDefaultAddress(ctx, req.(*GetUserDefaultAddressReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Address_ServiceDesc is the grpc.ServiceDesc for Address service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Address_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.Address",
	HandlerType: (*AddressServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindAdressById",
			Handler:    _Address_FindAdressById_Handler,
		},
		{
			MethodName: "GetUserDefaultAddress",
			Handler:    _Address_GetUserDefaultAddress_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "address.proto",
}
