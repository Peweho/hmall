package logic

import (
	"context"
	"hmall/application/order/rpc/internal/utils"

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
	if err := l.svcCtx.OrderModel.UpdateOrderStatusById(l.ctx, in.Id, utils.NotPayment); err != nil {
		logx.Errorf("OrderModel.UpdateOrderStatusById: %v , error: %v", in.Id, err)
		return nil, err
	}

	return &pb.UpdateOrderStatusResp{}, nil
}
