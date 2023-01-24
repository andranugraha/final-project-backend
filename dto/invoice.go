package dto

type CheckoutRequest struct {
	VoucherCode string `json:"voucherCode" binding:"required"`
}

type ConfirmInvoiceRequest struct {
	ApprovalType string `json:"approvalType" binding:"required"`
}
