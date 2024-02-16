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

// 扣减库存
func (m *ItemModel) DecutStock(ctx context.Context, id int, num int) error {
	return m.db.WithContext(ctx).
		Model(&ItemDTO{}).
		Where("id = ?", id).
		Update("stock", gorm.Expr("stock - ?", num)).Error
}

// 根据ID删除商品
func (m *ItemModel) DelItemById(ctx context.Context, id int) error {
	return m.db.WithContext(ctx).
		Where("id = ?", id).
		Delete(&ItemDTO{}).Error
}

// 分页查询
func (m *ItemModel) QueryItemPage(ctx context.Context, page, pageSize int, sortBy, isAsc string) ([]ItemDTO, error) {
	var res []ItemDTO
	err := m.db.WithContext(ctx).
		Limit(pageSize).
		Offset((page - 1) * pageSize).
		Order(sortBy + " " + isAsc).
		Find(&res).Error
	return res, err
}

// 根据id更新商品
func (m *ItemModel) UpdateItemById(ctx context.Context, item ItemDTO) error {
	return m.db.WithContext(ctx).Updates(item).Error
}

// 根据id更新商品状态
func (m *ItemModel) UpdateItemStatusById(ctx context.Context, id int, status int) error {
	return m.db.WithContext(ctx).
		Model(&ItemDTO{}).
		Where("id = ?", id).
		Updates(map[string]any{"status": status}).Error
}
