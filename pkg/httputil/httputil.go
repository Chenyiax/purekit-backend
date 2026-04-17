package httputil

import (
	"net/http"

	"purekit-backend/errors"

	"github.com/gin-gonic/gin"
)

// Response 统一响应结构
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// Success 成功响应
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    http.StatusOK,
		Message: "success",
		Data:    data,
	})
}

// Error 错误响应
func Error(c *gin.Context, err error) {
	if appErr := errors.GetAppError(err); appErr != nil {
		var statusCode int
		switch appErr.Code {
		case errors.ErrBadRequest:
			statusCode = http.StatusBadRequest
		case errors.ErrUnsupportedFormat:
			statusCode = http.StatusBadRequest
		case errors.ErrImageTooLarge:
			statusCode = http.StatusBadRequest
		case errors.ErrImageProcessing:
			statusCode = http.StatusInternalServerError
		default:
			statusCode = http.StatusInternalServerError
		}

		c.JSON(statusCode, Response{
			Code:    statusCode,
			Message: appErr.Message,
		})
		return
	}

	// 未知错误
	c.JSON(http.StatusInternalServerError, Response{
		Code:    http.StatusInternalServerError,
		Message: "Internal server error",
	})
}
