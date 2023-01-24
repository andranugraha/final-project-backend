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

	res, err := h.courseUsecase.CreateCourse(req)
	if err != nil {
		if errors.Is(err, errResp.ErrDuplicateTitle) {
			response.SendError(c, http.StatusBadRequest, errResp.ErrCodeDuplicate, err.Error())
			return
		}
		response.SendError(c, http.StatusInternalServerError, errResp.ErrCodeInternalServerError, err.Error())
		return
	}

	response.SendSuccess(c, http.StatusOK, res)
}

func (h *Handler) GetCourses(c *gin.Context) {
	roleId := c.GetInt("roleId")
	categoryId, _ := strconv.Atoi(c.Query("categoryId"))
	splitTagIds := strings.Split(c.Query("tagIds"), ",")
	var tagIds []int
	for _, tagId := range splitTagIds {
		id, _ := strconv.Atoi(tagId)
		tagIds = append(tagIds, id)
	}

	limit, _ := strconv.Atoi(c.Query("limit"))
	page, _ := strconv.Atoi(c.Query("page"))
	params := entity.NewCourseParams(c.Query("title"), categoryId, tagIds, c.Query("sort"), limit, page, roleId, c.Query("status"))

	res, totalRows, totalPages, err := h.courseUsecase.GetCourses(params)
	if err != nil {
		response.SendError(c, http.StatusInternalServerError, errResp.ErrCodeInternalServerError, errResp.ErrInternalServerError.Error())
		return
	}

	response.SendSuccessWithPagination(c, http.StatusOK, res, totalRows, totalPages)
}

func (h *Handler) GetCourse(c *gin.Context) {
	slug := c.Param("slug")

	res, err := h.courseUsecase.GetCourse(slug)
	if err != nil {
		if errors.Is(err, errResp.ErrCourseNotFound) {
			response.SendError(c, http.StatusNotFound, errResp.ErrCodeNotFound, err.Error())
			return
		}

		response.SendError(c, http.StatusInternalServerError, errResp.ErrCodeInternalServerError, errResp.ErrInternalServerError.Error())
		return
	}

	response.SendSuccess(c, http.StatusOK, res)
}

func (h *Handler) UpdateCourse(c *gin.Context) {
	var req dto.UpdateCourseRequest
	if err := c.ShouldBind(&req); err != nil {
		response.SendError(c, http.StatusBadRequest, errResp.ErrCodeBadRequest, err.Error())
		return
	}

	slug := c.Param("slug")
	res, err := h.courseUsecase.UpdateCourse(slug, req)
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

	response.SendSuccess(c, http.StatusOK, res)
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
