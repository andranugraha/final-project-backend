package entity

import "gorm.io/gorm"

type UserGift struct {
	ID       int
	UserId   int
	User     User `gorm:"foreignKey:UserId"`
	GiftId   int
	Gift     Gift `gorm:"foreignKey:GiftId"`
	TrackId  int
	Track    Track `gorm:"foreignKey:TrackId"`
	Subtotal int
	gorm.Model
}