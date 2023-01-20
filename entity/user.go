package entity

import "gorm.io/gorm"

type User struct {
	ID          int     `json:"id" gorm:"primaryKey"`
	Email       string  `json:"email"`
	Password    string  `json:"-"`
	Username    string  `json:"username"`
	Fullname    string  `json:"fullname"`
	Address     string  `json:"address"`
	PhoneNo     string  `json:"phoneNo"`
	Referral    string  `json:"referral"`
	RefReferral *string `json:"refReferral,omitempty"`
	RoleId      int     `json:"roleId"`
	Role        *Role   `json:"role,omitempty" gorm:"foreignKey:RoleId"`
	LevelId     int     `json:"levelId"`
	Level       *Level  `json:"level,omitempty" gorm:"foreignKey:LevelId"`
	gorm.Model  `json:"-"`
}
