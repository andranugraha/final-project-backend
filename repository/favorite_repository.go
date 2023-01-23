package repository

import (
	"errors"
	"final-project-backend/entity"
	errResp "final-project-backend/utils/errors"

	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type FavoriteRepository interface {
	FindByUserId(userId int) ([]*entity.Course, error)
	Insert(favorite entity.Favorite) error
	Delete(favorite entity.Favorite) error
}

type favoriteRepositoryImpl struct {
	db *gorm.DB
}

type FavoriteRConfig struct {
	DB *gorm.DB
}

func NewFavoriteRepository(cfg *FavoriteRConfig) FavoriteRepository {
	return &favoriteRepositoryImpl{
		db: cfg.DB,
	}
}

func (r *favoriteRepositoryImpl) FindByUserId(userId int) ([]*entity.Course, error) {
	var courses []*entity.Course
	err := r.db.Where("user_id = ?", userId).Joins("JOIN favorites ON favorites.course_id = courses.id").Find(&courses).Error
	if err != nil {
		return nil, err
	}

	return courses, nil
}

func (r *favoriteRepositoryImpl) Insert(favorite entity.Favorite) error {
	err := r.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}, {Name: "course_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"updated_at", "deleted_at"}),
	}).Create(&favorite).Error
	if err != nil {
		if pgError := err.(*pgconn.PgError); errors.Is(err, pgError) {
			if pgError.Code == errResp.ErrSqlUniqueViolation {
				err = errResp.ErrDuplicateFavorite
			}
		}
		return err
	}

	return nil
}

func (r *favoriteRepositoryImpl) Delete(favorite entity.Favorite) error {
	err := r.db.Where("user_id = ?", favorite.UserId).Where("course_id = ?", favorite.CourseId).Delete(&favorite).Error
	if err != nil {
		return err
	}

	return nil
}
