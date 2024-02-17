package logic

import (
	"context"

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
	// todo: add your logic here and delete this line

	return nil
}
