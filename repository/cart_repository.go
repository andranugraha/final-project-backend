package repository

import (
	"errors"
	"final-project-backend/entity"
	errResp "final-project-backend/utils/errors"

	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
)

type CartRepository interface {
	FindByUserId(userId int) ([]*entity.Cart, error)
	Insert(cart entity.Cart) error
	Delete(cart entity.Cart) error
	EmptyCart(tx *gorm.DB, userId int) error
}

type cartRepositoryImpl struct {
	db *gorm.DB
}

type CartRConfig struct {
	DB *gorm.DB
}

func NewCartRepository(cfg *CartRConfig) CartRepository {
	return &cartRepositoryImpl{
		db: cfg.DB,
	}
}

func (r *cartRepositoryImpl) FindByUserId(userId int) ([]*entity.Cart, error) {
	var carts []*entity.Cart
	err := r.db.Where("user_id = ?", userId).Joins("Course").Find(&carts).Error
	if err != nil {
		return nil, err
	}
	return carts, nil
}

func (r *cartRepositoryImpl) Insert(cart entity.Cart) error {
	err := r.db.Create(&cart).Error
	if err != nil {
		if pgError := err.(*pgconn.PgError); errors.Is(err, pgError) {
			if pgError.Code == errResp.ErrSqlUniqueViolation {
				err = errResp.ErrDuplicateFavorite
			}
		}
		return err
	}

	return nil
}

func (r *cartRepositoryImpl) Delete(cart entity.Cart) error {
	err := r.db.Unscoped().Where("user_id = ?", cart.UserId).Where("course_id = ?", cart.CourseId).Delete(&cart)
	if err.Error != nil {
		return err.Error
	}

	if err.RowsAffected == 0 {
		return errResp.ErrCartNotFound
	}

	return nil
}

func (r *cartRepositoryImpl) EmptyCart(tx *gorm.DB, userId int) error {
	err := tx.Unscoped().Where("user_id = ?", userId).Delete(&entity.Cart{}).Error
	if err != nil {
		return err
	}

	return nil
}
