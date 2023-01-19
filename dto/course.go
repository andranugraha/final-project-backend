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
		Title:        dto.Title,
		Slug:         slug,
		Summary:      dto.Summary,
		Content:      dto.Content,
		AuthorName:   dto.AuthorName,
		Status:       dto.Status,
		CategoryId:   dto.CategoryId,
		Price:        dto.Price,
		ImgUrl:       image.Url,
		ImgThumbnail: image.ThumbnailUrl,
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
