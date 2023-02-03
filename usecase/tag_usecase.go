package usecase

import (
	"final-project-backend/entity"
	"final-project-backend/repository"
)

type TagUsecase interface {
	GetTags() ([]*entity.Tag, error)
}

type tagUsecaseImpl struct {
	tagRepo repository.TagRepository
}

type TagUConfig struct {
	TagRepo repository.TagRepository
}

func NewTagUsecase(cfg *TagUConfig) TagUsecase {
	return &tagUsecaseImpl{
		tagRepo: cfg.TagRepo,
	}
}

func (t *tagUsecaseImpl) GetTags() ([]*entity.Tag, error) {
	return t.tagRepo.FindAll()
}
