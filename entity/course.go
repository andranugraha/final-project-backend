package entity

import "gorm.io/gorm"

type Course struct {
	ID           int
	Title        string
	Slug         string
	Summary      string
	Content      string
	ImgThumbnail string
	ImgUrl       string
	AuthorName   string
	Status       string
	CategoryId   int
	Category     Category `gorm:"foreignKey:CategoryId"`
	gorm.Model
}
