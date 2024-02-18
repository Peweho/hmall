package logic

import (
	"context"

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
	//1、
	return nil
}
