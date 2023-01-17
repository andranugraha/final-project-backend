package entity

import "gorm.io/gorm"

type UserCourse struct {
	ID       int
	UserId   int
	User     User `gorm:"foreignKey:UserId"`
	CourseId int
	Course   Course `gorm:"foreignKey:CourseId"`
	Status   string
	gorm.Model
}
