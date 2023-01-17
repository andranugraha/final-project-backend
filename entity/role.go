package entity

import "gorm.io/gorm"

type Role struct {
	ID       int
	Name     string
	gorm.Model
}