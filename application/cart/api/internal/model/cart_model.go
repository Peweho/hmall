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

// 查询购物车
func (m *CartModel) QueryCatr(ctx context.Context, usr int) ([]CartPO, error) {
	var res []CartPO
	err := m.db.WithContext(ctx).Where("user_id = ?", usr).Find(&res).Error
	return res, err
}
