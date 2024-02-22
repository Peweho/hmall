package model

import "time"

type CartPO struct {
	Id        int
	Image     string
	UserId    int `gorm:"column:user_id"`
	ItemId    int `gorm:"column:item_id"`
	Name      string
	Num       int
	Price     int
	Spec      string
	CreatedAt *time.Time `gorm:"column:create_time"`
	UpdatedAt *time.Time `gorm:"column:update_time"`
}

func (m *CartPO) TableName() string {
	return "cart"
}
