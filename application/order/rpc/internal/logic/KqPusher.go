package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"hmall/application/order/rpc/internal/svc"
	"strconv"
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

func (l *PusherLogic) UpdateStatus(id int) error {
	if err := l.svcCtx.KqPusherClient.Push(strconv.Itoa(id)); err != nil {
		logx.Errorf("KqPusherClient Push Error , err :%v", err)
		return err
	}
	return nil
}
