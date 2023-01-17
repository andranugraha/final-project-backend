package server

import (
	"final-project-backend/handler"
	"final-project-backend/usecase"

	"github.com/gin-gonic/gin"
)

type RouterConfig struct {
	UserUsecase usecase.UserUsecase
}

func NewRouter(cfg *RouterConfig) *gin.Engine {
	router := gin.Default()
	h := handler.New(&handler.Config{
		UserUsecase: cfg.UserUsecase,
	})

	router.Static("/docs", "swagger-ui")

	router.POST("/sign-in", h.SignIn)

	return router
}
