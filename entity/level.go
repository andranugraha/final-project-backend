package entity

import "gorm.io/gorm"

type Level struct {
	ID              int     `json:"id" gorm:"primaryKey"`
	Name            string  `json:"name"`
	Discount        float64 `json:"discount"`
	AvatarUrl       string  `json:"avatarUrl"`
	MinTransanction int     `json:"minTransanction"`
	Point           int     `json:"point"`
	gorm.Model      `json:"-"`
}
