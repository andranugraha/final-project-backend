package handler

import (
	"errors"
	"final-project-backend/entity"
	errResp "final-project-backend/utils/errors"
	"final-project-backend/utils/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetFavoriteCourses(c *gin.Context) {
	userId := c.GetInt("userId")
	limit, _ := strconv.Atoi(c.Query("limit"))
	page, _ := strconv.Atoi(c.Query("page"))
	params := entity.NewFavoritesParams(userId, limit, page)
	courses, totalRows, totalPages, err := h.favoriteUsecase.GetFavoriteCourses(params)
	if err != nil {
		response.SendError(c, http.StatusInternalServerError, errResp.ErrCodeInternalServerError, errResp.ErrInternalServerError.Error())
		return
	}

	response.SendSuccessWithPagination(c, http.StatusOK, courses, totalRows, totalPages)
}

func (h *Handler) SaveUnsaveFavoriteCourse(c *gin.Context) {
	userId := c.GetInt("userId")
	courseId, err := strconv.Atoi(c.Param("courseId"))
	action := c.Param("action")
	if err != nil {
		response.SendError(c, http.StatusBadRequest, errResp.ErrCodeBadRequest, errResp.ErrInvalidParamFormat.Error())
		return
	}

	err = h.favoriteUsecase.SaveUnsaveFavoriteCourse(userId, courseId, action)
	if err != nil {
		if errors.Is(err, errResp.ErrFavoriteNotFound) {
			response.SendError(c, http.StatusNotFound, errResp.ErrCodeNotFound, err.Error())
			return
		}

		if errors.Is(err, errResp.ErrDuplicateFavorite) || errors.Is(err, errResp.ErrUnknownAction) {
			response.SendError(c, http.StatusBadRequest, errResp.ErrCodeBadRequest, err.Error())
			return
		}

		response.SendError(c, http.StatusInternalServerError, errResp.ErrCodeInternalServerError, errResp.ErrInternalServerError.Error())
		return
	}

	response.SendSuccess(c, http.StatusOK, nil)
}
