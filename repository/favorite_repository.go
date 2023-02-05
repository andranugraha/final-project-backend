package repository

import (
	"errors"
	"final-project-backend/entity"
	errResp "final-project-backend/utils/errors"
	"math"

	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type FavoriteRepository interface {
	FindByUserId(entity.GetFavoritesParams) ([]*entity.Course, int64, int, error)
	FindByUserIdAndCourseId(userId, courseId int) (*entity.Favorite, error)
	Insert(favorite entity.Favorite) error
	Delete(favorite entity.Favorite) error
	CountByCourseId(courseId int) int
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

func (r *favoriteRepositoryImpl) FindByUserId(params entity.GetFavoritesParams) ([]*entity.Course, int64, int, error) {
	var courses []*entity.Course
	var count int64

	db := r.db.Scopes(params.Scope()).Joins("JOIN favorites ON favorites.course_id = courses.id")
	db.Model(&courses).Count(&count)
	totalPages := int(math.Ceil(float64(count) / float64(params.Limit)))

	err := db.Preload(clause.Associations).Offset(params.Offset()).Limit(params.Limit).Find(&courses).Error
	if err != nil {
		return nil, 0, 0, err
	}

	return courses, count, totalPages, nil
}

func (r *favoriteRepositoryImpl) FindByUserIdAndCourseId(userId, courseId int) (*entity.Favorite, error) {
	var favorite entity.Favorite
	err := r.db.Where("user_id = ?", userId).Where("course_id = ?", courseId).First(&favorite).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errResp.ErrFavoriteNotFound
		}
		return nil, err
	}

	return &favorite, nil
}

func (r *favoriteRepositoryImpl) Insert(favorite entity.Favorite) error {
	err := r.db.Create(&favorite).Error
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
	err := r.db.Unscoped().Where("user_id = ?", favorite.UserId).Where("course_id = ?", favorite.CourseId).Delete(&favorite)
	if err.Error != nil {
		return err.Error
	}

	if err.RowsAffected == 0 {
		return errResp.ErrFavoriteNotFound
	}

	return nil
}

func (r *favoriteRepositoryImpl) CountByCourseId(courseId int) int {
	var count int64
	r.db.Model(&entity.Favorite{}).Where("course_id = ?", courseId).Count(&count)

	return int(count)
}
