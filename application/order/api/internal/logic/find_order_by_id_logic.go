package logic

import (
	"context"

	"hmall/application/order/api/internal/svc"
	"hmall/application/order/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindOrderByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindOrderByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindOrderByIdLogic {
	return &FindOrderByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindOrderByIdLogic) FindOrderById(req *types.FindOrderByIdReq) (resp *types.FindOrderByIdResp, err error) {
	// todo: add your logic here and delete this line

	return
}
