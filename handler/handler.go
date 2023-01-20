package handler

import "final-project-backend/usecase"

type Handler struct {
	authUsecase   usecase.AuthUsecase
	userUsecase   usecase.UserUsecase
	courseUsecase usecase.CourseUsecase
}

type Config struct {
	AuthUsecase   usecase.AuthUsecase
	UserUsecase   usecase.UserUsecase
	CourseUsecase usecase.CourseUsecase
}

func New(cfg *Config) *Handler {
	return &Handler{
		authUsecase:   cfg.AuthUsecase,
		userUsecase:   cfg.UserUsecase,
		courseUsecase: cfg.CourseUsecase,
	}
}
