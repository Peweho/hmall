package model

import "time"

type OrderPO struct {
	Id           int64
	Total_fee    int
	Payment_type int
	User_id      int64
	Status       int
	CreatedAt    time.Time `gorm:"column:create_time"`
	UpdatedAt    time.Time `gorm:"column:update_time"`
	Pay_time     time.Time
	Consign_time time.Time
	End_time     time.Time
	Close_time   time.Time
	Comment_time time.Time
}

func (m *OrderPO) TableName() string {
	return "order"
}
