package entity

import (
	"time"

	"gorm.io/gorm"
)

type Invoice struct {
	ID          int
	UserId      int
	User        User `gorm:"foreignKey:UserId"`
	Status      string
	Total       float64
	PaymentDate time.Time
	VoucherId   int
	Voucher     Voucher `gorm:"foreignKey:VoucherId"`
	gorm.Model
}
