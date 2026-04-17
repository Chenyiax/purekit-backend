package middleware

import (
	"net/http"

	"purekit-backend/config"
	"purekit-backend/errors"
	"purekit-backend/pkg/httputil"

	"github.com/gin-gonic/gin"
)

// SizeLimit 请求体大小限制中间件
func SizeLimit() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, config.AppConfig.MaxImageSize)
		c.Next()

		// 检查是否因为请求体过大而导致的错误
		if c.Errors.Last() != nil {
			if c.Errors.Last().Type == gin.ErrorTypePrivate {
				httputil.Error(c, errors.NewImageTooLargeError("Image size exceeds limit"))
				c.Abort()
			}
		}
	}
}
