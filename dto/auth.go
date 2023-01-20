package dto

import (
	"final-project-backend/entity"
	"final-project-backend/utils/constant"
	"strings"
)

type SignInRequest struct {
	Identifier string `json:"identifier" binding:"required"`
	Password   string `json:"password" binding:"required,min=8,max=20"`
}

type SignUpRequest struct {
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required,min=8,max=20"`
	Username    string `json:"username" binding:"required,min=3,max=20"`
	Fullname    string `json:"fullname" binding:"required"`
	Address     string `json:"address" binding:"required"`
	PhoneNo     string `json:"phoneNo" binding:"required,numeric,min=10,max=14"`
	RefReferral string `json:"refReferral"`
}

type SignInResponse struct {
	AccessToken string `json:"accessToken"`
}

type SignUpResponse struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Fullname string `json:"fullname"`
	Address  string `json:"address"`
	PhoneNo  string `json:"phoneNo"`
	Referral string `json:"referral"`
}

func (s *SignUpRequest) ToUser() entity.User {
	replacer := strings.NewReplacer("@", "", " ", "")

	return entity.User{
		Email:    s.Email,
		Username: replacer.Replace(s.Username),
		Fullname: s.Fullname,
		Address:  s.Address,
		PhoneNo:  s.PhoneNo,
		RefReferral: func() *string {
			if s.RefReferral != "" {
				return &s.RefReferral
			}
			return nil
		}(),
		RoleId:  constant.UserRoleId,
		LevelId: constant.NewbieLevelId,
	}
}

func (s *SignUpResponse) FromUser(u entity.User) {
	s.Id = u.ID
	s.Email = u.Email
	s.Username = u.Username
	s.Fullname = u.Fullname
	s.Address = u.Address
	s.PhoneNo = u.PhoneNo
	s.Referral = u.Referral
}
