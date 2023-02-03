package repository

import (
	"final-project-backend/entity"

	"gorm.io/gorm"
)

type RedeemableRepository interface {
	FindByUserId(int) (*entity.Redeemable, error)
	AddPoint(*gorm.DB, int, int) error
	ReducePoint(*gorm.DB, int, int) error
}

type redeemableRepositoryImpl struct {
	db *gorm.DB
}

type RedeemableRConfig struct {
	DB *gorm.DB
}

func NewRedeemableRepository(cfg *RedeemableRConfig) RedeemableRepository {
	return &redeemableRepositoryImpl{
		db: cfg.DB,
	}
}

func (r *redeemableRepositoryImpl) FindByUserId(userId int) (*entity.Redeemable, error) {
	var redeemable entity.Redeemable
	err := r.db.Where("user_id = ?", userId).First(&redeemable).Error
	if err != nil {
		return nil, err
	}

	return &redeemable, nil
}

func (r *redeemableRepositoryImpl) AddPoint(tx *gorm.DB, userId int, point int) error {
	err := tx.Model(&entity.Redeemable{}).Where("user_id = ?", userId).Update("point", gorm.Expr("point + ?", point)).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *redeemableRepositoryImpl) ReducePoint(tx *gorm.DB, userId int, point int) error {
	err := tx.Model(&entity.Redeemable{}).Where("user_id = ?", userId).Update("point", gorm.Expr("point - ?", point)).Error
	if err != nil {
		return err
	}

	return nil
}
