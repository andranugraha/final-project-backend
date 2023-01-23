package repository

import (
	"errors"
	"final-project-backend/entity"
	errResp "final-project-backend/utils/errors"
	"math"

	"gorm.io/gorm"
)

type InvoiceRepository interface {
	FindById(id int) (*entity.Invoice, error)
	FindByUserId(userId int, params entity.InvoiceParams) ([]*entity.Invoice, int64, int, error)
	Insert(invoice entity.Invoice) (*entity.Invoice, error)
	Update(invoice entity.Invoice) (*entity.Invoice, error)
}

type invoiceRepositoryImpl struct {
	db              *gorm.DB
	cartRepo        CartRepository
	userVoucherRepo UserVoucherRepository
}

type InvoiceRConfig struct {
	DB              *gorm.DB
	CartRepo        CartRepository
	UserVoucherRepo UserVoucherRepository
}

func NewInvoiceRepository(cfg *InvoiceRConfig) InvoiceRepository {
	return &invoiceRepositoryImpl{
		db:              cfg.DB,
		cartRepo:        cfg.CartRepo,
		userVoucherRepo: cfg.UserVoucherRepo,
	}
}

func (r *invoiceRepositoryImpl) FindById(id int) (*entity.Invoice, error) {
	var invoice entity.Invoice
	err := r.db.Preload("Transactions.Course").Preload("Voucher").First(&invoice, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errResp.ErrInvoiceNotFound
		}
		return nil, err
	}
	return &invoice, nil
}

func (r *invoiceRepositoryImpl) FindByUserId(userId int, params entity.InvoiceParams) ([]*entity.Invoice, int64, int, error) {
	var invoices []*entity.Invoice
	var count int64

	db := r.db.Where("user_id = ?", userId).Scopes(params.Scope())
	db.Model(&invoices).Count(&count)
	totalPages := int(math.Ceil(float64(count) / float64(params.Limit)))

	err := db.Preload("Voucher").Preload("Transactions.Course").Order(params.Sort).Offset(params.Offset()).Limit(params.Limit).Find(&invoices).Error
	if err != nil {
		return nil, 0, 0, err
	}
	return invoices, count, totalPages, nil
}

func (r *invoiceRepositoryImpl) Insert(invoice entity.Invoice) (*entity.Invoice, error) {
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return nil, err
	}

	err := tx.Omit("Voucher").Create(&invoice).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = r.cartRepo.EmptyCart(tx, invoice.UserId)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = r.userVoucherRepo.ConsumeVoucher(tx, invoice.UserVoucherId)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return &invoice, nil
}

func (r *invoiceRepositoryImpl) Update(invoice entity.Invoice) (*entity.Invoice, error) {
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return nil, err
	}

	err := tx.Save(&invoice).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return &invoice, nil
}
