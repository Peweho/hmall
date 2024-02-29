package logic

import (
	"context"
	"hmall/application/cart/api/internal/model"
	"hmall/application/cart/api/internal/svc"
	"hmall/application/cart/api/internal/types"
	"hmall/application/cart/api/internal/utils"
	"hmall/pkg/util"
	"hmall/pkg/xcode"
	"strconv"

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

	//2、删除缓存，如果失败调用mq服务删除缓存
	uid, err := util.GetUsr(l.ctx, types.JwtKey)
	if err != nil {
		logx.Errorf("util.GetUsr, error: %v", err)
		return err
	}

	key := util.CacheKey(types.CacheCartKey, strconv.Itoa(uid))
	if _, err = l.svcCtx.BizRedis.Hdel(key, strconv.Itoa(req.Id)); err != nil {
		logx.Errorf("BizRedis.Zadd: %v, error: %v", key, err)
		//调用mq服务删除缓存
		msg := &utils.KqMsg{
			Category: types.MSgDelCache,
			Data:     &model.CartPO{Id: req.Id, UserId: uid},
		}
		pusher := utils.NewPusherLogic(l.ctx, l.svcCtx)
		if err1 := pusher.UpdateCache(msg); err1 != nil {
			logx.Errorf("pusher.UpdateCache: %v, error: %v", *msg, err1)
			return err1
		}
	}
	return xcode.New(types.OK, "")
}
