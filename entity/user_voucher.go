package entity

import (
	"time"

	"gorm.io/gorm"
)

type UserVoucher struct {
	ID         int       `json:"id" gorm:"primaryKey"`
	UserId     int       `json:"user_id"`
	User       *User     `json:"user,omitempty" gorm:"foreignKey:UserId"`
	VoucherId  int       `json:"voucherId"`
	Voucher    Voucher   `json:"voucher" gorm:"foreignKey:VoucherId"`
	ExpiryDate time.Time `json:"expiryDate"`
	IsConsumed bool      `json:"isConsumed"`
	gorm.Model `json:"-"`
}
