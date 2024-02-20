package logic

import (
	"context"

	"hmall/application/order/rpc/internal/svc"
	"hmall/application/order/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateOrderStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateOrderStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOrderStatusLogic {
	return &UpdateOrderStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateOrderStatusLogic) UpdateOrderStatus(in *pb.UpdateOrderStatusReq) (*pb.UpdateOrderStatusResp, error) {
	// todo: add your logic here and delete this line

	return &pb.UpdateOrderStatusResp{}, nil
}
