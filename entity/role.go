package entity

import "gorm.io/gorm"

type Role struct {
	ID         int    `json:"id" gorm:"primaryKey"`
	Name       string `json:"name"`
	gorm.Model `json:"-"`
}
