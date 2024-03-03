package model

import "time"

type ItemDTO struct {
	Brand        string
	Category     string
	CommentCount int64
	Id           int64
	Image        string
	IsAD         bool `gorm:"column:isAD"`
	Name         string
	Price        int64
	Sold         int64
	Spec         string
	Status       int64
	Stock        int64
	CreatedAt    *time.Time `gorm:"column:create_time"`
	UpdateTime   *time.Time `gorm:"column:update_time"`
	Creater      int64
	Updater      int64
}

func (m *ItemDTO) TableName() string {
	return "item"
}
