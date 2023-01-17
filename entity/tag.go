package entity

import "gorm.io/gorm"

type Tag struct {
	ID   int
	Name string
	gorm.Model
}
