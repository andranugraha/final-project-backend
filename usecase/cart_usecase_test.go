package usecase_test

import (
	"final-project-backend/dto"
	"final-project-backend/entity"
	mocks "final-project-backend/mocks/repository"
	"final-project-backend/usecase"
	errResp "final-project-backend/utils/errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCart(t *testing.T) {
	var (
		userId = 1
		cart   = []*entity.Cart{
			{
				UserId: userId,
				Course: &entity.Course{
					Price: 10000,
				},
			},
		}
	)

	tests := map[string]struct {
		findRes     []*entity.Cart
		expectedRes *dto.GetCartResponse
		expectedErr error
	}{
		"should return cart detail when given valid request": {
			findRes: cart,
			expectedRes: func() *dto.GetCartResponse {
				res := &dto.GetCartResponse{}
				res.FromCart(cart)
				return res
			}(),
			expectedErr: nil,
		},
		"should return error when find failed": {
			findRes:     []*entity.Cart{},
			expectedRes: nil,
			expectedErr: errResp.ErrInternalServerError,
		},
	}

	for scenario, test := range tests {
		t.Run(scenario, func(t *testing.T) {
			mockRepo := mocks.NewCartRepository(t)
			mockRepo.On("FindByUserId", userId).Return(test.findRes, test.expectedErr).Once()
			u := usecase.NewCartUsecase(&usecase.CartUConfig{
				CartRepo: mockRepo,
			})

			res, err := u.GetCart(userId)

			assert.Equal(t, test.expectedRes, res)
			assert.ErrorIs(t, test.expectedErr, err)
		})
	}
}

func TestAddToCart(t *testing.T) {
	var (
		userId   = 1
		courseId = 1
		cart     = entity.Cart{
			UserId:   userId,
			CourseId: courseId,
		}
	)

	tests := map[string]struct {
		expectedErr error
		beforeTest  func(*mocks.CartRepository, *mocks.CourseRepository, *mocks.TransactionRepository)
	}{
		"should return no error when given valid request": {
			expectedErr: nil,
			beforeTest: func(mock *mocks.CartRepository, mockCourse *mocks.CourseRepository, mockTransaction *mocks.TransactionRepository) {
				mockCourse.On("FindPublishedById", courseId).Return(nil, nil)
				mockTransaction.On("FindBoughtByUserIdAndCourseId", userId, courseId).Return(nil, nil)
				mock.On("Insert", cart).Return(nil)
			},
		},
		"should return error when find course failed": {
			expectedErr: errResp.ErrInternalServerError,
			beforeTest: func(mock *mocks.CartRepository, mockCourse *mocks.CourseRepository, mockTransaction *mocks.TransactionRepository) {
				mockCourse.On("FindPublishedById", courseId).Return(nil, errResp.ErrInternalServerError)
			},
		},
		"should return error when find transaction failed": {
			expectedErr: errResp.ErrInternalServerError,
			beforeTest: func(mock *mocks.CartRepository, mockCourse *mocks.CourseRepository, mockTransaction *mocks.TransactionRepository) {
				mockCourse.On("FindPublishedById", courseId).Return(nil, nil)
				mockTransaction.On("FindBoughtByUserIdAndCourseId", userId, courseId).Return(nil, errResp.ErrInternalServerError)
			},
		},
		"should return error when course already bought": {
			expectedErr: errResp.ErrCourseAlreadyBought,
			beforeTest: func(mock *mocks.CartRepository, mockCourse *mocks.CourseRepository, mockTransaction *mocks.TransactionRepository) {
				mockCourse.On("FindPublishedById", courseId).Return(nil, nil)
				mockTransaction.On("FindBoughtByUserIdAndCourseId", userId, courseId).Return(&entity.Transaction{
					CourseId: courseId,
				}, nil)
			},
		},
		"should return error when insert failed": {
			expectedErr: errResp.ErrInternalServerError,
			beforeTest: func(mock *mocks.CartRepository, mockCourse *mocks.CourseRepository, mockTransaction *mocks.TransactionRepository) {
				mockCourse.On("FindPublishedById", courseId).Return(nil, nil)
				mockTransaction.On("FindBoughtByUserIdAndCourseId", userId, courseId).Return(nil, nil)
				mock.On("Insert", cart).Return(errResp.ErrInternalServerError)
			},
		},
	}

	for scenario, test := range tests {
		t.Run(scenario, func(t *testing.T) {
			mockCourseRepo := mocks.NewCourseRepository(t)
			mockTransactionRepo := mocks.NewTransactionRepository(t)
			mockRepo := mocks.NewCartRepository(t)
			test.beforeTest(mockRepo, mockCourseRepo, mockTransactionRepo)
			u := usecase.NewCartUsecase(&usecase.CartUConfig{
				CourseRepo:      mockCourseRepo,
				TransactionRepo: mockTransactionRepo,
				CartRepo:        mockRepo,
			})

			err := u.AddToCart(userId, courseId)

			assert.ErrorIs(t, test.expectedErr, err)
		})
	}
}

func TestRemoveFromCart(t *testing.T) {
	var (
		userId   = 1
		courseId = 1
		cart     = entity.Cart{
			UserId:   userId,
			CourseId: courseId,
		}
	)

	tests := map[string]struct {
		expectedErr error
		beforeTest  func(*mocks.CartRepository)
	}{
		"should return no error when given valid request": {
			expectedErr: nil,
			beforeTest: func(mock *mocks.CartRepository) {
				mock.On("Delete", cart).Return(nil)
			},
		},
		"should return error when delete failed": {
			expectedErr: errResp.ErrInternalServerError,
			beforeTest: func(mock *mocks.CartRepository) {
				mock.On("Delete", cart).Return(errResp.ErrInternalServerError)
			},
		},
	}

	for scenario, test := range tests {
		t.Run(scenario, func(t *testing.T) {
			mockRepo := mocks.NewCartRepository(t)
			test.beforeTest(mockRepo)
			u := usecase.NewCartUsecase(&usecase.CartUConfig{
				CartRepo: mockRepo,
			})

			err := u.RemoveFromCart(userId, courseId)

			assert.ErrorIs(t, test.expectedErr, err)
		})
	}
}
