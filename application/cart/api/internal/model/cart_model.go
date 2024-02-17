package model

import "gorm.io/gorm"

type CartModel struct {
	db *gorm.DB
}

func NewCartModel(db *gorm.DB) *CartModel {
	return &CartModel{
		db: db,
	}
}
