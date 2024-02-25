package logic

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"hmall/application/pay/rpc/internal/svc"
	"hmall/application/pay/rpc/internal/utils"
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
	update := map[string]any{
		"status":           utils.NotPay,
		"pay_success_time": nil,
		"pay_type":         nil,
	}
	if err := l.svcCtx.PayModel.UpdatePayOrder(l.ctx, update); err != nil {
		logx.Errorf("PayModel.UpdatePayOrder: %v,error : %v", update, err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.UpdatePayOrderResp{}, nil
}
