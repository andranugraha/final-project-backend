package server

import (
	"final-project-backend/handler"
	"final-project-backend/middleware"
	"final-project-backend/usecase"
	"final-project-backend/utils/errors"
	"final-project-backend/utils/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RouterConfig struct {
	UserUsecase   usecase.UserUsecase
	CourseUsecase usecase.CourseUsecase
}

func NewRouter(cfg *RouterConfig) *gin.Engine {
	router := gin.Default()
	h := handler.New(&handler.Config{
		UserUsecase:   cfg.UserUsecase,
		CourseUsecase: cfg.CourseUsecase,
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
	router.POST("/sign-up", h.SignUp)

	course := router.Group("/courses", middleware.Authenticated)
	{
		course.GET("/:slug", h.GetCourse)
		course.POST("/", middleware.Admin, h.CreateCourse)
		course.PUT("/:slug", middleware.Admin, h.UpdateCourse)
		course.DELETE("/:slug", middleware.Admin, h.DeleteCourse)
	}

	return router
}
