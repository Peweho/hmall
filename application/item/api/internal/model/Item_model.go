package model

import (
	"context"
	"gorm.io/gorm"
	"hmall/application/item/api/internal/types"
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
func (m *ItemModel) DecutStock(ctx context.Context, id int, num int) (*ItemDTO, error) {
	var item ItemDTO
	err := m.db.WithContext(ctx).
		Model(&ItemDTO{}).
		Where("id = ? and stock >= ?", id, num).
		Update("stock", gorm.Expr("stock - ?", num)).Find(&item).Error
	return &item, err
}

// 根据ID删除商品
func (m *ItemModel) DelItemById(ctx context.Context, id int) error {
	return m.db.WithContext(ctx).
		Where("id = ?", id).
		Update("status", types.ItemStatusDeleted).Error
}

// 分页查询
func (m *ItemModel) QueryItemPage(ctx context.Context, page, pageSize int, sortBy, isAsc string) ([]ItemDTO, error) {
	var res []ItemDTO
	err := m.db.WithContext(ctx).
		Omit(types.SelOmit).
		Where("status = ?", types.ItemStatusNormal).
		Limit(pageSize).
		Offset((page - 1) * pageSize).
		Order(sortBy + " " + isAsc).
		Find(&res).Error
	return res, err
}

// 查询商品总数量
func (m *ItemModel) GetItemTotal(ctx context.Context, total *int64) error {
	return m.db.WithContext(ctx).
		Where("status = ?", types.ItemStatusNormal).
		Model(&ItemDTO{}).
		Count(total).Error
}

// 根据id更新商品
func (m *ItemModel) UpdateItemById(ctx context.Context, item ItemDTO) error {
	return m.db.WithContext(ctx).Updates(item).Error
}

// 根据id更新商品状态
func (m *ItemModel) UpdateItemStatusById(ctx context.Context, id int, status int) (*ItemDTO, error) {
	var item ItemDTO
	err := m.db.WithContext(ctx).
		Model(&ItemDTO{}).
		Where("id = ?", id).
		Updates(map[string]any{"status": status}).Error
	if err != nil {
		return nil, err
	}
	err = m.db.WithContext(ctx).
		Omit(types.Field...).
		Where("id = ?", id).
		Find(&item).Error
	return &item, err
}
