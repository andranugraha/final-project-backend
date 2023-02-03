package handler

import (
	"errors"
	"final-project-backend/dto"
	"final-project-backend/entity"
	errResp "final-project-backend/utils/errors"
	"final-project-backend/utils/response"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateCourse(c *gin.Context) {
	var req dto.CreateCourseRequest
	if err := c.ShouldBind(&req); err != nil {
		response.SendError(c, http.StatusBadRequest, errResp.ErrCodeBadRequest, err.Error())
		return
	}

	course, err := h.courseUsecase.CreateCourse(req)
	if err != nil {
		if errors.Is(err, errResp.ErrDuplicateTitle) {
			response.SendError(c, http.StatusBadRequest, errResp.ErrCodeDuplicate, err.Error())
			return
		}
		response.SendError(c, http.StatusInternalServerError, errResp.ErrCodeInternalServerError, err.Error())
		return
	}

	response.SendSuccess(c, http.StatusOK, course)
}

func (h *Handler) GetTrendingCourses(c *gin.Context) {
	courses, err := h.courseUsecase.GetTrendingCourses()
	if err != nil {
		response.SendError(c, http.StatusInternalServerError, errResp.ErrCodeInternalServerError, errResp.ErrInternalServerError.Error())
		return
	}

	response.SendSuccess(c, http.StatusOK, courses)
}

func (h *Handler) GetCourses(c *gin.Context) {
	roleId := c.GetInt("roleId")
	categoryId, _ := strconv.Atoi(c.Query("categoryId"))
	splitTagIds := strings.Split(c.Query("tagIds"), ",")
	var tagIds []int
	for _, tagId := range splitTagIds {
		if tagId != "" {
			id, _ := strconv.Atoi(tagId)
			tagIds = append(tagIds, id)
		}
	}

	limit, _ := strconv.Atoi(c.Query("limit"))
	page, _ := strconv.Atoi(c.Query("page"))
	params := entity.NewCourseParams(c.Query("title"), categoryId, tagIds, c.Query("sort"), limit, page, roleId, c.Query("status"))

	courses, totalRows, totalPages, err := h.courseUsecase.GetCourses(params)
	if err != nil {
		response.SendError(c, http.StatusInternalServerError, errResp.ErrCodeInternalServerError, errResp.ErrInternalServerError.Error())
		return
	}

	response.SendSuccessWithPagination(c, http.StatusOK, courses, totalRows, totalPages)
}

func (h *Handler) GetCourse(c *gin.Context) {
	slug := c.Param("slug")
	userId := c.GetInt("userId")

	course, err := h.courseUsecase.GetCourse(slug, userId)
	if err != nil {
		if errors.Is(err, errResp.ErrCourseNotFound) {
			response.SendError(c, http.StatusNotFound, errResp.ErrCodeNotFound, err.Error())
			return
		}

		response.SendError(c, http.StatusInternalServerError, errResp.ErrCodeInternalServerError, errResp.ErrInternalServerError.Error())
		return
	}

	response.SendSuccess(c, http.StatusOK, course)
}

func (h *Handler) UpdateCourse(c *gin.Context) {
	var req dto.UpdateCourseRequest
	if err := c.ShouldBind(&req); err != nil {
		response.SendError(c, http.StatusBadRequest, errResp.ErrCodeBadRequest, err.Error())
		return
	}

	slug := c.Param("slug")
	course, err := h.courseUsecase.UpdateCourse(slug, req)
	if err != nil {
		if errors.Is(err, errResp.ErrCourseNotFound) {
			response.SendError(c, http.StatusNotFound, errResp.ErrCodeNotFound, err.Error())
			return
		}

		if err == errResp.ErrDuplicateTitle {
			response.SendError(c, http.StatusBadRequest, errResp.ErrCodeDuplicate, err.Error())
			return
		}

		response.SendError(c, http.StatusInternalServerError, errResp.ErrCodeInternalServerError, errResp.ErrInternalServerError.Error())
		return
	}

	response.SendSuccess(c, http.StatusOK, course)
}

func (h *Handler) DeleteCourse(c *gin.Context) {
	slug := c.Param("slug")
	err := h.courseUsecase.DeleteCourse(slug)
	if err != nil {
		if errors.Is(err, errResp.ErrCourseNotFound) {
			response.SendError(c, http.StatusNotFound, errResp.ErrCodeNotFound, err.Error())
			return
		}

		response.SendError(c, http.StatusInternalServerError, errResp.ErrCodeInternalServerError, errResp.ErrInternalServerError.Error())
		return
	}

	response.SendSuccess(c, http.StatusOK, nil)
}

func (h *Handler) GetUserCourses(c *gin.Context) {
	roleId := c.GetInt("roleId")
	categoryId, _ := strconv.Atoi(c.Query("categoryId"))
	splitTagIds := strings.Split(c.Query("tagIds"), ",")
	var tagIds []int
	for _, tagId := range splitTagIds {
		if tagId != "" {
			id, _ := strconv.Atoi(tagId)
			tagIds = append(tagIds, id)
		}
	}

	limit, _ := strconv.Atoi(c.Query("limit"))
	page, _ := strconv.Atoi(c.Query("page"))
	params := entity.NewCourseParams(c.Query("title"), categoryId, tagIds, c.Query("sort"), limit, page, roleId, c.Query("status"))

	userId := c.GetInt("userId")

	courses, totalRows, totalPages, err := h.courseUsecase.GetUserCourses(userId, params)
	if err != nil {
		response.SendError(c, http.StatusInternalServerError, errResp.ErrCodeInternalServerError, errResp.ErrInternalServerError.Error())
		return
	}

	response.SendSuccessWithPagination(c, http.StatusOK, courses, totalRows, totalPages)
}

func (h *Handler) GetUserCourse(c *gin.Context) {
	slug := c.Param("slug")
	userId := c.GetInt("userId")

	course, err := h.courseUsecase.GetUserCourse(userId, slug)
	if err != nil {
		if errors.Is(err, errResp.ErrCourseNotFound) {
			response.SendError(c, http.StatusNotFound, errResp.ErrCodeNotFound, err.Error())
			return
		}

		response.SendError(c, http.StatusInternalServerError, errResp.ErrCodeInternalServerError, errResp.ErrInternalServerError.Error())
		return
	}

	response.SendSuccess(c, http.StatusOK, course)
}

func (h *Handler) CompleteCourse(c *gin.Context) {
	userId := c.GetInt("userId")
	slug := c.Param("slug")
	err := h.courseUsecase.CompleteCourse(userId, slug)
	if err != nil {
		if errors.Is(err, errResp.ErrCourseNotFound) {
			response.SendError(c, http.StatusNotFound, errResp.ErrCodeNotFound, err.Error())
			return
		}

		response.SendError(c, http.StatusInternalServerError, errResp.ErrCodeInternalServerError, errResp.ErrInternalServerError.Error())
		return
	}

	response.SendSuccess(c, http.StatusOK, nil)
}
