package repository

import (
	"final-project-backend/entity"

	"gorm.io/gorm"
)

type UserCourseRepository interface {
	Insert(*gorm.DB, []*entity.UserCourse) ([]*entity.UserCourse, error)
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
