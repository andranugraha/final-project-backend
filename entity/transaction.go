package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Transaction struct {
	ID         int       `json:"id" gorm:"primaryKey"`
	InvoiceId  uuid.UUID `json:"invoiceId"`
	Invoice    *Invoice  `json:"invoice,omitempty" gorm:"foreignKey:InvoiceId"`
	CourseId   int       `json:"courseId"`
	Course     *Course   `json:"course,omitempty" gorm:"foreignKey:CourseId"`
	Price      float64   `json:"price"`
	gorm.Model `json:"-"`
}

func (t *Transaction) FromCart(c *Cart) {
	t.CourseId = c.CourseId
	t.Price = c.Course.Price
}
