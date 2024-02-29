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

	//2、删除缓存
	uid, err := util.GetUsr(l.ctx, types.JwtKey)
	if err != nil {
		logx.Errorf("util.GetUsr, error: %v", err)
		panic(err)
	}
	pusher := utils.NewPusherLogic(l.ctx, l.svcCtx)
	for _, id := range req.Ids {
		key := util.CacheKey(types.CacheCartKey, strconv.Itoa(uid))
		if _, err := l.svcCtx.BizRedis.Hdel(key, id); err != nil {
			//调用mq服务删除缓存
			intId, err1 := strconv.Atoi(id)
			if err1 != nil {
				logx.Errorf("strconv.Atoi: %v, error : %v", id, err1)
				return err1
			}
			msg := &utils.KqMsg{
				Category: types.MSgDelCache,
				Data:     &model.CartPO{Id: intId, UserId: uid},
			}
			if err2 := pusher.UpdateCache(msg); err2 != nil {
				logx.Errorf("pusher.UpdateCache: %v, error: %v", *msg, err2)
				return err2
			}
		}
	}
	return xcode.New(types.OK, "")
}
