package usecase

import (
	"final-project-backend/entity"
	"final-project-backend/repository"
)

type CategoryUsecase interface {
	GetCategories() ([]*entity.Category, error)
}

type categoryUsecaseImpl struct {
	categoryRepo repository.CategoryRepository
}

type CategoryUConfig struct {
	CategoryRepo repository.CategoryRepository
}

func NewCategoryUsecase(cfg *CategoryUConfig) CategoryUsecase {
	return &categoryUsecaseImpl{
		categoryRepo: cfg.CategoryRepo,
	}
}

func (u *categoryUsecaseImpl) GetCategories() ([]*entity.Category, error) {
	return u.categoryRepo.FindAll()
}
