package entity

import "gorm.io/gorm"

type UserVoucher struct {
	ID        int
	UserId    int
	User      User `gorm:"foreignKey:UserId"`
	VoucherId int
	Voucher   Voucher `gorm:"foreignKey:VoucherId"`
	ExpiryDate string
	Status	string
	gorm.Model
}