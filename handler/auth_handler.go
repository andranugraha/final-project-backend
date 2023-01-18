package handler

import (
	"errors"
	"net/http"

	"final-project-backend/dto"
	"final-project-backend/utils/constant"
	errResp "final-project-backend/utils/errors"

	"github.com/gin-gonic/gin"
)

func (h *Handler) SignIn(c *gin.Context) {
	var req dto.SignInRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    errResp.ErrCodeBadRequest,
			"message": errResp.ErrInvalidBody.Error(),
		})
		return
	}

	res, err := h.userUsecase.SignIn(req, constant.UserRoleId)
	if err != nil {
		if errors.Is(err, errResp.ErrUserNotFound) || errors.Is(err, errResp.ErrWrongPassword) {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    errResp.ErrCodeBadRequest,
				"message": err.Error(),
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    errResp.ErrCodeInternalServerError,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *Handler) AdminSignIn(c *gin.Context) {
	var req dto.SignInRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    errResp.ErrCodeBadRequest,
			"message": errResp.ErrInvalidBody.Error(),
		})
		return
	}

	res, err := h.userUsecase.SignIn(req, constant.AdminRoleId)
	if err != nil {
		if errors.Is(err, errResp.ErrUserNotFound) || errors.Is(err, errResp.ErrWrongPassword) {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    errResp.ErrCodeBadRequest,
				"message": err.Error(),
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    errResp.ErrCodeInternalServerError,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, res)
}
