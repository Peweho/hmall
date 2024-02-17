package logic

import (
	"context"

	"hmall/application/cart/api/internal/svc"
	"hmall/application/cart/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryCartLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryCartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryCartLogic {
	return &QueryCartLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryCartLogic) QueryCart() (resp *types.QueryCartResp, err error) {
	// todo: add your logic here and delete this line

	return
}
