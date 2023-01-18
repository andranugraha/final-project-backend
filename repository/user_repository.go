package repository

import (
	"errors"

	"final-project-backend/entity"
	errResp "final-project-backend/utils/errors"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetById(int) (*entity.User, error)
	GetByIdentifierAndRole(string, int) (*entity.User, error)
	Insert(entity.User) (*entity.User, error)
}

type userRepositoryImpl struct {
	db *gorm.DB
}

type UserRConfig struct {
	DB *gorm.DB
}

func NewUserRepository(cfg *UserRConfig) UserRepository {
	return &userRepositoryImpl{db: cfg.DB}
}

func (r *userRepositoryImpl) GetById(id int) (*entity.User, error) {
	var res *entity.User
	err := r.db.First(&res, id)
	if err.Error != nil {
		if errors.Is(err.Error, gorm.ErrRecordNotFound) {
			return nil, errResp.ErrUserNotFound
		}
		return nil, err.Error
	}

	return res, nil
}

func (r *userRepositoryImpl) GetByIdentifierAndRole(identifier string, roleId int) (*entity.User, error) {
	var res *entity.User
	err := r.db.Where(
		r.db.Where("email = ?", identifier).Or("username = ?", identifier),
	).Where("role_id = ?", roleId).First(&res)
	if err.Error != nil {
		if errors.Is(err.Error, gorm.ErrRecordNotFound) {
			return nil, errResp.ErrUserNotFound
		}
		return nil, err.Error
	}

	return res, nil
}

func (r *userRepositoryImpl) Insert(req entity.User) (*entity.User, error) {
	err := r.db.Create(&req)
	if err.Error != nil {
		return nil, err.Error
	}

	return &req, nil
}
