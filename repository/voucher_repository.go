package repository

import (
	"errors"
	"final-project-backend/entity"

	"gorm.io/gorm"
)

type VoucherRepository interface {
	FindValidByAmount(int64) (*entity.Voucher, error)
}

type voucherRepositoryImpl struct {
	db *gorm.DB
}

type VoucherRConfig struct {
	DB *gorm.DB
}

func NewVoucherRepository(cfg *VoucherRConfig) VoucherRepository {
	return &voucherRepositoryImpl{
		db: cfg.DB,
	}
}

func (r *voucherRepositoryImpl) FindValidByAmount(amount int64) (*entity.Voucher, error) {
	var voucher entity.Voucher
	err := r.db.Where("min_amount <= ?", amount).Order("min_amount desc").First(&voucher).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &voucher, nil
}
