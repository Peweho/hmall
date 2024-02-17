package mqs

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"hmall/application/item/mq/internal/svc"
	"hmall/application/item/rpc/item"
	"hmall/application/item/rpc/types"
	"log"
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

	key := CacheIds(id)
	//1、删除缓存
	_, err := l.svcCtx.BizRedis.Del(key)
	log.Println("删除缓存: %v", key)
	//2、添加缓存
	_, err = l.svcCtx.ItemRPC.FindItemByIds(l.ctx, &item.FindItemByIdsReq{Ids: []string{id}})
	if err != nil {
		logx.Errorf("ItemRPC.FindItemByIds: %v, error: %v", id, err)
		return err
	}
	return nil
}

func CacheIds(id string) string {
	return fmt.Sprintf("%s#%s", types.CacheItemKey, id)
}
