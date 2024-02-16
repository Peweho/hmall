package logic

import (
	"context"
	"hmall/application/item/api/internal/model"
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
	item := &model.ItemDTO{
		Id:           int64(req.Id),
		Brand:        req.Brand,
		Category:     req.Category,
		CommentCount: int64(req.CommentCount),
		Image:        req.Image,
		IsAD:         req.IsAD,
		Name:         req.Name,
		Price:        int64(req.Price),
		Sold:         int64(req.Sold),
		Spec:         req.Spec,
		Status:       int64(req.Status),
		Stock:        int64(req.Stock),
	}
	err := l.svcCtx.ItemModel.InserItem(l.ctx, item)
	if err != nil {
		logx.Errorf("ItemModel.InserItem: %v, error: %v", *item, err)
		return err
	}

	return nil
}
