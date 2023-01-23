package handler

import (
	"errors"
	errResp "final-project-backend/utils/errors"
	"final-project-backend/utils/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetCart(c *gin.Context) {
	userId := c.GetInt("userId")
	courses, err := h.cartUsecase.GetCart(userId)
	if err != nil {
		response.SendError(c, http.StatusInternalServerError, errResp.ErrCodeInternalServerError, err.Error())
		return
	}

	response.SendSuccess(c, http.StatusOK, courses)
}

func (h *Handler) AddToCart(c *gin.Context) {
	userId := c.GetInt("userId")
	courseId, err := strconv.Atoi(c.Param("courseId"))
	if err != nil {
		response.SendError(c, http.StatusBadRequest, errResp.ErrCodeBadRequest, err.Error())
		return
	}
	err = h.cartUsecase.AddToCart(userId, courseId)
	if err != nil {
		if errors.Is(err, errResp.ErrDuplicateCart) {
			response.SendError(c, http.StatusBadRequest, errResp.ErrCodeBadRequest, err.Error())
			return
		}

		response.SendError(c, http.StatusInternalServerError, errResp.ErrCodeInternalServerError, err.Error())
		return
	}

	response.SendSuccess(c, http.StatusOK, nil)
}

func (h *Handler) RemoveFromCart(c *gin.Context) {
	userId := c.GetInt("userId")
	courseId, err := strconv.Atoi(c.Param("courseId"))
	if err != nil {
		response.SendError(c, http.StatusBadRequest, errResp.ErrCodeBadRequest, err.Error())
		return
	}
	err = h.cartUsecase.RemoveFromCart(userId, courseId)
	if err != nil {
		if errors.Is(err, errResp.ErrCartNotFound) {
			response.SendError(c, http.StatusNotFound, errResp.ErrCodeNotFound, err.Error())
			return
		}

		response.SendError(c, http.StatusInternalServerError, errResp.ErrCodeInternalServerError, err.Error())
		return
	}

	response.SendSuccess(c, http.StatusOK, nil)
}
