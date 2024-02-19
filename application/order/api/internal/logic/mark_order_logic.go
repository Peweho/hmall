package logic

import (
	"context"

	"hmall/application/order/api/internal/svc"
	"hmall/application/order/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MarkOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMarkOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MarkOrderLogic {
	return &MarkOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MarkOrderLogic) MarkOrder(req *types.MarkOrderReq) error {
	// todo: add your logic here and delete this line

	return nil
}
