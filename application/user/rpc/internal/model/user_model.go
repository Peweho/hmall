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

// 根据id查询密码
func (m *UserModel) FindUserPwdById(ctx context.Context, id int64) (string, error) {
	var res string
	err := m.db.WithContext(ctx).Model(&UserDTO{}).Select("password").Where("id = ?", id).Find(&res).Error
	return res, err
}

// 修改金额
func (m *UserModel) UpdateBalance(ctx context.Context, id int64, balance int64) (int, error) {
	res := 0
	err := m.db.WithContext(ctx).
		Model(&UserDTO{}).Select("balance").
		Where("id = ?", id).
		Update("balance", gorm.Expr("balance - ?", balance)).Find(&res).Error
	return res, err
}

// 恢复金额
func (m *UserModel) UpdateBalanceRollBack(ctx context.Context, id int64, balance int64) error {
	return m.db.WithContext(ctx).
		Model(&UserDTO{}).
		Where("id = ?", id).
		Update("balance", gorm.Expr("balance + ?", balance)).Error
}
