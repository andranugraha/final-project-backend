package repository

import (
	"errors"
	"final-project-backend/entity"
	errResp "final-project-backend/utils/errors"
	"time"

	"gorm.io/gorm"
)

type UserVoucherRepository interface {
	FindValidByCode(string, int) (*entity.UserVoucher, error)
	ConsumeVoucher(*gorm.DB, int) error
}

type userVoucherRepositoryImpl struct {
	db *gorm.DB
}

type UserVoucherRConfig struct {
	DB *gorm.DB
}

func NewUserVoucherRepository(cfg *UserVoucherRConfig) UserVoucherRepository {
	return &userVoucherRepositoryImpl{
		db: cfg.DB,
	}
}

func (r *userVoucherRepositoryImpl) FindValidByCode(code string, userId int) (*entity.UserVoucher, error) {
	var userVoucher entity.UserVoucher
	err := r.db.Where("user_id = ?", userId).Where("expiry_date > ?", time.Now()).Where("is_consumed = ?", false).Joins("Voucher", r.db.Where(&entity.Voucher{Code: code})).First(&userVoucher).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errResp.ErrVoucherNotFound
		}

		return nil, err
	}
	return &userVoucher, nil
}

func (r *userVoucherRepositoryImpl) ConsumeVoucher(tx *gorm.DB, id int) error {
	err := tx.Model(&entity.UserVoucher{}).Where("id = ?", id).Update("is_consumed", true).Error
	if err != nil {
		return err
	}
	return nil
}
