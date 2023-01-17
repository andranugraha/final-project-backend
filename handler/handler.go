package handler

import "final-project-backend/usecase"

type Handler struct {
	userUsecase usecase.UserUsecase
}

type Config struct {
	UserUsecase usecase.UserUsecase
}

func New(cfg *Config) *Handler {
	return &Handler{
		userUsecase: cfg.UserUsecase,
	}
}
