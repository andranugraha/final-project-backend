package handler_test

import (
	"encoding/json"
	"final-project-backend/entity"
	mocks "final-project-backend/mocks/usecase"
	"final-project-backend/server"
	"final-project-backend/testutils"
	"final-project-backend/utils/errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetFavoriteCourses(t *testing.T) {
	tests := map[string]struct {
		code        int
		expectedRes gin.H
		expectedErr error
		beforeTest  func(*mocks.FavoriteUsecase)
	}{
		"should return favorite courses when sucess": {
			code: http.StatusOK,
			expectedRes: gin.H{
				"data": []*entity.Course{
					{
						ID: 1,
					},
				},
			},
			expectedErr: nil,
			beforeTest: func(mock *mocks.FavoriteUsecase) {
				mock.On("GetFavoriteCourses", 1).Return([]*entity.Course{
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
			beforeTest: func(mock *mocks.FavoriteUsecase) {
				mock.On("GetFavoriteCourses", 1).Return(nil, errors.ErrInternalServerError)
			},
		},
	}

	for scenario, test := range tests {
		t.Run(scenario, func(t *testing.T) {
			jsonResult, _ := json.Marshal(test.expectedRes)
			mock := mocks.NewFavoriteUsecase(t)
			test.beforeTest(mock)
			cfg := &server.RouterConfig{
				FavoriteUsecase: mock,
			}

			req, _ := http.NewRequest("GET", "/api/v1/favorites", nil)
			_, rec := testutils.ServeReq(cfg, req)

			assert.Equal(t, test.code, rec.Code)
			assert.Equal(t, string(jsonResult), rec.Body.String())
		})
	}
}

func TestSaveUnsaveFavoriteCourse(t *testing.T) {
	tests := map[string]struct {
		param       string
		code        int
		expectedRes gin.H
		expectedErr error
		beforeTest  func(*mocks.FavoriteUsecase)
	}{
		"should return favorite courses when sucess": {
			param: "1",
			code:  http.StatusOK,
			expectedRes: gin.H{
				"data": nil,
			},
			expectedErr: nil,
			beforeTest: func(mock *mocks.FavoriteUsecase) {
				mock.On("SaveUnsaveFavoriteCourse", 1, 1, "save").Return(nil)
			},
		},
		"should return bad request when course id is not integer": {
			param: "a",
			code:  http.StatusBadRequest,
			expectedRes: gin.H{
				"code":    errors.ErrCodeBadRequest,
				"message": errors.ErrInvalidParamFormat.Error(),
			},
			expectedErr: errors.ErrInvalidParamFormat,
			beforeTest:  func(mock *mocks.FavoriteUsecase) {},
		},
		"should return not found when course id is not found": {
			param: "1",
			code:  http.StatusNotFound,
			expectedRes: gin.H{
				"code":    errors.ErrCodeNotFound,
				"message": errors.ErrFavoriteNotFound.Error(),
			},
			expectedErr: errors.ErrFavoriteNotFound,
			beforeTest: func(mock *mocks.FavoriteUsecase) {
				mock.On("SaveUnsaveFavoriteCourse", 1, 1, "save").Return(errors.ErrFavoriteNotFound)
			},
		},
		"should return bad request when action is not save or unsave": {
			param: "1",
			code:  http.StatusBadRequest,
			expectedRes: gin.H{
				"code":    errors.ErrCodeBadRequest,
				"message": errors.ErrUnknownAction.Error(),
			},
			expectedErr: errors.ErrUnknownAction,
			beforeTest: func(mockUsecase *mocks.FavoriteUsecase) {
				mockUsecase.On("SaveUnsaveFavoriteCourse", 1, 1, mock.Anything).Return(errors.ErrUnknownAction)
			},
		},
		"should return bad request when duplicate favorite": {
			param: "1",
			code:  http.StatusBadRequest,
			expectedRes: gin.H{
				"code":    errors.ErrCodeBadRequest,
				"message": errors.ErrDuplicateFavorite.Error(),
			},
			expectedErr: errors.ErrDuplicateFavorite,
			beforeTest: func(mock *mocks.FavoriteUsecase) {
				mock.On("SaveUnsaveFavoriteCourse", 1, 1, "save").Return(errors.ErrDuplicateFavorite)
			},
		},
		"should return internal server error when failed": {
			param: "1",
			code:  http.StatusInternalServerError,
			expectedRes: gin.H{
				"code":    errors.ErrCodeInternalServerError,
				"message": errors.ErrInternalServerError.Error(),
			},
			expectedErr: errors.ErrInternalServerError,
			beforeTest: func(mock *mocks.FavoriteUsecase) {
				mock.On("SaveUnsaveFavoriteCourse", 1, 1, "save").Return(errors.ErrInternalServerError)
			},
		},
	}

	for scenario, test := range tests {
		t.Run(scenario, func(t *testing.T) {
			jsonResult, _ := json.Marshal(test.expectedRes)
			mock := mocks.NewFavoriteUsecase(t)
			test.beforeTest(mock)
			cfg := &server.RouterConfig{
				FavoriteUsecase: mock,
			}

			req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/api/v1/favorites/%s/save", test.param), nil)
			_, rec := testutils.ServeReq(cfg, req)

			assert.Equal(t, test.code, rec.Code)
			assert.Equal(t, string(jsonResult), rec.Body.String())
		})
	}
}
