package logic

import (
	"context"
	"hmall/application/cart/api/internal/model"
	"hmall/application/cart/api/internal/svc"
	"hmall/application/cart/api/internal/types"
	"hmall/application/cart/api/internal/utils"
	"hmall/pkg/util"
	"hmall/pkg/xcode"

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
	//3、调用mq服务写缓存
	msg := &utils.KqMsg{
		Category: types.MSgAddCache,
		Data:     cart,
	}
	pusher := utils.NewPusherLogic(l.ctx, l.svcCtx)
	err = pusher.UpdateCache(msg)
	if err != nil {
		logx.Errorf("pusher.UpdateCache: %v, error: %v", *msg, err)
		return err
	}

	//4、返回
	return xcode.New(types.OK, "")
}
