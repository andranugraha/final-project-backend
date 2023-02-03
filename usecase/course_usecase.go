package usecase

import (
	"final-project-backend/dto"
	"final-project-backend/entity"
	"final-project-backend/repository"
	"final-project-backend/utils/constant"
	errResp "final-project-backend/utils/errors"
	"final-project-backend/utils/storage"
	"strings"
)

type CourseUsecase interface {
	CreateCourse(dto.CreateCourseRequest) (*entity.Course, error)
	GetCourse(string, int) (*dto.GetCourseResponse, error)
	GetCourses(entity.CourseParams) ([]entity.Course, int64, int, error)
	GetUserCourses(int, entity.CourseParams) ([]entity.Course, int64, int, error)
	GetUserCourse(int, string) (*entity.UserCourse, error)
	GetTrendingCourses() ([]entity.Course, error)
	UpdateCourse(string, dto.UpdateCourseRequest) (*entity.Course, error)
	DeleteCourse(string) error
	CompleteCourse(int, string) error
}

type courseUsecaseImpl struct {
	courseRepo      repository.CourseRepository
	userCourseRepo  repository.UserCourseRepository
	favoriteUsecase FavoriteUsecase
	cartRepo        repository.CartRepository
	transactionRepo repository.TransactionRepository
}

type CourseUConfig struct {
	CourseRepo      repository.CourseRepository
	UserCourseRepo  repository.UserCourseRepository
	FavoriteUsecase FavoriteUsecase
	CartRepo        repository.CartRepository
	TransactionRepo repository.TransactionRepository
}

func NewCourseUsecase(cfg *CourseUConfig) CourseUsecase {
	return &courseUsecaseImpl{
		courseRepo:      cfg.CourseRepo,
		userCourseRepo:  cfg.UserCourseRepo,
		favoriteUsecase: cfg.FavoriteUsecase,
		cartRepo:        cfg.CartRepo,
		transactionRepo: cfg.TransactionRepo,
	}
}

func (u *courseUsecaseImpl) GetCourse(slug string, userId int) (*dto.GetCourseResponse, error) {
	course, err := u.courseRepo.FindBySlug(slug)
	if err != nil {
		return nil, err
	}

	res := &dto.GetCourseResponse{}
	res.FromCourse(*course)

	transaction, err := u.transactionRepo.FindBoughtByUserIdAndCourseId(userId, course.ID)
	if err != nil {
		return nil, err
	}

	res.IsBought = transaction != nil

	_, err = u.userCourseRepo.FindByUserIdAndCourseId(userId, course.ID)
	if err != nil && err != errResp.ErrUserCourseNotFound {
		return nil, err
	}

	res.IsEnrolled = err == nil

	_, err = u.cartRepo.FindByUserIdAndCourseId(userId, course.ID)
	if err != nil && err != errResp.ErrCartNotFound {
		return nil, err
	}

	res.IsInCart = err == nil

	res.IsFavorite = u.favoriteUsecase.CheckIsFavoriteCourse(userId, course.ID)
	res.TotalFavorite = u.favoriteUsecase.GetTotalFavorited(course.ID)
	res.TotalCompleted = u.userCourseRepo.CountByCourseIdAndStatus(course.ID, constant.CourseStatusCompleted)

	return res, nil
}

func (u *courseUsecaseImpl) GetCourses(params entity.CourseParams) ([]entity.Course, int64, int, error) {
	return u.courseRepo.FindAll(params)
}

func (u *courseUsecaseImpl) GetTrendingCourses() ([]entity.Course, error) {
	return u.courseRepo.FindTrending()
}

func (u *courseUsecaseImpl) GetUserCourses(userId int, params entity.CourseParams) ([]entity.Course, int64, int, error) {
	return u.userCourseRepo.FindByUserId(userId, params)
}

func (u *courseUsecaseImpl) GetUserCourse(userId int, slug string) (*entity.UserCourse, error) {
	return u.userCourseRepo.FindByUserIdAndCourseSlug(userId, slug)
}

func (u *courseUsecaseImpl) CreateCourse(req dto.CreateCourseRequest) (*entity.Course, error) {
	courseSlug := u.generateSlug(req.Title)
	imgUrl, err := storage.Upload(&req.Image, courseSlug)
	if err != nil {
		return nil, err
	}

	course := req.ToCourse(courseSlug, *imgUrl)

	createdCourse, err := u.courseRepo.Insert(course)
	if err != nil {
		return nil, err
	}

	return createdCourse, nil
}

func (u *courseUsecaseImpl) UpdateCourse(slug string, req dto.UpdateCourseRequest) (*entity.Course, error) {
	course, err := u.courseRepo.FindBySlug(slug)
	if err != nil {
		return nil, err
	}

	var newSlug string = course.Slug
	if req.Title != course.Title {
		course.Title = req.Title
		newSlug = u.generateSlug(req.Title)

		if req.Image == nil {
			imgUrl, err := storage.Rename(course.Slug, newSlug)
			if err != nil {
				return nil, err
			}

			course.ImgThumbnail = imgUrl.ThumbnailUrl
			course.ImgUrl = imgUrl.Url
		}
	}

	if req.Image != nil {
		storage.Delete(course.Slug)

		imgUrl, err := storage.Upload(req.Image, newSlug)
		if err != nil {
			return nil, err
		}

		course.ImgThumbnail = imgUrl.ThumbnailUrl
		course.ImgUrl = imgUrl.Url
	}

	course.Slug = newSlug
	course.Summary = req.Summary
	course.Price = req.Price
	course.Content = req.Content
	course.AuthorName = req.AuthorName
	course.Status = req.Status

	course.Tags = func() []*entity.Tag {
		var tags []*entity.Tag
		for _, tag := range req.Tags {
			tags = append(tags, &entity.Tag{
				Name: tag,
			})
		}

		return tags
	}()

	updatedCourse, err := u.courseRepo.Update(*course)
	if err != nil {
		return nil, err
	}

	return updatedCourse, nil
}

func (u *courseUsecaseImpl) DeleteCourse(slug string) error {
	err := u.courseRepo.Delete(slug)
	if err != nil {
		return err
	}

	return nil
}

func (u *courseUsecaseImpl) CompleteCourse(userId int, slug string) error {
	course, err := u.courseRepo.FindBySlug(slug)
	if err != nil {
		return err
	}

	userCourse, err := u.userCourseRepo.FindByUserIdAndCourseId(userId, course.ID)
	if err != nil {
		return err
	}

	if userCourse.Status == constant.CourseStatusCompleted {
		return errResp.ErrCourseAlreadyCompleted
	}

	userCourse.Status = constant.CourseStatusCompleted

	_, err = u.userCourseRepo.Complete(*userCourse)
	if err != nil {
		return err
	}

	return nil
}

func (u *courseUsecaseImpl) generateSlug(title string) string {
	trimmed := strings.TrimSpace(title)
	slug := strings.ReplaceAll(strings.ToLower(trimmed), " ", "-")

	return slug
}
