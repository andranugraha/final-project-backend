package entity

import (
	"time"

	"gorm.io/gorm"
)

type Favorite struct {
	ID       int       `gorm:"primaryKey"`
	UserId   int       `gorm:"uniqueIndex:idx_favorite"`
	User     User      `gorm:"foreignKey:UserId"`
	CourseId int       `gorm:"uniqueIndex:idx_favorite"`
	Course   Course    `gorm:"foreignKey:CourseId"`
	Date     time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	gorm.Model
}

type GetFavoritesParams struct {
	UserId int
	Limit  int
	Page   int
}

func (g *GetFavoritesParams) Scope() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("user_id = ?", g.UserId)
	}
}

func (g *GetFavoritesParams) Offset() int {
	return (g.Page - 1) * g.Limit
}

func NewFavoritesParams(userId int, limit int, page int) GetFavoritesParams {
	return GetFavoritesParams{
		UserId: userId,
		Limit: func() int {
			if limit > 0 {
				return limit
			}

			return 10
		}(),
		Page: func() int {
			if page > 1 {
				return page
			}

			return 1
		}(),
	}
}
