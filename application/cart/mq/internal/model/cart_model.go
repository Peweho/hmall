package model

import (
	"context"
	"gorm.io/gorm"
)

type CartModel struct {
	db *gorm.DB
}

func NewCartModel(db *gorm.DB) *CartModel {
	return &CartModel{
		db: db,
	}
}

// 根据Id查询购物车
func (m *CartModel) FindCatrById(ctx context.Context, id string) (*CartPO, error) {
	res := &CartPO{}
	return res, m.db.WithContext(ctx).Where("id = ?", id).Find(res).Error
}
