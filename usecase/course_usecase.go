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
	GetCourse(string) (*entity.Course, error)
	GetCourses(entity.CourseParams) ([]entity.Course, int64, int, error)
	GetTrendingCourses() ([]entity.Course, error)
	UpdateCourse(string, dto.UpdateCourseRequest) (*entity.Course, error)
	DeleteCourse(string) error
	CompleteCourse(int, string) error
}

type courseUsecaseImpl struct {
	courseRepo     repository.CourseRepository
	userCourseRepo repository.UserCourseRepository
}

type CourseUConfig struct {
	CourseRepo     repository.CourseRepository
	UserCourseRepo repository.UserCourseRepository
}

func NewCourseUsecase(cfg *CourseUConfig) CourseUsecase {
	return &courseUsecaseImpl{
		courseRepo:     cfg.CourseRepo,
		userCourseRepo: cfg.UserCourseRepo,
	}
}

func (u *courseUsecaseImpl) GetCourse(slug string) (*entity.Course, error) {
	course, err := u.courseRepo.FindBySlug(slug)
	if err != nil {
		return nil, err
	}

	return course, nil
}

func (u *courseUsecaseImpl) GetCourses(params entity.CourseParams) ([]entity.Course, int64, int, error) {
	return u.courseRepo.FindAll(params)
}

func (u *courseUsecaseImpl) GetTrendingCourses() ([]entity.Course, error) {
	return u.courseRepo.FindTrending()
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
		err := storage.Delete(course.Slug)
		if err != nil {
			return nil, err
		}

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
