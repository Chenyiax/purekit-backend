package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// Logger 访问日志中间件
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 开始时间
		start := time.Now()
		// 请求路径
		path := c.Request.URL.Path
		// 请求方法
		method := c.Request.Method
		// 客户端IP
		clientIP := c.ClientIP()

		// 处理请求
		c.Next()

		// 结束时间
		end := time.Now()
		// 执行时间
		latency := end.Sub(start)
		// 状态码
		status := c.Writer.Status()

		// 打印格式化访问日志，用于后续观察访问量
		log.Printf("[ACCESS] %s | %3d | %13v | %15s | %-7s %s",
			end.Format("2006/01/02 15:04:05"),
			status,
			latency,
			clientIP,
			method,
			path,
		)
	}
}
