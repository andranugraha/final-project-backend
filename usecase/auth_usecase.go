package usecase

import (
	"final-project-backend/dto"
	"final-project-backend/entity"
	"final-project-backend/repository"
	"final-project-backend/utils/auth"
	"final-project-backend/utils/errors"
)

type AuthUsecase interface {
	SignIn(dto.SignInRequest, int) (*dto.SignInResponse, error)
	SignUp(dto.SignUpRequest) (*dto.SignUpResponse, error)
}

type authUsecaseImpl struct {
	userRepo    repository.UserRepository
	utilUsecase auth.AuthUtil
}

type AuthUConfig struct {
	UserRepo    repository.UserRepository
	UtilUsecase auth.AuthUtil
}

func NewAuthUsecase(cfg *AuthUConfig) AuthUsecase {
	return &authUsecaseImpl{
		userRepo:    cfg.UserRepo,
		utilUsecase: cfg.UtilUsecase,
	}
}

func (u *authUsecaseImpl) SignUp(req dto.SignUpRequest) (*dto.SignUpResponse, error) {
	user := req.ToUser()
	user.Password = u.utilUsecase.HashAndSalt(req.Password)
	user.Referral = u.utilUsecase.GenerateReferralCode()
	user.Redeemable = &entity.Redeemable{}

	registeredUser, err := u.userRepo.Insert(user)
	if err != nil {
		return nil, err
	}

	res := dto.SignUpResponse{}
	res.FromUser(*registeredUser)

	return &res, nil
}

func (u *authUsecaseImpl) SignIn(req dto.SignInRequest, roleId int) (*dto.SignInResponse, error) {
	user, err := u.userRepo.FindByIdentifierAndRole(req.Identifier, roleId)
	if err != nil {
		return nil, err
	}

	if !u.utilUsecase.ComparePassword(user.Password, req.Password) {
		return nil, errors.ErrWrongPassword
	}

	res := u.utilUsecase.GenerateAccessToken(*user)

	return &res, nil
}
