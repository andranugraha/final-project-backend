package handler_test

import (
	"encoding/json"
	"final-project-backend/dto"
	mocks "final-project-backend/mocks/usecase"
	"final-project-backend/server"
	"final-project-backend/testutils"
	"final-project-backend/utils/constant"
	"final-project-backend/utils/errors"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestSignIn(t *testing.T) {
	var (
		req = dto.SignInRequest{
			Identifier: "test@mail.com",
			Password:   "password",
		}
	)

	tests := map[string]struct {
		req         dto.SignInRequest
		expectedRes gin.H
		code        int
		beforeTest  func(*mocks.AuthUsecase)
	}{
		"should return access token when status code 200": {
			req: req,
			expectedRes: gin.H{
				"data": dto.SignInResponse{
					AccessToken: "accessToken",
				},
			},
			code: http.StatusOK,
			beforeTest: func(mock *mocks.AuthUsecase) {
				mock.On("SignIn", req, constant.UserRoleId).Return(&dto.SignInResponse{
					AccessToken: "accessToken",
				}, nil)
			},
		},
		"should return error bad request when identifier is empty": {
			req: dto.SignInRequest{
				Identifier: "",
				Password:   req.Password,
			},
			expectedRes: gin.H{
				"code":    errors.ErrCodeBadRequest,
				"message": errors.ErrInvalidBody.Error(),
			},
			code:       http.StatusBadRequest,
			beforeTest: func(mock *mocks.AuthUsecase) {},
		},
		"should return error bad request when password is less than 8 characters": {
			req: dto.SignInRequest{
				Identifier: req.Identifier,
				Password:   "1234567",
			},
			expectedRes: gin.H{
				"code":    errors.ErrCodeBadRequest,
				"message": errors.ErrInvalidBody.Error(),
			},
			code:       http.StatusBadRequest,
			beforeTest: func(mock *mocks.AuthUsecase) {},
		},
		"should return error bad request when password is more than 20 characters": {
			req: dto.SignInRequest{
				Identifier: req.Identifier,
				Password:   "123456789012345678901",
			},
			expectedRes: gin.H{
				"code":    errors.ErrCodeBadRequest,
				"message": errors.ErrInvalidBody.Error(),
			},
			code:       http.StatusBadRequest,
			beforeTest: func(mock *mocks.AuthUsecase) {},
		},
		"should return error bad request when user not found": {
			req: req,
			expectedRes: gin.H{
				"code":    errors.ErrCodeBadRequest,
				"message": errors.ErrUserNotFound.Error(),
			},
			code: http.StatusBadRequest,
			beforeTest: func(mock *mocks.AuthUsecase) {
				mock.On("SignIn", req, constant.UserRoleId).Return(nil, errors.ErrUserNotFound)
			},
		},
		"should return error bad request when password is wrong": {
			req: req,
			expectedRes: gin.H{
				"code":    errors.ErrCodeBadRequest,
				"message": errors.ErrWrongPassword.Error(),
			},
			code: http.StatusBadRequest,
			beforeTest: func(mock *mocks.AuthUsecase) {
				mock.On("SignIn", req, constant.UserRoleId).Return(nil, errors.ErrWrongPassword)
			},
		},
		"should return error internal server when usecase return error": {
			req: req,
			expectedRes: gin.H{
				"code":    errors.ErrCodeInternalServerError,
				"message": errors.ErrInternalServerError.Error(),
			},
			code: http.StatusInternalServerError,
			beforeTest: func(mock *mocks.AuthUsecase) {
				mock.On("SignIn", req, constant.UserRoleId).Return(nil, errors.ErrInternalServerError)
			},
		},
	}

	for scenario, test := range tests {
		t.Run(scenario, func(t *testing.T) {
			jsonResult, _ := json.Marshal(test.expectedRes)
			mockUsecase := mocks.NewAuthUsecase(t)
			test.beforeTest(mockUsecase)
			cfg := &server.RouterConfig{
				AuthUsecase: mockUsecase,
			}
			payload := testutils.MakeRequestBody(test.req)

			req, _ := http.NewRequest(http.MethodPost, "/api/v1/sign-in", payload)
			_, rec := testutils.ServeReq(cfg, req)

			assert.Equal(t, test.code, rec.Code)
			assert.Equal(t, string(jsonResult), rec.Body.String())
		})
	}
}

func TestAdminSignIn(t *testing.T) {
	var (
		req = dto.SignInRequest{
			Identifier: "test@mail.com",
			Password:   "password",
		}
	)

	tests := map[string]struct {
		req         dto.SignInRequest
		expectedRes gin.H
		code        int
		beforeTest  func(*mocks.AuthUsecase)
	}{
		"should return access token when status code 200": {
			req: req,
			expectedRes: gin.H{
				"data": dto.SignInResponse{
					AccessToken: "accessToken",
				},
			},
			code: http.StatusOK,
			beforeTest: func(mock *mocks.AuthUsecase) {
				mock.On("SignIn", req, constant.AdminRoleId).Return(&dto.SignInResponse{
					AccessToken: "accessToken",
				}, nil)
			},
		},
		"should return error bad request when identifier is empty": {
			req: dto.SignInRequest{
				Identifier: "",
				Password:   req.Password,
			},
			expectedRes: gin.H{
				"code":    errors.ErrCodeBadRequest,
				"message": errors.ErrInvalidBody.Error(),
			},
			code:       http.StatusBadRequest,
			beforeTest: func(mock *mocks.AuthUsecase) {},
		},
		"should return error bad request when password is less than 8 characters": {
			req: dto.SignInRequest{
				Identifier: req.Identifier,
				Password:   "1234567",
			},
			expectedRes: gin.H{
				"code":    errors.ErrCodeBadRequest,
				"message": errors.ErrInvalidBody.Error(),
			},
			code:       http.StatusBadRequest,
			beforeTest: func(mock *mocks.AuthUsecase) {},
		},
		"should return error bad request when password is more than 20 characters": {
			req: dto.SignInRequest{
				Identifier: req.Identifier,
				Password:   "123456789012345678901",
			},
			expectedRes: gin.H{
				"code":    errors.ErrCodeBadRequest,
				"message": errors.ErrInvalidBody.Error(),
			},
			code:       http.StatusBadRequest,
			beforeTest: func(mock *mocks.AuthUsecase) {},
		},
		"should return error bad request when user not found": {
			req: req,
			expectedRes: gin.H{
				"code":    errors.ErrCodeBadRequest,
				"message": errors.ErrUserNotFound.Error(),
			},
			code: http.StatusBadRequest,
			beforeTest: func(mock *mocks.AuthUsecase) {
				mock.On("SignIn", req, constant.AdminRoleId).Return(nil, errors.ErrUserNotFound)
			},
		},
		"should return error bad request when password is wrong": {
			req: req,
			expectedRes: gin.H{
				"code":    errors.ErrCodeBadRequest,
				"message": errors.ErrWrongPassword.Error(),
			},
			code: http.StatusBadRequest,
			beforeTest: func(mock *mocks.AuthUsecase) {
				mock.On("SignIn", req, constant.AdminRoleId).Return(nil, errors.ErrWrongPassword)
			},
		},
		"should return error internal server when usecase return error": {
			req: req,
			expectedRes: gin.H{
				"code":    errors.ErrCodeInternalServerError,
				"message": errors.ErrInternalServerError.Error(),
			},
			code: http.StatusInternalServerError,
			beforeTest: func(mock *mocks.AuthUsecase) {
				mock.On("SignIn", req, constant.AdminRoleId).Return(nil, errors.ErrInternalServerError)
			},
		},
	}

	for scenario, test := range tests {
		t.Run(scenario, func(t *testing.T) {
			jsonResult, _ := json.Marshal(test.expectedRes)
			mockUsecase := mocks.NewAuthUsecase(t)
			test.beforeTest(mockUsecase)
			cfg := &server.RouterConfig{
				AuthUsecase: mockUsecase,
			}
			payload := testutils.MakeRequestBody(test.req)

			req, _ := http.NewRequest(http.MethodPost, "/api/v1/admin/sign-in", payload)
			_, rec := testutils.ServeReq(cfg, req)

			assert.Equal(t, test.code, rec.Code)
			assert.Equal(t, string(jsonResult), rec.Body.String())
		})
	}
}

func TestSignUp(t *testing.T) {
	var (
		req = dto.SignUpRequest{
			Email:       "test@mail.com",
			Password:    "password",
			Username:    "username",
			Fullname:    "fullname",
			Address:     "address",
			PhoneNo:     "081234567890",
			RefReferral: "referral",
		}
	)

	tests := map[string]struct {
		req         dto.SignUpRequest
		expectedRes gin.H
		code        int
		beforeTest  func(*mocks.AuthUsecase)
	}{
		"should return registered user when status code 200": {
			req: req,
			expectedRes: gin.H{
				"data": dto.SignUpResponse{
					Id:       1,
					Email:    req.Email,
					Username: req.Username,
					Fullname: req.Fullname,
					Address:  req.Address,
					PhoneNo:  req.PhoneNo,
					Referral: "referral",
				},
			},
			code: http.StatusOK,
			beforeTest: func(mock *mocks.AuthUsecase) {
				mock.On("SignUp", req).Return(&dto.SignUpResponse{
					Id:       1,
					Email:    req.Email,
					Username: req.Username,
					Fullname: req.Fullname,
					Address:  req.Address,
					PhoneNo:  req.PhoneNo,
					Referral: "referral",
				}, nil)
			},
		},
		"should return error bad request when email is empty": {
			req: dto.SignUpRequest{
				Email:       "",
				Password:    req.Password,
				Username:    req.Username,
				Fullname:    req.Fullname,
				Address:     req.Address,
				PhoneNo:     req.PhoneNo,
				RefReferral: req.RefReferral,
			},
			expectedRes: gin.H{
				"code":    errors.ErrCodeBadRequest,
				"message": errors.ErrInvalidBody.Error(),
			},
			code:       http.StatusBadRequest,
			beforeTest: func(mock *mocks.AuthUsecase) {},
		},
		"should return error bad request when email is not in email format": {
			req: dto.SignUpRequest{
				Email:       "test",
				Password:    req.Password,
				Username:    req.Username,
				Fullname:    req.Fullname,
				Address:     req.Address,
				PhoneNo:     req.PhoneNo,
				RefReferral: req.RefReferral,
			},
			expectedRes: gin.H{
				"code":    errors.ErrCodeBadRequest,
				"message": errors.ErrInvalidBody.Error(),
			},
			code:       http.StatusBadRequest,
			beforeTest: func(mock *mocks.AuthUsecase) {},
		},
		"should return error bad request when password is less than 8 characters": {
			req: dto.SignUpRequest{
				Email:       req.Email,
				Password:    "1234567",
				Username:    req.Username,
				Fullname:    req.Fullname,
				Address:     req.Address,
				PhoneNo:     req.PhoneNo,
				RefReferral: req.RefReferral,
			},
			expectedRes: gin.H{
				"code":    errors.ErrCodeBadRequest,
				"message": errors.ErrInvalidBody.Error(),
			},
			code:       http.StatusBadRequest,
			beforeTest: func(mock *mocks.AuthUsecase) {},
		},
		"should return error bad request when password is more than 20 characters": {
			req: dto.SignUpRequest{
				Email:       req.Email,
				Password:    "123456789012345678901",
				Username:    req.Username,
				Fullname:    req.Fullname,
				Address:     req.Address,
				PhoneNo:     req.PhoneNo,
				RefReferral: req.RefReferral,
			},
			expectedRes: gin.H{
				"code":    errors.ErrCodeBadRequest,
				"message": errors.ErrInvalidBody.Error(),
			},
			code:       http.StatusBadRequest,
			beforeTest: func(mock *mocks.AuthUsecase) {},
		},
		"should return error bad request when username is less than 3 characters": {
			req: dto.SignUpRequest{
				Email:       req.Email,
				Password:    req.Password,
				Username:    "12",
				Fullname:    req.Fullname,
				Address:     req.Address,
				PhoneNo:     req.PhoneNo,
				RefReferral: req.RefReferral,
			},
			expectedRes: gin.H{
				"code":    errors.ErrCodeBadRequest,
				"message": errors.ErrInvalidBody.Error(),
			},
			code:       http.StatusBadRequest,
			beforeTest: func(mock *mocks.AuthUsecase) {},
		},
		"should return error bad request when username is more than 20 characters": {
			req: dto.SignUpRequest{
				Email:       req.Email,
				Password:    req.Password,
				Username:    "123456789012345678901",
				Fullname:    req.Fullname,
				Address:     req.Address,
				PhoneNo:     req.PhoneNo,
				RefReferral: req.RefReferral,
			},
			expectedRes: gin.H{
				"code":    errors.ErrCodeBadRequest,
				"message": errors.ErrInvalidBody.Error(),
			},
			code:       http.StatusBadRequest,
			beforeTest: func(mock *mocks.AuthUsecase) {},
		},
		"should return error bad request when fullname is empty": {
			req: dto.SignUpRequest{
				Email:       req.Email,
				Password:    req.Password,
				Username:    req.Username,
				Fullname:    "",
				Address:     req.Address,
				PhoneNo:     req.PhoneNo,
				RefReferral: req.RefReferral,
			},
			expectedRes: gin.H{
				"code":    errors.ErrCodeBadRequest,
				"message": errors.ErrInvalidBody.Error(),
			},
			code:       http.StatusBadRequest,
			beforeTest: func(mock *mocks.AuthUsecase) {},
		},
		"should return error bad request when address is empty": {
			req: dto.SignUpRequest{
				Email:       req.Email,
				Password:    req.Password,
				Username:    req.Username,
				Fullname:    req.Fullname,
				Address:     "",
				PhoneNo:     req.PhoneNo,
				RefReferral: req.RefReferral,
			},
			expectedRes: gin.H{
				"code":    errors.ErrCodeBadRequest,
				"message": errors.ErrInvalidBody.Error(),
			},
			code:       http.StatusBadRequest,
			beforeTest: func(mock *mocks.AuthUsecase) {},
		},
		"should return error bad request when phone number is less than 10 characters": {
			req: dto.SignUpRequest{
				Email:       req.Email,
				Password:    req.Password,
				Username:    req.Username,
				Fullname:    req.Fullname,
				Address:     req.Address,
				PhoneNo:     "123456789",
				RefReferral: req.RefReferral,
			},
			expectedRes: gin.H{
				"code":    errors.ErrCodeBadRequest,
				"message": errors.ErrInvalidBody.Error(),
			},
			code:       http.StatusBadRequest,
			beforeTest: func(mock *mocks.AuthUsecase) {},
		},
		"should return error bad request when phone number is more than 14 characters": {
			req: dto.SignUpRequest{
				Email:       req.Email,
				Password:    req.Password,
				Username:    req.Username,
				Fullname:    req.Fullname,
				Address:     req.Address,
				PhoneNo:     "123456789012345",
				RefReferral: req.RefReferral,
			},
			expectedRes: gin.H{
				"code":    errors.ErrCodeBadRequest,
				"message": errors.ErrInvalidBody.Error(),
			},
			code:       http.StatusBadRequest,
			beforeTest: func(mock *mocks.AuthUsecase) {},
		},
		"should return error bad request when phone number is not numeric": {
			req: dto.SignUpRequest{
				Email:       req.Email,
				Password:    req.Password,
				Username:    req.Username,
				Fullname:    req.Fullname,
				Address:     req.Address,
				PhoneNo:     "1234567890a",
				RefReferral: req.RefReferral,
			},
			expectedRes: gin.H{
				"code":    errors.ErrCodeBadRequest,
				"message": errors.ErrInvalidBody.Error(),
			},
			code:       http.StatusBadRequest,
			beforeTest: func(mock *mocks.AuthUsecase) {},
		},
		"should return error bad request when user already exist": {
			req: req,
			expectedRes: gin.H{
				"code":    errors.ErrCodeBadRequest,
				"message": errors.ErrUserAlreadyExist.Error(),
			},
			code: http.StatusBadRequest,
			beforeTest: func(mock *mocks.AuthUsecase) {
				mock.On("SignUp", req).Return(nil, errors.ErrDuplicateRecord)
			},
		},
		"should return error internal server when error from usecase": {
			req: req,
			expectedRes: gin.H{
				"code":    errors.ErrCodeInternalServerError,
				"message": errors.ErrInternalServerError.Error(),
			},
			code: http.StatusInternalServerError,
			beforeTest: func(mock *mocks.AuthUsecase) {
				mock.On("SignUp", req).Return(nil, errors.ErrInternalServerError)
			},
		},
	}

	for scenario, test := range tests {
		t.Run(scenario, func(t *testing.T) {
			jsonResult, _ := json.Marshal(test.expectedRes)
			mockUsecase := mocks.NewAuthUsecase(t)
			test.beforeTest(mockUsecase)
			cfg := &server.RouterConfig{
				AuthUsecase: mockUsecase,
			}
			payload := testutils.MakeRequestBody(test.req)

			req, _ := http.NewRequest(http.MethodPost, "/api/v1/sign-up", payload)
			_, rec := testutils.ServeReq(cfg, req)

			assert.Equal(t, test.code, rec.Code)
			assert.Equal(t, string(jsonResult), rec.Body.String())
		})
	}
}
