package logic

import (
	"context"
	"hmall/application/order/rpc/order"

	"hmall/application/order/api/internal/svc"
	"hmall/application/order/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindOrderByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindOrderByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindOrderByIdLogic {
	return &FindOrderByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindOrderByIdLogic) FindOrderById(req *types.FindOrderByIdReq) (resp *types.FindOrderByIdResp, err error) {
	OrderResp, err := l.svcCtx.OrderRPC.FindOrderById(l.ctx, &order.FindOrderByIdReq{Id: int64(req.Id)})
	if err != nil {
		logx.Errorf(": %v, error: %v", req.Id, err)
		return nil, err
	}

	return &types.FindOrderByIdResp{
		FindOrderByIdVO: types.FindOrderByIdVO{
			Id:          int(OrderResp.Id),
			PayTime:     OrderResp.PayTime,
			PaymentType: int(OrderResp.PaymentType),
			Status:      int(OrderResp.Status),
			TotalFee:    int(OrderResp.Status),
			UserId:      int(OrderResp.UserId),
			CloseTime:   OrderResp.CloseTime,
			CommentTime: OrderResp.CommentTime,
			ConsignTime: OrderResp.ConsignTime,
			CreateTime:  OrderResp.CreateTime,
			EndTime:     OrderResp.EndTime,
		},
	}, nil
}
