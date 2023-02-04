package usecase_test

import (
	"final-project-backend/entity"
	mocks "final-project-backend/mocks/repository"
	"final-project-backend/usecase"
	errResp "final-project-backend/utils/errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTags(t *testing.T) {
	tests := map[string]struct {
		expectedRes []*entity.Tag
		expectedErr error
	}{
		"should return all tags when success": {
			expectedRes: []*entity.Tag{
				{
					ID:   1,
					Name: "tag1",
				},
			},
			expectedErr: nil,
		},
		"should return error when failed": {
			expectedRes: nil,
			expectedErr: errResp.ErrInternalServerError,
		},
	}

	for scenario, test := range tests {
		t.Run(scenario, func(t *testing.T) {
			tagRepo := mocks.NewTagRepository(t)
			tagRepo.On("FindAll").Return(test.expectedRes, test.expectedErr)
			tagUsecase := usecase.NewTagUsecase(&usecase.TagUConfig{
				TagRepo: tagRepo,
			})

			res, err := tagUsecase.GetTags()

			assert.Equal(t, test.expectedRes, res)
			assert.ErrorIs(t, test.expectedErr, err)
		})
	}
}
