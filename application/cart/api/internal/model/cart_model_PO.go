package model

type CartPO struct {
	Id        int
	Image     string
	userId    int `gorm:"column:user_id"`
	ItemId    int `gorm:"column:item_id"`
	Name      string
	Num       int
	Price     int
	Spec      string
	CreatedAt string `gorm:"column:create_time"`
	UpdatedAt string `gorm:"column:update_time"`
}

func (m *CartPO) TableName() string {
	return "cart"
}
