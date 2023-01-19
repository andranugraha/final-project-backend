package handler

import (
	"final-project-backend/dto"
	errResp "final-project-backend/utils/errors"
	"final-project-backend/utils/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateCourse(c *gin.Context) {
	var req dto.CreateCourseRequest
	err := c.ShouldBind(&req)
	if err != nil {
		response.SendError(c, http.StatusBadRequest, errResp.ErrCodeBadRequest, err.Error())
		return
	}

	res, err := h.courseUsecase.CreateCourse(req)
	if err != nil {
		if err == errResp.ErrDuplicateTitle {
			response.SendError(c, http.StatusBadRequest, errResp.ErrCodeDuplicate, err.Error())
			return
		}
		response.SendError(c, http.StatusInternalServerError, errResp.ErrCodeInternalServerError, err.Error())
		return
	}

	response.SendSuccess(c, http.StatusOK, res)
}

func (h *Handler) GetCourse(c *gin.Context) {
	slug := c.Param("slug")

	res, err := h.courseUsecase.GetCourse(slug)
	if err != nil {
		if err == errResp.ErrCourseNotFound {
			response.SendError(c, http.StatusNotFound, errResp.ErrCodeNotFound, err.Error())
			return
		}

		response.SendError(c, http.StatusInternalServerError, errResp.ErrCodeInternalServerError, errResp.ErrInternalServerError.Error())
		return
	}

	response.SendSuccess(c, http.StatusOK, res)
}

func (h *Handler) UpdateCourse(c *gin.Context) {
	slug := c.Param("slug")

	var req dto.UpdateCourseRequest
	err := c.ShouldBind(&req)
	if err != nil {
		response.SendError(c, http.StatusBadRequest, errResp.ErrCodeBadRequest, err.Error())
		return
	}

	res, err := h.courseUsecase.UpdateCourse(slug, req)
	if err != nil {
		if err == errResp.ErrCourseNotFound {
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
		if err == errResp.ErrCourseNotFound {
			response.SendError(c, http.StatusNotFound, errResp.ErrCodeNotFound, err.Error())
			return
		}

		response.SendError(c, http.StatusInternalServerError, errResp.ErrCodeInternalServerError, errResp.ErrInternalServerError.Error())
		return
	}

	response.SendSuccess(c, http.StatusOK, nil)
}
