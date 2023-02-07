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
	FindById(int) (*entity.User, error)
	FindByIdentifierAndRole(string, int) (*entity.User, error)
	FindDetailById(int) (*entity.User, error)
	FindByReferral(string) (*entity.User, error)
	Insert(entity.User) (*entity.User, error)
	UpdateDetail(entity.User) (*entity.User, error)
	LevelUp(*gorm.DB, int, int64) (*entity.User, error)
	AddRedeemablePoint(*gorm.DB, int) error
}

type userRepositoryImpl struct {
	db             *gorm.DB
	levelRepo      LevelRepository
	redeemableRepo RedeemableRepository
}

type UserRConfig struct {
	DB             *gorm.DB
	LevelRepo      LevelRepository
	RedeemableRepo RedeemableRepository
}

func NewUserRepository(cfg *UserRConfig) UserRepository {
	return &userRepositoryImpl{
		db:             cfg.DB,
		levelRepo:      cfg.LevelRepo,
		redeemableRepo: cfg.RedeemableRepo,
	}
}

func (r *userRepositoryImpl) FindById(id int) (*entity.User, error) {
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

func (r *userRepositoryImpl) FindByIdentifierAndRole(identifier string, roleId int) (*entity.User, error) {
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

func (r *userRepositoryImpl) FindDetailById(id int) (*entity.User, error) {
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

func (r *userRepositoryImpl) FindByReferral(referral string) (*entity.User, error) {
	var res *entity.User
	err := r.db.Where("referral = ?", referral).First(&res)
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

func (r *userRepositoryImpl) LevelUp(tx *gorm.DB, id int, totalTransaction int64) (*entity.User, error) {
	user, err := r.FindById(id)
	if err != nil {
		return nil, err
	}

	level, err := r.levelRepo.FindValidByTotalTransaction(totalTransaction)
	if err != nil {
		return nil, err
	}

	if user.LevelId != level.ID {
		err = tx.Model(&entity.User{}).Where("id = ?", user.ID).Update("level_id", level.ID).Error
		if err != nil {
			return nil, err
		}
	}

	return user, nil
}

func (r *userRepositoryImpl) AddRedeemablePoint(tx *gorm.DB, id int) error {
	user, err := r.FindById(id)
	if err != nil {
		return err
	}

	if user.Level.Point == 0 {
		return nil
	}

	err = r.redeemableRepo.AddPoint(tx, user.ID, user.Level.Point)
	if err != nil {
		return err
	}

	return nil
}
