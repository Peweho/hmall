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
	FindOrderByIdReq  = pb.FindOrderByIdReq
	FindOrderByIdResp = pb.FindOrderByIdResp

	Order interface {
		FindOrderById(ctx context.Context, in *FindOrderByIdReq, opts ...grpc.CallOption) (*FindOrderByIdResp, error)
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
