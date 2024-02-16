package model

import "time"

type ItemDTO struct {
	Brand        string
	Category     string
	CommentCount int64
	Id           int64
	Image        string
	IsAD         bool
	Name         string
	Price        int64
	Sold         int64
	Spec         string
	Status       int64
	Stock        int64
	CreatedAt    *time.Time `goem:"create_time"`
	UpdateTime   *time.Time `goem:"update_time"`
	Creater      int64
	Updater      int64
}

func (m *ItemDTO) TableName() string {
	return "item"
}
