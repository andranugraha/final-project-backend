package entity

import (
	"time"

	"gorm.io/gorm"
)

type Favorite struct {
	ID       int
	UserId   int
	User     User `gorm:"foreignKey:UserId"`
	CourseId int
	Course   Course `gorm:"foreignKey:CourseId"`
	Date     time.Time
	gorm.Model
}
