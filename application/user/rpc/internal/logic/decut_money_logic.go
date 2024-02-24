package logic

import (
	"context"

	"hmall/application/user/rpc/internal/svc"
	"hmall/application/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DecutMoneyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDecutMoneyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DecutMoneyLogic {
	return &DecutMoneyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DecutMoneyLogic) DecutMoney(in *pb.DecutMoneyReq) (*pb.DecutMoneyResp, error) {
	// todo: add your logic here and delete this line

	return &pb.DecutMoneyResp{}, nil
}
