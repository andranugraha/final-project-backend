package handler_test

import (
	"encoding/json"
	"final-project-backend/dto"
	"final-project-backend/entity"
	mocks "final-project-backend/mocks/usecase"
	"final-project-backend/server"
	"final-project-backend/testutils"
	"final-project-backend/utils/errors"
	"fmt"
	"mime/multipart"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateCourse(t *testing.T) {
	var (
		defaultReq = dto.CreateCourseRequest{
			Title:      "test",
			Summary:    "test",
			Content:    "test",
			AuthorName: "test",
			Status:     "test",
			CategoryId: 1,
			Tags:       []string{"test"},
			Price:      10000,
			Image:      multipart.FileHeader{},
		}
	)

	tests := map[string]struct {
		courseReq   map[string]string
		expectedRes gin.H
		code        int
		beforeTest  func(*mocks.CourseUsecase)
	}{
		"should return 200 when create course success": {
			courseReq: map[string]string{
				"title":      defaultReq.Title,
				"summary":    defaultReq.Summary,
				"content":    defaultReq.Content,
				"authorName": defaultReq.AuthorName,
				"status":     defaultReq.Status,
				"categoryId": fmt.Sprint(defaultReq.CategoryId),
				"tags":       defaultReq.Tags[0],
				"price":      fmt.Sprint(int(defaultReq.Price)),
			},
			expectedRes: gin.H{
				"data": &entity.Course{
					ID: 1,
				},
			},
			code: http.StatusOK,
			beforeTest: func(m *mocks.CourseUsecase) {
				m.On("CreateCourse", defaultReq).Return(&entity.Course{
					ID: 1,
				}, nil)
			},
		},
		"should return 400 when body invalid": {
			courseReq: map[string]string{},
			expectedRes: gin.H{
				"code":    errors.ErrCodeBadRequest,
				"message": errors.ErrInvalidBody.Error(),
			},
			code:       http.StatusBadRequest,
			beforeTest: func(m *mocks.CourseUsecase) {},
		},
		"should return 400 when title used": {
			courseReq: map[string]string{
				"title":      defaultReq.Title,
				"summary":    defaultReq.Summary,
				"content":    defaultReq.Content,
				"authorName": defaultReq.AuthorName,
				"status":     defaultReq.Status,
				"categoryId": fmt.Sprint(defaultReq.CategoryId),
				"tags":       defaultReq.Tags[0],
				"price":      fmt.Sprint(int(defaultReq.Price)),
			},
			expectedRes: gin.H{
				"code":    errors.ErrCodeDuplicate,
				"message": errors.ErrDuplicateTitle.Error(),
			},
			code: http.StatusBadRequest,
			beforeTest: func(m *mocks.CourseUsecase) {
				m.On("CreateCourse", defaultReq).Return(nil, errors.ErrDuplicateTitle)
			},
		},
		"should return 500 when create course failed": {
			courseReq: map[string]string{
				"title":      defaultReq.Title,
				"summary":    defaultReq.Summary,
				"content":    defaultReq.Content,
				"authorName": defaultReq.AuthorName,
				"status":     defaultReq.Status,
				"categoryId": fmt.Sprint(defaultReq.CategoryId),
				"tags":       defaultReq.Tags[0],
				"price":      fmt.Sprint(int(defaultReq.Price)),
			},
			expectedRes: gin.H{
				"code":    errors.ErrCodeInternalServerError,
				"message": errors.ErrInternalServerError.Error(),
			},
			code: http.StatusInternalServerError,
			beforeTest: func(m *mocks.CourseUsecase) {
				m.On("CreateCourse", defaultReq).Return(nil, errors.ErrInternalServerError)
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			jsonResult, _ := json.Marshal(test.expectedRes)
			mockUsecase := mocks.NewCourseUsecase(t)
			test.beforeTest(mockUsecase)
			cfg := &server.RouterConfig{
				CourseUsecase: mockUsecase,
			}
			payload, writer := testutils.MakeRequestBodyMultiPartFormData(test.courseReq)

			req, _ := http.NewRequest(http.MethodPost, "/api/v1/admin/courses", payload)
			req.Header.Set("Content-Type", writer.FormDataContentType())
			_, rec := testutils.ServeReq(cfg, req)

			assert.Equal(t, test.code, rec.Code)
			assert.Equal(t, string(jsonResult), rec.Body.String())
		})
	}
}

func TestGetCourses(t *testing.T) {
	tests := map[string]struct {
		tagId       string
		expectedRes gin.H
		code        int
		beforeTest  func(*mocks.CourseUsecase)
	}{
		"should return 200 when get courses success": {
			tagId: "1",
			expectedRes: gin.H{
				"data": []entity.Course{
					{
						ID: 1,
					},
				},
				"totalRows":  1,
				"totalPages": 1,
			},
			code: http.StatusOK,
			beforeTest: func(m *mocks.CourseUsecase) {
				m.On("GetCourses", entity.CourseParams{
					TagIds: []int{1},
					Sort:   "created_at DESC",
					Limit:  10,
					Page:   1,
					Status: "publish",
				}).Return([]entity.Course{
					{
						ID: 1,
					},
				}, int64(1), 1, nil)
			},
		},
		"should return 500 when get courses failed": {
			tagId: "1",
			expectedRes: gin.H{
				"code":    errors.ErrCodeInternalServerError,
				"message": errors.ErrInternalServerError.Error(),
			},
			code: http.StatusInternalServerError,
			beforeTest: func(m *mocks.CourseUsecase) {
				m.On("GetCourses", entity.CourseParams{
					TagIds: []int{1},
					Sort:   "created_at DESC",
					Limit:  10,
					Page:   1,
					Status: "publish",
				}).Return(nil, int64(0), 0, errors.ErrInternalServerError)
			},
		},
	}

	for scenario, test := range tests {
		t.Run(scenario, func(t *testing.T) {
			jsonResult, _ := json.Marshal(test.expectedRes)
			mockUsecase := mocks.NewCourseUsecase(t)
			test.beforeTest(mockUsecase)
			cfg := &server.RouterConfig{
				CourseUsecase: mockUsecase,
			}

			req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/courses?tagIds=%s", test.tagId), nil)
			_, rec := testutils.ServeReq(cfg, req)

			assert.Equal(t, test.code, rec.Code)
			assert.Equal(t, string(jsonResult), rec.Body.String())
		})
	}
}

func TestGetTrendingCourses(t *testing.T) {
	tests := map[string]struct {
		expectedRes gin.H
		code        int
		beforeTest  func(*mocks.CourseUsecase)
	}{
		"should return 200 when get courses success": {
			expectedRes: gin.H{
				"data": []entity.Course{
					{
						ID: 1,
					},
				},
			},
			code: http.StatusOK,
			beforeTest: func(m *mocks.CourseUsecase) {
				m.On("GetTrendingCourses").Return([]entity.Course{
					{
						ID: 1,
					},
				}, nil)
			},
		},
		"should return 500 when get courses failed": {
			expectedRes: gin.H{
				"code":    errors.ErrCodeInternalServerError,
				"message": errors.ErrInternalServerError.Error(),
			},
			code: http.StatusInternalServerError,
			beforeTest: func(m *mocks.CourseUsecase) {
				m.On("GetTrendingCourses").Return(nil, errors.ErrInternalServerError)
			},
		},
	}

	for scenario, test := range tests {
		t.Run(scenario, func(t *testing.T) {
			jsonResult, _ := json.Marshal(test.expectedRes)
			mockUsecase := mocks.NewCourseUsecase(t)
			test.beforeTest(mockUsecase)
			cfg := &server.RouterConfig{
				CourseUsecase: mockUsecase,
			}

			req, _ := http.NewRequest(http.MethodGet, "/api/v1/courses/trending", nil)
			_, rec := testutils.ServeReq(cfg, req)

			assert.Equal(t, test.code, rec.Code)
			assert.Equal(t, string(jsonResult), rec.Body.String())
		})
	}
}

func TestGetCourse(t *testing.T) {
	tests := map[string]struct {
		courseSlug  string
		expectedRes gin.H
		code        int
		beforeTest  func(*mocks.CourseUsecase)
	}{
		"should return 200 when get course success": {
			courseSlug: "1",
			expectedRes: gin.H{
				"data": &dto.GetCourseResponse{
					ID: 1,
				},
			},
			code: http.StatusOK,
			beforeTest: func(m *mocks.CourseUsecase) {
				m.On("GetCourse", "1", 1).Return(&dto.GetCourseResponse{
					ID: 1,
				}, nil)
			},
		},
		"should return 404 when course not found": {
			courseSlug: "1",
			expectedRes: gin.H{
				"code":    errors.ErrCodeNotFound,
				"message": errors.ErrCourseNotFound.Error(),
			},
			code: http.StatusNotFound,
			beforeTest: func(m *mocks.CourseUsecase) {
				m.On("GetCourse", "1", 1).Return(nil, errors.ErrCourseNotFound)
			},
		},
		"should return 500 when get course failed": {
			courseSlug: "1",
			expectedRes: gin.H{
				"code":    errors.ErrCodeInternalServerError,
				"message": errors.ErrInternalServerError.Error(),
			},
			code: http.StatusInternalServerError,
			beforeTest: func(m *mocks.CourseUsecase) {
				m.On("GetCourse", "1", 1).Return(nil, errors.ErrInternalServerError)
			},
		},
	}

	for scenario, test := range tests {
		t.Run(scenario, func(t *testing.T) {
			jsonResult, _ := json.Marshal(test.expectedRes)
			mockUsecase := mocks.NewCourseUsecase(t)
			test.beforeTest(mockUsecase)
			cfg := &server.RouterConfig{
				CourseUsecase: mockUsecase,
			}

			req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/courses/%s", test.courseSlug), nil)
			_, rec := testutils.ServeReq(cfg, req)

			assert.Equal(t, test.code, rec.Code)
			assert.Equal(t, string(jsonResult), rec.Body.String())
		})
	}
}

func TestUpdateCourse(t *testing.T) {
	var (
		defaultReq = dto.UpdateCourseRequest{
			Title:      "test",
			Summary:    "test",
			Content:    "test",
			AuthorName: "test",
			Status:     "test",
			CategoryId: 1,
			Tags:       []string{"test"},
			Price:      10000,
			Image:      nil,
		}
	)

	tests := map[string]struct {
		courseReq   map[string]string
		expectedRes gin.H
		code        int
		beforeTest  func(*mocks.CourseUsecase)
	}{
		"should return 200 when update course success": {
			courseReq: map[string]string{
				"title":      defaultReq.Title,
				"summary":    defaultReq.Summary,
				"content":    defaultReq.Content,
				"authorName": defaultReq.AuthorName,
				"status":     defaultReq.Status,
				"categoryId": fmt.Sprint(defaultReq.CategoryId),
				"tags":       defaultReq.Tags[0],
				"price":      fmt.Sprint(int(defaultReq.Price)),
			},
			expectedRes: gin.H{
				"data": &entity.Course{
					ID: 1,
				},
			},
			code: http.StatusOK,
			beforeTest: func(m *mocks.CourseUsecase) {
				m.On("UpdateCourse", "1", defaultReq).Return(&entity.Course{
					ID: 1,
				}, nil)
			},
		},
		"should return 400 when invalid request": {
			courseReq: map[string]string{},
			expectedRes: gin.H{
				"code":    errors.ErrCodeBadRequest,
				"message": errors.ErrInvalidBody.Error(),
			},
			code:       http.StatusBadRequest,
			beforeTest: func(m *mocks.CourseUsecase) {},
		},
		"should return 404 when course not found": {
			courseReq: map[string]string{
				"title":      defaultReq.Title,
				"summary":    defaultReq.Summary,
				"content":    defaultReq.Content,
				"authorName": defaultReq.AuthorName,
				"status":     defaultReq.Status,
				"categoryId": fmt.Sprint(defaultReq.CategoryId),
				"tags":       defaultReq.Tags[0],
				"price":      fmt.Sprint(int(defaultReq.Price)),
			},
			expectedRes: gin.H{
				"code":    errors.ErrCodeNotFound,
				"message": errors.ErrCourseNotFound.Error(),
			},
			code: http.StatusNotFound,
			beforeTest: func(m *mocks.CourseUsecase) {
				m.On("UpdateCourse", "1", defaultReq).Return(nil, errors.ErrCourseNotFound)
			},
		},
		"should return 400 when title used": {
			courseReq: map[string]string{
				"title":      defaultReq.Title,
				"summary":    defaultReq.Summary,
				"content":    defaultReq.Content,
				"authorName": defaultReq.AuthorName,
				"status":     defaultReq.Status,
				"categoryId": fmt.Sprint(defaultReq.CategoryId),
				"tags":       defaultReq.Tags[0],
				"price":      fmt.Sprint(int(defaultReq.Price)),
			},
			expectedRes: gin.H{
				"code":    errors.ErrCodeDuplicate,
				"message": errors.ErrDuplicateTitle.Error(),
			},
			code: http.StatusBadRequest,
			beforeTest: func(m *mocks.CourseUsecase) {
				m.On("UpdateCourse", "1", defaultReq).Return(nil, errors.ErrDuplicateTitle)
			},
		},
		"should return 500 when update course failed": {
			courseReq: map[string]string{
				"title":      defaultReq.Title,
				"summary":    defaultReq.Summary,
				"content":    defaultReq.Content,
				"authorName": defaultReq.AuthorName,
				"status":     defaultReq.Status,
				"categoryId": fmt.Sprint(defaultReq.CategoryId),
				"tags":       defaultReq.Tags[0],
				"price":      fmt.Sprint(int(defaultReq.Price)),
			},
			expectedRes: gin.H{
				"code":    errors.ErrCodeInternalServerError,
				"message": errors.ErrInternalServerError.Error(),
			},
			code: http.StatusInternalServerError,
			beforeTest: func(m *mocks.CourseUsecase) {
				m.On("UpdateCourse", "1", defaultReq).Return(nil, errors.ErrInternalServerError)
			},
		},
	}

	for scenario, test := range tests {
		t.Run(scenario, func(t *testing.T) {
			jsonResult, _ := json.Marshal(test.expectedRes)
			mockUsecase := mocks.NewCourseUsecase(t)
			test.beforeTest(mockUsecase)
			cfg := &server.RouterConfig{
				CourseUsecase: mockUsecase,
			}
			payload, writer := testutils.MakeRequestBodyMultiPartFormData(test.courseReq)

			req, _ := http.NewRequest(http.MethodPut, "/api/v1/admin/courses/1", payload)
			req.Header.Set("Content-Type", writer.FormDataContentType())
			_, rec := testutils.ServeReq(cfg, req)

			assert.Equal(t, test.code, rec.Code)
			assert.Equal(t, string(jsonResult), rec.Body.String())
		})
	}
}

func TestDeleteCourse(t *testing.T) {
	tests := map[string]struct {
		courseId    string
		expectedRes gin.H
		code        int
		beforeTest  func(m *mocks.CourseUsecase)
	}{
		"should return 200 when delete course success": {
			courseId: "1",
			expectedRes: gin.H{
				"message": "Course deleted successfully",
				"data":    nil,
			},
			code: http.StatusOK,
			beforeTest: func(m *mocks.CourseUsecase) {
				m.On("DeleteCourse", "1").Return(nil)
			},
		},
		"should return 404 when course not found": {
			courseId: "1",
			expectedRes: gin.H{
				"code":    errors.ErrCodeNotFound,
				"message": errors.ErrCourseNotFound.Error(),
			},
			code: http.StatusNotFound,
			beforeTest: func(m *mocks.CourseUsecase) {
				m.On("DeleteCourse", "1").Return(errors.ErrCourseNotFound)
			},
		},
		"should return 500 when delete course failed": {
			courseId: "1",
			expectedRes: gin.H{
				"code":    errors.ErrCodeInternalServerError,
				"message": errors.ErrInternalServerError.Error(),
			},
			code: http.StatusInternalServerError,
			beforeTest: func(m *mocks.CourseUsecase) {
				m.On("DeleteCourse", "1").Return(errors.ErrInternalServerError)
			},
		},
	}

	for scenario, test := range tests {
		t.Run(scenario, func(t *testing.T) {
			jsonResult, _ := json.Marshal(test.expectedRes)
			mockUsecase := mocks.NewCourseUsecase(t)
			test.beforeTest(mockUsecase)
			cfg := &server.RouterConfig{
				CourseUsecase: mockUsecase,
			}

			req, _ := http.NewRequest(http.MethodDelete, "/api/v1/admin/courses/1", nil)
			_, rec := testutils.ServeReq(cfg, req)

			assert.Equal(t, test.code, rec.Code)
			assert.Equal(t, string(jsonResult), rec.Body.String())
		})
	}
}

func TestGetUserCourses(t *testing.T) {
	var (
		userId = 1
	)

	tests := map[string]struct {
		tagId       string
		expectedRes gin.H
		code        int
		beforeTest  func(*mocks.CourseUsecase)
	}{
		"should return 200 when get user courses success": {
			tagId: "1",
			expectedRes: gin.H{
				"data": []entity.Course{
					{
						ID: 1,
					},
				},
				"totalRows":  1,
				"totalPages": 1,
			},
			code: http.StatusOK,
			beforeTest: func(m *mocks.CourseUsecase) {
				m.On("GetUserCourses", userId, entity.CourseParams{
					TagIds: []int{1},
					Sort:   "created_at DESC",
					Limit:  10,
					Page:   1,
					Status: "publish",
				}).Return([]entity.Course{
					{
						ID: 1,
					},
				}, int64(1), 1, nil)
			},
		},
		"should return 500 when get courses failed": {
			tagId: "1",
			expectedRes: gin.H{
				"code":    errors.ErrCodeInternalServerError,
				"message": errors.ErrInternalServerError.Error(),
			},
			code: http.StatusInternalServerError,
			beforeTest: func(m *mocks.CourseUsecase) {
				m.On("GetUserCourses", userId, entity.CourseParams{
					TagIds: []int{1},
					Sort:   "created_at DESC",
					Limit:  10,
					Page:   1,
					Status: "publish",
				}).Return(nil, int64(0), 0, errors.ErrInternalServerError)
			},
		},
	}

	for scenario, test := range tests {
		t.Run(scenario, func(t *testing.T) {
			jsonResult, _ := json.Marshal(test.expectedRes)
			mockUsecase := mocks.NewCourseUsecase(t)
			test.beforeTest(mockUsecase)
			cfg := &server.RouterConfig{
				CourseUsecase: mockUsecase,
			}

			req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/user/courses?tagIds=%s", test.tagId), nil)
			_, rec := testutils.ServeReq(cfg, req)

			assert.Equal(t, test.code, rec.Code)
			assert.Equal(t, string(jsonResult), rec.Body.String())
		})
	}
}

func TestGetUserCourse(t *testing.T) {
	var (
		userId = 1
	)

	tests := map[string]struct {
		courseId    string
		expectedRes gin.H
		code        int
		beforeTest  func(*mocks.CourseUsecase)
	}{
		"should return 200 when get user course success": {
			courseId: "1",
			expectedRes: gin.H{
				"data": entity.UserCourse{
					ID: 1,
				},
			},
			code: http.StatusOK,
			beforeTest: func(m *mocks.CourseUsecase) {
				m.On("GetUserCourse", userId, "1").Return(&entity.UserCourse{
					ID: 1,
				}, nil)
			},
		},
		"should return 404 when course not found": {
			courseId: "1",
			expectedRes: gin.H{
				"code":    errors.ErrCodeNotFound,
				"message": errors.ErrCourseNotFound.Error(),
			},
			code: http.StatusNotFound,
			beforeTest: func(m *mocks.CourseUsecase) {
				m.On("GetUserCourse", userId, "1").Return(nil, errors.ErrCourseNotFound)
			},
		},
		"should return 500 when get course failed": {
			courseId: "1",
			expectedRes: gin.H{
				"code":    errors.ErrCodeInternalServerError,
				"message": errors.ErrInternalServerError.Error(),
			},
			code: http.StatusInternalServerError,
			beforeTest: func(m *mocks.CourseUsecase) {
				m.On("GetUserCourse", userId, "1").Return(nil, errors.ErrInternalServerError)
			},
		},
	}

	for scenario, test := range tests {
		t.Run(scenario, func(t *testing.T) {
			jsonResult, _ := json.Marshal(test.expectedRes)
			mockUsecase := mocks.NewCourseUsecase(t)
			test.beforeTest(mockUsecase)
			cfg := &server.RouterConfig{
				CourseUsecase: mockUsecase,
			}

			req, _ := http.NewRequest(http.MethodGet, "/api/v1/user/courses/1", nil)
			_, rec := testutils.ServeReq(cfg, req)

			assert.Equal(t, test.code, rec.Code)
			assert.Equal(t, string(jsonResult), rec.Body.String())
		})
	}
}

func TestCompleteCourse(t *testing.T) {
	var (
		userId = 1
	)

	tests := map[string]struct {
		courseId    string
		expectedRes gin.H
		code        int
		beforeTest  func(*mocks.CourseUsecase)
	}{
		"should return 200 when complete course success": {
			courseId: "1",
			expectedRes: gin.H{
				"message": "Course completed successfully",
				"data":    nil,
			},
			code: http.StatusOK,
			beforeTest: func(m *mocks.CourseUsecase) {
				m.On("CompleteCourse", userId, "1").Return(nil)
			},
		},
		"should return 404 when course not found": {
			courseId: "1",
			expectedRes: gin.H{
				"code":    errors.ErrCodeNotFound,
				"message": errors.ErrCourseNotFound.Error(),
			},
			code: http.StatusNotFound,
			beforeTest: func(m *mocks.CourseUsecase) {
				m.On("CompleteCourse", userId, "1").Return(errors.ErrCourseNotFound)
			},
		},
		"should return 500 when complete course failed": {
			courseId: "1",
			expectedRes: gin.H{
				"code":    errors.ErrCodeInternalServerError,
				"message": errors.ErrInternalServerError.Error(),
			},
			code: http.StatusInternalServerError,
			beforeTest: func(m *mocks.CourseUsecase) {
				m.On("CompleteCourse", userId, "1").Return(errors.ErrInternalServerError)
			},
		},
	}

	for scenario, test := range tests {
		t.Run(scenario, func(t *testing.T) {
			jsonResult, _ := json.Marshal(test.expectedRes)
			mockUsecase := mocks.NewCourseUsecase(t)
			test.beforeTest(mockUsecase)
			cfg := &server.RouterConfig{
				CourseUsecase: mockUsecase,
			}

			req, _ := http.NewRequest(http.MethodPost, "/api/v1/user/courses/1", nil)
			_, rec := testutils.ServeReq(cfg, req)

			assert.Equal(t, test.code, rec.Code)
			assert.Equal(t, string(jsonResult), rec.Body.String())
		})
	}
}
