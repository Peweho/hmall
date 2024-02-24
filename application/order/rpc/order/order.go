// Code generated by goctl. DO NOT EDIT.
// Source: order.proto

package order

import (
	"context"

	"hmall/application/order/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CreateOrderReq        = pb.CreateOrderReq
	CreateOrderResp       = pb.CreateOrderResp
	DetailDTO             = pb.DetailDTO
	FindOrderByIdReq      = pb.FindOrderByIdReq
	FindOrderByIdResp     = pb.FindOrderByIdResp
	UpdateOrderStatusReq  = pb.UpdateOrderStatusReq
	UpdateOrderStatusResp = pb.UpdateOrderStatusResp

	Order interface {
		FindOrderById(ctx context.Context, in *FindOrderByIdReq, opts ...grpc.CallOption) (*FindOrderByIdResp, error)
		UpdateOrderStatus(ctx context.Context, in *UpdateOrderStatusReq, opts ...grpc.CallOption) (*UpdateOrderStatusResp, error)
		UpdateOrderStatusRollBack(ctx context.Context, in *UpdateOrderStatusReq, opts ...grpc.CallOption) (*UpdateOrderStatusResp, error)
		CreateOrder(ctx context.Context, in *CreateOrderReq, opts ...grpc.CallOption) (*CreateOrderResp, error)
		CreateOrderRollBack(ctx context.Context, in *CreateOrderReq, opts ...grpc.CallOption) (*CreateOrderResp, error)
	}

	defaultOrder struct {
		cli zrpc.Client
	}
)

func NewOrder(cli zrpc.Client) Order {
	return &defaultOrder{
		cli: cli,
	}
}

func (m *defaultOrder) FindOrderById(ctx context.Context, in *FindOrderByIdReq, opts ...grpc.CallOption) (*FindOrderByIdResp, error) {
	client := pb.NewOrderClient(m.cli.Conn())
	return client.FindOrderById(ctx, in, opts...)
}

func (m *defaultOrder) UpdateOrderStatus(ctx context.Context, in *UpdateOrderStatusReq, opts ...grpc.CallOption) (*UpdateOrderStatusResp, error) {
	client := pb.NewOrderClient(m.cli.Conn())
	return client.UpdateOrderStatus(ctx, in, opts...)
}

func (m *defaultOrder) UpdateOrderStatusRollBack(ctx context.Context, in *UpdateOrderStatusReq, opts ...grpc.CallOption) (*UpdateOrderStatusResp, error) {
	client := pb.NewOrderClient(m.cli.Conn())
	return client.UpdateOrderStatusRollBack(ctx, in, opts...)
}

func (m *defaultOrder) CreateOrder(ctx context.Context, in *CreateOrderReq, opts ...grpc.CallOption) (*CreateOrderResp, error) {
	client := pb.NewOrderClient(m.cli.Conn())
	return client.CreateOrder(ctx, in, opts...)
}

func (m *defaultOrder) CreateOrderRollBack(ctx context.Context, in *CreateOrderReq, opts ...grpc.CallOption) (*CreateOrderResp, error) {
	client := pb.NewOrderClient(m.cli.Conn())
	return client.CreateOrderRollBack(ctx, in, opts...)
}
