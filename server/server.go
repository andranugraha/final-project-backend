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
	userUsecase := usecase.NewUserUsecase(&usecase.UserUConfig{
		UserRepo:      userRepo,
		BcryptUsecase: auth.AuthUtilImpl{},
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

	return NewRouter(&RouterConfig{
		UserUsecase:   userUsecase,
		CourseUsecase: courseUsecase,
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
