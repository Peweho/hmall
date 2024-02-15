package logic

import (
	"context"

	"hmall/application/item/api/internal/svc"
	"hmall/application/item/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindItemsByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindItemsByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindItemsByIdLogic {
	return &FindItemsByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindItemsByIdLogic) FindItemsById(req *types.FindItemsByIdReq) (resp *types.FindItemsByIdResp, err error) {
	// todo: add your logic here and delete this line

	return
}
