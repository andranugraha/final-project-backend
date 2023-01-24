package handler

import (
	"errors"
	"final-project-backend/dto"
	"final-project-backend/entity"
	"final-project-backend/utils/constant"
	errResp "final-project-backend/utils/errors"
	"final-project-backend/utils/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetInvoices(c *gin.Context) {
	userId := c.GetInt("userId")
	roleId := c.GetInt("roleId")
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))

	params := entity.NewInvoiceParams(c.Query("status"), c.Query("sort"), page, limit, userId, roleId)

	invoices, totalRows, totalPages, err := h.invoiceUsecase.GetInvoices(params)
	if err != nil {
		response.SendError(c, http.StatusInternalServerError, errResp.ErrCodeInternalServerError, err.Error())
		return
	}

	response.SendSuccessWithPagination(c, http.StatusOK, invoices, totalRows, totalPages)
}

func (h *Handler) GetInvoice(c *gin.Context) {
	userId := c.GetInt("userId")
	roleId := c.GetInt("roleId")
	invoiceId := c.Param("invoiceId")

	if roleId == constant.AdminRoleId {
		userId = 0
	}

	invoice, err := h.invoiceUsecase.GetInvoiceDetail(userId, invoiceId)
	if err != nil {
		if errors.Is(err, errResp.ErrInvoiceNotFound) {
			response.SendError(c, http.StatusNotFound, errResp.ErrCodeNotFound, err.Error())
			return
		}

		response.SendError(c, http.StatusInternalServerError, errResp.ErrCodeInternalServerError, err.Error())
		return
	}

	response.SendSuccess(c, http.StatusOK, invoice)
}

func (h *Handler) Checkout(c *gin.Context) {
	userId := c.GetInt("userId")
	var req dto.CheckoutRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.SendError(c, http.StatusBadRequest, errResp.ErrCodeBadRequest, err.Error())
		return
	}

	invoice, err := h.invoiceUsecase.Checkout(userId, req)
	if err != nil {
		if errors.Is(err, errResp.ErrVoucherNotFound) || errors.Is(err, errResp.ErrCartEmpty) {
			response.SendError(c, http.StatusBadRequest, errResp.ErrCodeBadRequest, err.Error())
			return
		}

		response.SendError(c, http.StatusInternalServerError, errResp.ErrCodeInternalServerError, err.Error())
		return
	}

	response.SendSuccess(c, http.StatusOK, invoice)
}

func (h *Handler) PayInvoice(c *gin.Context) {
	userId := c.GetInt("userId")
	invoiceId := c.Param("invoiceId")

	invoice, err := h.invoiceUsecase.PayInvoice(userId, invoiceId)
	if err != nil {
		if errors.Is(err, errResp.ErrInvoiceNotFound) {
			response.SendError(c, http.StatusNotFound, errResp.ErrCodeNotFound, err.Error())
			return
		}

		if errors.Is(err, errResp.ErrInvoiceAlreadyPaid) {
			response.SendError(c, http.StatusBadRequest, errResp.ErrCodeBadRequest, err.Error())
			return
		}

		response.SendError(c, http.StatusInternalServerError, errResp.ErrCodeInternalServerError, err.Error())
		return
	}

	response.SendSuccess(c, http.StatusOK, invoice)
}

func (h *Handler) ConfirmInvoice(c *gin.Context) {
	invoiceId := c.Param("invoiceId")

	var req dto.ConfirmInvoiceRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.SendError(c, http.StatusBadRequest, errResp.ErrCodeBadRequest, err.Error())
		return
	}

	if req.ApprovalType != "confirm" && req.ApprovalType != "reject" {
		response.SendError(c, http.StatusBadRequest, errResp.ErrCodeBadRequest, errResp.ErrInvalidInvoiceAction.Error())
		return
	}

	var status string
	switch req.ApprovalType {
	case "confirm":
		status = constant.InvoiceStatusCompleted
	case "reject":
		status = constant.InvoiceStatusCancelled
	}

	invoice, err := h.invoiceUsecase.ConfirmInvoice(invoiceId, status)
	if err != nil {
		if errors.Is(err, errResp.ErrInvoiceNotFound) {
			response.SendError(c, http.StatusNotFound, errResp.ErrCodeNotFound, err.Error())
			return
		}

		if errors.Is(err, errResp.ErrInvoiceStatusNotWaitingForConfirmation) {
			response.SendError(c, http.StatusBadRequest, errResp.ErrCodeBadRequest, err.Error())
			return
		}

		response.SendError(c, http.StatusInternalServerError, errResp.ErrCodeInternalServerError, err.Error())
		return
	}

	response.SendSuccess(c, http.StatusOK, invoice)
}
