package model

import (
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
