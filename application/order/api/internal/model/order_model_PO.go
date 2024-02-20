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

type OrderDetailPO struct {
	Id        int64
	Order_id  int64
	Item_id   int64
	Num       int
	Name      string
	Spec      string
	Price     int
	Image     string
	CreatedAt time.Time `gorm:"column:create_time"`
	UpdatedAt time.Time `gorm:"column:update_time"`
}

func (m *OrderDetailPO) TableName() string {
	return "order_detail"
}

type OrderLogisticsPO struct {
	Order_id          int64
	Logistics_number  string
	Logistics_company string
	Contact           string
	Mobile            string
	Province          string
	City              string
	Town              string
	Street            string
	CreatedAt         time.Time `gorm:"column:create_time"`
	UpdatedAt         time.Time `gorm:"column:update_time"`
}

func (m *OrderLogisticsPO) TableName() string {
	return "order_logistics"
}
