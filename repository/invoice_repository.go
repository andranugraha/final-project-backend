package repository

import (
	"errors"
	"final-project-backend/entity"
	"final-project-backend/utils/constant"
	errResp "final-project-backend/utils/errors"
	"math"

	"gorm.io/gorm"
)

type InvoiceRepository interface {
	FindById(id string) (*entity.Invoice, error)
	FindByIdAndUserId(id string, userId int) (*entity.Invoice, error)
	FindAll(params entity.InvoiceParams) ([]*entity.Invoice, int64, int, error)
	Insert(invoice entity.Invoice) (*entity.Invoice, error)
	Update(invoice entity.Invoice) (*entity.Invoice, error)
}

type invoiceRepositoryImpl struct {
	db              *gorm.DB
	cartRepo        CartRepository
	userVoucherRepo UserVoucherRepository
	userCourseRepo  UserCourseRepository
	userRepo        UserRepository
}

type InvoiceRConfig struct {
	DB              *gorm.DB
	CartRepo        CartRepository
	UserVoucherRepo UserVoucherRepository
	UserCourseRepo  UserCourseRepository
	UserRepo        UserRepository
}

func NewInvoiceRepository(cfg *InvoiceRConfig) InvoiceRepository {
	return &invoiceRepositoryImpl{
		db:              cfg.DB,
		cartRepo:        cfg.CartRepo,
		userVoucherRepo: cfg.UserVoucherRepo,
		userCourseRepo:  cfg.UserCourseRepo,
		userRepo:        cfg.UserRepo,
	}
}

func (r *invoiceRepositoryImpl) FindById(id string) (*entity.Invoice, error) {
	var invoice entity.Invoice
	err := r.db.Preload("Transactions.Course").Where("id = ?", id).Preload("Voucher").First(&invoice).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errResp.ErrInvoiceNotFound
		}
		return nil, err
	}
	return &invoice, nil
}

func (r *invoiceRepositoryImpl) FindByIdAndUserId(id string, userId int) (*entity.Invoice, error) {
	var invoice entity.Invoice
	err := r.db.Preload("Transactions.Course").Where("id = ? AND user_id = ?", id, userId).Preload("Voucher").First(&invoice).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errResp.ErrInvoiceNotFound
		}
		return nil, err
	}
	return &invoice, nil
}

func (r *invoiceRepositoryImpl) FindAll(params entity.InvoiceParams) ([]*entity.Invoice, int64, int, error) {
	var invoices []*entity.Invoice
	var count int64

	db := r.db.Scopes(params.Scope())
	db.Model(&invoices).Count(&count)
	totalPages := int(math.Ceil(float64(count) / float64(params.Limit)))

	err := db.Preload("Voucher").Preload("Transactions.Course").Preload("User").Order(params.Sort).Offset(params.Offset()).Limit(params.Limit).Find(&invoices).Error
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

	if invoice.UserVoucherId != nil {
		err = r.userVoucherRepo.ConsumeVoucher(tx, *invoice.UserVoucherId)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
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

	err := tx.Omit("Transactions").Save(&invoice).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if invoice.Status == constant.InvoiceStatusCompleted {
		err = r.completeInvoice(tx, &invoice)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	if invoice.Status == constant.InvoiceStatusCancelled && invoice.UserVoucherId != nil {
		err = r.userVoucherRepo.UnconsumeVoucher(tx, *invoice.UserVoucherId)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	tx.Commit()

	return &invoice, nil
}

func (r *invoiceRepositoryImpl) completeInvoice(tx *gorm.DB, invoice *entity.Invoice) error {
	var userCourses []*entity.UserCourse
	for _, transaction := range invoice.Transactions {
		userCourses = append(userCourses, &entity.UserCourse{
			UserId:   invoice.UserId,
			CourseId: transaction.CourseId,
			Status:   constant.CourseStatusInProgress,
		})
	}

	_, err := r.userCourseRepo.Insert(tx, userCourses)
	if err != nil {
		tx.Rollback()
		return err
	}

	var totalTransaction int64
	err = tx.Model(&entity.Invoice{}).Where("user_id = ?", invoice.UserId).Where("status = ?", constant.InvoiceStatusCompleted).Count(&totalTransaction).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	user, err := r.userRepo.LevelUp(tx, invoice.UserId, totalTransaction)
	if err != nil {
		tx.Rollback()
		return err
	}

	if user.RefReferral != nil && *user.RefReferral != "" {
		referrer, err := r.userRepo.FindByReferral(*user.RefReferral)
		if err != nil {
			tx.Rollback()
			return err
		}

		_, err = r.userVoucherRepo.Insert(tx, int64(invoice.Total), referrer.ID)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return nil
}
