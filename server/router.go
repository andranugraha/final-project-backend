package server

import (
	"final-project-backend/handler"
	"final-project-backend/usecase"
	"final-project-backend/utils/errors"
	"final-project-backend/utils/response"
	"net/http"

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

	router.NoRoute(func(c *gin.Context) {
		response.SendError(c, http.StatusNotFound, errors.ErrCodeRouteNotFound, errors.ErrRouteNotFound.Error())
	})

	admin := router.Group("/admin")
	{
		admin.POST("/sign-in", h.AdminSignIn)
	}

	router.POST("/sign-in", h.SignIn)

	return router
}
