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

// 根据id批量删除
func (m *CartModel) DelCartsByIds(ctx context.Context, ids []string) error {
	return m.db.WithContext(ctx).Where("id in ?", ids).Delete(&CartPO{}).Error
}
