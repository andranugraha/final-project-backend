package entity

import (
	"final-project-backend/utils/constant"

	"gorm.io/gorm"
)

type Course struct {
	ID           int       `json:"id" gorm:"primaryKey"`
	Title        string    `json:"title"`
	Slug         string    `json:"slug"`
	Summary      string    `json:"summary"`
	Content      string    `json:"content"`
	ImgThumbnail string    `json:"imgThumbnail"`
	ImgUrl       string    `json:"imgUrl"`
	AuthorName   string    `json:"authorName"`
	Status       string    `json:"status"`
	Price        float64   `json:"price"`
	CategoryId   int       `json:"categoryId"`
	Category     *Category `json:"category,omitempty" gorm:"foreignKey:CategoryId"`
	Tags         []*Tag    `json:"tags" gorm:"many2many:course_tags;"`
	gorm.Model   `json:"-"`
}

type CourseParams struct {
	Keyword        string
	CategoryId     int
	TagIds         []int
	Sort           string
	Limit          int
	Page           int
	Status         string
	ProgressStatus string
}

func (c *CourseParams) Scope() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if c.Keyword != "" {
			db = db.Where("title ILIKE ?", "%"+c.Keyword+"%")
		}
		if c.CategoryId > 0 {
			db = db.Where("category_id = ?", c.CategoryId)
		}
		if len(c.TagIds) > 0 {
			db = db.Joins("JOIN course_tags ON course_tags.course_id = courses.id AND course_tags.tag_id IN (?)", c.TagIds).
				Group("courses.id").
				Having("COUNT(course_tags.id) = ?", len(c.TagIds))
		}
		if c.Status != "" {
			db = db.Where("courses.status = ?", c.Status)
		}
		if c.ProgressStatus != "" {
			db = db.Where("user_courses.status = ?", c.ProgressStatus)
		}
		return db
	}
}

func (c *CourseParams) Offset() int {
	return (c.Page - 1) * c.Limit
}

func NewCourseParams(s string, categoryId int, tagIds []int, sort string, limit, page int, roleId int, status string) CourseParams {
	return CourseParams{
		Keyword: s,
		CategoryId: func() int {
			if categoryId > 0 {
				return categoryId
			}
			return 0
		}(),
		TagIds: tagIds,
		Sort: func() string {
			if sort != "" && sort == "oldest" {
				return "created_at ASC"
			}
			return "created_at DESC"
		}(),
		Limit: func() int {
			if limit > 0 {
				return limit
			}
			return 10
		}(),
		Page: func() int {
			if page > 1 {
				return page
			}
			return 1
		}(),
		Status: func() string {
			if roleId == constant.AdminRoleId {
				if status == constant.DraftStatus || status == constant.PublishStatus {
					return status
				}
				return ""
			}
			return constant.PublishStatus
		}(),
		ProgressStatus: func() string {
			if roleId == constant.UserRoleId && (status == constant.CourseStatusCompleted || status == constant.CourseStatusInProgress) {
				return status
			}
			return ""
		}(),
	}
}
