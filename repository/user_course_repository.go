package repository

import (
	"final-project-backend/entity"

	"gorm.io/gorm"
)

type UserCourseRepository interface {
	Insert(*gorm.DB, []*entity.UserCourse) ([]*entity.UserCourse, error)
	FindByUserIdAndCourseId(int, int) (*entity.UserCourse, error)
	Complete(entity.UserCourse) (*entity.UserCourse, error)
}

type userCourseRepositoryImpl struct {
	db *gorm.DB
}

type UserCourseRConfig struct {
	DB *gorm.DB
}

func NewUserCourseRepository(cfg *UserCourseRConfig) UserCourseRepository {
	return &userCourseRepositoryImpl{
		db: cfg.DB,
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
		return nil, err
	}

	return &userCourse, nil
}

func (r *userCourseRepositoryImpl) Complete(userCourse entity.UserCourse) (*entity.UserCourse, error) {
	err := r.db.Save(&userCourse).Error
	if err != nil {
		return nil, err
	}

	return &userCourse, nil
}
