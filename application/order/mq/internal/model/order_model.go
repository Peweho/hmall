package model

import (
	"context"
	"gorm.io/gorm"
	"hmall/application/order/mq/internal/types"
)

type OrderModel struct {
	db *gorm.DB
}

func NewOrderModel(db *gorm.DB) *OrderModel {
	return &OrderModel{
		db: db,
	}
}

func (m *OrderModel) UpdateOrderStatusById(ctx context.Context, id string) error {
	return m.db.WithContext(ctx).Model(&OrderPO{}).Where("id = ?", id).Updates(map[string]any{"status": types.Paied}).Error
}
