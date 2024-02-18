package logic

import (
	"context"
	"hmall/application/cart/api/internal/model"
	"hmall/application/cart/api/internal/utils"
	"hmall/pkg/util"
	"hmall/pkg/xcode"

	"hmall/application/cart/api/internal/svc"
	"hmall/application/cart/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelCartItemLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelCartItemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelCartItemLogic {
	return &DelCartItemLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelCartItemLogic) DelCartItem(req *types.DelCartItemReq) error {
	//1、删除数据库
	if err := l.svcCtx.CartModel.DelCartById(l.ctx, req.Id); err != nil {
		logx.Errorf("CartModel.DelCartById: %V, error: %v", req.Id, err)
		return err
	}
	//2、调用mq服务删除缓存
	usr, err := util.GetUsr(l.ctx, types.JwtKey)
	if err != nil {
		logx.Errorf("util.GetUsr, error: %v", err)
		return err
	}
	msg := &utils.KqMsg{
		Category: types.MSgDelCache,
		Data:     &model.CartPO{Id: req.Id, UserId: usr},
	}
	pusher := utils.NewPusherLogic(l.ctx, l.svcCtx)
	if err = pusher.UpdateCache(msg); err != nil {
		logx.Errorf("pusher.UpdateCache: %v, error: %v", *msg, err)
		return err
	}
	return xcode.New(types.OK, "")
}
