package util

import (
	"context"
	"encoding/json"
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

type KqCacheMsg struct {
	Code   int //补偿类型：0、全字段添加；1、更新部分字段；2，更新库存；3、更新状态
	Field  string
	Stock  string
	Status string
	Key    string
}

func (l *PusherLogic) Pusher(msg *KqCacheMsg) error {

	marshal, err := json.Marshal(msg)
	if err != nil {
		logx.Errorf("KqPusherClient Push Error , err :%v", err)
		return err
	}

	if err := l.svcCtx.KqPusherClientUpdateCache.Push(string(marshal)); err != nil {
		logx.Errorf("KqPusherClient Push: %v , error :%v", string(marshal), err)
		return err
	}

	return nil
}
