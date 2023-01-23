package usecase

import (
	"final-project-backend/dto"
	"final-project-backend/entity"
	"final-project-backend/repository"
	"final-project-backend/utils/constant"
	errResp "final-project-backend/utils/errors"
)

type InvoiceUsecase interface {
	GetInvoices(userId int, params entity.InvoiceParams) ([]*entity.Invoice, int64, int, error)
	Checkout(userId int, req dto.CheckoutRequest) (*entity.Invoice, error)
	PayInvoice(userId int, invoiceId int) (*entity.Invoice, error)
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

func (u *invoiceUsecaseImpl) GetInvoices(userId int, params entity.InvoiceParams) ([]*entity.Invoice, int64, int, error) {
	return u.invoiceRepo.FindByUserId(userId, params)
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

	userVoucher, err := u.userVoucherUsecase.FindValidByCode(req.VoucherCode, userId)
	if err != nil {
		return nil, err
	}

	invoice := entity.Invoice{
		UserId:        userId,
		Status:        constant.InvoiceStatusWaitingPayment,
		UserVoucherId: userVoucher.ID,
		VoucherId:     userVoucher.VoucherId,
		Voucher:       &userVoucher.Voucher,
		Discount:      user.Level.Discount,
	}

	var subtotal float64
	for _, item := range cart {
		transaction := entity.Transaction{}
		transaction.FromCart(item)
		invoice.Transactions = append(invoice.Transactions, &transaction)
		subtotal += transaction.Price
	}

	invoice.Subtotal = subtotal
	invoice.Total = subtotal - invoice.Discount - userVoucher.Voucher.Amount

	return u.invoiceRepo.Insert(invoice)
}

func (u *invoiceUsecaseImpl) PayInvoice(userId int, invoiceId int) (*entity.Invoice, error) {
	invoice, err := u.invoiceRepo.FindById(invoiceId)
	if err != nil {
		return nil, err
	}

	if invoice.UserId != userId {
		return nil, errResp.ErrForbidden
	}

	if invoice.Status != constant.InvoiceStatusWaitingPayment {
		return nil, errResp.ErrInvoiceAlreadyPaid
	}

	invoice.Status = constant.InvoiceStatusWaitingConfirmation

	return u.invoiceRepo.Update(*invoice)
}
