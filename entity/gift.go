package entity

import "gorm.io/gorm"

type Gift struct {
	ID        int
	Name	  string
	Stock	  int
	gorm.Model
}