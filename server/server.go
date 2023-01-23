package server

import (
	"log"

	"final-project-backend/db"
	"final-project-backend/repository"
	"final-project-backend/usecase"
	"final-project-backend/utils/auth"

	"github.com/gin-gonic/gin"
)

func createRouter() *gin.Engine {
	userRepo := repository.NewUserRepository(&repository.UserRConfig{
		DB: db.Get(),
	})
	authUsecase := usecase.NewAuthUsecase(&usecase.AuthUConfig{
		UserRepo:      userRepo,
		BcryptUsecase: auth.AuthUtilImpl{},
	})
	userUsecase := usecase.NewUserUsecase(&usecase.UserUConfig{
		UserRepo: userRepo,
	})

	tagRepo := repository.NewTagRepository(&repository.TagRConfig{
		DB: db.Get(),
	})

	courseRepo := repository.NewCourseRepository(&repository.CourseRConfig{
		DB:      db.Get(),
		TagRepo: tagRepo,
	})
	courseUsecase := usecase.NewCourseUsecase(&usecase.CourseUConfig{
		CourseRepo: courseRepo,
	})

	favoriteRepo := repository.NewFavoriteRepository(&repository.FavoriteRConfig{
		DB: db.Get(),
	})
	favoriteUsecase := usecase.NewFavoriteUsecase(&usecase.FavoriteUConfig{
		FavoriteRepo: favoriteRepo,
	})

	cartRepo := repository.NewCartRepository(&repository.CartRConfig{
		DB: db.Get(),
	})

	cartUsecase := usecase.NewCartUsecase(&usecase.CartUConfig{
		CartRepo: cartRepo,
	})

	invoiceRepo := repository.NewInvoiceRepository(&repository.InvoiceRConfig{
		DB: db.Get(),
	})
	invoiceUsecase := usecase.NewInvoiceUsecase(&usecase.InvoiceUConfig{
		InvoiceRepo: invoiceRepo,
	})

	return NewRouter(&RouterConfig{
		AuthUsecase:     authUsecase,
		UserUsecase:     userUsecase,
		CourseUsecase:   courseUsecase,
		FavoriteUsecase: favoriteUsecase,
		CartUsecase:     cartUsecase,
		InvoiceUsecase:  invoiceUsecase,
	})
}

func Init() {
	r := createRouter()
	err := r.Run()
	if err != nil {
		log.Println("error while running server", err)
		return
	}
}
