package usecase

import (
	"final-project-backend/entity"
	"final-project-backend/repository"
)

type UserVoucherUsecase interface {
	GetUserVouchers(int) ([]entity.UserVoucher, error)
	FindValidByCode(string, int) (*entity.UserVoucher, error)
}

type userVoucherUsecaseImpl struct {
	userVoucherRepo repository.UserVoucherRepository
}

type UserVoucherUConfig struct {
	UserVoucherRepo repository.UserVoucherRepository
}

func NewUserVoucherUsecase(cfg *UserVoucherUConfig) UserVoucherUsecase {
	return &userVoucherUsecaseImpl{
		userVoucherRepo: cfg.UserVoucherRepo,
	}
}

func (u *userVoucherUsecaseImpl) GetUserVouchers(userId int) ([]entity.UserVoucher, error) {
	return u.userVoucherRepo.FindAll(userId)
}

func (u *userVoucherUsecaseImpl) FindValidByCode(code string, userId int) (*entity.UserVoucher, error) {
	voucher, err := u.userVoucherRepo.FindValidByCode(code, userId)
	if err != nil {
		return nil, err
	}

	return voucher, nil
}
