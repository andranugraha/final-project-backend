package entity

import "gorm.io/gorm"

type Course struct {
	ID           int       `json:"id" gorm:"primaryKey"`
	Title        string    `json:"title"`
	Slug         string    `json:"slug"`
	Summary      string    `json:"summary"`
	Content      string    `json:"content"`
	ImgThumbnail string    `json:"imgThumbnail"`
	ImgUrl       string    `json:"imgUrl"`
	AuthorName   string    `json:"authorName"`
	Status       string    `json:"status"`
	Price        float64   `json:"price"`
	CategoryId   int       `json:"categoryId"`
	Category     *Category `json:"category,omitempty" gorm:"foreignKey:CategoryId"`
	Tags         []*Tag    `json:"tags" gorm:"many2many:course_tags;"`
	gorm.Model   `json:"-"`
}
