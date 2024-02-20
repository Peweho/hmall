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

type CartPOArr []CartPO

func (m CartPOArr) Len() int { return len(m) }

func (m CartPOArr) Less(i, j int) bool { return m[i].ItemId < m[j].ItemId }

func (m CartPOArr) Swap(i, j int) { m[i], m[j] = m[j], m[i] }
