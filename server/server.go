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
	levelRepo := repository.NewLevelRepository(&repository.LevelRConfig{
		DB: db.Get(),
	})

	userRepo := repository.NewUserRepository(&repository.UserRConfig{
		DB:        db.Get(),
		LevelRepo: levelRepo,
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

	transactionRepo := repository.NewTransactionRepository(&repository.TransactionRConfig{
		DB: db.Get(),
	})

	cartRepo := repository.NewCartRepository(&repository.CartRConfig{
		DB: db.Get(),
	})
	cartUsecase := usecase.NewCartUsecase(&usecase.CartUConfig{
		CartRepo:        cartRepo,
		CourseRepo:      courseRepo,
		TransactionRepo: transactionRepo,
	})

	voucherRepo := repository.NewVoucherRepository(&repository.VoucherRConfig{
		DB: db.Get(),
	})

	userVoucherRepo := repository.NewUserVoucherRepository(&repository.UserVoucherRConfig{
		DB:          db.Get(),
		VoucherRepo: voucherRepo,
	})
	UserVoucherUsecase := usecase.NewUserVoucherUsecase(&usecase.UserVoucherUConfig{
		UserVoucherRepo: userVoucherRepo,
	})

	userCourseRepo := repository.NewUserCourseRepository(&repository.UserCourseRConfig{
		DB: db.Get(),
	})

	invoiceRepo := repository.NewInvoiceRepository(&repository.InvoiceRConfig{
		DB:              db.Get(),
		CartRepo:        cartRepo,
		UserVoucherRepo: userVoucherRepo,
		UserCourseRepo:  userCourseRepo,
		UserRepo:        userRepo,
	})
	invoiceUsecase := usecase.NewInvoiceUsecase(&usecase.InvoiceUConfig{
		InvoiceRepo:        invoiceRepo,
		CartRepo:           cartRepo,
		UserVoucherUsecase: UserVoucherUsecase,
		UserUsecase:        userUsecase,
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
