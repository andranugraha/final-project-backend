package repository

import (
	"errors"
	"final-project-backend/entity"
	"final-project-backend/utils/constant"
	errResp "final-project-backend/utils/errors"
	"math"

	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CourseRepository interface {
	Insert(entity.Course) (*entity.Course, error)
	FindBySlug(string) (*entity.Course, error)
	FindAll(entity.CourseParams) ([]entity.Course, int64, int, error)
	FindPublishedById(int) (*entity.Course, error)
	FindTrending() ([]entity.Course, error)
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

func (r *courseRepositoryImpl) FindAll(params entity.CourseParams) ([]entity.Course, int64, int, error) {
	var courses []entity.Course
	var count int64

	db := r.db.Preload(clause.Associations).Scopes(params.Scope())
	db.Model(&courses).Count(&count)
	totalPages := int(math.Ceil(float64(count) / float64(params.Limit)))

	err := db.Order(params.Sort).Offset(params.Offset()).Limit(params.Limit).Find(&courses).Error
	if err != nil {
		return nil, 0, 0, err
	}

	return courses, count, totalPages, nil
}

func (r *courseRepositoryImpl) FindTrending() ([]entity.Course, error) {
	var courses []entity.Course
	defaultLimit := 5

	err := r.db.Preload(clause.Associations).
		Joins("left join transactions t ON courses.id = t.course_id and t.created_at >= date_trunc('week', now())").
		Joins("left join invoices i on t.invoice_id = i.id and i.status = ?", constant.CourseStatusCompleted).
		Where("courses.status = ?", constant.PublishStatus).
		Group("1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14").
		Order("count(t.id) DESC").Limit(defaultLimit).Find(&courses).Error
	if err != nil {
		return nil, err
	}

	return courses, nil
}

func (r *courseRepositoryImpl) FindBySlug(slug string) (*entity.Course, error) {
	var course *entity.Course

	err := r.db.Preload(clause.Associations).Where("slug = ?", slug).First(&course).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errResp.ErrCourseNotFound
		}
		return nil, err
	}

	return course, nil
}

func (r *courseRepositoryImpl) FindPublishedById(id int) (*entity.Course, error) {
	var course *entity.Course

	err := r.db.Preload(clause.Associations).Where("status = ?", constant.PublishStatus).First(&course, id).Error
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

	err := tx.Model(&req).Association("Tags").Replace(req.Tags)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Save(&req).Error
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
