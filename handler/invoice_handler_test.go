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
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetInvoices(t *testing.T) {
	tests := map[string]struct {
		code        int
		expectedRes gin.H
		beforeTest  func(*mocks.InvoiceUsecase)
	}{
		"should return all invoices when success": {
			code: http.StatusOK,
			expectedRes: gin.H{
				"data": []*entity.Invoice{
					{
						UserId: 1,
						User: &entity.User{
							ID: 1,
						},
					},
				},
				"totalRows":  1,
				"totalPages": 1,
			},
			beforeTest: func(mockUsecase *mocks.InvoiceUsecase) {
				mockUsecase.On("GetInvoices", mock.Anything).Return([]*entity.Invoice{
					{
						UserId: 1,
						User: &entity.User{
							ID: 1,
						},
					},
				}, int64(1), 1, nil)
			},
		},
		"should return internal server error when failed": {
			code: http.StatusInternalServerError,
			expectedRes: gin.H{
				"code":    errors.ErrCodeInternalServerError,
				"message": errors.ErrInternalServerError.Error(),
			},
			beforeTest: func(mockUsecase *mocks.InvoiceUsecase) {
				mockUsecase.On("GetInvoices", mock.Anything).Return(nil, int64(0), 0, errors.ErrInternalServerError)
			},
		},
	}

	for scenario, test := range tests {
		t.Run(scenario, func(t *testing.T) {
			jsonResult, _ := json.Marshal(test.expectedRes)
			mockUsecase := mocks.NewInvoiceUsecase(t)
			test.beforeTest(mockUsecase)
			cfg := &server.RouterConfig{
				InvoiceUsecase: mockUsecase,
			}

			req, _ := http.NewRequest(http.MethodGet, "/api/v1/invoices", nil)
			_, rec := testutils.ServeReq(cfg, req)

			assert.Equal(t, test.code, rec.Code)
			assert.Equal(t, string(jsonResult), rec.Body.String())
		})
	}
}

func TestGetInvoice(t *testing.T) {
	tests := map[string]struct {
		roleId      int
		code        int
		expectedRes gin.H
		beforeTest  func(*mocks.InvoiceUsecase)
	}{
		"should return invoice when success as user": {
			roleId: 2,
			code:   http.StatusOK,
			expectedRes: gin.H{
				"data": &entity.Invoice{
					UserId: 1,
					User: &entity.User{
						ID: 1,
					},
				},
			},
			beforeTest: func(mockUsecase *mocks.InvoiceUsecase) {
				mockUsecase.On("GetInvoiceDetail", mock.Anything, mock.Anything).Return(&entity.Invoice{
					UserId: 1,
					User: &entity.User{
						ID: 1,
					},
				}, nil)
			},
		},
		"should return invoice when success as admin": {
			roleId: 1,
			code:   http.StatusOK,
			expectedRes: gin.H{
				"data": &entity.Invoice{
					UserId: 1,
					User: &entity.User{
						ID: 1,
					},
				},
			},
			beforeTest: func(mockUsecase *mocks.InvoiceUsecase) {
				mockUsecase.On("GetInvoiceDetail", mock.Anything, mock.Anything).Return(&entity.Invoice{
					UserId: 1,
					User: &entity.User{
						ID: 1,
					},
				}, nil)
			},
		},
		"should return not found when invoice not found": {
			roleId: 2,
			code:   http.StatusNotFound,
			expectedRes: gin.H{
				"code":    errors.ErrCodeNotFound,
				"message": errors.ErrInvoiceNotFound.Error(),
			},
			beforeTest: func(mockUsecase *mocks.InvoiceUsecase) {
				mockUsecase.On("GetInvoiceDetail", mock.Anything, mock.Anything).Return(nil, errors.ErrInvoiceNotFound)
			},
		},
		"should return internal server error when failed": {
			roleId: 2,
			code:   http.StatusInternalServerError,
			expectedRes: gin.H{
				"code":    errors.ErrCodeInternalServerError,
				"message": errors.ErrInternalServerError.Error(),
			},
			beforeTest: func(mockUsecase *mocks.InvoiceUsecase) {
				mockUsecase.On("GetInvoiceDetail", mock.Anything, mock.Anything).Return(nil, errors.ErrInternalServerError)
			},
		},
	}

	for scenario, test := range tests {
		t.Run(scenario, func(t *testing.T) {
			jsonResult, _ := json.Marshal(test.expectedRes)
			mockUsecase := mocks.NewInvoiceUsecase(t)
			test.beforeTest(mockUsecase)
			cfg := &handler.Config{
				InvoiceUsecase: mockUsecase,
			}
			h := handler.New(cfg)
			rec := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(rec)
			c.Set("roleId", test.roleId)

			c.Request, _ = http.NewRequest(http.MethodGet, "/api/v1/invoices/1", nil)
			h.GetInvoice(c)

			assert.Equal(t, test.code, rec.Code)
			assert.Equal(t, string(jsonResult), rec.Body.String())
		})
	}
}

