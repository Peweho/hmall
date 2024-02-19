package model

import (
	"context"
	"gorm.io/gorm"
)

type AddressModel struct {
	db *gorm.DB
}

func NewAddressModel(db *gorm.DB) *AddressModel {
	return &AddressModel{
		db: db,
	}
}

// 查询当前用户地址列表
func (m *AddressModel) QueryAddresses(ctx context.Context, usr int) ([]AddressPO, error) {
	var res []AddressPO
	err := m.db.WithContext(ctx).Where("user_id = ?", usr).Find(&res).Error
	return res, err
}
