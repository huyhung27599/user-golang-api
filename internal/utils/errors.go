package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrorCode string

const (
	ErrCodeBadRequest         ErrorCode = "BAD_REQUEST"
	ErrCodeInternal           ErrorCode = "INTERNAL_SERVER_ERROR"
	ErrCodeNotFound           ErrorCode = "NOT_FOUND"
	ErrCodeUnauthorized       ErrorCode = "UNAUTHORIZED"
	ErrCodeForbidden          ErrorCode = "FORBIDDEN"
	ErrCodeConflict           ErrorCode = "CONFLICT"
	ErrCodeBadGateway         ErrorCode = "BAD_GATEWAY"
	ErrCodeServiceUnavailable ErrorCode = "SERVICE_UNAVAILABLE"
	ErrCodeGatewayTimeout     ErrorCode = "GATEWAY_TIMEOUT"
)

type AppError struct {
	Message string
	Code    ErrorCode
	Err     error
}

func (ae *AppError) Error() string {
	return ""
}

func WrapError(message string, code ErrorCode, err error) error {
	return &AppError{
		Message: message,
		Code:    code,
		Err:     err,
	}
}

func ResponseError(c *gin.Context, err error) {
	if appErr, ok := err.(*AppError); ok {
		httpStatus := httpStatusFromCode(appErr.Code)
	 response := gin.H{
		"error": appErr.Message,
		"code": appErr.Code,
		"message": appErr.Message,
	 }
	  if appErr.Err != nil {
		response["error_detail"] = appErr.Err.Error()
	  }
	  c.JSON(httpStatus	, response)
	} else {
		
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}


func ResponseSuccess(c *gin.Context,status int, data any) {
	response := gin.H{
		"status": status,
		"data": data,
	}
	c.JSON(status, response)
}

func httpStatusFromCode(code ErrorCode) int {
	switch code {
	case ErrCodeBadRequest:
		return http.StatusBadRequest
	case ErrCodeInternal:
		return http.StatusInternalServerError
	case ErrCodeNotFound:
		return http.StatusNotFound
	case ErrCodeUnauthorized:
		return http.StatusUnauthorized
	case ErrCodeForbidden:
		return http.StatusForbidden
	case ErrCodeConflict:
		return http.StatusConflict
	case ErrCodeBadGateway:
		return http.StatusBadGateway
	case ErrCodeServiceUnavailable:
		return http.StatusServiceUnavailable
	case ErrCodeGatewayTimeout:
		return http.StatusGatewayTimeout
	}
	return http.StatusInternalServerError
}
