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
	AuthUsecase     usecase.AuthUsecase
	UserUsecase     usecase.UserUsecase
	CourseUsecase   usecase.CourseUsecase
	FavoriteUsecase usecase.FavoriteUsecase
	CartUsecase     usecase.CartUsecase
	InvoiceUsecase  usecase.InvoiceUsecase
}

func NewRouter(cfg *RouterConfig) *gin.Engine {
	router := gin.Default()
	h := handler.New(&handler.Config{
		AuthUsecase:     cfg.AuthUsecase,
		UserUsecase:     cfg.UserUsecase,
		CourseUsecase:   cfg.CourseUsecase,
		FavoriteUsecase: cfg.FavoriteUsecase,
		CartUsecase:     cfg.CartUsecase,
		InvoiceUsecase:  cfg.InvoiceUsecase,
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

				invoice := authenticated.Group("/invoices")
				{
					invoice.POST("/:invoiceId/:action", h.ConfirmInvoice)
				}
			}
		}

		v1.POST("/sign-in", h.SignIn)
		v1.POST("/sign-up", h.SignUp)

		user := v1.Group("/user", middleware.Authenticated, middleware.User)
		{
			user.GET("/", h.GetUserDetail)
			user.PUT("/", h.UpdateUserDetail)
		}

		course := v1.Group("/courses")
		{
			course.GET("/", h.GetCourses)
			course.GET("/trending", h.GetTrendingCourses)
			authenticatedCourse := course.Group("/", middleware.Authenticated)
			{
				authenticatedCourse.GET("/:slug", h.GetCourse)
			}
		}

		favorite := v1.Group("/favorites", middleware.Authenticated)
		{
			favorite.GET("/", h.GetFavoriteCourses)
			favorite.GET("/:courseId", h.CheckIsFavoriteCourse)
			favorite.POST("/:courseId/:action", h.SaveUnsaveFavoriteCourse)
		}

		cart := v1.Group("/carts", middleware.Authenticated, middleware.User)
		{
			cart.GET("/", h.GetCart)
			cart.POST("/:courseId", h.AddToCart)
			cart.DELETE("/:courseId", h.RemoveFromCart)
		}

		invoice := v1.Group("/invoices", middleware.Authenticated)
		{
			invoice.GET("/", h.GetInvoices)
			invoice.GET("/:invoiceId", h.GetInvoice)
			userInvoice := invoice.Group("/", middleware.User)
			{
				userInvoice.POST("/", h.Checkout)
				userInvoice.POST("/:invoiceId/pay", h.PayInvoice)
			}

			invoice.POST("/:invoiceId/confirm", middleware.Admin, h.ConfirmInvoice)
		}
	}

	return router
}
