package handler_test

import (
	"encoding/json"
	"final-project-backend/dto"
	"final-project-backend/entity"
	"final-project-backend/handler"
	mocks "final-project-backend/mocks/usecase"
	"final-project-backend/server"
	"final-project-backend/testutils"
	"final-project-backend/utils/errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetCart(t *testing.T) {
	tests := map[string]struct {
		expectedRes gin.H
		code        int
		beforeTest  func(*mocks.CartUsecase)
	}{
		"should return 200 when get cart success": {
			expectedRes: gin.H{
				"data": dto.GetCartResponse{
					CartItems: []*entity.Cart{
						{
							ID: 1,
						},
					},
					TotalPrice: 0,
				},
			},
			code: http.StatusOK,
			beforeTest: func(m *mocks.CartUsecase) {
				m.On("GetCart", 1).Return(&dto.GetCartResponse{
					CartItems: []*entity.Cart{
						{
							ID: 1,
						},
					},
					TotalPrice: 0,
				}, nil)
			},
		},
		"should return 500 when get cart failed": {
			expectedRes: gin.H{
				"code":    errors.ErrCodeInternalServerError,
				"message": errors.ErrInternalServerError.Error(),
			},
			code: http.StatusInternalServerError,
			beforeTest: func(m *mocks.CartUsecase) {
				m.On("GetCart", 1).Return(nil, errors.ErrInternalServerError)
			},
		},
	}

	for scenario, test := range tests {
		t.Run(scenario, func(t *testing.T) {
			jsonResult, _ := json.Marshal(test.expectedRes)
			mockUsecase := mocks.NewCartUsecase(t)
			test.beforeTest(mockUsecase)
			cfg := &handler.Config{
				CartUsecase: mockUsecase,
			}
			h := handler.New(cfg)
			rec := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(rec)
			c.Set("userId", 1)

			c.Request, _ = http.NewRequest(http.MethodGet, "/api/v1/carts", nil)
			h.GetCart(c)

			assert.Equal(t, test.code, rec.Code)
			assert.Equal(t, string(jsonResult), rec.Body.String())
		})
	}
}

func TestAddToCart(t *testing.T) {
	var (
		userId   = 1
		courseId = 1
	)

	tests := map[string]struct {
		param       string
		expectedRes gin.H
		code        int
		beforeTest  func(*mocks.CartUsecase)
	}{
		"should return 200 when add to cart success": {
			param: "1",
			expectedRes: gin.H{
				"message": "Course added to cart",
				"data":    nil,
			},
			code: http.StatusOK,
			beforeTest: func(m *mocks.CartUsecase) {
				m.On("AddToCart", userId, courseId).Return(nil)
			},
		},
		"should return 400 when course id param is not integer": {
			param: "abc",
			expectedRes: gin.H{
				"code":    errors.ErrCodeBadRequest,
				"message": errors.ErrInvalidParamFormat.Error(),
			},
			code:       http.StatusBadRequest,
			beforeTest: func(m *mocks.CartUsecase) {},
		},
		"should return 400 when course already exist in cart": {
			param: "1",
			expectedRes: gin.H{
				"code":    errors.ErrCodeBadRequest,
				"message": errors.ErrDuplicateCart.Error(),
			},
			code: http.StatusBadRequest,
			beforeTest: func(m *mocks.CartUsecase) {
				m.On("AddToCart", userId, courseId).Return(errors.ErrDuplicateCart)
			},
		},
		"should return 400 when course not found": {
			param: "1",
			expectedRes: gin.H{
				"code":    errors.ErrCodeBadRequest,
				"message": errors.ErrCourseNotFound.Error(),
			},
			code: http.StatusBadRequest,
			beforeTest: func(m *mocks.CartUsecase) {
				m.On("AddToCart", userId, courseId).Return(errors.ErrCourseNotFound)
			},
		},
		"should return 400 when course already bought": {
			param: "1",
			expectedRes: gin.H{
				"code":    errors.ErrCodeBadRequest,
				"message": errors.ErrCourseAlreadyBought.Error(),
			},
			code: http.StatusBadRequest,
			beforeTest: func(m *mocks.CartUsecase) {
				m.On("AddToCart", userId, courseId).Return(errors.ErrCourseAlreadyBought)
			},
		},
		"should return 500 when add to cart failed": {
			param: "1",
			expectedRes: gin.H{
				"code":    errors.ErrCodeInternalServerError,
				"message": errors.ErrInternalServerError.Error(),
			},
			code: http.StatusInternalServerError,
			beforeTest: func(m *mocks.CartUsecase) {
				m.On("AddToCart", userId, courseId).Return(errors.ErrInternalServerError)
			},
		},
	}

	for scenario, test := range tests {
		t.Run(scenario, func(t *testing.T) {
			jsonResult, _ := json.Marshal(test.expectedRes)
			mockUsecase := mocks.NewCartUsecase(t)
			test.beforeTest(mockUsecase)
			cfg := &server.RouterConfig{
				CartUsecase: mockUsecase,
			}

			req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/api/v1/carts/%v", test.param), nil)
			_, rec := testutils.ServeReq(cfg, req)

			assert.Equal(t, test.code, rec.Code)
			assert.Equal(t, string(jsonResult), rec.Body.String())
		})
	}
}

