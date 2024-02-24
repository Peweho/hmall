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
	// todo: add your logic here and delete this line

	return &pb.DecutMoneyResp{}, nil
}
