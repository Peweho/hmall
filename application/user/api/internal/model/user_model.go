package model

import (
	"context"
	"gorm.io/gorm"
)

type UserModel struct {
	db *gorm.DB
}

func NewEmployeeGormModel(db *gorm.DB) *UserModel {
	return &UserModel{
		db: db,
	}
}

// 根据username查询信息
func (m *UserModel) FindUserByName(ctx context.Context, userName string) (UserDTO, error) {
	var res UserDTO
	err := m.db.WithContext(ctx).Where("username = ?", userName).Find(&res).Error
	return res, err
}

// 根据id查询信息
func (m *UserModel) FindUserById(ctx context.Context, id int) (UserDTO, error) {
	var res UserDTO
	err := m.db.WithContext(ctx).Where("id = ?", id).Find(&res).Error
	return res, err
}

// 修改金额
func (m *UserModel) UpdateBalance(ctx context.Context, id int, balance int) error {
	return m.db.WithContext(ctx).
		Model(&UserDTO{}).
		Where("id = ?", id).
		Updates(map[string]any{"balance": balance}).Error
}
