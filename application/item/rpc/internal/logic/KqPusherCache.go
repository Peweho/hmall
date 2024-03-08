package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"hmall/application/item/rpc/internal/svc"
)

type PusherCacheLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPusherCacheLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PusherCacheLogic {
	return &PusherCacheLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PusherCacheLogic) PusherCache(id string) error {
	if err := l.svcCtx.KqPusherCache.Push(id); err != nil {
		logx.Errorf("KqPusherCache.Push: %v, error: %v", id, err)
		return err
	}
	return nil
}
