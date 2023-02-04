package usecase_test

import (
	"final-project-backend/dto"
	"final-project-backend/entity"
	mocks "final-project-backend/mocks/repository"
	"final-project-backend/usecase"
	errResp "final-project-backend/utils/errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserDetail(t *testing.T) {
	var (
		userId = 1
		user   = entity.User{
			Fullname:   "test",
			Email:      "test@mail.com",
			Username:   "test",
			PhoneNo:    "081234567890",
			Password:   "password",
			Address:    "test address",
			Referral:   "abcdefg",
			Redeemable: &entity.Redeemable{},
		}
	)

	tests := map[string]struct {
		findRes     *entity.User
		expectedRes *entity.User
		expectedErr error
	}{
		"should return user detail when given valid request": {
			findRes:     &user,
			expectedRes: &user,
			expectedErr: nil,
		},
		"should return error when find failed": {
			findRes:     nil,
			expectedRes: nil,
			expectedErr: errResp.ErrInternalServerError,
		},
	}

	for scenario, test := range tests {
		t.Run(scenario, func(t *testing.T) {
			mockRepo := mocks.NewUserRepository(t)
			mockRepo.On("FindDetailById", userId).Return(test.findRes, test.expectedErr).Once()

			u := usecase.NewUserUsecase(&usecase.UserUConfig{
				UserRepo: mockRepo,
			})

			res, err := u.GetUserDetail(userId)

			assert.Equal(t, test.expectedRes, res)
			assert.ErrorIs(t, test.expectedErr, err)
		})
	}
}

func TestUpdateUserDetail(t *testing.T) {
	var (
		userId = 1
		req    = dto.UpdateUserDetailRequest{
			Fullname: "test",
			Address:  "test address",
			PhoneNo:  "081234567890",
		}
	)

	tests := map[string]struct {
		updateRes   *entity.User
		expectedRes *entity.User
		expectedErr error
	}{
		"should return updated user detail when given valid request": {
			updateRes: func() *entity.User {
				user := req.ToUser(userId)
				return &user
			}(),
			expectedRes: func() *entity.User {
				user := req.ToUser(userId)
				return &user
			}(),
			expectedErr: nil,
		},
		"should return error when update failed": {
			updateRes:   nil,
			expectedRes: nil,
			expectedErr: errResp.ErrInternalServerError,
		},
	}

	for scenario, test := range tests {
		t.Run(scenario, func(t *testing.T) {
			mockRepo := mocks.NewUserRepository(t)
			mockRepo.On("UpdateDetail", req.ToUser(userId)).Return(test.updateRes, test.expectedErr).Once()

			u := usecase.NewUserUsecase(&usecase.UserUConfig{
				UserRepo: mockRepo,
			})

			res, err := u.UpdateUserDetail(userId, req)

			assert.Equal(t, test.expectedRes, res)
			assert.ErrorIs(t, test.expectedErr, err)
		})
	}
}
