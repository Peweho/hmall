package model

import (
	"context"
	"gorm.io/gorm"
)

type OrderModel struct {
	db *gorm.DB
}

func NewOrderModel(db *gorm.DB) *OrderModel {
	return &OrderModel{
		db: db,
	}
}

func (m *OrderModel) AddOrder(ctx context.Context, order *OrderPO) error {
	return m.db.WithContext(ctx).Omit("Pay_time", "Consign_time", "End_time", "Close_time", "Comment_time").Create(order).Error
}

func (m *OrderModel) AddOrderDetail(ctx context.Context, orderDetail []map[string]any) error {
	return m.db.WithContext(ctx).Model(&OrderDetailPO{}).Create(orderDetail).Error
}

func (m *OrderModel) AddOrderLogistics(ctx context.Context, orderLogistics map[string]any) error {
	return m.db.WithContext(ctx).Model(&OrderLogisticsPO{}).Create(orderLogistics).Error
}
