package usecase_test

import (
	"final-project-backend/entity"
	mocks "final-project-backend/mocks/repository"
	"final-project-backend/usecase"
	errResp "final-project-backend/utils/errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserVouchers(t *testing.T) {
	var userId int = 1

	tests := map[string]struct {
		expectedRes []entity.UserVoucher
		expectedErr error
	}{
		"should return all user vouchers when success": {
			expectedRes: []entity.UserVoucher{
				{
					ID:        1,
					UserId:    1,
					VoucherId: 1,
				},
			},
			expectedErr: nil,
		},
		"should return error when failed": {
			expectedRes: nil,
			expectedErr: errResp.ErrInternalServerError,
		},
	}

	for scenario, test := range tests {
		t.Run(scenario, func(t *testing.T) {
			userVoucherRepo := mocks.NewUserVoucherRepository(t)
			userVoucherRepo.On("FindAll", userId).Return(test.expectedRes, test.expectedErr)
			userVoucherUsecase := usecase.NewUserVoucherUsecase(&usecase.UserVoucherUConfig{
				UserVoucherRepo: userVoucherRepo,
			})

			res, err := userVoucherUsecase.GetUserVouchers(userId)

			assert.Equal(t, test.expectedRes, res)
			assert.ErrorIs(t, test.expectedErr, err)
		})
	}
}

func TestFindValidByCode(t *testing.T) {
	var (
		code   string = "code"
		userId int    = 1
	)

	tests := map[string]struct {
		expectedRes *entity.UserVoucher
		expectedErr error
	}{
		"should return user voucher when success": {
			expectedRes: &entity.UserVoucher{
				ID:        1,
				UserId:    1,
				VoucherId: 1,
			},
			expectedErr: nil,
		},
		"should return error when failed": {
			expectedRes: nil,
			expectedErr: errResp.ErrInternalServerError,
		},
	}

	for scenario, test := range tests {
		t.Run(scenario, func(t *testing.T) {
			userVoucherRepo := mocks.NewUserVoucherRepository(t)
			userVoucherRepo.On("FindValidByCode", code, userId).Return(test.expectedRes, test.expectedErr)
			userVoucherUsecase := usecase.NewUserVoucherUsecase(&usecase.UserVoucherUConfig{
				UserVoucherRepo: userVoucherRepo,
			})

			res, err := userVoucherUsecase.FindValidByCode(code, userId)

			assert.Equal(t, test.expectedRes, res)
			assert.ErrorIs(t, test.expectedErr, err)
		})
	}
}
