package logic

import (
	"context"

	"hmall/application/item/api/internal/svc"
	"hmall/application/item/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindItemByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindItemByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindItemByIdLogic {
	return &FindItemByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindItemByIdLogic) FindItemById(req *types.FindItemByIdReq) (resp *types.ItemReqAndResp, err error) {
	// todo: add your logic here and delete this line

	return
}
