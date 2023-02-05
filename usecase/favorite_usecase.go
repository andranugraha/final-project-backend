package usecase

import (
	"final-project-backend/entity"
	"final-project-backend/repository"
	"final-project-backend/utils/errors"
)

type FavoriteUsecase interface {
	GetFavoriteCourses(entity.GetFavoritesParams) ([]*entity.Course, int64, int, error)
	GetTotalFavorited(courseId int) int
	CheckIsFavoriteCourse(userId, courseId int) bool
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

func (u *favoriteUsecaseImpl) GetFavoriteCourses(params entity.GetFavoritesParams) ([]*entity.Course, int64, int, error) {
	return u.favoriteRepo.FindByUserId(params)
}

func (u *favoriteUsecaseImpl) CheckIsFavoriteCourse(userId, courseId int) bool {
	_, err := u.favoriteRepo.FindByUserIdAndCourseId(userId, courseId)
	return err == nil
}

func (u *favoriteUsecaseImpl) GetTotalFavorited(courseId int) int {
	return u.favoriteRepo.CountByCourseId(courseId)
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
