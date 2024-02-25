package model

import (
	"context"
	"gorm.io/gorm"
)

type PayModel struct {
	db *gorm.DB
}

func NewPayModel(db *gorm.DB) *PayModel {
	return &PayModel{
		db: db,
	}
}

// 更新支付单
func (m *PayModel) UpdatePayOrder(ctx context.Context, pay map[string]any) error {
	return m.db.WithContext(ctx).Model(&PayPO{}).Updates(pay).Error
}
