package entity

import (
	"time"

	"gorm.io/gorm"
)

type Favorite struct {
	ID       int       `gorm:"primaryKey"`
	UserId   int       `gorm:"uniqueIndex:idx_favorite"`
	User     User      `gorm:"foreignKey:UserId"`
	CourseId int       `gorm:"uniqueIndex:idx_favorite"`
	Course   Course    `gorm:"foreignKey:CourseId"`
	Date     time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	gorm.Model
}
