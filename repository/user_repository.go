package repository

import (
	"errors"

	"final-project-backend/entity"
	errResp "final-project-backend/utils/errors"

	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepository interface {
	GetById(int) (*entity.User, error)
	GetByIdentifierAndRole(string, int) (*entity.User, error)
	GetDetailById(int) (*entity.User, error)
	Insert(entity.User) (*entity.User, error)
	UpdateDetail(entity.User) (*entity.User, error)
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

func (r *userRepositoryImpl) GetDetailById(id int) (*entity.User, error) {
	var res *entity.User
	err := r.db.Preload(clause.Associations).First(&res, id)
	if err.Error != nil {
		if errors.Is(err.Error, gorm.ErrRecordNotFound) {
			return nil, errResp.ErrUserNotFound
		}
		return nil, err.Error
	}

	return res, nil
}

func (r *userRepositoryImpl) Insert(req entity.User) (*entity.User, error) {
	err := r.db.Create(&req).Error
	if err != nil {
		if pgError := err.(*pgconn.PgError); errors.Is(err, pgError) {
			if pgError.Code == errResp.ErrSqlUniqueViolation {
				err = errResp.ErrDuplicateRecord
			}
		}
		return nil, err
	}

	return &req, nil
}

func (r *userRepositoryImpl) UpdateDetail(req entity.User) (*entity.User, error) {
	res := r.db.Model(&req).Select("Fullname", "Address", "PhoneNo").Clauses(clause.Returning{}).Updates(req)
	if res.Error != nil {
		if pgError := res.Error.(*pgconn.PgError); errors.Is(res.Error, pgError) {
			if pgError.Code == errResp.ErrSqlUniqueViolation {
				res.Error = errResp.ErrDuplicatePhoneNo
			}
		}
		return nil, res.Error
	}

	if res.RowsAffected == 0 {
		return nil, errResp.ErrUserNotFound
	}

	return &req, res.Error
}
