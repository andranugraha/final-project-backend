package usecase_test

import (
	"final-project-backend/entity"
	mocks "final-project-backend/mocks/repository"
	"final-project-backend/usecase"
	errResp "final-project-backend/utils/errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFavoriteCourses(t *testing.T) {
	var (
		userId  = 1
		courses = []*entity.Course{
			{
				ID: 1,
			},
		}
	)

	tests := map[string]struct {
		findRes     []*entity.Course
		expectedRes []*entity.Course
		expectedErr error
	}{
		"should return favorite courses when given valid request": {
			findRes:     courses,
			expectedRes: courses,
			expectedErr: nil,
		},
		"should return error when find failed": {
			findRes:     []*entity.Course{},
			expectedRes: []*entity.Course{},
			expectedErr: errResp.ErrInternalServerError,
		},
	}

	for scenario, test := range tests {
		t.Run(scenario, func(t *testing.T) {
			mockRepo := mocks.NewFavoriteRepository(t)
			mockRepo.On("FindByUserId", userId).Return(test.findRes, test.expectedErr).Once()
			u := usecase.NewFavoriteUsecase(&usecase.FavoriteUConfig{
				FavoriteRepo: mockRepo,
			})

			res, err := u.GetFavoriteCourses(userId)

			assert.Equal(t, test.expectedRes, res)
			assert.ErrorIs(t, test.expectedErr, err)
		})
	}
}

func TestCheckIsFavoriteCourse(t *testing.T) {
	var (
		userId   = 1
		courseId = 1
	)

	tests := map[string]struct {
		findRes     *entity.Favorite
		expectedRes bool
		expectedErr error
	}{
		"should return true when given valid request": {
			findRes:     &entity.Favorite{ID: courseId},
			expectedRes: true,
			expectedErr: nil,
		},
		"should return false when find failed": {
			findRes:     nil,
			expectedRes: false,
			expectedErr: errResp.ErrInternalServerError,
		},
	}

	for scenario, test := range tests {
		t.Run(scenario, func(t *testing.T) {
			mockRepo := mocks.NewFavoriteRepository(t)
			mockRepo.On("FindByUserIdAndCourseId", userId, courseId).Return(test.findRes, test.expectedErr).Once()
			u := usecase.NewFavoriteUsecase(&usecase.FavoriteUConfig{
				FavoriteRepo: mockRepo,
			})

			res := u.CheckIsFavoriteCourse(userId, courseId)

			assert.Equal(t, test.expectedRes, res)
		})
	}
}

func TestGetTotalFavorited(t *testing.T) {
	var (
		courseId = 1
	)

	tests := map[string]struct {
		countRes    int
		expectedRes int
		expectedErr error
	}{
		"should return total favorited when given valid request": {
			countRes:    1,
			expectedRes: 1,
			expectedErr: nil,
		},
	}

	for scenario, test := range tests {
		t.Run(scenario, func(t *testing.T) {
			mockRepo := mocks.NewFavoriteRepository(t)
			mockRepo.On("CountByCourseId", courseId).Return(test.countRes).Once()
			u := usecase.NewFavoriteUsecase(&usecase.FavoriteUConfig{
				FavoriteRepo: mockRepo,
			})

			res := u.GetTotalFavorited(courseId)

			assert.Equal(t, test.expectedRes, res)
		})
	}
}

func TestSaveFavoriteCourse(t *testing.T) {
	var (
		userId   = 1
		courseId = 1
		action   = "save"
	)

	tests := map[string]struct {
		expectedErr error
	}{
		"should return no error when save favorite course error": {
			expectedErr: nil,
		},
		"should return error when save failed": {
			expectedErr: errResp.ErrInternalServerError,
		},
	}

	for scenario, test := range tests {
		t.Run(scenario, func(t *testing.T) {
			mockRepo := mocks.NewFavoriteRepository(t)
			mockRepo.On("Insert", entity.Favorite{
				UserId:   userId,
				CourseId: courseId,
			}).Return(test.expectedErr).Once()
			u := usecase.NewFavoriteUsecase(&usecase.FavoriteUConfig{
				FavoriteRepo: mockRepo,
			})

			err := u.SaveUnsaveFavoriteCourse(userId, courseId, action)

			assert.ErrorIs(t, test.expectedErr, err)
		})
	}
}

func TestUnsaveFavoriteCourse(t *testing.T) {
	var (
		userId   = 1
		courseId = 1
		action   = "unsave"
	)

	tests := map[string]struct {
		expectedErr error
	}{
		"should return no error when unsave favorite course error": {
			expectedErr: nil,
		},
		"should return error when unsave favorite course failed": {
			expectedErr: errResp.ErrInternalServerError,
		},
	}

	for scenario, test := range tests {
		t.Run(scenario, func(t *testing.T) {
			mockRepo := mocks.NewFavoriteRepository(t)
			mockRepo.On("Delete", entity.Favorite{
				UserId:   userId,
				CourseId: courseId,
			}).Return(test.expectedErr).Once()
			u := usecase.NewFavoriteUsecase(&usecase.FavoriteUConfig{
				FavoriteRepo: mockRepo,
			})

			err := u.SaveUnsaveFavoriteCourse(userId, courseId, action)

			assert.ErrorIs(t, test.expectedErr, err)
		})
	}
}

func TestUnknownActionSaveUnsaveFavoriteCourse(t *testing.T) {
	t.Run("should return error when unknown action", func(t *testing.T) {
		mockRepo := mocks.NewFavoriteRepository(t)
		u := usecase.NewFavoriteUsecase(&usecase.FavoriteUConfig{
			FavoriteRepo: mockRepo,
		})

		err := u.SaveUnsaveFavoriteCourse(1, 1, "unknown")

		assert.ErrorIs(t, err, errResp.ErrUnknownAction)
	},
	)
}
