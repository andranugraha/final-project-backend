package entity

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type Track struct {
	ID        int
	UserId	  int
	User      User `gorm:"foreignKey:UserId"`
	Status	  string
	DepartDate time.Time
	ArriveDate sql.NullTime
	Name 	string
	Description string
	Total int
	gorm.Model
}
