package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"hmall/application/order/api/internal/svc"
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
