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

func (m *ItemModel) InserItem(ctx context.Context) (*[]ItemDTO, error) {
	res := make([]ItemDTO, 0, 1000)
	err := m.db.WithContext(ctx).Limit(1000).Find(&res).Error
	return &res, err
}

// 查询商品（导入数据使用，生产环境删除）
func (m *ItemModel) FindItem(ctx context.Context, num int, offset int) ([]ItemDTO, error) {
	items := make([]ItemDTO, num)
	err := m.db.WithContext(ctx).Limit(num).Offset(offset).Find(&items).Error
	return items, err
}
