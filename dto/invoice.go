package dto

type CheckoutRequest struct {
	VoucherCode string `json:"voucherCode" binding:"required"`
}
