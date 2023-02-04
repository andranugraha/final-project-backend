package usecase

import (
	"final-project-backend/dto"
	"final-project-backend/entity"
	"final-project-backend/repository"
	"final-project-backend/utils/constant"
	errResp "final-project-backend/utils/errors"
	"time"
)

type InvoiceUsecase interface {
	GetInvoices(params entity.InvoiceParams) ([]*entity.Invoice, int64, int, error)
	GetInvoiceDetail(userId int, invoiceId string) (*entity.Invoice, error)
	Checkout(userId int, req dto.CheckoutRequest) (*entity.Invoice, error)
	PayInvoice(userId int, invoiceId string) (*entity.Invoice, error)
	ConfirmInvoice(invoiceId, status string) (*entity.Invoice, error)
}

type invoiceUsecaseImpl struct {
	invoiceRepo        repository.InvoiceRepository
	cartRepo           repository.CartRepository
	userVoucherUsecase UserVoucherUsecase
	userUsecase        UserUsecase
}

type InvoiceUConfig struct {
	InvoiceRepo        repository.InvoiceRepository
	CartRepo           repository.CartRepository
	UserVoucherUsecase UserVoucherUsecase
	UserUsecase        UserUsecase
}

func NewInvoiceUsecase(cfg *InvoiceUConfig) InvoiceUsecase {
	return &invoiceUsecaseImpl{
		invoiceRepo:        cfg.InvoiceRepo,
		cartRepo:           cfg.CartRepo,
		userVoucherUsecase: cfg.UserVoucherUsecase,
		userUsecase:        cfg.UserUsecase,
	}
}

func (u *invoiceUsecaseImpl) GetInvoices(params entity.InvoiceParams) ([]*entity.Invoice, int64, int, error) {
	return u.invoiceRepo.FindAll(params)
}

func (u *invoiceUsecaseImpl) GetInvoiceDetail(userId int, invoiceId string) (*entity.Invoice, error) {
	if userId == 0 {
		return u.invoiceRepo.FindById(invoiceId)
	}

	return u.invoiceRepo.FindByIdAndUserId(invoiceId, userId)
}

func (u *invoiceUsecaseImpl) Checkout(userId int, req dto.CheckoutRequest) (*entity.Invoice, error) {
	cart, err := u.cartRepo.FindByUserId(userId)
	if err != nil {
		return nil, err
	}

	if len(cart) == 0 {
		return nil, errResp.ErrCartEmpty
	}

	user, err := u.userUsecase.GetUserDetail(userId)
	if err != nil {
		return nil, err
	}

	invoice := entity.Invoice{
		UserId: userId,
		Status: constant.InvoiceStatusWaitingPayment,
	}

	var subtotal float64
	for _, item := range cart {
		transaction := entity.Transaction{}
		transaction.FromCart(item)
		invoice.Transactions = append(invoice.Transactions, &transaction)
		subtotal += transaction.Price
	}

	invoice.Subtotal = subtotal
	total := subtotal

	if req.VoucherCode != "" {
		userVoucher, err := u.userVoucherUsecase.FindValidByCode(req.VoucherCode, userId)
		if err != nil {
			return nil, err
		}

		invoice.UserVoucherId = &userVoucher.ID
		invoice.VoucherId = &userVoucher.VoucherId
		invoice.Voucher = &userVoucher.Voucher

		total -= userVoucher.Voucher.Amount

		if total < 0 {
			total = 0
		}
	}

	if user.Level.Discount > 0 {
		discount := total * user.Level.Discount
		total = total - discount
		invoice.Discount = discount
	}

	invoice.Total = total

	return u.invoiceRepo.Insert(invoice)
}

func (u *invoiceUsecaseImpl) PayInvoice(userId int, invoiceId string) (*entity.Invoice, error) {
	invoice, err := u.invoiceRepo.FindByIdAndUserId(invoiceId, userId)
	if err != nil {
		return nil, err
	}

	if invoice.Status != constant.InvoiceStatusWaitingPayment {
		return nil, errResp.ErrInvoiceAlreadyPaid
	}

	invoice.Status = constant.InvoiceStatusWaitingConfirmation
	now := time.Now()
	invoice.PaymentDate = &now

	return u.invoiceRepo.Update(*invoice)
}

func (u *invoiceUsecaseImpl) ConfirmInvoice(invoiceId, status string) (*entity.Invoice, error) {
	invoice, err := u.invoiceRepo.FindById(invoiceId)
	if err != nil {
		return nil, err
	}

	if status != constant.InvoiceStatusCompleted && status != constant.InvoiceStatusCancelled {
		return nil, errResp.ErrInvalidInvoiceStatus
	}

	if invoice.Status != constant.InvoiceStatusWaitingConfirmation {
		return nil, errResp.ErrInvoiceStatusNotWaitingForConfirmation
	}

	invoice.Status = status

	return u.invoiceRepo.Update(*invoice)
}
