package api

import (
	"net/http"
	"net/http/httptest"
	"os"
	"sync"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestAuthRequired(t *testing.T) {
	gin.SetMode(gin.TestMode)
	os.Setenv("API_KEY", "testkey")
	r := gin.New()
	r.GET("/private", AuthRequired(), func(c *gin.Context) {
		c.String(200, "ok")
	})

	req := httptest.NewRequest("GET", "/private", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusUnauthorized {
		t.Errorf("Expected 401 for missing API key, got %d", w.Code)
	}

	req2 := httptest.NewRequest("GET", "/private", nil)
	req2.Header.Set("X-API-Key", "testkey")
	w2 := httptest.NewRecorder()
	r.ServeHTTP(w2, req2)
	if w2.Code != http.StatusOK {
		t.Errorf("Expected 200 for valid API key, got %d", w2.Code)
	}
}

func TestRateLimit(t *testing.T) {
	gin.SetMode(gin.TestMode)
	rateLimiters = make(map[string]*rateLimiter) // Reset global rate limiter state

	r := gin.New()
	r.GET("/limited", RateLimit(2), func(c *gin.Context) {
		c.String(200, "ok")
	})

	for i := 0; i < 2; i++ {
		req := httptest.NewRequest("GET", "/limited", nil)
		req.RemoteAddr = "1.2.3.4:12345"
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		if w.Code != http.StatusOK {
			t.Errorf("Expected 200 for request %d, got %d", i+1, w.Code)
		}
	}
	// Third request should be rate limited
	req := httptest.NewRequest("GET", "/limited", nil)
	req.RemoteAddr = "1.2.3.4:12345"
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusTooManyRequests {
		t.Errorf("Expected 429 for rate limited request, got %d", w.Code)
	}
}

func TestRateLimitMultipleIPs(t *testing.T) {
	gin.SetMode(gin.TestMode)
	rateLimiters = make(map[string]*rateLimiter)

	r := gin.New()
	r.GET("/limited", RateLimit(2), func(c *gin.Context) {
		c.String(200, "ok")
	})

	ips := []string{"1.2.3.4:12345", "5.6.7.8:12345"}
	for _, ip := range ips {
		for i := 0; i < 2; i++ {
			req := httptest.NewRequest("GET", "/limited", nil)
			req.RemoteAddr = ip
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)
			if w.Code != http.StatusOK {
				t.Errorf("Expected 200 for IP %s request %d, got %d", ip, i+1, w.Code)
			}
		}
		// Third request for this IP should be limited
		req := httptest.NewRequest("GET", "/limited", nil)
		req.RemoteAddr = ip
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		if w.Code != http.StatusTooManyRequests {
			t.Errorf("Expected 429 for IP %s rate limited request, got %d", ip, w.Code)
		}
	}
}

func TestRateLimitConcurrent(t *testing.T) {
	gin.SetMode(gin.TestMode)
	rateLimiters = make(map[string]*rateLimiter)

	r := gin.New()
	r.GET("/limited", RateLimit(10), func(c *gin.Context) {
		c.String(200, "ok")
	})

	wg := sync.WaitGroup{}
	wg.Add(20)
	results := make([]int, 20)
	for i := 0; i < 20; i++ {
		go func(idx int) {
			defer wg.Done()
			req := httptest.NewRequest("GET", "/limited", nil)
			req.RemoteAddr = "9.9.9.9:12345"
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)
			results[idx] = w.Code
		}(i)
	}
	wg.Wait()
	var okCount, limitedCount int
	for _, code := range results {
		if code == http.StatusOK {
			okCount++
		} else if code == http.StatusTooManyRequests {
			limitedCount++
		}
	}
	if okCount != 10 {
		t.Errorf("Expected 10 OK responses, got %d", okCount)
	}
	if limitedCount != 10 {
		t.Errorf("Expected 10 rate limited responses, got %d", limitedCount)
	}
}
