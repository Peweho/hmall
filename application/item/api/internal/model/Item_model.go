package model

import (
	"context"
	"gorm.io/gorm"
)

type ItemModel struct {
	db *gorm.DB
}

func NewItemModel(db *gorm.DB) *ItemModel {
	return &ItemModel{
		db: db,
	}
}

func (m *ItemModel) InserItem(ctx context.Context, item *ItemDTO) error {
	return m.db.WithContext(ctx).Create(item).Error
}
