package mqs

import (
	"context"
	"fmt"
	"hmall/application/order/mq/internal/svc"
	"hmall/application/order/mq/internal/types"
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
	return nil
}

func CacheIds(id string) string {
	return fmt.Sprintf("%s#%s", types.CacheOrderKey, id)
}
