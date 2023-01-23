package entity

import "gorm.io/gorm"

type Voucher struct {
	ID         int     `json:"id" gorm:"primaryKey"`
	Name       string  `json:"name"`
	Amount     float64 `json:"amount"`
	Code       string  `json:"code"`
	gorm.Model `json:"-"`
}
