package usecase

import (
	"final-project-backend/dto"
	"final-project-backend/repository"
	"final-project-backend/utils/auth"
	"final-project-backend/utils/errors"
	"math/rand"
	"time"
)

type UserUsecase interface {
	SignIn(dto.SignInRequest, int) (*dto.SignInResponse, error)
	SignUp(dto.SignUpRequest) (*dto.SignUpResponse, error)
}

type userUsecaseImpl struct {
	userRepo      repository.UserRepository
	bcryptUsecase auth.AuthUtil
}

type UserUConfig struct {
	UserRepo      repository.UserRepository
	BcryptUsecase auth.AuthUtil
}

func NewUserUsecase(cfg *UserUConfig) UserUsecase {
	return &userUsecaseImpl{
		userRepo:      cfg.UserRepo,
		bcryptUsecase: cfg.BcryptUsecase,
	}
}

func (u *userUsecaseImpl) SignUp(req dto.SignUpRequest) (*dto.SignUpResponse, error) {
	user := req.ToUser()
	user.Password = u.bcryptUsecase.HashAndSalt(req.Password)
	user.Referral = u.generateReferralCode()

	registeredUser, err := u.userRepo.Insert(user)
	if err != nil {
		return nil, err
	}

	res := dto.SignUpResponse{}
	res.FromUser(*registeredUser)

	return &res, nil
}

func (u *userUsecaseImpl) SignIn(req dto.SignInRequest, roleId int) (*dto.SignInResponse, error) {
	user, err := u.userRepo.GetByIdentifierAndRole(req.Identifier, roleId)
	if err != nil {
		return nil, err
	}

	if !u.bcryptUsecase.ComparePassword(user.Password, req.Password) {
		return nil, errors.ErrWrongPassword
	}

	res := u.bcryptUsecase.GenerateAccessToken(*user)

	return &res, nil
}

func (u *userUsecaseImpl) generateReferralCode() string {
	rand.Seed(time.Now().UnixNano())
	alphabet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	referralCode := ""

	for i := 0; i < 7; i++ {
		referralCode += string(alphabet[rand.Intn(len(alphabet))])
	}

	return referralCode
}
