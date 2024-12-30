package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// RateLimiter limits the number of requests per minute
func RateLimiter() gin.HandlerFunc {
	limiter := make(map[string]time.Time)

	return func(c *gin.Context) {
		clientIP := c.ClientIP()
		lastRequestTime, exists := limiter[clientIP]

		if exists && time.Since(lastRequestTime) < time.Minute {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "Rate limit exceeded"})
			c.Abort()
			return
		}

		limiter[clientIP] = time.Now()
		c.Next()
	}
}
