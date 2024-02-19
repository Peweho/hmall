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

// 根据ID查询地址
func (m *AddressModel) QueryAddressFindById(ctx context.Context, id int64) (AddressPO, error) {
	var res AddressPO
	err := m.db.WithContext(ctx).Where("id = ?", id).Find(&res).Error
	return res, err
}
