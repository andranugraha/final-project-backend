package handler_test

import (
	"encoding/json"
	"final-project-backend/dto"
	"final-project-backend/entity"
	mocks "final-project-backend/mocks/usecase"
	"final-project-backend/server"
	"final-project-backend/testutils"
	"final-project-backend/utils/errors"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetUserDetail(t *testing.T) {
	tests := map[string]struct {
		code        int
		expectedRes gin.H
		expectedErr error
		beforeTest  func(*mocks.UserUsecase)
	}{
		"should return user detail when sucess": {
			code: http.StatusOK,
			expectedRes: gin.H{
				"data": &entity.User{
					ID: 1,
				},
			},
			expectedErr: nil,
			beforeTest: func(mock *mocks.UserUsecase) {
				mock.On("GetUserDetail", 1).Return(&entity.User{
					ID: 1,
				}, nil)
			},
		},
		"should return error when failed": {
			code: http.StatusInternalServerError,
			expectedRes: gin.H{
				"code":    errors.ErrCodeInternalServerError,
				"message": errors.ErrInternalServerError.Error(),
			},
			expectedErr: errors.ErrInternalServerError,
			beforeTest: func(mock *mocks.UserUsecase) {
				mock.On("GetUserDetail", 1).Return(nil, errors.ErrInternalServerError)
			},
		},
	}

	for scenario, test := range tests {
		t.Run(scenario, func(t *testing.T) {
			jsonResult, _ := json.Marshal(test.expectedRes)
			mock := mocks.NewUserUsecase(t)
			test.beforeTest(mock)
			cfg := &server.RouterConfig{
				UserUsecase: mock,
			}

			req, _ := http.NewRequest(http.MethodGet, "/api/v1/user", nil)
			_, rec := testutils.ServeReq(cfg, req)

			assert.Equal(t, test.code, rec.Code)
			assert.Equal(t, string(jsonResult), rec.Body.String())
		})
	}
}

func TestUpdateUserDetail(t *testing.T) {
	var (
		defaultReq = dto.UpdateUserDetailRequest{
			Fullname: "test",
			Address:  "test",
			PhoneNo:  "081234567890",
		}
	)

	tests := map[string]struct {
		req         dto.UpdateUserDetailRequest
		code        int
		expectedRes gin.H
		expectedErr error
		beforeTest  func(*mocks.UserUsecase)
	}{
		"should return user detail when sucess": {
			req:  defaultReq,
			code: http.StatusOK,
			expectedRes: gin.H{
				"data": &entity.User{
					ID: 1,
				},
			},
			expectedErr: nil,
			beforeTest: func(mock *mocks.UserUsecase) {
				mock.On("UpdateUserDetail", 1, defaultReq).Return(&entity.User{
					ID: 1,
				}, nil)
			},
		},
		"should return bad request when invalid body request": {
			req:  dto.UpdateUserDetailRequest{},
			code: http.StatusBadRequest,
			expectedRes: gin.H{
				"code":    errors.ErrCodeBadRequest,
				"message": errors.ErrInvalidBody.Error(),
			},
			expectedErr: errors.ErrInvalidBody,
			beforeTest:  func(mock *mocks.UserUsecase) {},
		},
		"should return not found when user not found": {
			req:  defaultReq,
			code: http.StatusNotFound,
			expectedRes: gin.H{
				"code":    errors.ErrCodeNotFound,
				"message": errors.ErrUserNotFound.Error(),
			},
			expectedErr: errors.ErrUserNotFound,
			beforeTest: func(mock *mocks.UserUsecase) {
				mock.On("UpdateUserDetail", 1, defaultReq).Return(nil, errors.ErrUserNotFound)
			},
		},
		"should return bad request when duplicate phone number": {
			req:  defaultReq,
			code: http.StatusBadRequest,
			expectedRes: gin.H{
				"code":    errors.ErrCodeBadRequest,
				"message": errors.ErrDuplicatePhoneNo.Error(),
			},
			expectedErr: errors.ErrDuplicatePhoneNo,
			beforeTest: func(mock *mocks.UserUsecase) {
				mock.On("UpdateUserDetail", 1, defaultReq).Return(nil, errors.ErrDuplicatePhoneNo)
			},
		},
		"should return internal server error when failed": {
			req:  defaultReq,
			code: http.StatusInternalServerError,
			expectedRes: gin.H{
				"code":    errors.ErrCodeInternalServerError,
				"message": errors.ErrInternalServerError.Error(),
			},
			expectedErr: errors.ErrInternalServerError,
			beforeTest: func(mock *mocks.UserUsecase) {
				mock.On("UpdateUserDetail", 1, defaultReq).Return(nil, errors.ErrInternalServerError)
			},
		},
	}

	for scenario, test := range tests {
		t.Run(scenario, func(t *testing.T) {
			jsonResult, _ := json.Marshal(test.expectedRes)
			mock := mocks.NewUserUsecase(t)
			test.beforeTest(mock)
			cfg := &server.RouterConfig{
				UserUsecase: mock,
			}
			payload := testutils.MakeRequestBody(test.req)

			req, _ := http.NewRequest(http.MethodPut, "/api/v1/user", payload)
			_, rec := testutils.ServeReq(cfg, req)

			assert.Equal(t, test.code, rec.Code)
			assert.Equal(t, string(jsonResult), rec.Body.String())
		})
	}
}