func TestCheckout(t *testing.T) {
	tests := map[string]struct {
		req         dto.CheckoutRequest
		code        int
		expectedRes gin.H
		beforeTest  func(*mocks.InvoiceUsecase)
	}{
		"should return invoice when success": {
			req: dto.CheckoutRequest{
				VoucherCode: "voucher",
			},
			code: http.StatusOK,
			expectedRes: gin.H{
				"data": &entity.Invoice{
					UserId: 1,
					User: &entity.User{
						ID: 1,
					},
				},
			},
			beforeTest: func(mockUsecase *mocks.InvoiceUsecase) {
				mockUsecase.On("Checkout", mock.Anything, mock.Anything, mock.Anything).Return(&entity.Invoice{
					UserId: 1,
					User: &entity.User{
						ID: 1,
					},
				}, nil)
			},
		},
		"should return not found when voucher not found": {
			req: dto.CheckoutRequest{
				VoucherCode: "voucher",
			},
			code: http.StatusBadRequest,
			expectedRes: gin.H{
				"code":    errors.ErrCodeBadRequest,
				"message": errors.ErrVoucherNotFound.Error(),
			},
			beforeTest: func(mockUsecase *mocks.InvoiceUsecase) {
				mockUsecase.On("Checkout", mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.ErrVoucherNotFound)
			},
		},
		"should return internal server error when failed": {
			req: dto.CheckoutRequest{
				VoucherCode: "voucher",
			},
			code: http.StatusInternalServerError,
			expectedRes: gin.H{
				"code":    errors.ErrCodeInternalServerError,
				"message": errors.ErrInternalServerError.Error(),
			},
			beforeTest: func(mockUsecase *mocks.InvoiceUsecase) {
				mockUsecase.On("Checkout", mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.ErrInternalServerError)
			},
		},
	}

	for scenario, test := range tests {
		t.Run(scenario, func(t *testing.T) {
			jsonResult, _ := json.Marshal(test.expectedRes)
			mockUsecase := mocks.NewInvoiceUsecase(t)
			test.beforeTest(mockUsecase)
			cfg := &server.RouterConfig{
				InvoiceUsecase: mockUsecase,
			}
			payload := testutils.MakeRequestBody(test.req)

			req, _ := http.NewRequest(http.MethodPost, "/api/v1/invoices", payload)
			_, rec := testutils.ServeReq(cfg, req)

			assert.Equal(t, test.code, rec.Code)
			assert.Equal(t, string(jsonResult), rec.Body.String())
		})
	}
}

func TestPayInvoice(t *testing.T) {
	tests := map[string]struct {
		code        int
		expectedRes gin.H
		beforeTest  func(*mocks.InvoiceUsecase)
	}{
		"should return invoice when success": {
			code: http.StatusOK,
			expectedRes: gin.H{
				"data": &entity.Invoice{
					UserId: 1,
					User: &entity.User{
						ID: 1,
					},
				},
			},
			beforeTest: func(mockUsecase *mocks.InvoiceUsecase) {
				mockUsecase.On("PayInvoice", mock.Anything, mock.Anything).Return(&entity.Invoice{
					UserId: 1,
					User: &entity.User{
						ID: 1,
					},
				}, nil)
			},
		},
		"should return not found when invoice not found": {
			code: http.StatusNotFound,
			expectedRes: gin.H{
				"code":    errors.ErrCodeNotFound,
				"message": errors.ErrInvoiceNotFound.Error(),
			},
			beforeTest: func(mockUsecase *mocks.InvoiceUsecase) {
				mockUsecase.On("PayInvoice", mock.Anything, mock.Anything).Return(nil, errors.ErrInvoiceNotFound)
			},
		},
		"should return bad request when invoice already paid": {
			code: http.StatusBadRequest,
			expectedRes: gin.H{
				"code":    errors.ErrCodeBadRequest,
				"message": errors.ErrInvoiceAlreadyPaid.Error(),
			},
			beforeTest: func(mockUsecase *mocks.InvoiceUsecase) {
				mockUsecase.On("PayInvoice", mock.Anything, mock.Anything).Return(nil, errors.ErrInvoiceAlreadyPaid)
			},
		},
		"should return internal server error when failed": {
			code: http.StatusInternalServerError,
			expectedRes: gin.H{
				"code":    errors.ErrCodeInternalServerError,
				"message": errors.ErrInternalServerError.Error(),
			},
			beforeTest: func(mockUsecase *mocks.InvoiceUsecase) {
				mockUsecase.On("PayInvoice", mock.Anything, mock.Anything).Return(nil, errors.ErrInternalServerError)
			},
		},
	}

	for scenario, test := range tests {
		t.Run(scenario, func(t *testing.T) {
			jsonResult, _ := json.Marshal(test.expectedRes)
			mockUsecase := mocks.NewInvoiceUsecase(t)
			test.beforeTest(mockUsecase)
			cfg := &server.RouterConfig{
				InvoiceUsecase: mockUsecase,
			}

			req, _ := http.NewRequest(http.MethodPost, "/api/v1/invoices/1/pay", nil)
			_, rec := testutils.ServeReq(cfg, req)

			assert.Equal(t, test.code, rec.Code)
			assert.Equal(t, string(jsonResult), rec.Body.String())
		})
	}
}

