package auth

import (
	"math/rand"
	"time"

	"final-project-backend/config"
	"final-project-backend/dto"
	"final-project-backend/entity"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type AuthUtil interface {
	HashAndSalt(pwd string) string
	ComparePassword(hashedPwd string, inputPwd string) bool
	GenerateAccessToken(req entity.User) dto.SignInResponse
	GenerateReferralCode() string
}

type authUtilImpl struct{}

func NewAuthUtil() AuthUtil {
	return authUtilImpl{}
}

func (d authUtilImpl) HashAndSalt(pwd string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)

	return string(hash)
}

func (d authUtilImpl) ComparePassword(hashedPwd string, inputPwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(inputPwd))
	return err == nil
}

func (d authUtilImpl) GenerateReferralCode() string {
	rand.Seed(time.Now().UnixNano())
	alphabet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	referralCode := ""

	for i := 0; i < 7; i++ {
		referralCode += string(alphabet[rand.Intn(len(alphabet))])
	}

	return referralCode
}

type accessTokenClaims struct {
	UserId int    `json:"userId"`
	Email  string `json:"email"`
	RoleId int    `json:"roleId"`
	jwt.RegisteredClaims
}

func (d authUtilImpl) GenerateAccessToken(req entity.User) dto.SignInResponse {
	claims := accessTokenClaims{
		req.ID,
		req.Email,
		req.RoleId,
		jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    config.AppName,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 6)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString([]byte(config.Secret))

	return dto.SignInResponse{AccessToken: tokenString}
}
