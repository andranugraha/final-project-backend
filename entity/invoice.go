package entity

import (
	"final-project-backend/utils/constant"
	"time"

	"gorm.io/gorm"
)

type Invoice struct {
	ID            int            `json:"id" gorm:"primaryKey"`
	UserId        int            `json:"userId"`
	User          *User          `json:"user,omitempty" gorm:"foreignKey:UserId"`
	Status        string         `json:"status"`
	PaymentDate   *time.Time     `json:"paymentDate,omitempty"`
	UserVoucherId int            `json:"userVoucherId"`
	VoucherId     int            `json:"voucherId"`
	Voucher       *Voucher       `json:"voucher,omitempty" gorm:"foreignKey:VoucherId"`
	Discount      float64        `json:"discount"`
	Subtotal      float64        `json:"subtotal"`
	Total         float64        `json:"total"`
	Transactions  []*Transaction `json:"transactions,omitempty" gorm:"foreignKey:InvoiceId"`
	gorm.Model    `json:"-"`
}

type InvoiceParams struct {
	Status string `json:"status"`
	Limit  int    `json:"limit"`
	Page   int    `json:"page"`
	Sort   string `json:"sort"`
}

func (i *InvoiceParams) Scope() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if i.Status != "" {
			db = db.Where("status = ?", i.Status)
		}
		return db
	}
}

func (i *InvoiceParams) Offset() int {
	return (i.Page - 1) * i.Limit
}

func NewInvoiceParams(status, sort string, limit, page int) InvoiceParams {
	return InvoiceParams{
		Status: func() string {
			if status == constant.InvoiceStatusWaitingPayment || status == constant.InvoiceStatusWaitingConfirmation || status == constant.InvoiceStatusCompleted || status == constant.InvoiceStatusCancelled {
				return status
			}
			return ""
		}(),
		Sort: func() string {
			if sort != "" && sort == "oldest" {
				return "created_at ASC"
			}
			return "created_at DESC"
		}(),
		Limit: func() int {
			if limit > 0 {
				return limit
			}
			return 10
		}(),
		Page: func() int {
			if page > 1 {
				return page
			}
			return 1
		}(),
	}
}
