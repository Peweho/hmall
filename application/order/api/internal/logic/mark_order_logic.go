package logic

import (
	"context"
	"strconv"

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
	id, err := strconv.Atoi(req.OrderId)
	if err != nil {
		logx.Errorf("strconv.Atoi: %v, error: %v", id, err)
		return err
	}
	pusherLogic := NewPusherLogic(l.ctx, l.svcCtx)
	if err = pusherLogic.UpdateStatus(id); err != nil {
		return err
	}

	return nil
}
