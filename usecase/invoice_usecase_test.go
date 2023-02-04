package usecase_test

import (
	"final-project-backend/dto"
	"final-project-backend/entity"
	mocks "final-project-backend/mocks/repository"
	mockUsecases "final-project-backend/mocks/usecase"
	"final-project-backend/usecase"
	"final-project-backend/utils/constant"
	errResp "final-project-backend/utils/errors"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetInvoices(t *testing.T) {
	var (
		params = entity.InvoiceParams{
			UserId: 1,
		}
	)

	tests := map[string]struct {
		expectedRes        []*entity.Invoice
		expectedTotalRows  int64
		expectedTotalPages int
		expectedErr        error
	}{
		"should return invoices when given valid request": {
			expectedRes: []*entity.Invoice{
				{
					ID: uuid.New(),
				},
			},
			expectedTotalRows:  1,
			expectedTotalPages: 1,
			expectedErr:        nil,
		},
		"should return error when find failed": {
			expectedRes:        []*entity.Invoice{},
			expectedTotalRows:  0,
			expectedTotalPages: 0,
			expectedErr:        errResp.ErrInternalServerError,
		},
	}

	for scenario, test := range tests {
		t.Run(scenario, func(t *testing.T) {
			mockRepo := mocks.NewInvoiceRepository(t)
			mockRepo.On("FindAll", params).Return(test.expectedRes, test.expectedTotalRows, test.expectedTotalPages, test.expectedErr)
			u := usecase.NewInvoiceUsecase(&usecase.InvoiceUConfig{
				InvoiceRepo: mockRepo,
			})

			res, totalRows, totalPages, err := u.GetInvoices(params)

			assert.Equal(t, test.expectedRes, res)
			assert.Equal(t, test.expectedTotalRows, totalRows)
			assert.Equal(t, test.expectedTotalPages, totalPages)
			assert.ErrorIs(t, test.expectedErr, err)
		})
	}
}

func TestGetInvoiceDetail(t *testing.T) {
	var (
		userId      = 1
		adminUserId = 0
		invoiceId   = uuid.New()
	)

	tests := map[string]struct {
		userId      int
		expectedRes *entity.Invoice
		expectedErr error
		beforeTest  func(*mocks.InvoiceRepository)
	}{
		"should return invoice when get as user success": {
			userId: userId,
			expectedRes: &entity.Invoice{
				ID:     invoiceId,
				UserId: userId,
			},
			expectedErr: nil,
			beforeTest: func(mockRepo *mocks.InvoiceRepository) {
				mockRepo.On("FindByIdAndUserId", invoiceId.String(), userId).Return(&entity.Invoice{
					ID:     invoiceId,
					UserId: userId,
				}, nil)
			},
		},
		"should return invoice when get as admin success": {
			userId: adminUserId,
			expectedRes: &entity.Invoice{
				ID:     invoiceId,
				UserId: adminUserId,
			},
			expectedErr: nil,
			beforeTest: func(mockRepo *mocks.InvoiceRepository) {
				mockRepo.On("FindById", invoiceId.String()).Return(&entity.Invoice{
					ID:     invoiceId,
					UserId: adminUserId,
				}, nil)
			},
		},
		"should return error when get as user failed": {
			userId:      userId,
			expectedRes: nil,
			expectedErr: errResp.ErrInternalServerError,
			beforeTest: func(mockRepo *mocks.InvoiceRepository) {
				mockRepo.On("FindByIdAndUserId", invoiceId.String(), userId).Return(nil, errResp.ErrInternalServerError)
			},
		},
		"should return error when get as admin failed": {
			userId:      adminUserId,
			expectedRes: nil,
			expectedErr: errResp.ErrInternalServerError,
			beforeTest: func(mockRepo *mocks.InvoiceRepository) {
				mockRepo.On("FindById", invoiceId.String()).Return(nil, errResp.ErrInternalServerError)
			},
		},
	}

	for scenario, test := range tests {
		t.Run(scenario, func(t *testing.T) {
			mockRepo := mocks.NewInvoiceRepository(t)
			test.beforeTest(mockRepo)
			u := usecase.NewInvoiceUsecase(&usecase.InvoiceUConfig{
				InvoiceRepo: mockRepo,
			})

			res, err := u.GetInvoiceDetail(test.userId, invoiceId.String())

			assert.Equal(t, test.expectedRes, res)
			assert.ErrorIs(t, test.expectedErr, err)
		})
	}
}

func TestCheckout(t *testing.T) {
	var (
		userId = 1
		req    = dto.CheckoutRequest{
			VoucherCode: "VOUCHER_CODE",
		}
		user = &entity.User{
			ID: userId,
			Level: &entity.Level{
				Discount: 10,
			},
		}
		cart = []*entity.Cart{
			{
				ID:       1,
				UserId:   userId,
				CourseId: 1,
				Course: &entity.Course{
					Price: 10000,
				},
			},
		}
		userVoucher = &entity.UserVoucher{
			ID:        1,
			UserId:    userId,
			VoucherId: 1,
			Voucher: entity.Voucher{
				Amount: 100000,
			},
		}
		invoice = &entity.Invoice{
			ID:     uuid.New(),
			UserId: userId,
		}
	)

	tests := map[string]struct {
		expectedRes *entity.Invoice
		expectedErr error
		beforeTest  func(*mocks.CartRepository, *mocks.InvoiceRepository, *mockUsecases.UserUsecase, *mockUsecases.UserVoucherUsecase)
	}{
		"should return invoice when given valid request": {
			expectedRes: invoice,
			expectedErr: nil,
			beforeTest: func(mockCartRepo *mocks.CartRepository, mockInvoiceRepo *mocks.InvoiceRepository, mockUserUsecase *mockUsecases.UserUsecase, mockUserVoucherUsecase *mockUsecases.UserVoucherUsecase) {
				mockCartRepo.On("FindByUserId", userId).Return(cart, nil).Once()
				mockUserUsecase.On("GetUserDetail", userId).Return(user, nil).Once()
				mockUserVoucherUsecase.On("FindValidByCode", req.VoucherCode, userId).Return(userVoucher, nil).Once()
				mockInvoiceRepo.On("Insert", mock.Anything).Return(invoice, nil).Once()
			},
		},
		"should set total price to 0 when voucher discount is more than total price": {
			expectedRes: invoice,
			expectedErr: nil,
			beforeTest: func(mockCartRepo *mocks.CartRepository, mockInvoiceRepo *mocks.InvoiceRepository, mockUserUsecase *mockUsecases.UserUsecase, mockUserVoucherUsecase *mockUsecases.UserVoucherUsecase) {
				mockCartRepo.On("FindByUserId", userId).Return(cart, nil).Once()
				mockUserUsecase.On("GetUserDetail", userId).Return(user, nil).Once()
				mockUserVoucherUsecase.On("FindValidByCode", req.VoucherCode, userId).Return(userVoucher, nil).Once()
				mockInvoiceRepo.On("Insert", mock.Anything).Return(invoice, nil).Once()
			},
		},
		"should return error when find cart failed": {
			expectedRes: nil,
			expectedErr: errResp.ErrInternalServerError,
			beforeTest: func(mockCartRepo *mocks.CartRepository, mockInvoiceRepo *mocks.InvoiceRepository, mockUserUsecase *mockUsecases.UserUsecase, mockUserVoucherUsecase *mockUsecases.UserVoucherUsecase) {
				mockCartRepo.On("FindByUserId", userId).Return(nil, errResp.ErrInternalServerError).Once()
			},
		},
		"should return error when cart is empty": {
			expectedRes: nil,
			expectedErr: errResp.ErrCartEmpty,
			beforeTest: func(mockCartRepo *mocks.CartRepository, mockInvoiceRepo *mocks.InvoiceRepository, mockUserUsecase *mockUsecases.UserUsecase, mockUserVoucherUsecase *mockUsecases.UserVoucherUsecase) {
				mockCartRepo.On("FindByUserId", userId).Return(nil, nil).Once()
			},
		},
		"should return error when find user failed": {
			expectedRes: nil,
			expectedErr: errResp.ErrInternalServerError,
			beforeTest: func(mockCartRepo *mocks.CartRepository, mockInvoiceRepo *mocks.InvoiceRepository, mockUserUsecase *mockUsecases.UserUsecase, mockUserVoucherUsecase *mockUsecases.UserVoucherUsecase) {
				mockCartRepo.On("FindByUserId", userId).Return(cart, nil).Once()
				mockUserUsecase.On("GetUserDetail", userId).Return(nil, errResp.ErrInternalServerError).Once()
			},
		},
		"should return error when find user voucher failed": {
			expectedRes: nil,
			expectedErr: errResp.ErrInternalServerError,
			beforeTest: func(mockCartRepo *mocks.CartRepository, mockInvoiceRepo *mocks.InvoiceRepository, mockUserUsecase *mockUsecases.UserUsecase, mockUserVoucherUsecase *mockUsecases.UserVoucherUsecase) {
				mockCartRepo.On("FindByUserId", userId).Return(cart, nil).Once()
				mockUserUsecase.On("GetUserDetail", userId).Return(user, nil).Once()
				mockUserVoucherUsecase.On("FindValidByCode", req.VoucherCode, userId).Return(nil, errResp.ErrInternalServerError).Once()
			},
		},
		"should return error when insert invoice failed": {
			expectedRes: nil,
			expectedErr: errResp.ErrInternalServerError,
			beforeTest: func(mockCartRepo *mocks.CartRepository, mockInvoiceRepo *mocks.InvoiceRepository, mockUserUsecase *mockUsecases.UserUsecase, mockUserVoucherUsecase *mockUsecases.UserVoucherUsecase) {
				mockCartRepo.On("FindByUserId", userId).Return(cart, nil).Once()
				mockUserUsecase.On("GetUserDetail", userId).Return(user, nil).Once()
				mockUserVoucherUsecase.On("FindValidByCode", req.VoucherCode, userId).Return(userVoucher, nil).Once()
				mockInvoiceRepo.On("Insert", mock.Anything).Return(nil, errResp.ErrInternalServerError).Once()
			},
		},
	}

	for scenario, test := range tests {
		t.Run(scenario, func(t *testing.T) {
			mockCartRepo := mocks.NewCartRepository(t)
			mockInvoiceRepo := mocks.NewInvoiceRepository(t)
			mockUserUsecase := mockUsecases.NewUserUsecase(t)
			mockUserVoucherUsecase := mockUsecases.NewUserVoucherUsecase(t)

			test.beforeTest(mockCartRepo, mockInvoiceRepo, mockUserUsecase, mockUserVoucherUsecase)

			u := usecase.NewInvoiceUsecase(&usecase.InvoiceUConfig{
				CartRepo:           mockCartRepo,
				InvoiceRepo:        mockInvoiceRepo,
				UserUsecase:        mockUserUsecase,
				UserVoucherUsecase: mockUserVoucherUsecase,
			})

			res, err := u.Checkout(userId, req)

			assert.Equal(t, test.expectedRes, res)
			assert.ErrorIs(t, test.expectedErr, err)
		})
	}
}

func TestPayInvoice(t *testing.T) {
	var (
		userId    = 1
		invoiceId = uuid.New()
	)

	tests := map[string]struct {
		expectedRes *entity.Invoice
		expectedErr error
		beforeTest  func(mockInvoiceRepo *mocks.InvoiceRepository)
	}{
		"should return invoice when pay invoice success": {
			expectedRes: &entity.Invoice{
				ID:     invoiceId,
				Status: constant.InvoiceStatusWaitingConfirmation,
			},
			expectedErr: nil,
			beforeTest: func(mockInvoiceRepo *mocks.InvoiceRepository) {
				mockInvoiceRepo.On("FindByIdAndUserId", invoiceId.String(), userId).Return(&entity.Invoice{
					ID:     invoiceId,
					Status: constant.InvoiceStatusWaitingPayment,
				}, nil).Once()
				mockInvoiceRepo.On("Update", mock.Anything).Return(&entity.Invoice{
					ID:     invoiceId,
					Status: constant.InvoiceStatusWaitingConfirmation,
				}, nil).Once()
			},
		},
		"should return error when find invoice failed": {
			expectedRes: nil,
			expectedErr: errResp.ErrInternalServerError,
			beforeTest: func(mockInvoiceRepo *mocks.InvoiceRepository) {
				mockInvoiceRepo.On("FindByIdAndUserId", invoiceId.String(), userId).Return(nil, errResp.ErrInternalServerError).Once()
			},
		},
		"should return error when invoice status is not waiting payment": {
			expectedRes: nil,
			expectedErr: errResp.ErrInvoiceAlreadyPaid,
			beforeTest: func(mockInvoiceRepo *mocks.InvoiceRepository) {
				mockInvoiceRepo.On("FindByIdAndUserId", invoiceId.String(), userId).Return(&entity.Invoice{
					ID:     invoiceId,
					Status: constant.InvoiceStatusWaitingConfirmation,
				}, nil).Once()
			},
		},
		"should return error when update invoice failed": {
			expectedRes: nil,
			expectedErr: errResp.ErrInternalServerError,
			beforeTest: func(mockInvoiceRepo *mocks.InvoiceRepository) {
				mockInvoiceRepo.On("FindByIdAndUserId", invoiceId.String(), userId).Return(&entity.Invoice{
					ID:     invoiceId,
					Status: constant.InvoiceStatusWaitingPayment,
				}, nil).Once()
				mockInvoiceRepo.On("Update", mock.Anything).Return(nil, errResp.ErrInternalServerError).Once()
			},
		},
	}

	for scenario, test := range tests {
		t.Run(scenario, func(t *testing.T) {
			mockInvoiceRepo := mocks.NewInvoiceRepository(t)
			test.beforeTest(mockInvoiceRepo)
			u := usecase.NewInvoiceUsecase(&usecase.InvoiceUConfig{
				InvoiceRepo: mockInvoiceRepo,
			})

			res, err := u.PayInvoice(userId, invoiceId.String())

			assert.Equal(t, test.expectedRes, res)
			assert.ErrorIs(t, test.expectedErr, err)
		})
	}
}

func TestConfirmInvoice(t *testing.T) {
	var (
		invoiceId = uuid.New()
		status    = constant.InvoiceStatusCompleted
	)

	tests := map[string]struct {
		status      string
		expectedRes *entity.Invoice
		expectedErr error
		beforeTest  func(mockInvoiceRepo *mocks.InvoiceRepository)
	}{
		"should return invoice when confirm invoice success": {
			status: status,
			expectedRes: &entity.Invoice{
				ID:     invoiceId,
				Status: status,
			},
			expectedErr: nil,
			beforeTest: func(mockInvoiceRepo *mocks.InvoiceRepository) {
				mockInvoiceRepo.On("FindById", invoiceId.String()).Return(&entity.Invoice{
					ID:     invoiceId,
					Status: constant.InvoiceStatusWaitingConfirmation,
				}, nil).Once()
				mockInvoiceRepo.On("Update", mock.Anything).Return(&entity.Invoice{
					ID:     invoiceId,
					Status: status,
				}, nil).Once()
			},
		},
		"should return error when find invoice failed": {
			status:      status,
			expectedRes: nil,
			expectedErr: errResp.ErrInternalServerError,
			beforeTest: func(mockInvoiceRepo *mocks.InvoiceRepository) {
				mockInvoiceRepo.On("FindById", invoiceId.String()).Return(nil, errResp.ErrInternalServerError).Once()
			},
		},
		"should return error when input status is not completed or cancelled": {
			status:      "invalid",
			expectedRes: nil,
			expectedErr: errResp.ErrInvalidInvoiceStatus,
			beforeTest: func(mockInvoiceRepo *mocks.InvoiceRepository) {
				mockInvoiceRepo.On("FindById", invoiceId.String()).Return(&entity.Invoice{
					ID:     invoiceId,
					Status: constant.InvoiceStatusWaitingConfirmation,
				}, nil).Once()
			},
		},
		"should return error when invoice status is not waiting confirmation": {
			status:      status,
			expectedRes: nil,
			expectedErr: errResp.ErrInvoiceStatusNotWaitingForConfirmation,
			beforeTest: func(mockInvoiceRepo *mocks.InvoiceRepository) {
				mockInvoiceRepo.On("FindById", invoiceId.String()).Return(&entity.Invoice{
					ID:     invoiceId,
					Status: status,
				}, nil).Once()
			},
		},
		"should return error when update invoice failed": {
			status:      status,
			expectedRes: nil,
			expectedErr: errResp.ErrInternalServerError,
			beforeTest: func(mockInvoiceRepo *mocks.InvoiceRepository) {
				mockInvoiceRepo.On("FindById", invoiceId.String()).Return(&entity.Invoice{
					ID:     invoiceId,
					Status: constant.InvoiceStatusWaitingConfirmation,
				}, nil).Once()
				mockInvoiceRepo.On("Update", mock.Anything).Return(nil, errResp.ErrInternalServerError).Once()
			},
		},
	}

	for scenario, test := range tests {
		t.Run(scenario, func(t *testing.T) {
			mockInvoiceRepo := mocks.NewInvoiceRepository(t)
			test.beforeTest(mockInvoiceRepo)
			u := usecase.NewInvoiceUsecase(&usecase.InvoiceUConfig{
				InvoiceRepo: mockInvoiceRepo,
			})

			res, err := u.ConfirmInvoice(invoiceId.String(), test.status)

			assert.Equal(t, test.expectedRes, res)
			assert.ErrorIs(t, test.expectedErr, err)
		})
	}
}
