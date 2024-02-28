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

// 扣减库存
func (m *ItemModel) DecutStock(ctx context.Context, id string, num int64) (*ItemDTO, error) {
	var res ItemDTO
	err := m.db.WithContext(ctx).
		Model(&ItemDTO{}).
		Where("id = ?", id).
		Update("stock", gorm.Expr("stock - ?", num)).Find(&res).Error
	return &res, err
}

// 恢复库存
func (m *ItemModel) AddStock(ctx context.Context, id string, num int64) (*ItemDTO, error) {
	var res ItemDTO
	err := m.db.WithContext(ctx).
		Model(&ItemDTO{}).
		Where("id = ?", id).
		Update("stock", gorm.Expr("stock + ?", num)).Find(&res).Error
	return &res, err
}
