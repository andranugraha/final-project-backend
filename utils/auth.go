package utils

import (
	"time"

	"final-project-backend/config"
	"final-project-backend/dto"
	"final-project-backend/entity"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func HashAndSalt(pwd string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func ComparePassword(hashedPwd string, inputPwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(inputPwd))
	return err == nil
}

type accessTokenClaims struct {
	UserId int    `json:"userId"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

func GenerateAccessToken(req *entity.User) (*dto.SignInResponse, error) {
	claims := accessTokenClaims{
		req.ID,
		req.Fullname,
		req.Email,
		jwt.RegisteredClaims{
			IssuedAt: jwt.NewNumericDate(time.Now()),
			Issuer:   req.Email,
			// ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 1)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.Secret))
	if err != nil {
		return nil, err
	}

	return &dto.SignInResponse{AccessToken: tokenString}, nil
}
