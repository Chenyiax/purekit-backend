package middleware

import (
	"context"
	"time"

	"purekit-backend/config"

	"github.com/gin-gonic/gin"
)

// Timeout 请求超时中间件
func Timeout() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 创建带超时的上下文
		timeout := time.Duration(config.AppConfig.RequestTimeout) * time.Second
		ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)
		defer cancel()

		// 替换请求上下文
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}
