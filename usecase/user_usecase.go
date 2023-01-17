package usecase

import (
	"final-project-backend/dto"
	"final-project-backend/repository"
	"final-project-backend/utils"
)

type UserUsecase interface {
	SignIn(req dto.SignInRequest) (*dto.SignInResponse, error)
}

type userUsecaseImpl struct {
	userRepository repository.UserRepository
}

type UserUConfig struct {
	UserRepository repository.UserRepository
}

func NewUserUsecase(cfg *UserUConfig) UserUsecase {
	return &userUsecaseImpl{userRepository: cfg.UserRepository}
}

func (u *userUsecaseImpl) SignIn(req dto.SignInRequest) (*dto.SignInResponse, error) {
	user, err := u.userRepository.GetByEmail(req.Email)
	if err != nil {
		return nil, err
	}

	if !utils.ComparePassword(user.Password, req.Password) {
		return nil, utils.ErrWrongPassword
	}

	res, err := utils.GenerateAccessToken(user)
	if err != nil {
		return nil, err
	}

	return res, nil
}
