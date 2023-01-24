package errors

import "errors"

const (
	ErrSqlUniqueViolation      = "23505"
	ErrCodeDuplicate           = "DUPLICATE_RECORD"
	ErrCodeInternalServerError = "INTERNAL_SERVER_ERROR"
	ErrCodeBadRequest          = "BAD_REQUEST"
	ErrCodeUnauthorized        = "UNAUTHORIZED"
	ErrCodeForbidden           = "FORBIDDEN_ACCESS"
	ErrCodeNotFound            = "NOT_FOUND"
	ErrCodeRouteNotFound       = "ROUTE_NOT_FOUND"
)

var (
	ErrDuplicateRecord   = errors.New("duplicate record")
	ErrDuplicateTitle    = errors.New("duplicate title")
	ErrDuplicatePhoneNo  = errors.New("duplicate phone number")
	ErrDuplicateFavorite = errors.New("duplicate favorite")
	ErrDuplicateCart     = errors.New("duplicate cart")

	ErrInvalidBody  = errors.New("invalid body request")
	ErrInvalidParam = errors.New("invalid params")

	ErrUserNotFound     = errors.New("user not found")
	ErrRecordNotFound   = errors.New("record not found")
	ErrCourseNotFound   = errors.New("course not found")
	ErrCartNotFound     = errors.New("cart item not found")
	ErrFavoriteNotFound = errors.New("favorite item not found")
	ErrVoucherNotFound  = errors.New("voucher not found")
	ErrInvoiceNotFound  = errors.New("invoice not found")

	ErrCartEmpty                              = errors.New("cart is empty")
	ErrInvoiceAlreadyPaid                     = errors.New("invoice already paid")
	ErrInvoiceStatusNotWaitingForConfirmation = errors.New("invoice status is not waiting for confirmation")
	ErrInvalidInvoiceAction                   = errors.New("invalid invoice action")

	ErrWrongPassword         = errors.New("password mismatch")
	ErrForbidden             = errors.New("forbidden access to resources")
	ErrRouteNotFound         = errors.New("the requested route is not exist")
	ErrFailedToHash          = errors.New("failed to hash")
	ErrFailedToGenerateToken = errors.New("failed to generate token")
	ErrInternalServerError   = errors.New("internal server error")
	ErrUnknownAction         = errors.New("unknown action")
)
