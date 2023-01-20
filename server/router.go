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
	AuthUsecase   usecase.AuthUsecase
	UserUsecase   usecase.UserUsecase
	CourseUsecase usecase.CourseUsecase
}

func NewRouter(cfg *RouterConfig) *gin.Engine {
	router := gin.Default()
	h := handler.New(&handler.Config{
		AuthUsecase:   cfg.AuthUsecase,
		UserUsecase:   cfg.UserUsecase,
		CourseUsecase: cfg.CourseUsecase,
	})

	router.Static("/docs", "swagger-ui")

	router.NoRoute(func(c *gin.Context) {
		response.SendError(c, http.StatusNotFound, errors.ErrCodeRouteNotFound, errors.ErrRouteNotFound.Error())
	})

	v1 := router.Group("/api/v1")
	{
		admin := v1.Group("/admin")
		{
			admin.POST("/sign-in", h.AdminSignIn)
			authenticated := admin.Group("/", middleware.Authenticated, middleware.Admin)
			{
				course := authenticated.Group("/courses")
				{
					course.GET("/", h.GetCourses)
					course.POST("/", h.CreateCourse)
					course.PUT("/:slug", h.UpdateCourse)
					course.DELETE("/:slug", h.DeleteCourse)
				}
			}
		}

		v1.POST("/sign-in", h.SignIn)
		v1.POST("/sign-up", h.SignUp)

		user := v1.Group("/user", middleware.Authenticated)
		{
			user.GET("/", h.GetUserDetail)
			user.PUT("/", h.UpdateUserDetail)
		}

		course := v1.Group("/courses")
		{
			course.GET("/", h.GetCourses)
			authenticatedCourse := course.Group("/", middleware.Authenticated)
			{
				authenticatedCourse.GET("/:slug", h.GetCourse)
			}
		}
	}

	return router
}
