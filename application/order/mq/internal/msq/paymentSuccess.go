package mqs

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"hmall/application/order/mq/internal/svc"
)

type PaymentSuccess struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPaymentSuccess(ctx context.Context, svcCtx *svc.ServiceContext) *PaymentSuccess {
	return &PaymentSuccess{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 更新缓存
func (l *PaymentSuccess) Consume(_, id string) error {
	if err := l.svcCtx.OrderModel.UpdateOrderStatusById(l.ctx, id); err != nil {
		logx.Errorf("OrderModel.UpdateOrderStatusById: %v, error: %v", id, err)
		return err
	}
	return nil
}
