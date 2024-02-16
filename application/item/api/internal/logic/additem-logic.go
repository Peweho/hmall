package logic

import (
	"context"
	"hmall/application/item/api/internal/svc"
	"hmall/application/item/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdditemLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdditemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdditemLogic {
	return &AdditemLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AdditemLogic) Additem(req *types.ItemReqAndResp) error {
	// todo: add your logic here and delete this line

	return nil
}
