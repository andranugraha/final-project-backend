package usecase

import (
	"final-project-backend/entity"
	"final-project-backend/repository"
	"final-project-backend/utils/errors"
)

type FavoriteUsecase interface {
	GetFavoriteCourse(userId int) ([]*entity.Course, error)
	SaveUnsaveFavoriteCourse(userId, courseId int, action string) error
}

type favoriteUsecaseImpl struct {
	favoriteRepo repository.FavoriteRepository
}

type FavoriteUConfig struct {
	FavoriteRepo repository.FavoriteRepository
}

func NewFavoriteUsecase(cfg *FavoriteUConfig) FavoriteUsecase {
	return &favoriteUsecaseImpl{
		favoriteRepo: cfg.FavoriteRepo,
	}
}

func (u *favoriteUsecaseImpl) GetFavoriteCourse(userId int) ([]*entity.Course, error) {
	return u.favoriteRepo.FindByUserId(userId)
}

func (u *favoriteUsecaseImpl) SaveUnsaveFavoriteCourse(userId, courseId int, action string) error {
	if action == "save" {
		return u.saveFavoriteCourse(userId, courseId)
	} else if action == "unsave" {
		return u.unsaveFavoriteCourse(userId, courseId)
	}

	return errors.ErrUnknownAction
}

func (u *favoriteUsecaseImpl) saveFavoriteCourse(userId, courseId int) error {
	favorite := entity.Favorite{
		UserId:   userId,
		CourseId: courseId,
	}
	return u.favoriteRepo.Insert(favorite)
}

func (u *favoriteUsecaseImpl) unsaveFavoriteCourse(userId, courseId int) error {
	favorite := entity.Favorite{
		UserId:   userId,
		CourseId: courseId,
	}
	return u.favoriteRepo.Delete(favorite)
}
