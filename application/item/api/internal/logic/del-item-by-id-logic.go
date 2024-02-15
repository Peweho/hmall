package logic

import (
	"context"

	"hmall/application/item/api/internal/svc"
	"hmall/application/item/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelItemByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelItemByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelItemByIdLogic {
	return &DelItemByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelItemByIdLogic) DelItemById(req *types.DeductItemsReq) error {
	// todo: add your logic here and delete this line

	return nil
}
