package model

import (
	"context"
	"gorm.io/gorm"
	"hmall/application/cart/api/internal/types"
)

type CartModel struct {
	db *gorm.DB
}

func NewCartModel(db *gorm.DB) *CartModel {
	return &CartModel{
		db: db,
	}
}

// 添加购物车
func (m *CartModel) AddCatr(ctx context.Context, po *CartPO) error {
	return m.db.WithContext(ctx).Create(po).Error
}

// 查询购物车
func (m *CartModel) QueryCatr(ctx context.Context, usr int) ([]CartPO, error) {
	var res []CartPO
	err := m.db.WithContext(ctx).Where("user_id = ? and is_del = ?", usr, types.NotDeleted).Find(&res).Error
	return res, err
}

// 根据id删除
func (m *CartModel) DelCartById(ctx context.Context, id int) error {
	return m.db.WithContext(ctx).Where("id = ?", id).Delete(&CartPO{}).Error
}

// 根据id批量删除
func (m *CartModel) DelCartsByIds(ctx context.Context, ids []string) error {
	return m.db.WithContext(ctx).Where("id in ?", ids).Delete(&CartPO{}).Error
}

// 更新购物车
func (m *CartModel) UpdateCart(ctx context.Context, po *CartPO) error {
	return m.db.WithContext(ctx).Updates(po).Error
}
