package util

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"hmall/application/item/api/internal/svc"
)

type PusherLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPusherLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PusherLogic {
	return &PusherLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PusherLogic) Pusher(id string) error {

	//......业务逻辑....

	if err := l.svcCtx.KqPusherClientUpdateCache.Push(id); err != nil {
		logx.Errorf("KqPusherClient Push Error , err :%v", err)
		return err
	}

	return nil
}
