package logic

import (
	"context"

	"hmall/application/item/api/internal/svc"
	"hmall/application/item/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateItemStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateItemStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateItemStatusLogic {
	return &UpdateItemStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateItemStatusLogic) UpdateItemStatus(req *types.UpdateItemStatusReq) error {
	// todo: add your logic here and delete this line

	return nil
}
