package handler_test

import (
	"encoding/json"
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

func TestGetUserVouchers(t *testing.T) {
	tests := map[string]struct {
		code        int
		expectedRes gin.H
		expectedErr error
		beforeTest  func(*mocks.UserVoucherUsecase)
	}{
		"should return user vouchers when sucess": {
			code: http.StatusOK,
			expectedRes: gin.H{
				"data": []entity.UserVoucher{
					{
						ID: 1,
					},
				},
			},
			expectedErr: nil,
			beforeTest: func(mock *mocks.UserVoucherUsecase) {
				mock.On("GetUserVouchers", 1).Return([]entity.UserVoucher{
					{
						ID: 1,
					},
				}, nil)
			},
		},
		"should return not found when user not found": {
			code: http.StatusNotFound,
			expectedRes: gin.H{
				"code":    errors.ErrCodeNotFound,
				"message": errors.ErrUserNotFound.Error(),
			},
			expectedErr: errors.ErrUserNotFound,
			beforeTest: func(mock *mocks.UserVoucherUsecase) {
				mock.On("GetUserVouchers", 1).Return([]entity.UserVoucher{}, errors.ErrUserNotFound)
			},
		},
		"should return error when failed": {
			code: http.StatusInternalServerError,
			expectedRes: gin.H{
				"code":    errors.ErrCodeInternalServerError,
				"message": errors.ErrInternalServerError.Error(),
			},
			expectedErr: errors.ErrInternalServerError,
			beforeTest: func(mock *mocks.UserVoucherUsecase) {
				mock.On("GetUserVouchers", 1).Return([]entity.UserVoucher{}, errors.ErrInternalServerError)
			},
		},
	}

	for scenario, test := range tests {
		t.Run(scenario, func(t *testing.T) {
			jsonResult, _ := json.Marshal(test.expectedRes)
			mock := mocks.NewUserVoucherUsecase(t)
			test.beforeTest(mock)
			cfg := &server.RouterConfig{
				UserVoucherUsecase: mock,
			}

			req, _ := http.NewRequest(http.MethodGet, "/api/v1/vouchers", nil)
			_, rec := testutils.ServeReq(cfg, req)

			assert.Equal(t, test.code, rec.Code)
			assert.Equal(t, string(jsonResult), rec.Body.String())
		})
	}
}
