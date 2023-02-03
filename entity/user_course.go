package entity

import "gorm.io/gorm"

type UserCourse struct {
	ID         int     `json:"id" gorm:"primaryKey"`
	UserId     int     `json:"userId"`
	User       *User   `json:"user,omitempty" gorm:"foreignKey:UserId"`
	CourseId   int     `json:"courseId"`
	Course     *Course `json:"course,omitempty" gorm:"foreignKey:CourseId"`
	Status     string  `json:"status"`
	gorm.Model `json:"-"`
}
