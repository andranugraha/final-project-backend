package handler

import "final-project-backend/usecase"

type Handler struct {
	authUsecase        usecase.AuthUsecase
	userUsecase        usecase.UserUsecase
	courseUsecase      usecase.CourseUsecase
	favoriteUsecase    usecase.FavoriteUsecase
	cartUsecase        usecase.CartUsecase
	invoiceUsecase     usecase.InvoiceUsecase
	userVoucherUsecase usecase.UserVoucherUsecase
	categoryUsecase    usecase.CategoryUsecase
	tagUsecase         usecase.TagUsecase
}

type Config struct {
	AuthUsecase        usecase.AuthUsecase
	UserUsecase        usecase.UserUsecase
	CourseUsecase      usecase.CourseUsecase
	FavoriteUsecase    usecase.FavoriteUsecase
	CartUsecase        usecase.CartUsecase
	InvoiceUsecase     usecase.InvoiceUsecase
	UserVoucherUsecase usecase.UserVoucherUsecase
	CategoryUsecase    usecase.CategoryUsecase
	TagUsecase         usecase.TagUsecase
}

func New(cfg *Config) *Handler {
	return &Handler{
		authUsecase:        cfg.AuthUsecase,
		userUsecase:        cfg.UserUsecase,
		courseUsecase:      cfg.CourseUsecase,
		favoriteUsecase:    cfg.FavoriteUsecase,
		cartUsecase:        cfg.CartUsecase,
		invoiceUsecase:     cfg.InvoiceUsecase,
		userVoucherUsecase: cfg.UserVoucherUsecase,
		categoryUsecase:    cfg.CategoryUsecase,
		tagUsecase:         cfg.TagUsecase,
	}
}
