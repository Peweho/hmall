// Code generated by goctl. DO NOT EDIT.
// Source: pay.proto

package server

import (
	"context"

	"hmall/application/pay/rpc/internal/logic"
	"hmall/application/pay/rpc/internal/svc"
	"hmall/application/pay/rpc/pb"
)

type PayServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedPayServer
}

func NewPayServer(svcCtx *svc.ServiceContext) *PayServer {
	return &PayServer{
		svcCtx: svcCtx,
	}
}

func (s *PayServer) UpdatePayOrder(ctx context.Context, in *pb.UpdatePayOrderReq) (*pb.UpdatePayOrderResp, error) {
	l := logic.NewUpdatePayOrderLogic(ctx, s.svcCtx)
	return l.UpdatePayOrder(in)
}

func (s *PayServer) UpdatePayOrderRollBack(ctx context.Context, in *pb.UpdatePayOrderReq) (*pb.UpdatePayOrderResp, error) {
	l := logic.NewUpdatePayOrderRollBackLogic(ctx, s.svcCtx)
	return l.UpdatePayOrderRollBack(in)
}
