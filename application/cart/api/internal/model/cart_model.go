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

// 添加购物车
func (m *CartModel) AddCatr(ctx context.Context, po *CartPO) error {
	return m.db.WithContext(ctx).Create(po).Error
}
