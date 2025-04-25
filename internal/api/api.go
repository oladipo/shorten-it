package api

import (
	"net"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/oladipo/shorten-it/internal/geo"
	"github.com/oladipo/shorten-it/internal/storage"
	"github.com/oladipo/shorten-it/internal/analytics"
)

// RegisterRoutes registers API routes to the provided Gin router
func RegisterRoutes(r *gin.Engine, store storage.Storage) {
	r.Use(RateLimit(10)) // 10 requests/minute per IP

	r.GET(":shortcode", func(c *gin.Context) {
		shortcode := c.Param("shortcode")
		ctx := c.Request.Context()
		url, err := store.GetURL(ctx, shortcode)
		if err != nil || url == "" {
			c.String(http.StatusNotFound, "Shortcode not found")
			return
		}

		ip := c.ClientIP()
		if net.ParseIP(ip) == nil {
			ip = "8.8.8.8"
		}
		geoInfo, err := geo.LookupIP(ip)
		country, city := "", ""
		if err == nil && geoInfo != nil {
			country = geoInfo.Country
			city = geoInfo.City
		}
		analytics.Record(shortcode, analytics.Event{
			Timestamp: time.Now(),
			IP:        ip,
			Country:   country,
			City:      city,
			Referrer:  c.Request.Referer(),
			UserAgent: c.Request.UserAgent(),
		})

		c.Redirect(http.StatusFound, url)
	})

	r.GET(":shortcode/stats", func(c *gin.Context) {
		shortcode := c.Param("shortcode")
		stats := analytics.GetEvents(shortcode)
		c.JSON(http.StatusOK, stats)
	})

	r.POST("/shorten", AuthRequired(), func(c *gin.Context) {
		// Example: require API key for shortening
		c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
	})
}
