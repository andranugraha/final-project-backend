package entity

import "gorm.io/gorm"

type Voucher struct {
	ID          int
	Name        string
	Amount      float64
	VoucherCode string
	gorm.Model
}