func TestRemoveFromCart(t *testing.T) {
	var (
		userId   = 1
		courseId = 1
	)

	tests := map[string]struct {
		param       string
		expectedRes gin.H
		code        int
		beforeTest  func(*mocks.CartUsecase)
	}{
		"should return 200 when remove from cart success": {
			param: "1",
			expectedRes: gin.H{
				"message": "Course removed from cart",
				"data":    nil,
			},
			code: http.StatusOK,
			beforeTest: func(m *mocks.CartUsecase) {
				m.On("RemoveFromCart", userId, courseId).Return(nil)
			},
		},
		"should return 400 when course id param is not integer": {
			param: "abc",
			expectedRes: gin.H{
				"code":    errors.ErrCodeBadRequest,
				"message": errors.ErrInvalidParamFormat.Error(),
			},
			code:       http.StatusBadRequest,
			beforeTest: func(m *mocks.CartUsecase) {},
		},
		"should return 400 when course not found": {
			param: "1",
			expectedRes: gin.H{
				"code":    errors.ErrCodeBadRequest,
				"message": errors.ErrCourseNotFound.Error(),
			},
			code: http.StatusBadRequest,
			beforeTest: func(m *mocks.CartUsecase) {
				m.On("RemoveFromCart", userId, courseId).Return(errors.ErrCourseNotFound)
			},
		},
		"should return 500 when remove from cart failed": {
			param: "1",
			expectedRes: gin.H{
				"code":    errors.ErrCodeInternalServerError,
				"message": errors.ErrInternalServerError.Error(),
			},
			code: http.StatusInternalServerError,
			beforeTest: func(m *mocks.CartUsecase) {
				m.On("RemoveFromCart", userId, courseId).Return(errors.ErrInternalServerError)
			},
		},
	}

	for scenario, test := range tests {
		t.Run(scenario, func(t *testing.T) {
			jsonResult, _ := json.Marshal(test.expectedRes)
			mockUsecase := mocks.NewCartUsecase(t)
			test.beforeTest(mockUsecase)
			cfg := &server.RouterConfig{
				CartUsecase: mockUsecase,
			}

			req, _ := http.NewRequest(http.MethodDelete, fmt.Sprintf("/api/v1/carts/%v", test.param), nil)
			_, rec := testutils.ServeReq(cfg, req)

			assert.Equal(t, test.code, rec.Code)
			assert.Equal(t, string(jsonResult), rec.Body.String())
		})
	}
}