func TestConfirmInvoice(t *testing.T) {
	tests := map[string]struct {
		req         dto.ConfirmInvoiceRequest
		code        int
		expectedRes gin.H
		beforeTest  func(*mocks.InvoiceUsecase)
	}{
		"should return invoice when confirm success": {
			req: dto.ConfirmInvoiceRequest{
				ApprovalType: "confirm",
			},
			code: http.StatusOK,
			expectedRes: gin.H{
				"data": &entity.Invoice{
					UserId: 1,
					User: &entity.User{
						ID: 1,
					},
				},
			},
			beforeTest: func(mockUsecase *mocks.InvoiceUsecase) {
				mockUsecase.On("ConfirmInvoice", mock.Anything, mock.Anything).Return(&entity.Invoice{
					UserId: 1,
					User: &entity.User{
						ID: 1,
					},
				}, nil)
			},
		},
		"should return invoice when reject success": {
			req: dto.ConfirmInvoiceRequest{
				ApprovalType: "reject",
			},
			code: http.StatusOK,
			expectedRes: gin.H{
				"data": &entity.Invoice{
					UserId: 1,
					User: &entity.User{
						ID: 1,
					},
				},
			},
			beforeTest: func(mockUsecase *mocks.InvoiceUsecase) {
				mockUsecase.On("ConfirmInvoice", mock.Anything, mock.Anything).Return(&entity.Invoice{
					UserId: 1,
					User: &entity.User{
						ID: 1,
					},
				}, nil)
			},
		},
		"should return bad request when approval type is empty": {
			req:  dto.ConfirmInvoiceRequest{},
			code: http.StatusBadRequest,
			expectedRes: gin.H{
				"code":    errors.ErrCodeBadRequest,
				"message": errors.ErrInvalidBody.Error(),
			},
			beforeTest: func(mockUsecase *mocks.InvoiceUsecase) {},
		},
		"should return bad request when approval type is not confirm or reject": {
			req: dto.ConfirmInvoiceRequest{
				ApprovalType: "test",
			},
			code: http.StatusBadRequest,
			expectedRes: gin.H{
				"code":    errors.ErrCodeBadRequest,
				"message": errors.ErrInvalidInvoiceAction.Error(),
			},
			beforeTest: func(mockUsecase *mocks.InvoiceUsecase) {},
		},
		"should return not found when invoice not found": {
			req: dto.ConfirmInvoiceRequest{
				ApprovalType: "confirm",
			},
			code: http.StatusNotFound,
			expectedRes: gin.H{
				"code":    errors.ErrCodeNotFound,
				"message": errors.ErrInvoiceNotFound.Error(),
			},
			beforeTest: func(mockUsecase *mocks.InvoiceUsecase) {
				mockUsecase.On("ConfirmInvoice", mock.Anything, mock.Anything).Return(nil, errors.ErrInvoiceNotFound)
			},
		},
		"should return bad request when invoice already confirmed": {
			req: dto.ConfirmInvoiceRequest{
				ApprovalType: "confirm",
			},
			code: http.StatusBadRequest,
			expectedRes: gin.H{
				"code":    errors.ErrCodeBadRequest,
				"message": errors.ErrInvoiceStatusNotWaitingForConfirmation.Error(),
			},
			beforeTest: func(mockUsecase *mocks.InvoiceUsecase) {
				mockUsecase.On("ConfirmInvoice", mock.Anything, mock.Anything).Return(nil, errors.ErrInvoiceStatusNotWaitingForConfirmation)
			},
		},
		"should return internal server error when failed": {
			req: dto.ConfirmInvoiceRequest{
				ApprovalType: "confirm",
			},
			code: http.StatusInternalServerError,
			expectedRes: gin.H{
				"code":    errors.ErrCodeInternalServerError,
				"message": errors.ErrInternalServerError.Error(),
			},
			beforeTest: func(mockUsecase *mocks.InvoiceUsecase) {
				mockUsecase.On("ConfirmInvoice", mock.Anything, mock.Anything).Return(nil, errors.ErrInternalServerError)
			},
		},
	}

	for scenario, test := range tests {
		t.Run(scenario, func(t *testing.T) {
			jsonResult, _ := json.Marshal(test.expectedRes)
			mockUsecase := mocks.NewInvoiceUsecase(t)
			test.beforeTest(mockUsecase)
			cfg := &server.RouterConfig{
				InvoiceUsecase: mockUsecase,
			}
			payload := testutils.MakeRequestBody(test.req)

			req, _ := http.NewRequest(http.MethodPost, "/api/v1/invoices/1/confirm", payload)
			_, rec := testutils.ServeReq(cfg, req)

			assert.Equal(t, test.code, rec.Code)
			assert.Equal(t, string(jsonResult), rec.Body.String())
		})
	}
}
