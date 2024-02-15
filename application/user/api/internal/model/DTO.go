package model

type UserDTO struct {
	Id       int
	UserName string `gorm:"column:username"`
	PassWord string `gorm:"column:password"`
	Status   int
	Balance  int
}

func (u *UserDTO) TableName() string {
	return "user"
}
