package utils

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
	"hmall/application/cart/api/internal/model"
	"hmall/application/cart/api/internal/svc"
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

type KqMsg struct {
	Category string // 1,删缓存；0，加缓存
	Data     *model.CartPO
}

func (l *PusherLogic) UpdateCache(msg *KqMsg) error {
	marshal, err := json.Marshal(msg)
	if err != nil {
		logx.Errorf("json.Marshal: %v, error: %v", msg, err)
		return err
	}
	if err = l.svcCtx.KqPusherClient.Push(string(marshal)); err != nil {
		logx.Errorf("KqPusherClient Push Error , err :%v", err)
		return err
	}

	return nil
}
