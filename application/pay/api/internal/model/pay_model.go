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

func (m *PayModel) CreatePayOrder(ctx context.Context, pay *PayPO) error {
	return m.db.WithContext(ctx).Create(pay).Error
}
