package errors

import "errors"

const (
	ErrSqlUniqueViolation      = "23505"
	ErrCodeDuplicate           = "DUPLICATE_RECORD"
	ErrCodeInternalServerError = "INTERNAL_SERVER_ERROR"
	ErrCodeBadRequest          = "BAD_REQUEST"
	ErrCodeUnauthorized        = "UNAUTHORIZED"
	ErrCodeForbidden           = "FORBIDDEN_ACCESS"
	ErrCodeRouteNotFound       = "ROUTE_NOT_FOUND"
)

var (
	ErrInvalidBody           = errors.New("invalid body request")
	ErrInvalidParam          = errors.New("invalid params")
	ErrUserNotFound          = errors.New("user not found")
	ErrRecordNotFound        = errors.New("record not found")
	ErrWrongPassword         = errors.New("password mismatch")
	ErrForbidden             = errors.New("forbidden access to resources")
	ErrRouteNotFound         = errors.New("the requested route is not exist")
	ErrFailedToHash          = errors.New("failed to hash")
	ErrFailedToGenerateToken = errors.New("failed to generate token")
	ErrInternalServerError   = errors.New("internal server error")
)
