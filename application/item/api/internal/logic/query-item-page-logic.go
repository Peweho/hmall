package logic

import (
	"context"

	"hmall/application/item/api/internal/svc"
	"hmall/application/item/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryItemPageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryItemPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryItemPageLogic {
	return &QueryItemPageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryItemPageLogic) QueryItemPage(req *types.QueryItemPageReq) (resp *types.QueryItemPageResp, err error) {
	// todo: add your logic here and delete this line

	return
}
