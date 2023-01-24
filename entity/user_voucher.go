package entity

import (
	"time"

	"gorm.io/gorm"
)

type UserVoucher struct {
	ID         int       `json:"id" gorm:"primaryKey"`
	UserId     int       `json:"user_id"`
	User       *User     `json:"user,omitempty" gorm:"foreignKey:UserId"`
	VoucherId  int       `json:"voucher_id"`
	Voucher    Voucher   `json:"voucher" gorm:"foreignKey:VoucherId"`
	ExpiryDate time.Time `json:"expiry_date"`
	IsConsumed bool      `json:"is_consumed"`
	gorm.Model `json:"-"`
}
