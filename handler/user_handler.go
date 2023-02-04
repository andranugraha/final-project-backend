package handler

import (
	"errors"
	"final-project-backend/dto"
	errResp "final-project-backend/utils/errors"
	"final-project-backend/utils/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetUserDetail(c *gin.Context) {
	userId := c.GetInt("userId")
	res, err := h.userUsecase.GetUserDetail(userId)
	if err != nil {
		response.SendError(c, http.StatusInternalServerError, errResp.ErrCodeInternalServerError, err.Error())
		return
	}

	response.SendSuccess(c, http.StatusOK, res)
}

func (h *Handler) UpdateUserDetail(c *gin.Context) {
	userId := c.GetInt("userId")
	var req dto.UpdateUserDetailRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.SendError(c, http.StatusBadRequest, errResp.ErrCodeBadRequest, errResp.ErrInvalidBody.Error())
		return
	}

	res, err := h.userUsecase.UpdateUserDetail(userId, req)
	if err != nil {
		if errors.Is(err, errResp.ErrUserNotFound) {
			response.SendError(c, http.StatusNotFound, errResp.ErrCodeNotFound, err.Error())
			return
		}

		if errors.Is(err, errResp.ErrDuplicatePhoneNo) {
			response.SendError(c, http.StatusBadRequest, errResp.ErrCodeBadRequest, err.Error())
			return
		}

		response.SendError(c, http.StatusInternalServerError, errResp.ErrCodeInternalServerError, err.Error())
		return
	}

	response.SendSuccess(c, http.StatusOK, res)
}
