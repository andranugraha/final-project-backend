package usecase

import (
	"final-project-backend/dto"
	"final-project-backend/repository"
	"final-project-backend/utils/auth"
	"final-project-backend/utils/errors"
)

type UserUsecase interface {
	SignIn(dto.SignInRequest, int) (*dto.SignInResponse, error)
	SignUp(dto.SignUpRequest) (*dto.SignUpResponse, error)
}

type userUsecaseImpl struct {
	userRepository repository.UserRepository
	bcryptUsecase  auth.AuthUtil
}

type UserUConfig struct {
	UserRepository repository.UserRepository
	BcryptUsecase  auth.AuthUtil
}

func NewUserUsecase(cfg *UserUConfig) UserUsecase {
	return &userUsecaseImpl{
		userRepository: cfg.UserRepository,
		bcryptUsecase:  cfg.BcryptUsecase,
	}
}

func (u *userUsecaseImpl) SignUp(req dto.SignUpRequest) (res *dto.SignUpResponse, err error) {
	user := req.ToUser()
	user.Password = u.bcryptUsecase.HashAndSalt(req.Password)

	const userRoleId = 2
	user.RoleId = userRoleId

	const newbieLevelId = 1
	user.LevelId = newbieLevelId

	registeredUser, err := u.userRepository.Insert(user)
	if err != nil {
		return
	}

	res.UserToResponse(*registeredUser)
	return
}

func (u *userUsecaseImpl) SignIn(req dto.SignInRequest, roleId int) (*dto.SignInResponse, error) {
	user, err := u.userRepository.GetByIdentifierAndRole(req.Identifier, roleId)
	if err != nil {
		return nil, err
	}

	if !u.bcryptUsecase.ComparePassword(user.Password, req.Password) {
		return nil, errors.ErrWrongPassword
	}

	res := u.bcryptUsecase.GenerateAccessToken(*user)

	return &res, nil
}
