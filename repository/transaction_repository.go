package repository

import (
	"errors"
	"final-project-backend/entity"
	"final-project-backend/utils/constant"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	FindBoughtByUserIdAndCourseId(int, int) (*entity.Transaction, error)
}

type transactionRepositoryImpl struct {
	db *gorm.DB
}

type TransactionRConfig struct {
	DB *gorm.DB
}

func NewTransactionRepository(cfg *TransactionRConfig) TransactionRepository {
	return &transactionRepositoryImpl{
		db: cfg.DB,
	}
}

func (r *transactionRepositoryImpl) FindBoughtByUserIdAndCourseId(userId, courseId int) (*entity.Transaction, error) {
	var transaction entity.Transaction
	err := r.db.Where("user_id = ?", userId).Where("course_id = ?", courseId).Joins("join invoices on invoices.id = transactions.invoice_id and invoices.status != ?", constant.InvoiceStatusCancelled).First(&transaction).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}
	return &transaction, nil
}
