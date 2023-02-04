package usecase_test

import (
	"final-project-backend/dto"
	"final-project-backend/entity"
	mocks "final-project-backend/mocks/repository"
	mockUsecase "final-project-backend/mocks/usecase"
	mockStorageUtil "final-project-backend/mocks/utils/storage"
	"final-project-backend/usecase"
	"final-project-backend/utils/constant"
	errResp "final-project-backend/utils/errors"
	"final-project-backend/utils/storage"
	"mime/multipart"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetCourses(t *testing.T) {
	var (
		params = entity.CourseParams{
			Keyword:    "Go",
			CategoryId: 1,
			TagIds:     []int{1, 2, 3},
			Sort:       "",
			Limit:      10,
			Page:       1,
			Status:     constant.PublishStatus,
		}
	)

	tests := map[string]struct {
		expectedRes        []entity.Course
		expectedTotalRows  int64
		expectedTotalPages int
		expectedErr        error
	}{
		"should return no error when given valid request": {
			expectedRes: []entity.Course{
				{
					ID:    1,
					Title: "Go",
					Price: 10000,
					Category: &entity.Category{
						ID:   1,
						Name: "Programming",
					},
				},
			},
			expectedTotalRows:  1,
			expectedTotalPages: 1,
			expectedErr:        nil,
		},
		"should return error when find all failed": {
			expectedRes:        []entity.Course{},
			expectedTotalRows:  0,
			expectedTotalPages: 0,
			expectedErr:        errResp.ErrInternalServerError,
		},
	}

	for scenario, test := range tests {
		t.Run(scenario, func(t *testing.T) {
			mockRepo := mocks.NewCourseRepository(t)
			mockRepo.On("FindAll", params).Return(test.expectedRes, test.expectedTotalRows, test.expectedTotalPages, test.expectedErr).Once()
			u := usecase.NewCourseUsecase(&usecase.CourseUConfig{
				CourseRepo: mockRepo,
			})

			res, totalRows, totalPages, err := u.GetCourses(params)

			assert.Equal(t, test.expectedRes, res)
			assert.Equal(t, test.expectedTotalRows, totalRows)
			assert.Equal(t, test.expectedTotalPages, totalPages)
			assert.ErrorIs(t, test.expectedErr, err)
		})
	}
}

func TestGetCourse(t *testing.T) {
	var (
		slug   = "go"
		userId = 1
		course = entity.Course{
			ID:    1,
			Title: "Go",
			Price: 10000,
			Category: &entity.Category{
				ID:   1,
				Name: "Programming",
			},
		}
	)

	tests := map[string]struct {
		expectedRes *dto.GetCourseResponse
		expectedErr error
		beforeTest  func(*mocks.CourseRepository, *mocks.TransactionRepository, *mocks.UserCourseRepository, *mocks.CartRepository, *mockUsecase.FavoriteUsecase)
	}{
		"should return no error when given valid request": {
			expectedRes: &dto.GetCourseResponse{
				ID:         course.ID,
				Title:      course.Title,
				Price:      course.Price,
				Category:   course.Category,
				IsBought:   true,
				IsEnrolled: true,
			},
			expectedErr: nil,
			beforeTest: func(mockCourseRepo *mocks.CourseRepository, mockTransactionRepo *mocks.TransactionRepository, mockUserCourseRepo *mocks.UserCourseRepository, mockCartRepo *mocks.CartRepository, mockFavoriteUsecase *mockUsecase.FavoriteUsecase) {
				mockCourseRepo.On("FindBySlug", slug).Return(&course, nil).Once()
				mockTransactionRepo.On("FindBoughtByUserIdAndCourseId", course.ID, userId).Return(&entity.Transaction{}, nil).Once()
				mockUserCourseRepo.On("FindByUserIdAndCourseId", userId, course.ID).Return(&entity.UserCourse{}, nil).Once()
				mockCartRepo.On("FindByUserIdAndCourseId", userId, course.ID).Return(nil, errResp.ErrCartNotFound).Once()
				mockFavoriteUsecase.On("CheckIsFavoriteCourse", userId, course.ID).Return(false).Once()
				mockFavoriteUsecase.On("GetTotalFavorited", course.ID).Return(0).Once()
				mockUserCourseRepo.On("CountByCourseIdAndStatus", course.ID, constant.CourseStatusCompleted).Return(0).Once()
			},
		},
		"should return error when find by slug failed": {
			expectedRes: nil,
			expectedErr: errResp.ErrInternalServerError,
			beforeTest: func(mockCourseRepo *mocks.CourseRepository, mockTransactionRepo *mocks.TransactionRepository, mockUserCourseRepo *mocks.UserCourseRepository, mockCartRepo *mocks.CartRepository, mockFavoriteUsecase *mockUsecase.FavoriteUsecase) {
				mockCourseRepo.On("FindBySlug", slug).Return(nil, errResp.ErrInternalServerError).Once()
			},
		},
		"should return error when find bought by user id and course id failed": {
			expectedRes: nil,
			expectedErr: errResp.ErrInternalServerError,
			beforeTest: func(mockCourseRepo *mocks.CourseRepository, mockTransactionRepo *mocks.TransactionRepository, mockUserCourseRepo *mocks.UserCourseRepository, mockCartRepo *mocks.CartRepository, mockFavoriteUsecase *mockUsecase.FavoriteUsecase) {
				mockCourseRepo.On("FindBySlug", slug).Return(&course, nil).Once()
				mockTransactionRepo.On("FindBoughtByUserIdAndCourseId", course.ID, userId).Return(nil, errResp.ErrInternalServerError).Once()
			},
		},
		"should return error when find user course failed": {
			expectedRes: nil,
			expectedErr: errResp.ErrInternalServerError,
			beforeTest: func(mockCourseRepo *mocks.CourseRepository, mockTransactionRepo *mocks.TransactionRepository, mockUserCourseRepo *mocks.UserCourseRepository, mockCartRepo *mocks.CartRepository, mockFavoriteUsecase *mockUsecase.FavoriteUsecase) {
				mockCourseRepo.On("FindBySlug", slug).Return(&course, nil).Once()
				mockTransactionRepo.On("FindBoughtByUserIdAndCourseId", course.ID, userId).Return(&entity.Transaction{}, nil).Once()
				mockUserCourseRepo.On("FindByUserIdAndCourseId", userId, course.ID).Return(nil, errResp.ErrInternalServerError).Once()
			},
		},
		"should return error when find course in cart failed": {
			expectedRes: nil,
			expectedErr: errResp.ErrInternalServerError,
			beforeTest: func(mockCourseRepo *mocks.CourseRepository, mockTransactionRepo *mocks.TransactionRepository, mockUserCourseRepo *mocks.UserCourseRepository, mockCartRepo *mocks.CartRepository, mockFavoriteUsecase *mockUsecase.FavoriteUsecase) {
				mockCourseRepo.On("FindBySlug", slug).Return(&course, nil).Once()
				mockTransactionRepo.On("FindBoughtByUserIdAndCourseId", course.ID, userId).Return(&entity.Transaction{}, nil).Once()
				mockUserCourseRepo.On("FindByUserIdAndCourseId", userId, course.ID).Return(&entity.UserCourse{}, nil).Once()
				mockCartRepo.On("FindByUserIdAndCourseId", userId, course.ID).Return(nil, errResp.ErrInternalServerError).Once()
			},
		},
	}

	for scenario, test := range tests {
		t.Run(scenario, func(t *testing.T) {
			mockCourseRepo := new(mocks.CourseRepository)
			mockTransactionRepo := new(mocks.TransactionRepository)
			mockUserCourseRepo := new(mocks.UserCourseRepository)
			mockCartRepo := new(mocks.CartRepository)
			mockFavoriteUsecase := new(mockUsecase.FavoriteUsecase)

			test.beforeTest(mockCourseRepo, mockTransactionRepo, mockUserCourseRepo, mockCartRepo, mockFavoriteUsecase)

			uc := usecase.NewCourseUsecase(&usecase.CourseUConfig{
				CourseRepo:      mockCourseRepo,
				TransactionRepo: mockTransactionRepo,
				UserCourseRepo:  mockUserCourseRepo,
				CartRepo:        mockCartRepo,
				FavoriteUsecase: mockFavoriteUsecase,
			})

			res, err := uc.GetCourse(slug, userId)

			assert.Equal(t, test.expectedRes, res)
			assert.ErrorIs(t, test.expectedErr, err)
		})
	}
}

func TestGetTrendingCourses(t *testing.T) {
	tests := map[string]struct {
		expectedRes []entity.Course
		expectedErr error
	}{
		"should return trending courses": {
			expectedRes: []entity.Course{
				{
					ID:    1,
					Title: "title",
					Slug:  "slug",
					Price: 10000,
				},
			},
			expectedErr: nil,
		},
		"should return error when find trending courses failed": {
			expectedRes: nil,
			expectedErr: errResp.ErrInternalServerError,
		},
	}

	for scenario, test := range tests {
		t.Run(scenario, func(t *testing.T) {
			mockCourseRepo := new(mocks.CourseRepository)

			mockCourseRepo.On("FindTrending").Return(test.expectedRes, test.expectedErr).Once()

			uc := usecase.NewCourseUsecase(&usecase.CourseUConfig{
				CourseRepo: mockCourseRepo,
			})

			res, err := uc.GetTrendingCourses()

			assert.Equal(t, test.expectedRes, res)
			assert.ErrorIs(t, test.expectedErr, err)
		})
	}
}

func TestGetUserCourses(t *testing.T) {
	var (
		userId = 1
		params = entity.CourseParams{
			Keyword:        "",
			CategoryId:     0,
			TagIds:         []int{},
			Sort:           "latest",
			Page:           1,
			Limit:          10,
			Status:         constant.PublishStatus,
			ProgressStatus: "",
		}
	)

	tests := map[string]struct {
		expectedRes        []entity.Course
		expectedTotalRows  int64
		expectedTotalPages int
		expectedErr        error
	}{
		"should return user courses": {
			expectedRes: []entity.Course{
				{
					ID:    1,
					Title: "title",
					Slug:  "slug",
					Price: 10000,
				},
			},
			expectedTotalRows:  1,
			expectedTotalPages: 1,
			expectedErr:        nil,
		},
		"should return error when find user courses failed": {
			expectedRes:        nil,
			expectedTotalRows:  0,
			expectedTotalPages: 0,
			expectedErr:        errResp.ErrInternalServerError,
		},
	}

	for scenario, test := range tests {
		t.Run(scenario, func(t *testing.T) {
			mockUserCourseRepo := new(mocks.UserCourseRepository)
			mockUserCourseRepo.On("FindByUserId", userId, params).Return(test.expectedRes, test.expectedTotalRows, test.expectedTotalPages, test.expectedErr).Once()
			uc := usecase.NewCourseUsecase(&usecase.CourseUConfig{
				UserCourseRepo: mockUserCourseRepo,
			})

			res, totalRows, totalPages, err := uc.GetUserCourses(userId, params)

			assert.Equal(t, test.expectedRes, res)
			assert.Equal(t, test.expectedTotalRows, totalRows)
			assert.Equal(t, test.expectedTotalPages, totalPages)
			assert.ErrorIs(t, test.expectedErr, err)
		})
	}
}

func TestGetUserCourse(t *testing.T) {
	var (
		userId = 1
		slug   = "slug"
	)

	tests := map[string]struct {
		expectedRes *entity.UserCourse
		expectedErr error
	}{
		"should return user course": {
			expectedRes: &entity.UserCourse{
				ID:       1,
				UserId:   userId,
				CourseId: 1,
			},
			expectedErr: nil,
		},
		"should return error when find user course failed": {
			expectedRes: nil,
			expectedErr: errResp.ErrInternalServerError,
		},
	}

	for scenario, test := range tests {
		t.Run(scenario, func(t *testing.T) {
			mockUserCourseRepo := new(mocks.UserCourseRepository)
			mockUserCourseRepo.On("FindByUserIdAndCourseSlug", userId, slug).Return(test.expectedRes, test.expectedErr).Once()
			uc := usecase.NewCourseUsecase(&usecase.CourseUConfig{
				UserCourseRepo: mockUserCourseRepo,
			})

			res, err := uc.GetUserCourse(userId, slug)

			assert.Equal(t, test.expectedRes, res)
			assert.ErrorIs(t, test.expectedErr, err)
		})
	}
}

func TestCreateCourse(t *testing.T) {
	var (
		req = dto.CreateCourseRequest{
			Title: "title",
			Image: multipart.FileHeader{},
		}
	)

	tests := map[string]struct {
		expectedRes *entity.Course
		expectedErr error
		beforeTest  func(*mocks.CourseRepository, *mockStorageUtil.StorageUtil)
	}{
		"should return error when upload to storage failed": {
			expectedRes: nil,
			expectedErr: errResp.ErrInternalServerError,
			beforeTest: func(mockCourseRepo *mocks.CourseRepository, mockStorageUtil *mockStorageUtil.StorageUtil) {
				mockStorageUtil.On("Upload", mock.Anything, mock.Anything).Return(nil, errResp.ErrInternalServerError).Once()
			},
		},
		"should return error when create course failed": {
			expectedRes: nil,
			expectedErr: errResp.ErrInternalServerError,
			beforeTest: func(mockCourseRepo *mocks.CourseRepository, mockStorageUtil *mockStorageUtil.StorageUtil) {
				mockStorageUtil.On("Upload", mock.Anything, mock.Anything).Return(&storage.StoredImage{}, nil).Once()
				mockCourseRepo.On("Insert", mock.Anything).Return(nil, errResp.ErrInternalServerError).Once()
			},
		},
		"should return course when create course success": {
			expectedRes: &entity.Course{
				ID:    1,
				Title: req.Title,
				Slug:  req.Title,
			},
			expectedErr: nil,
			beforeTest: func(mockCourseRepo *mocks.CourseRepository, mockStorageUtil *mockStorageUtil.StorageUtil) {
				mockStorageUtil.On("Upload", mock.Anything, mock.Anything).Return(&storage.StoredImage{}, nil).Once()
				mockCourseRepo.On("Insert", mock.Anything).Return(&entity.Course{
					ID:    1,
					Title: req.Title,
					Slug:  req.Title,
				}, nil).Once()
			},
		},
	}

	for scenario, test := range tests {
		t.Run(scenario, func(t *testing.T) {
			mockCourseRepo := new(mocks.CourseRepository)
			mockStorageUtil := new(mockStorageUtil.StorageUtil)
			test.beforeTest(mockCourseRepo, mockStorageUtil)
			uc := usecase.NewCourseUsecase(&usecase.CourseUConfig{
				CourseRepo:  mockCourseRepo,
				StorageUtil: mockStorageUtil,
			})

			res, err := uc.CreateCourse(req)

			assert.Equal(t, test.expectedRes, res)
			assert.ErrorIs(t, test.expectedErr, err)
		})
	}
}

func TestUpdateCourse(t *testing.T) {
	var (
		slug   = "slug"
		course = entity.Course{
			ID:    1,
			Title: "title",
			Slug:  slug,
		}
	)

	tests := map[string]struct {
		image       *multipart.FileHeader
		expectedRes *entity.Course
		expectedErr error
		beforeTest  func(*mocks.CourseRepository, *mockStorageUtil.StorageUtil)
	}{
		"should return error when find course failed": {
			image:       nil,
			expectedRes: nil,
			expectedErr: errResp.ErrInternalServerError,
			beforeTest: func(mockCourseRepo *mocks.CourseRepository, mockStorageUtil *mockStorageUtil.StorageUtil) {
				mockCourseRepo.On("FindBySlug", slug).Return(nil, errResp.ErrInternalServerError).Once()
			},
		},
		"should return error when rename image failed": {
			image:       nil,
			expectedRes: nil,
			expectedErr: errResp.ErrInternalServerError,
			beforeTest: func(mockCourseRepo *mocks.CourseRepository, mockStorageUtil *mockStorageUtil.StorageUtil) {
				mockCourseRepo.On("FindBySlug", slug).Return(&entity.Course{
					Title: "test title",
				}, nil).Once()
				mockStorageUtil.On("Rename", mock.Anything, mock.Anything).Return(nil, errResp.ErrInternalServerError).Once()
			},
		},
		"should return error when upload to storage failed": {
			image:       &multipart.FileHeader{},
			expectedRes: nil,
			expectedErr: errResp.ErrInternalServerError,
			beforeTest: func(mockCourseRepo *mocks.CourseRepository, mockStorageUtil *mockStorageUtil.StorageUtil) {
				mockCourseRepo.On("FindBySlug", slug).Return(&course, nil).Once()
				mockStorageUtil.On("Delete", slug).Return(nil).Once()
				mockStorageUtil.On("Upload", mock.Anything, mock.Anything).Return(nil, errResp.ErrInternalServerError).Once()
			},
		},
		"should return error when update course failed": {
			image:       &multipart.FileHeader{},
			expectedRes: nil,
			expectedErr: errResp.ErrInternalServerError,
			beforeTest: func(mockCourseRepo *mocks.CourseRepository, mockStorageUtil *mockStorageUtil.StorageUtil) {
				mockCourseRepo.On("FindBySlug", slug).Return(&course, nil).Once()
				mockStorageUtil.On("Delete", slug).Return(nil).Once()
				mockStorageUtil.On("Upload", mock.Anything, mock.Anything).Return(&storage.StoredImage{}, nil).Once()
				mockCourseRepo.On("Update", mock.Anything).Return(nil, errResp.ErrInternalServerError).Once()
			},
		},
		"should return course when update course without image success": {
			image:       nil,
			expectedRes: &course,
			expectedErr: nil,
			beforeTest: func(mockCourseRepo *mocks.CourseRepository, mockStorageUtil *mockStorageUtil.StorageUtil) {
				mockCourseRepo.On("FindBySlug", slug).Return(&entity.Course{
					Title: "test title",
				}, nil).Once()
				mockStorageUtil.On("Rename", mock.Anything, mock.Anything).Return(&storage.StoredImage{}, nil).Once()
				mockCourseRepo.On("Update", mock.Anything).Return(&course, nil).Once()
			},
		},
		"should return course when update course with image success": {
			image:       &multipart.FileHeader{},
			expectedRes: &course,
			expectedErr: nil,
			beforeTest: func(mockCourseRepo *mocks.CourseRepository, mockStorageUtil *mockStorageUtil.StorageUtil) {
				mockCourseRepo.On("FindBySlug", slug).Return(&course, nil).Once()
				mockStorageUtil.On("Delete", slug).Return(nil).Once()
				mockStorageUtil.On("Upload", mock.Anything, mock.Anything).Return(&storage.StoredImage{}, nil).Once()
				mockCourseRepo.On("Update", mock.Anything).Return(&course, nil).Once()
			},
		},
	}

	for scenario, test := range tests {
		t.Run(scenario, func(t *testing.T) {
			mockCourseRepo := new(mocks.CourseRepository)
			mockStorageUtil := new(mockStorageUtil.StorageUtil)
			test.beforeTest(mockCourseRepo, mockStorageUtil)
			uc := usecase.NewCourseUsecase(&usecase.CourseUConfig{
				CourseRepo:  mockCourseRepo,
				StorageUtil: mockStorageUtil,
			})

			res, err := uc.UpdateCourse(slug, dto.UpdateCourseRequest{
				Title: "title",
				Image: test.image,
				Tags:  []string{"tag1", "tag2"},
			})

			assert.Equal(t, test.expectedRes, res)
			assert.ErrorIs(t, test.expectedErr, err)
		})
	}
}

func TestDeleteCourse(t *testing.T) {
	var (
		slug = "slug"
	)

	tests := map[string]struct {
		expectedErr error
	}{
		"should return error when delete course failed": {
			expectedErr: errResp.ErrInternalServerError,
		},
		"should return nil when delete course success": {
			expectedErr: nil,
		},
	}

	for scenario, test := range tests {
		t.Run(scenario, func(t *testing.T) {
			mockCourseRepo := new(mocks.CourseRepository)
			mockCourseRepo.On("Delete", slug).Return(test.expectedErr).Once()
			uc := usecase.NewCourseUsecase(&usecase.CourseUConfig{
				CourseRepo: mockCourseRepo,
			})

			err := uc.DeleteCourse(slug)

			assert.ErrorIs(t, test.expectedErr, err)
		})
	}
}

func TestCompleteCourse(t *testing.T) {
	var (
		userId     = 1
		slug       = "slug"
		userCourse = entity.UserCourse{
			Status: constant.CourseStatusCompleted,
		}
	)

	tests := map[string]struct {
		expectedErr error
		beforeTest  func(mockUserCourseRepo *mocks.UserCourseRepository)
	}{
		"should return error when find user course failed": {
			expectedErr: errResp.ErrInternalServerError,
			beforeTest: func(mockUserCourseRepo *mocks.UserCourseRepository) {
				mockUserCourseRepo.On("FindByUserIdAndCourseSlug", userId, slug).Return(nil, errResp.ErrInternalServerError).Once()
			},
		},
		"should return error when course already completed": {
			expectedErr: errResp.ErrCourseAlreadyCompleted,
			beforeTest: func(mockUserCourseRepo *mocks.UserCourseRepository) {
				mockUserCourseRepo.On("FindByUserIdAndCourseSlug", userId, slug).Return(&entity.UserCourse{
					Status: constant.CourseStatusCompleted,
				}, nil).Once()
			},
		},
		"should return error when complete user course failed": {
			expectedErr: errResp.ErrInternalServerError,
			beforeTest: func(mockUserCourseRepo *mocks.UserCourseRepository) {
				mockUserCourseRepo.On("FindByUserIdAndCourseSlug", userId, slug).Return(&entity.UserCourse{
					Status: constant.CourseStatusInProgress,
				}, nil).Once()
				mockUserCourseRepo.On("Complete", userCourse).Return(nil, errResp.ErrInternalServerError).Once()
			},
		},
		"should return nil when complete course success": {
			expectedErr: nil,
			beforeTest: func(mockUserCourseRepo *mocks.UserCourseRepository) {
				mockUserCourseRepo.On("FindByUserIdAndCourseSlug", userId, slug).Return(&entity.UserCourse{
					Status: constant.CourseStatusInProgress,
				}, nil).Once()
				mockUserCourseRepo.On("Complete", userCourse).Return(&userCourse, nil).Once()
			},
		},
	}

	for scenario, test := range tests {
		t.Run(scenario, func(t *testing.T) {
			mockUserCourseRepo := new(mocks.UserCourseRepository)
			test.beforeTest(mockUserCourseRepo)
			uc := usecase.NewCourseUsecase(&usecase.CourseUConfig{
				UserCourseRepo: mockUserCourseRepo,
			})

			err := uc.CompleteCourse(userId, slug)

			assert.ErrorIs(t, test.expectedErr, err)
		})
	}
}
