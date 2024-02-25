package logic

import (
	"context"
	"hmall/application/user/rpc/internal/svc"
	"hmall/application/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DecutMoneyRollBackLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDecutMoneyRollBackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DecutMoneyRollBackLogic {
	return &DecutMoneyRollBackLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DecutMoneyRollBackLogic) DecutMoneyRollBack(in *pb.DecutMoneyReq) (*pb.DecutMoneyResp, error) {
	if err := l.svcCtx.UserModel.UpdateBalanceRollBack(l.ctx, in.Uid, in.Amount); err != nil {
		logx.Errorf("UserModel.UpdateBalanceRollBack: %v, error: %v", in.Uid, err)
		return nil, err
	}

	return &pb.DecutMoneyResp{}, nil
}
