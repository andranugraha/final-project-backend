package handler

import (
	"errors"
	"net/http"

	"final-project-backend/dto"
	"final-project-backend/utils"

	"github.com/gin-gonic/gin"
)

func (h *Handler) SignIn(c *gin.Context) {
	var req dto.SignInRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    utils.ErrCodeBadRequest,
			"message": utils.ErrInvalidBody.Error(),
		})
		return
	}

	res, err := h.userUsecase.SignIn(req)
	if err != nil {
		if errors.Is(err, utils.ErrUserNotFound) || errors.Is(err, utils.ErrWrongPassword) {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    utils.ErrCodeBadRequest,
				"message": err.Error(),
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    utils.ErrCodeInternalServerError,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, res)
}
