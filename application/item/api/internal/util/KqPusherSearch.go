package util

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
	"hmall/application/item/api/internal/model"
	"hmall/application/item/api/internal/svc"
)

type PusherSearchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPusherSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PusherSearchLogic {
	return &PusherSearchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type SearchkKqMsg struct {
	Code int
	Date SearchItemDTO
}

func (l *PusherSearchLogic) PusherSearch(code int, item *model.ItemDTO) error {
	searchItem := &SearchkKqMsg{
		Code: code,
		Date: SearchItemDTO{
			Brand:        item.Brand,
			Category:     item.Category,
			CommentCount: item.CommentCount,
			Image:        item.Image,
			IsAD:         item.IsAD,
			Id:           item.Id,
			Price:        item.Price,
			Sold:         item.Sold,
			Spec:         item.Spec,
			Status:       item.Status,
			Stock:        item.Stock,
			Name:         item.Name,
		},
	}
	marshal, err := json.Marshal(searchItem)
	if err != nil {
		logx.Errorf("json.Marshal: %v, error: %v", searchItem, err)
		return err
	}
	if err = l.svcCtx.KqPusherSearch.Push(string(marshal)); err != nil {
		logx.Errorf("KqPusherSearch.Push: %v, error: %v", string(marshal), err)
		return err
	}
	return nil
}

type SearchItemDTO struct {
	Brand        string `json:"brand"`
	Category     string `json:"category"`
	CommentCount int64  `json:"commentCount"`
	Id           int64  `json:"id"`
	Image        string `json:"image"`
	IsAD         bool   `json:"isAD"`
	Name         string `json:"name"`
	Price        int64  `json:"price"`
	Sold         int64  `json:"sold"`
	Spec         string `json:"spec"`
	Status       int64  `json:"status"`
	Stock        int64  `json:"stock"`
}
