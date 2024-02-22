package logic

import (
	"context"
	"hmall/application/item/rpc/internal/svc"
	"hmall/application/item/rpc/service"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelStockRollBackLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelStockRollBackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelStockRollBackLogic {
	return &DelStockRollBackLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelStockRollBackLogic) DelStockRollBack(in *service.DelStockReq) (*service.DelStockResp, error) {
	for _, val := range in.Detail {
		err := l.svcCtx.ItemModel.AddStock(l.ctx, val.ItemId, val.Num)
		if err != nil {
			logx.Errorf("ItemModel.DecutStock: %v, error: %v", val, err)
			return nil, err
		}
	}
	return &service.DelStockResp{}, nil
}
