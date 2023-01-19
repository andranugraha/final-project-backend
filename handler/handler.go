package handler

import "final-project-backend/usecase"

type Handler struct {
	userUsecase   usecase.UserUsecase
	courseUsecase usecase.CourseUsecase
}

type Config struct {
	UserUsecase   usecase.UserUsecase
	CourseUsecase usecase.CourseUsecase
}

func New(cfg *Config) *Handler {
	return &Handler{
		userUsecase:   cfg.UserUsecase,
		courseUsecase: cfg.CourseUsecase,
	}
}
