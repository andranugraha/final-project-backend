package entity

import (
	"time"

	"gorm.io/gorm"
)

type UserVoucher struct {
	ID         int
	UserId     int
	User       User `gorm:"foreignKey:UserId"`
	VoucherId  int
	Voucher    Voucher `gorm:"foreignKey:VoucherId"`
	ExpiryDate time.Time
	IsConsumed bool
	gorm.Model
}
