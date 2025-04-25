package api

import (
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

// AuthRequired is a middleware for API key authentication
func AuthRequired() gin.HandlerFunc {
	apiKey := os.Getenv("API_KEY")
	return func(c *gin.Context) {
		if apiKey == "" {
			c.Next()
			return // No API key set, allow all (dev mode)
		}
		key := c.GetHeader("X-API-Key")
		if key != apiKey {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid API key"})
			return
		}
		c.Next()
	}
}

// RateLimit is a simple in-memory rate limiter per IP
var rateLimiters = make(map[string]*rateLimiter)
var mu sync.Mutex

type rateLimiter struct {
	last time.Time
	count int
}

func RateLimit(maxPerMinute int) gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		mu.Lock()
		rl, ok := rateLimiters[ip]
		if !ok || time.Since(rl.last) > time.Minute {
			rl = &rateLimiter{last: time.Now(), count: 1}
			rateLimiters[ip] = rl
			mu.Unlock()
			c.Next()
			return
		}
		rl.count++
		if rl.count > maxPerMinute {
			mu.Unlock()
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "rate limit exceeded"})
			return
		}
		mu.Unlock()
		c.Next()
	}
}
