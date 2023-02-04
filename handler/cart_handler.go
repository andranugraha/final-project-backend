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
		response.SendError(c, http.StatusBadRequest, errResp.ErrCodeBadRequest, errResp.ErrInvalidParamFormat.Error())
		return
	}
	err = h.cartUsecase.AddToCart(userId, courseId)
	if err != nil {
		if errors.Is(err, errResp.ErrDuplicateCart) || errors.Is(err, errResp.ErrCourseNotFound) || errors.Is(err, errResp.ErrCourseAlreadyBought) {
			response.SendError(c, http.StatusBadRequest, errResp.ErrCodeBadRequest, err.Error())
			return
		}

		response.SendError(c, http.StatusInternalServerError, errResp.ErrCodeInternalServerError, err.Error())
		return
	}

	response.SendSuccessWithMessage(c, http.StatusOK, nil, "Course added to cart")
}

func (h *Handler) RemoveFromCart(c *gin.Context) {
	userId := c.GetInt("userId")
	courseId, err := strconv.Atoi(c.Param("courseId"))
	if err != nil {
		response.SendError(c, http.StatusBadRequest, errResp.ErrCodeBadRequest, errResp.ErrInvalidParamFormat.Error())
		return
	}
	err = h.cartUsecase.RemoveFromCart(userId, courseId)
	if err != nil {
		if errors.Is(err, errResp.ErrCartNotFound) || errors.Is(err, errResp.ErrCourseNotFound) {
			response.SendError(c, http.StatusBadRequest, errResp.ErrCodeBadRequest, err.Error())
			return
		}

		response.SendError(c, http.StatusInternalServerError, errResp.ErrCodeInternalServerError, err.Error())
		return
	}

	response.SendSuccessWithMessage(c, http.StatusOK, nil, "Course removed from cart")
}
