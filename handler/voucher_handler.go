package handler

import (
	"errors"
	errResp "final-project-backend/utils/errors"
	"final-project-backend/utils/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetUserVouchers(c *gin.Context) {
	userId := c.GetInt("userId")

	res, err := h.userVoucherUsecase.GetUserVouchers(userId)
	if err != nil {
		if errors.Is(err, errResp.ErrUserNotFound) {
			response.SendError(c, http.StatusNotFound, errResp.ErrCodeNotFound, err.Error())
			return
		}

		response.SendError(c, http.StatusInternalServerError, errResp.ErrCodeInternalServerError, errResp.ErrInternalServerError.Error())
		return
	}

	response.SendSuccess(c, http.StatusOK, res)
}
