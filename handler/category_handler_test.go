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

func TestGetCategories(t *testing.T) {
	tests := map[string]struct {
		expectedRes gin.H
		code        int
		beforeTest  func(*mocks.CategoryUsecase)
	}{
		"should return 200 when get categories success": {
			expectedRes: gin.H{
				"data": []*entity.Category{
					{
						ID:   1,
						Name: "category 1",
					},
				},
			},
			code: http.StatusOK,
			beforeTest: func(m *mocks.CategoryUsecase) {
				m.On("GetCategories").Return([]*entity.Category{
					{
						ID:   1,
						Name: "category 1",
					},
				}, nil)
			},
		},
		"should return 500 when get categories failed": {
			expectedRes: gin.H{
				"code":    errors.ErrCodeInternalServerError,
				"message": errors.ErrInternalServerError.Error(),
			},
			code: http.StatusInternalServerError,
			beforeTest: func(m *mocks.CategoryUsecase) {
				m.On("GetCategories").Return(nil, errors.ErrInternalServerError)
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			jsonResult, _ := json.Marshal(test.expectedRes)
			mockUsecase := mocks.NewCategoryUsecase(t)
			test.beforeTest(mockUsecase)
			cfg := &server.RouterConfig{
				CategoryUsecase: mockUsecase,
			}

			req, _ := http.NewRequest(http.MethodGet, "/api/v1/categories", nil)
			_, rec := testutils.ServeReq(cfg, req)

			assert.Equal(t, test.code, rec.Code)
			assert.Equal(t, string(jsonResult), rec.Body.String())
		})
	}
}
