package model

type AddressPO struct {
	City      string
	Contact   string
	Id        int
	UserId    int64 `gorm:"column:user_id"`
	IsDefault int   `gorm:"is_default"`
	Mobile    string
	Notes     string
	Province  string
	Street    string
	Town      string
}

func (m *AddressPO) TableName() string {
	return "address"
}
