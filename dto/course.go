package dto

import (
	"final-project-backend/entity"
	"final-project-backend/utils/storage"
	"mime/multipart"
)

type CreateCourseRequest struct {
	Title      string               `form:"title" binding:"required"`
	Summary    string               `form:"summary" binding:"required"`
	Content    string               `form:"content" binding:"required"`
	AuthorName string               `form:"authorName" binding:"required"`
	Status     string               `form:"status" binding:"required"`
	CategoryId int                  `form:"categoryId" binding:"required"`
	Tags       []string             `form:"tags" binding:"required"`
	Price      float64              `form:"price" binding:"required"`
	Image      multipart.FileHeader `form:"image" binding:"required"`
}

type CreateCourseResponse struct {
	ID int `json:"id"`
}

func (dto CreateCourseRequest) ToCourse(slug string, image storage.StoredImage) entity.Course {
	return entity.Course{
		Title:      dto.Title,
		Slug:       slug,
		Summary:    dto.Summary,
		Content:    dto.Content,
		AuthorName: dto.AuthorName,
		Status:     dto.Status,
		CategoryId: dto.CategoryId,
		Price:      dto.Price,
		Tags: func() []*entity.Tag {
			var tags []*entity.Tag
			for _, tag := range dto.Tags {
				tags = append(tags, &entity.Tag{
					Name: tag,
				})
			}

			return tags
		}(),
	}
}

type UpdateCourseRequest struct {
	Title      string                `form:"title" binding:"required"`
	Summary    string                `form:"summary" binding:"required"`
	Content    string                `form:"content" binding:"required"`
	AuthorName string                `form:"authorName" binding:"required"`
	Status     string                `form:"status" binding:"required"`
	CategoryId int                   `form:"categoryId" binding:"required"`
	Tags       []string              `form:"tags" binding:"required"`
	Price      float64               `form:"price" binding:"required"`
	Image      *multipart.FileHeader `form:"image"`
}

type GetCourseResponse struct {
	ID             int              `json:"id"`
	Title          string           `json:"title"`
	Slug           string           `json:"slug"`
	Summary        string           `json:"summary"`
	Content        string           `json:"content"`
	AuthorName     string           `json:"authorName"`
	Status         string           `json:"status"`
	CategoryId     int              `json:"categoryId"`
	Price          float64          `json:"price"`
	ImgUrl         string           `json:"imgUrl"`
	ImgThumbnail   string           `json:"imgThumbnail"`
	Category       *entity.Category `json:"category"`
	Tags           []*entity.Tag    `json:"tags"`
	IsBought       bool             `json:"isBought"`
	IsEnrolled     bool             `json:"isEnrolled"`
	IsInCart       bool             `json:"isInCart"`
	IsFavorite     bool             `json:"isFavorite"`
	TotalFavorite  int              `json:"totalFavorite"`
	TotalCompleted int              `json:"totalCompleted"`
}

func (dto *GetCourseResponse) FromCourse(course entity.Course) {
	dto.ID = course.ID
	dto.Title = course.Title
	dto.Slug = course.Slug
	dto.Summary = course.Summary
	dto.Content = course.Content
	dto.AuthorName = course.AuthorName
	dto.Status = course.Status
	dto.CategoryId = course.CategoryId
	dto.Price = course.Price
	dto.ImgUrl = course.ImgUrl
	dto.ImgThumbnail = course.ImgThumbnail
	dto.Category = course.Category
	dto.Tags = course.Tags
}
