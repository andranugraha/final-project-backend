package dto

import "final-project-backend/entity"

type UpdateUserDetailRequest struct {
	Fullname string `json:"fullname"`
	Address  string `json:"address"`
	PhoneNo  string `json:"phoneNo" binding:"numeric,min=10,max=14"`
}

func (dto *UpdateUserDetailRequest) ToUser(userId int) entity.User {
	return entity.User{
		ID:       userId,
		Fullname: dto.Fullname,
		Address:  dto.Address,
		PhoneNo:  dto.PhoneNo,
	}
}
