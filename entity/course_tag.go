package entity

import "gorm.io/gorm"

type CourseTag struct {
	ID       int
	CourseId int
	Course   Course `gorm:"foreignKey:CourseId"`
	TagId    int
	Tag      Tag `gorm:"foreignKey:TagId"`
	gorm.Model
}
