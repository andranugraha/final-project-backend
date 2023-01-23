package handler

import (
	"errors"
	errResp "final-project-backend/utils/errors"
	"final-project-backend/utils/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetFavoriteCourses(c *gin.Context) {
	userId := c.GetInt("userId")
	courses, err := h.favoriteUsecase.GetFavoriteCourse(userId)
	if err != nil {
		response.SendError(c, http.StatusInternalServerError, errResp.ErrCodeInternalServerError, errResp.ErrInternalServerError.Error())
		return
	}

	response.SendSuccess(c, http.StatusOK, courses)
}

func (h *Handler) SaveUnsaveFavoriteCourse(c *gin.Context) {
	userId := c.GetInt("userId")
	courseId, err := strconv.Atoi(c.Param("courseId"))
	action := c.Param("action")
	if err != nil {
		response.SendError(c, http.StatusBadRequest, errResp.ErrCodeBadRequest, err.Error())
		return
	}

	err = h.favoriteUsecase.SaveUnsaveFavoriteCourse(userId, courseId, action)
	if err != nil {
		if errors.Is(err, errResp.ErrDuplicateFavorite) {
			response.SendError(c, http.StatusBadRequest, errResp.ErrCodeBadRequest, err.Error())
			return
		}

		response.SendError(c, http.StatusInternalServerError, errResp.ErrCodeInternalServerError, errResp.ErrInternalServerError.Error())
		return
	}

	response.SendSuccess(c, http.StatusOK, nil)
}
