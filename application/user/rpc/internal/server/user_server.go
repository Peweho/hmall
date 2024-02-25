// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package server

import (
	"context"

	"hmall/application/user/rpc/internal/logic"
	"hmall/application/user/rpc/internal/svc"
	"hmall/application/user/rpc/pb"
)

type UserServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedUserServer
}

func NewUserServer(svcCtx *svc.ServiceContext) *UserServer {
	return &UserServer{
		svcCtx: svcCtx,
	}
}

func (s *UserServer) DecutMoney(ctx context.Context, in *pb.DecutMoneyReq) (*pb.DecutMoneyResp, error) {
	l := logic.NewDecutMoneyLogic(ctx, s.svcCtx)
	return l.DecutMoney(in)
}

func (s *UserServer) DecutMoneyRollBack(ctx context.Context, in *pb.DecutMoneyReq) (*pb.DecutMoneyResp, error) {
	l := logic.NewDecutMoneyRollBackLogic(ctx, s.svcCtx)
	return l.DecutMoneyRollBack(in)
}
