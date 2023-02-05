package server

import (
	"fmt"
	"log"

	"final-project-backend/config"
	"final-project-backend/db"
	"final-project-backend/repository"
	"final-project-backend/usecase"
	"final-project-backend/utils/auth"
	"final-project-backend/utils/storage"

	"github.com/gin-gonic/gin"
)

func createRouter() *gin.Engine {
	categoryRepo := repository.NewCategoryRepository(&repository.CategoryRConfig{
		DB: db.Get(),
	})
	categoryUsecase := usecase.NewCategoryUsecase(&usecase.CategoryUConfig{
		CategoryRepo: categoryRepo,
	})

	levelRepo := repository.NewLevelRepository(&repository.LevelRConfig{
		DB: db.Get(),
	})

	redeemableRepo := repository.NewRedeemableRepository(&repository.RedeemableRConfig{
		DB: db.Get(),
	})

	userRepo := repository.NewUserRepository(&repository.UserRConfig{
		DB:             db.Get(),
		LevelRepo:      levelRepo,
		RedeemableRepo: redeemableRepo,
	})
	authUtil := auth.NewAuthUtil()
	authUsecase := usecase.NewAuthUsecase(&usecase.AuthUConfig{
		UserRepo:    userRepo,
		UtilUsecase: authUtil,
	})
	userUsecase := usecase.NewUserUsecase(&usecase.UserUConfig{
		UserRepo: userRepo,
	})

	tagRepo := repository.NewTagRepository(&repository.TagRConfig{
		DB: db.Get(),
	})
	tagUsecase := usecase.NewTagUsecase(&usecase.TagUConfig{
		TagRepo: tagRepo,
	})

	favoriteRepo := repository.NewFavoriteRepository(&repository.FavoriteRConfig{
		DB: db.Get(),
	})
	favoriteUsecase := usecase.NewFavoriteUsecase(&usecase.FavoriteUConfig{
		FavoriteRepo: favoriteRepo,
	})

	userCourseRepo := repository.NewUserCourseRepository(&repository.UserCourseRConfig{
		DB:       db.Get(),
		UserRepo: userRepo,
	})

	transactionRepo := repository.NewTransactionRepository(&repository.TransactionRConfig{
		DB: db.Get(),
	})
	courseRepo := repository.NewCourseRepository(&repository.CourseRConfig{
		DB:      db.Get(),
		TagRepo: tagRepo,
	})
	cartRepo := repository.NewCartRepository(&repository.CartRConfig{
		DB: db.Get(),
	})
	storageUtil := storage.NewStorageUtil()
	courseUsecase := usecase.NewCourseUsecase(&usecase.CourseUConfig{
		CourseRepo:      courseRepo,
		UserCourseRepo:  userCourseRepo,
		FavoriteUsecase: favoriteUsecase,
		CartRepo:        cartRepo,
		TransactionRepo: transactionRepo,
		StorageUtil:     storageUtil,
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
		AuthUsecase:        authUsecase,
		UserUsecase:        userUsecase,
		CourseUsecase:      courseUsecase,
		FavoriteUsecase:    favoriteUsecase,
		CartUsecase:        cartUsecase,
		InvoiceUsecase:     invoiceUsecase,
		UserVoucherUsecase: UserVoucherUsecase,
		CategoryUsecase:    categoryUsecase,
		TagUsecase:         tagUsecase,
	})
}

func Init() {
	r := createRouter()
	err := r.Run(fmt.Sprintf(":%s", config.Port))
	if err != nil {
		log.Println("error while running server", err)
		return
	}
}
