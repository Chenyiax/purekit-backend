package middleware

import (
	"sync"
	"time"

	"purekit-backend/config"
	"purekit-backend/errors"
	"purekit-backend/pkg/httputil"

	"github.com/gin-gonic/gin"
)

// RateLimiter 速率限制器
type RateLimiter struct {
	mu            sync.Mutex
	ipCounts      map[string][]time.Time
	maxRequests   int
	timeWindow    time.Duration
	concurrentMu  sync.Mutex
	concurrent    int
	maxConcurrent int
}

// NewRateLimiter 创建速率限制器
func NewRateLimiter() *RateLimiter {
	return &RateLimiter{
		ipCounts:      make(map[string][]time.Time),
		maxRequests:   100, // 每个IP每分钟最多100个请求
		timeWindow:    time.Minute,
		maxConcurrent: config.AppConfig.MaxConcurrentRequests,
	}
}

// Limit 限流中间件
func (rl *RateLimiter) Limit() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 检查并发请求数
		rl.concurrentMu.Lock()
		if rl.concurrent >= rl.maxConcurrent {
			rl.concurrentMu.Unlock()
			httputil.Error(c, errors.NewBadRequestError("Too many concurrent requests"))
			c.Abort()
			return
		}
		rl.concurrent++
		rl.concurrentMu.Unlock()

		// 检查IP速率限制
		ip := c.ClientIP()
		rl.mu.Lock()
		now := time.Now()
		// 清理过期的请求记录
		var validRequests []time.Time
		for _, t := range rl.ipCounts[ip] {
			if now.Sub(t) < rl.timeWindow {
				validRequests = append(validRequests, t)
			}
		}
		rl.ipCounts[ip] = validRequests
		// 检查是否超过限制
		if len(rl.ipCounts[ip]) >= rl.maxRequests {
			rl.mu.Unlock()
			// 减少并发计数
			rl.concurrentMu.Lock()
			rl.concurrent--
			rl.concurrentMu.Unlock()
			httputil.Error(c, errors.NewBadRequestError("Rate limit exceeded"))
			c.Abort()
			return
		}
		// 记录新请求
		rl.ipCounts[ip] = append(rl.ipCounts[ip], now)
		rl.mu.Unlock()

		// 请求处理完成后减少并发计数
		c.Next()
		rl.concurrentMu.Lock()
		rl.concurrent--
		rl.concurrentMu.Unlock()
	}
}
