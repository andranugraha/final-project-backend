package dto

type CheckoutRequest struct {
	VoucherCode string `json:"voucherCode"`
}

type ConfirmInvoiceRequest struct {
	ApprovalType string `json:"approvalType" binding:"required"`
}
