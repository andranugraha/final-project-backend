package usecase_test

import (
	"final-project-backend/dto"
	"final-project-backend/entity"
	mocks "final-project-backend/mocks/repository"
	mocksAuth "final-project-backend/mocks/utils/auth"
	"final-project-backend/usecase"
	errResp "final-project-backend/utils/errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSignUp(t *testing.T) {
	var (
		req = dto.SignUpRequest{
			Email:    "test@mail.com",
			Password: "password",
		}
		entityReq    = req.ToUser()
		hashedPwd    = "testhashedpwd"
		referralCode = "abcdefg"
	)

	tests := map[string]struct {
		insertRes   *entity.User
		expectedRes *dto.SignUpResponse
		expectedErr error
	}{
		"should return registered user when given valid request": {
			insertRes: func() *entity.User {
				entityReq.Password = hashedPwd
				entityReq.Referral = referralCode
				entityReq.Redeemable = &entity.Redeemable{}
				return &entityReq
			}(),
			expectedRes: func() *dto.SignUpResponse {
				entityReq.Password = hashedPwd
				entityReq.Referral = referralCode
				entityReq.Redeemable = &entity.Redeemable{}

				var expectedRes dto.SignUpResponse
				expectedRes.FromUser(entityReq)
				return &expectedRes
			}(),
			expectedErr: nil,
		},
		"should return error when insert failed": {
			insertRes:   nil,
			expectedRes: nil,
			expectedErr: errResp.ErrInternalServerError,
		},
	}

	for scenario, test := range tests {
		t.Run(scenario, func(t *testing.T) {
			mockRepo := mocks.NewUserRepository(t)
			mockAuthUtil := mocksAuth.NewAuthUtil(t)
			mockAuthUtil.On("HashAndSalt", req.Password).Return(hashedPwd).Once()
			mockAuthUtil.On("GenerateReferralCode").Return(referralCode).Once()
			mockRepo.On("Insert", entityReq).Return(test.insertRes, test.expectedErr).Once()
			u := usecase.NewAuthUsecase(&usecase.AuthUConfig{
				UserRepo:    mockRepo,
				UtilUsecase: mockAuthUtil,
			})

			res, err := u.SignUp(req)

			assert.Equal(t, test.expectedRes, res)
			assert.ErrorIs(t, test.expectedErr, err)
		})
	}
}

func TestSignIn(t *testing.T) {
	var (
		req = dto.SignInRequest{
			Identifier: "test@mail.com",
			Password:   "password",
		}
		roleId = 1
		user   = entity.User{
			Password: "testhashedpwd",
		}
		token       = "testtoken"
		expectedRes = dto.SignInResponse{
			AccessToken: token,
		}
	)

	tests := map[string]struct {
		findRes     *entity.User
		findErr     error
		expectedRes *dto.SignInResponse
		expectedErr error
		mockUtilFn  func(*mocksAuth.AuthUtil)
	}{
		"should return access token when given valid request": {
			findRes: &user,
			findErr: nil,
			expectedRes: &dto.SignInResponse{
				AccessToken: token,
			},
			expectedErr: nil,
			mockUtilFn: func(mockAuthUtil *mocksAuth.AuthUtil) {
				mockAuthUtil.On("ComparePassword", user.Password, req.Password).Return(true).Once()
				mockAuthUtil.On("GenerateAccessToken", user).Return(expectedRes).Once()
			},
		},
		"should return error when user not found": {
			findRes:     nil,
			findErr:     errResp.ErrUserNotFound,
			expectedRes: nil,
			expectedErr: errResp.ErrUserNotFound,
			mockUtilFn:  func(mockAuthUtil *mocksAuth.AuthUtil) {},
		},
		"should return error when password is incorrect": {
			findRes:     &user,
			findErr:     nil,
			expectedRes: nil,
			expectedErr: errResp.ErrWrongPassword,
			mockUtilFn: func(mockAuthUtil *mocksAuth.AuthUtil) {
				mockAuthUtil.On("ComparePassword", user.Password, req.Password).Return(false).Once()
			},
		},
	}

	for scenario, test := range tests {
		t.Run(scenario, func(t *testing.T) {
			mockRepo := mocks.NewUserRepository(t)
			mockRepo.On("FindByIdentifierAndRole", req.Identifier, roleId).Return(test.findRes, test.findErr).Once()
			mockAuthUtil := mocksAuth.NewAuthUtil(t)
			test.mockUtilFn(mockAuthUtil)
			u := usecase.NewAuthUsecase(&usecase.AuthUConfig{
				UserRepo:    mockRepo,
				UtilUsecase: mockAuthUtil,
			})

			res, err := u.SignIn(req, roleId)

			assert.Equal(t, test.expectedRes, res)
			assert.ErrorIs(t, test.expectedErr, err)
		})
	}

}
