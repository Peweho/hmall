package logic

import (
	"context"
	"encoding/json"
	"hmall/application/cart/api/internal/model"
	"hmall/application/cart/api/internal/svc"
	"hmall/application/cart/api/internal/types"
	"hmall/application/cart/api/internal/utils"
	"hmall/pkg/util"
	"hmall/pkg/xcode"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddCartLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddCartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCartLogic {
	return &AddCartLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddCartLogic) AddCart(req *types.AddCartReq) error {
	// 1、构建PO
	usr, err := util.GetUsr(l.ctx, types.JwtKey)
	if err != nil {
		logx.Errorf("util.GetUsr, error: %v", err)
		return xcode.New(types.Unauthorized, "")
	}
	cart := &model.CartPO{
		Price:  req.Price,
		Image:  req.Image,
		ItemId: req.ItemId,
		Spec:   req.Spec,
		UserId: usr,
		Num:    req.Num,
		Name:   req.Name,
	}

	//2、加入数据库
	err = l.svcCtx.CartModel.AddCatr(l.ctx, cart)
	if err != nil {
		logx.Errorf("CartModel.AddCatr: %v, error: %v", cart, err)
		return err
	}

	//3、缓存存在，写缓存
	key := util.CacheKey(types.CacheCartKey, strconv.Itoa(cart.UserId))
	exists, _ := l.svcCtx.BizRedis.Exists(key)
	if !exists {
		return xcode.New(types.OK, "")
	}

	marshal, err := json.Marshal(cart)
	if err != nil {
		logx.Errorf("json.Marshal,error: %v", err)
		return err
	}

	if err = l.svcCtx.BizRedis.Hset(key, strconv.Itoa(cart.Id), string(marshal)); err != nil {
		logx.Errorf("BizRedis.Hset: %v, error: %v", string(marshal), err)
		//调用mq服务增加缓存
		msg := &utils.KqMsg{
			Category: types.MSgAddCache,
			Data:     cart,
		}
		pusher := utils.NewPusherLogic(l.ctx, l.svcCtx)
		if err1 := pusher.UpdateCache(msg); err1 != nil {
			logx.Errorf("pusher.UpdateCache: %v, error: %v", *msg, err1)
			return err1
		}
		return err
	}
	_ = l.svcCtx.BizRedis.Expire(key, types.CacheCartTime)

	//4、返回
	return xcode.New(types.OK, "")
}
