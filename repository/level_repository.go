package repository

import (
	"final-project-backend/entity"

	"gorm.io/gorm"
)

type LevelRepository interface {
	FindValidByTotalTransaction(int64) (*entity.Level, error)
}

type levelRepositoryImpl struct {
	db *gorm.DB
}

type LevelRConfig struct {
	DB *gorm.DB
}

func NewLevelRepository(cfg *LevelRConfig) LevelRepository {
	return &levelRepositoryImpl{
		db: cfg.DB,
	}
}

func (r *levelRepositoryImpl) FindValidByTotalTransaction(totalTransaction int64) (*entity.Level, error) {
	var level entity.Level
	err := r.db.Where("min_transaction <= ?", totalTransaction).Order("min_transaction desc").First(&level).Error
	if err != nil {
		return nil, err
	}

	return &level, nil
}
