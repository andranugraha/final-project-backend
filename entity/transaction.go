package entity

import "gorm.io/gorm"

type Transaction struct {
	ID        int
	InvoiceId int
	Invoice   Invoice `gorm:"foreignKey:InvoiceId"`
	CourseId  int
	Course    Course `gorm:"foreignKey:CourseId"`
	Price     float64
	gorm.Model
}
