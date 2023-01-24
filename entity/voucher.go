package entity

import "gorm.io/gorm"

type Voucher struct {
	ID         int     `json:"id" gorm:"primaryKey"`
	Name       string  `json:"name"`
	Amount     float64 `json:"amount"`
	Code       string  `json:"code"`
	MinAmount  float64 `json:"minAmount"`
	gorm.Model `json:"-"`
}
