package entity

import "gorm.io/gorm"

type User struct {
	ID       int
	Email    string
	Password string
	Fullname string
	Address string
	PhoneNo string
	Referral string
	RefReferral string
	RoleId   int
	Role     Role `gorm:"foreignKey:RoleId"`
	LevelId  int
	Level    Level `gorm:"foreignKey:LevelId"`
	gorm.Model
}

