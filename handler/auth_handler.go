package handler

import (
	"errors"
	"net/http"

	"final-project-backend/dto"
	"final-project-backend/utils/constant"
	errResp "final-project-backend/utils/errors"
	"final-project-backend/utils/response"

	"github.com/gin-gonic/gin"
)

func (h *Handler) SignIn(c *gin.Context) {
	var req dto.SignInRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.SendError(c, http.StatusBadRequest, errResp.ErrCodeBadRequest, errResp.ErrInvalidBody.Error())
		return
	}

	res, err := h.userUsecase.SignIn(req, constant.UserRoleId)
	if err != nil {
		if errors.Is(err, errResp.ErrUserNotFound) || errors.Is(err, errResp.ErrWrongPassword) {
			response.SendError(c, http.StatusBadRequest, errResp.ErrCodeBadRequest, err.Error())
			return
		}

		response.SendError(c, http.StatusInternalServerError, errResp.ErrCodeInternalServerError, err.Error())
		return
	}

	response.SendSuccess(c, http.StatusOK, res)
}

func (h *Handler) AdminSignIn(c *gin.Context) {
	var req dto.SignInRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.SendError(c, http.StatusBadRequest, errResp.ErrCodeBadRequest, errResp.ErrInvalidBody.Error())
		return
	}

	res, err := h.userUsecase.SignIn(req, constant.AdminRoleId)
	if err != nil {
		if errors.Is(err, errResp.ErrUserNotFound) || errors.Is(err, errResp.ErrWrongPassword) {
			response.SendError(c, http.StatusBadRequest, errResp.ErrCodeBadRequest, err.Error())
			return
		}

		response.SendError(c, http.StatusInternalServerError, errResp.ErrCodeInternalServerError, err.Error())
		return
	}

	response.SendSuccess(c, http.StatusOK, res)
}

func (h *Handler) SignUp(c *gin.Context) {
	var req dto.SignUpRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.SendError(c, http.StatusBadRequest, errResp.ErrCodeBadRequest, errResp.ErrInvalidBody.Error())
		return
	}

	res, err := h.userUsecase.SignUp(req)
	if err != nil {
		if errors.Is(err, errResp.ErrDuplicateRecord) {
			response.SendError(c, http.StatusBadRequest, errResp.ErrCodeBadRequest, err.Error())
			return
		}

		response.SendError(c, http.StatusInternalServerError, errResp.ErrCodeInternalServerError, err.Error())
		return
	}

	response.SendSuccess(c, http.StatusOK, res)
}
