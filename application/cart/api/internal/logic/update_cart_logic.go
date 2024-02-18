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

type UpdateCartLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateCartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCartLogic {
	return &UpdateCartLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateCartLogic) UpdateCart(req *types.UpdateCartReq) error {
	// 1、更新数据库
	cartPO := &model.CartPO{
		Id:  req.Id,
		Num: req.Num,
	}
	if err := l.svcCtx.CartModel.UpdateCart(l.ctx, cartPO); err != nil {
		logx.Errorf("pusher.UpdateCache: %v, error: %v", *cartPO, err)
		return err
	}
	// 2、更新缓存
	usr, err := util.GetUsr(l.ctx, types.JwtKey)
	if err != nil {
		logx.Errorf("util.GetUsr, error: %v", err)
		return err
	}
	//补充字段
	cartPO.UserId = usr
	cartPO.ItemId = req.ItemId
	cartPO.Image = req.Image
	cartPO.Spec = req.Spec
	cartPO.Name = req.Name
	cartPO.Price = req.Price

	msg := &utils.KqMsg{
		Category: types.MSgAddCache,
		Data:     cartPO,
	}
	pusher := utils.NewPusherLogic(l.ctx, l.svcCtx)
	if err = pusher.UpdateCache(msg); err != nil {
		logx.Errorf("pusher.UpdateCache: %v, error: %v", *msg, err)
		return err
	}
	return xcode.New(types.OK, "")
}
