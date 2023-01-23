package entity

import "gorm.io/gorm"

type Cart struct {
	ID         int     `json:"id" gorm:"primaryKey"`
	UserId     int     `json:"userId"`
	User       *User   `json:"user,omitempty" gorm:"foreignKey:UserId"`
	CourseId   int     `json:"courseId"`
	Course     *Course `json:"course,omitempty" gorm:"foreignKey:CourseId"`
	gorm.Model `json:"-"`
}
