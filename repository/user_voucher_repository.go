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
	Insert(*gorm.DB, int64, int) (*entity.UserVoucher, error)
}

type userVoucherRepositoryImpl struct {
	db          *gorm.DB
	voucherRepo VoucherRepository
}

type UserVoucherRConfig struct {
	DB          *gorm.DB
	VoucherRepo VoucherRepository
}

func NewUserVoucherRepository(cfg *UserVoucherRConfig) UserVoucherRepository {
	return &userVoucherRepositoryImpl{
		db:          cfg.DB,
		voucherRepo: cfg.VoucherRepo,
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

func (r *userVoucherRepositoryImpl) Insert(tx *gorm.DB, amount int64, userId int) (*entity.UserVoucher, error) {
	voucher, err := r.voucherRepo.FindValidByAmount(amount)
	if err != nil {
		return nil, err
	}

	if voucher == nil {
		return nil, nil
	}

	userVoucher := entity.UserVoucher{
		UserId:     userId,
		VoucherId:  voucher.ID,
		ExpiryDate: time.Now().AddDate(0, 1, 0),
	}

	err = tx.Create(&userVoucher).Error
	if err != nil {
		return nil, err
	}

	return &userVoucher, nil
}
