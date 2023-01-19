package entity

import (
	"time"

	"gorm.io/gorm"
)

type Tag struct {
	ID        int            `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name" gorm:"unique"`
	Courses   []*Course      `json:"courses,omitempty" gorm:"many2many:course_tags;"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
