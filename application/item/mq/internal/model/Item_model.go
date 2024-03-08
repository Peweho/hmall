package model

import (
	"context"
	"gorm.io/gorm"
	"hmall/application/item/rpc/types"
)

type ItemModel struct {
	db *gorm.DB
}

func NewItemModel(db *gorm.DB) *ItemModel {
	return &ItemModel{
		db: db,
	}
}

func (m *ItemModel) FindItemById(ctx context.Context, id string) (ItemDTO, error) {
	var item ItemDTO
	err := m.db.WithContext(ctx).
		Omit(types.SelOmit).
		Where("id = ? and status = ?", id, types.ItemStatusNormal).
		First(&item).Error
	return item, err
}
