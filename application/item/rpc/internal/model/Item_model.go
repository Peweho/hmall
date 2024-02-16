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

func (m *ItemModel) FindItemByIds(ctx context.Context, ids []string) ([]ItemDTO, error) {
	var items []ItemDTO
	err := m.db.WithContext(ctx).
		Where("id in ?", ids).
		Find(&items).Error
	return items, err
}
