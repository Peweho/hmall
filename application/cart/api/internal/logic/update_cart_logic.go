package logic

import (
	"context"

	"hmall/application/cart/api/internal/svc"
	"hmall/application/cart/api/internal/types"

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
	// todo: add your logic here and delete this line

	return nil
}
