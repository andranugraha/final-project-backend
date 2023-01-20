package handler

import (
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
