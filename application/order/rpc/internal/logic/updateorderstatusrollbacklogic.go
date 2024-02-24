package logic

import (
	"context"

	"hmall/application/order/rpc/internal/svc"
	"hmall/application/order/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateOrderStatusRollBackLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateOrderStatusRollBackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOrderStatusRollBackLogic {
	return &UpdateOrderStatusRollBackLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateOrderStatusRollBackLogic) UpdateOrderStatusRollBack(in *pb.UpdateOrderStatusReq) (*pb.UpdateOrderStatusResp, error) {
	// todo: add your logic here and delete this line

	return &pb.UpdateOrderStatusResp{}, nil
}
