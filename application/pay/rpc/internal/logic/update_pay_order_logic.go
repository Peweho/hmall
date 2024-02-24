package logic

import (
	"context"

	"hmall/application/pay/rpc/internal/svc"
	"hmall/application/pay/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePayOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdatePayOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePayOrderLogic {
	return &UpdatePayOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdatePayOrderLogic) UpdatePayOrder(in *pb.UpdatePayOrderReq) (*pb.UpdatePayOrderResp, error) {
	// todo: add your logic here and delete this line

	return &pb.UpdatePayOrderResp{}, nil
}
