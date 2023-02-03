package entity

import "gorm.io/gorm"

type Redeemable struct {
	ID         int   `json:"id"`
	UserId     int   `json:"user_id"`
	User       *User `json:"user,omitempty" gorm:"foreignKey:UserId"`
	Point      int   `json:"point"`
	gorm.Model `json:"-"`
}
