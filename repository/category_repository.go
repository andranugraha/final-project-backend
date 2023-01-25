package repository

import (
	"final-project-backend/entity"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	FindAll() ([]*entity.Category, error)
}

type categoryRepositoryImpl struct {
	db *gorm.DB
}

type CategoryRConfig struct {
	DB *gorm.DB
}

func NewCategoryRepository(cfg *CategoryRConfig) CategoryRepository {
	return &categoryRepositoryImpl{
		db: cfg.DB,
	}
}

func (r *categoryRepositoryImpl) FindAll() ([]*entity.Category, error) {
	var categories []*entity.Category
	err := r.db.Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return categories, nil
}
