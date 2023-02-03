package repository

import (
	"final-project-backend/entity"
	errResp "final-project-backend/utils/errors"
	"math"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserCourseRepository interface {
	Insert(*gorm.DB, []*entity.UserCourse) ([]*entity.UserCourse, error)
	FindByUserId(int, entity.CourseParams) ([]entity.Course, int64, int, error)
	FindByUserIdAndCourseSlug(int, string) (*entity.UserCourse, error)
	FindByUserIdAndCourseId(int, int) (*entity.UserCourse, error)
	CountByCourseIdAndStatus(int, string) int
	Complete(entity.UserCourse) (*entity.UserCourse, error)
}

type userCourseRepositoryImpl struct {
	db       *gorm.DB
	userRepo UserRepository
}

type UserCourseRConfig struct {
	DB       *gorm.DB
	UserRepo UserRepository
}

func NewUserCourseRepository(cfg *UserCourseRConfig) UserCourseRepository {
	return &userCourseRepositoryImpl{
		db:       cfg.DB,
		userRepo: cfg.UserRepo,
	}
}

func (r *userCourseRepositoryImpl) Insert(tx *gorm.DB, userCourses []*entity.UserCourse) ([]*entity.UserCourse, error) {
	err := tx.Create(&userCourses).Error
	if err != nil {
		return nil, err
	}

	return userCourses, nil
}

func (r *userCourseRepositoryImpl) FindByUserIdAndCourseId(userId int, courseId int) (*entity.UserCourse, error) {
	var userCourse entity.UserCourse
	err := r.db.Where("user_id = ? AND course_id = ?", userId, courseId).First(&userCourse).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errResp.ErrUserCourseNotFound
		}

		return nil, err
	}

	return &userCourse, nil
}

func (r *userCourseRepositoryImpl) FindByUserIdAndCourseSlug(userId int, slug string) (*entity.UserCourse, error) {
	var userCourse entity.UserCourse

	err := r.db.Joins("JOIN courses ON courses.id = user_courses.course_id").Where("user_courses.user_id = ?", userId).Where("courses.slug = ?", slug).First(&userCourse).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errResp.ErrUserCourseNotFound
		}

		return nil, err
	}

	return &userCourse, nil
}

func (r *userCourseRepositoryImpl) FindByUserId(userId int, params entity.CourseParams) ([]entity.Course, int64, int, error) {
	var courses []entity.Course
	var count int64

	db := r.db.Preload(clause.Associations).Joins("JOIN user_courses ON user_courses.course_id = courses.id").Where("user_courses.user_id = ?", userId).Scopes(params.Scope())
	db.Model(&courses).Count(&count)
	totalPages := int(math.Ceil(float64(count) / float64(params.Limit)))

	err := db.Order(params.Sort).Offset(params.Offset()).Limit(params.Limit).Find(&courses).Error
	if err != nil {
		return nil, 0, 0, err
	}

	return courses, count, totalPages, nil
}

func (r *userCourseRepositoryImpl) CountByCourseIdAndStatus(courseId int, status string) int {
	var count int64
	r.db.Model(&entity.UserCourse{}).Where("course_id = ?", courseId).Where("status = ?", status).Count(&count)

	return int(count)
}

func (r *userCourseRepositoryImpl) Complete(userCourse entity.UserCourse) (*entity.UserCourse, error) {
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return nil, err
	}

	err := tx.Save(&userCourse).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = r.userRepo.AddRedeemablePoint(tx, userCourse.UserId)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return &userCourse, nil
}
