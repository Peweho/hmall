package logic

import (
	"context"

	"hmall/application/item/api/internal/svc"
	"hmall/application/item/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeductItemsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeductItemsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeductItemsLogic {
	return &DeductItemsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeductItemsLogic) DeductItems(req *types.DeductItemsReq) error {
	// todo: add your logic here and delete this line

	return nil
}
