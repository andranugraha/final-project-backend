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

func TestGetTags(t *testing.T) {
	tests := map[string]struct {
		code        int
		expectedRes gin.H
		expectedErr error
		beforeTest  func(*mocks.TagUsecase)
	}{
		"should return tags when sucess": {
			code: http.StatusOK,
			expectedRes: gin.H{
				"data": []*entity.Tag{
					{
						ID: 1,
					},
				},
			},
			expectedErr: nil,
			beforeTest: func(mock *mocks.TagUsecase) {
				mock.On("GetTags").Return([]*entity.Tag{
					{
						ID: 1,
					},
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
			beforeTest: func(mock *mocks.TagUsecase) {
				mock.On("GetTags").Return(nil, errors.ErrInternalServerError)
			},
		},
	}

	for scenario, test := range tests {
		t.Run(scenario, func(t *testing.T) {
			jsonResult, _ := json.Marshal(test.expectedRes)
			mock := mocks.NewTagUsecase(t)
			test.beforeTest(mock)
			cfg := &server.RouterConfig{
				TagUsecase: mock,
			}

			req, _ := http.NewRequest(http.MethodGet, "/api/v1/tags", nil)
			_, rec := testutils.ServeReq(cfg, req)

			assert.Equal(t, test.code, rec.Code)
			assert.Equal(t, string(jsonResult), rec.Body.String())
		})
	}
}
