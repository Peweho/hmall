package logic

import (
	"context"

	"hmall/application/pay/rpc/internal/svc"
	"hmall/application/pay/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePayOrderRollBackLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdatePayOrderRollBackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePayOrderRollBackLogic {
	return &UpdatePayOrderRollBackLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdatePayOrderRollBackLogic) UpdatePayOrderRollBack(in *pb.UpdatePayOrderReq) (*pb.UpdatePayOrderResp, error) {
	// todo: add your logic here and delete this line

	return &pb.UpdatePayOrderResp{}, nil
}
