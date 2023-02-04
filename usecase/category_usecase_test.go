package usecase_test

import (
	"final-project-backend/entity"
	mocks "final-project-backend/mocks/repository"
	"final-project-backend/usecase"
	errResp "final-project-backend/utils/errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCategories(t *testing.T) {
	tests := map[string]struct {
		expectedRes []*entity.Category
		expectedErr error
	}{
		"should return no error when given valid request": {
			expectedRes: []*entity.Category{
				{
					ID:   1,
					Name: "Programming",
				},
			},
			expectedErr: nil,
		},
		"should return error when find all failed": {
			expectedRes: []*entity.Category{},
			expectedErr: errResp.ErrInternalServerError,
		},
	}

	for scenario, test := range tests {
		t.Run(scenario, func(t *testing.T) {
			mockRepo := mocks.NewCategoryRepository(t)
			mockRepo.On("FindAll").Return(test.expectedRes, test.expectedErr).Once()
			u := usecase.NewCategoryUsecase(&usecase.CategoryUConfig{
				CategoryRepo: mockRepo,
			})

			res, err := u.GetCategories()

			assert.Equal(t, test.expectedRes, res)
			assert.ErrorIs(t, test.expectedErr, err)
		})
	}
}
