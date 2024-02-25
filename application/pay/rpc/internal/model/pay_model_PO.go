package model

import "time"

type PayPO struct {
	Id               int64
	Biz_order_no     int64
	Pay_order_no     int64
	Biz_user_id      int64
	Pay_channel_code string
	Amount           int
	Pay_type         int
	Status           int
	Expand_json      string
	Result_code      string
	Result_msg       string
	Pay_success_time *time.Time
	Pay_over_time    time.Time
	Qr_code_url      string
	CreatedAt        time.Time `gorm:"column:create_time"`
	UpdatedAt        time.Time `gorm:"column:update_time"`
	Creater          int64
	Updater          int64
	Is_delete        int
}

func (m *PayPO) TableName() string {
	return "pay_order"
}
