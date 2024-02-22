package model

import (
	"context"
	"gorm.io/gorm"
	"hmall/application/cart/rpc/internal/utils"
)

type CartModel struct {
	db *gorm.DB
}

func NewCartModel(db *gorm.DB) *CartModel {
	return &CartModel{
		db: db,
	}
}

// 逻辑删除
func (m *CartModel) DelCartsByUidItemId(ctx context.Context, usr int64, itemIds []string) error {
	return m.db.WithContext(ctx).
		Model(&CartPO{}).
		Where("user_id =  ? and item_id in ?", usr, itemIds).
		Order("create_time desc").
		Limit(len(itemIds)).Updates(map[string]any{"is_del": utils.Deleted}).Error
}

// 逻恢复
func (m *CartModel) SetCartsByUidItemId(ctx context.Context, usr int64, itemIds []string) error {
	return m.db.WithContext(ctx).
		Model(&CartPO{}).
		Where("user_id =  ? and item_id in ?", usr, itemIds).
		Order("create_time desc").
		Limit(len(itemIds)).Updates(map[string]any{"is_del": utils.NotDeleted}).Error
}
