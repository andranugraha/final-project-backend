package usecase

import (
	"final-project-backend/entity"
	"final-project-backend/repository"
)

type UserUsecase interface {
	GetUserDetail(userId int) (*entity.User, error)
}

type userUsecaseImpl struct {
	userRepo repository.UserRepository
}

type UserUConfig struct {
	UserRepo repository.UserRepository
}

func NewUserUsecase(cfg *UserUConfig) UserUsecase {
	return &userUsecaseImpl{
		userRepo: cfg.UserRepo,
	}
}

func (u *userUsecaseImpl) GetUserDetail(userId int) (*entity.User, error) {
	return u.userRepo.GetDetailById(userId)
}
