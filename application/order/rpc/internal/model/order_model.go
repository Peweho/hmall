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

func (m *OrderModel) FindOrderById(ctx context.Context, id int64) (OrderPO, error) {
	var res OrderPO
	err := m.db.WithContext(ctx).Where("id = ?", id).Find(&res).Error
	return res, err
}
