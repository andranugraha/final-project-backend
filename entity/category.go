package entity

import "gorm.io/gorm"

type Category struct {
	ID         int    `json:"id" gorm:"primaryKey"`
	Name       string `json:"name" gorm:"unique"`
	gorm.Model `json:"-"`
}
