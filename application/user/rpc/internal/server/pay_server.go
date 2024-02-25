// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package server

import (
	"context"

	"hmall/application/user/rpc/internal/logic"
	"hmall/application/user/rpc/internal/svc"
	"hmall/application/user/rpc/pb"
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

func (s *PayServer) DecutMoney(ctx context.Context, in *pb.DecutMoneyReq) (*pb.DecutMoneyResp, error) {
	l := logic.NewDecutMoneyLogic(ctx, s.svcCtx)
	return l.DecutMoney(in)
}

func (s *PayServer) DecutMoneyRollBack(ctx context.Context, in *pb.DecutMoneyReq) (*pb.DecutMoneyResp, error) {
	l := logic.NewDecutMoneyRollBackLogic(ctx, s.svcCtx)
	return l.DecutMoneyRollBack(in)
}