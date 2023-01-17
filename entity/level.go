package entity

import "gorm.io/gorm"

type Level struct {
	ID       int
	Name     string
	Discount float32
	gorm.Model
}