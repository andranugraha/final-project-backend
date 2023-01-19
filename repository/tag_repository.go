package repository

import (
	"final-project-backend/entity"

	"gorm.io/gorm"
)

type TagRepository interface {
	FindOrCreate(*gorm.DB, *entity.Tag) (*entity.Tag, error)
}

type tagRepositoryImpl struct {
	db *gorm.DB
}

type TagRConfig struct {
	DB *gorm.DB
}

func NewTagRepository(cfg *TagRConfig) TagRepository {
	return &tagRepositoryImpl{
		db: cfg.DB,
	}
}

func (r *tagRepositoryImpl) FindOrCreate(tx *gorm.DB, req *entity.Tag) (*entity.Tag, error) {
	err := tx.Where("name = ?", req.Name).FirstOrCreate(&req).Error
	if err != nil {
		return nil, err
	}

	return req, nil
}
