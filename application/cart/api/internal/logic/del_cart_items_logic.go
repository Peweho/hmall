package logic

import (
	"context"
	"hmall/application/cart/api/internal/model"
	"hmall/application/cart/api/internal/utils"
	"hmall/pkg/util"
	"hmall/pkg/xcode"
	"strconv"

	"hmall/application/cart/api/internal/svc"
	"hmall/application/cart/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelCartItemsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelCartItemsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelCartItemsLogic {
	return &DelCartItemsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelCartItemsLogic) DelCartItems(req *types.DelCartItemsReq) error {
	//1、删除数据库
	if err := l.svcCtx.CartModel.DelCartsByIds(l.ctx, req.Ids); err != nil {
		logx.Errorf("CartModel.DelCartById: %V, error: %v", req.Ids, err)
		return err
	}
	//2、调用mq服务删除缓存
	usr, err := util.GetUsr(l.ctx, types.JwtKey)
	if err != nil {
		logx.Errorf("util.GetUsr, error: %v", err)
		return err
	}
	pusher := utils.NewPusherLogic(l.ctx, l.svcCtx)
	for _, id := range req.Ids {
		intId, err := strconv.Atoi(id)
		if err != nil {
			logx.Errorf("strconv.Atoi: %v, error : %v", id, err)
			return err
		}
		msg := &utils.KqMsg{
			Category: types.MSgDelCache,
			Data:     &model.CartPO{Id: intId, UserId: usr},
		}

		if err = pusher.UpdateCache(msg); err != nil {
			logx.Errorf("pusher.UpdateCache: %v, error: %v", *msg, err)
			return err
		}
	}
	return xcode.New(types.OK, "")
}
