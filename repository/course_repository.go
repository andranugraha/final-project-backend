package repository

import (
	"errors"
	"final-project-backend/entity"
	errResp "final-project-backend/utils/errors"

	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
)

type CourseRepository interface {
	Insert(entity.Course) (*entity.Course, error)
	FindBySlug(string) (*entity.Course, error)
	Update(entity.Course) (*entity.Course, error)
	Delete(string) error
}

type courseRepositoryImpl struct {
	db      *gorm.DB
	tagRepo TagRepository
}

type CourseRConfig struct {
	DB      *gorm.DB
	TagRepo TagRepository
}

func NewCourseRepository(cfg *CourseRConfig) CourseRepository {
	return &courseRepositoryImpl{
		db:      cfg.DB,
		tagRepo: cfg.TagRepo,
	}
}

func (r *courseRepositoryImpl) Insert(req entity.Course) (*entity.Course, error) {
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return nil, err
	}

	for _, tag := range req.Tags {
		_, err := r.tagRepo.FindOrCreate(tx, tag)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	err := tx.Create(&req).Error
	if err != nil {
		tx.Rollback()
		if pgError := err.(*pgconn.PgError); errors.Is(err, pgError) {
			if pgError.Code == errResp.ErrSqlUniqueViolation {
				err = errResp.ErrDuplicateTitle
			}
		}

		return nil, err
	}

	tx.Commit()

	return &req, nil
}

func (r *courseRepositoryImpl) FindBySlug(slug string) (*entity.Course, error) {
	var course *entity.Course

	err := r.db.Preload("Tags").Preload("Category").Where("slug = ?", slug).First(&course).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errResp.ErrCourseNotFound
		}
		return nil, err
	}

	return course, nil
}

func (r *courseRepositoryImpl) Update(req entity.Course) (*entity.Course, error) {
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return nil, err
	}

	for _, tag := range req.Tags {
		_, err := r.tagRepo.FindOrCreate(tx, tag)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	err := tx.Save(&req).Error
	if err != nil {
		tx.Rollback()
		if pgError := err.(*pgconn.PgError); errors.Is(err, pgError) {
			if pgError.Code == errResp.ErrSqlUniqueViolation {
				err = errResp.ErrDuplicateTitle
			}
		}

		return nil, err
	}

	tx.Commit()
	return &req, nil
}

func (r *courseRepositoryImpl) Delete(slug string) error {
	err := r.db.Where("slug = ?", slug).Delete(&entity.Course{})
	if err.Error != nil {
		return err.Error
	}

	if err.RowsAffected == 0 {
		return errResp.ErrCourseNotFound
	}

	return nil
}
