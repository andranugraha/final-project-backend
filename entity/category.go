package entity

import "gorm.io/gorm"

type Category struct {
	ID   int
	Name string
	gorm.Model
}
