package usecase

import (
	"final-project-backend/dto"
	"final-project-backend/entity"
	"final-project-backend/repository"
)

type CartUsecase interface {
	GetCart(userId int) (*dto.GetCartResponse, error)
	AddToCart(userId, courseId int) error
	RemoveFromCart(userId, courseId int) error
}

type cartUsecaseImpl struct {
	cartRepo   repository.CartRepository
	courseRepo repository.CourseRepository
}

type CartUConfig struct {
	CartRepo   repository.CartRepository
	CourseRepo repository.CourseRepository
}

func NewCartUsecase(cfg *CartUConfig) CartUsecase {
	return &cartUsecaseImpl{
		cartRepo:   cfg.CartRepo,
		courseRepo: cfg.CourseRepo,
	}
}

func (u *cartUsecaseImpl) GetCart(userId int) (*dto.GetCartResponse, error) {
	cart, err := u.cartRepo.FindByUserId(userId)
	if err != nil {
		return nil, err
	}

	res := &dto.GetCartResponse{}
	res.FromCart(cart)

	return res, nil
}

func (u *cartUsecaseImpl) AddToCart(userId, courseId int) error {
	_, err := u.courseRepo.FindPublishedById(courseId)
	if err != nil {
		return err
	}

	cart := entity.Cart{
		UserId:   userId,
		CourseId: courseId,
	}

	return u.cartRepo.Insert(cart)
}

func (u *cartUsecaseImpl) RemoveFromCart(userId, courseId int) error {
	cart := entity.Cart{
		UserId:   userId,
		CourseId: courseId,
	}
	return u.cartRepo.Delete(cart)
}
