package api

import (
	"log"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oladipo/shorten-it/internal/geo"
	"github.com/oladipo/shorten-it/internal/storage"
)

// RegisterRoutes registers API routes to the provided Gin router
func RegisterRoutes(r *gin.Engine, store storage.Storage) {
	r.GET(":shortcode", func(c *gin.Context) {
		shortcode := c.Param("shortcode")
		ctx := c.Request.Context()
		url, err := store.GetURL(ctx, shortcode)
		if err != nil || url == "" {
			c.String(http.StatusNotFound, "Shortcode not found")
			return
		}

		// Get client IP
		ip := c.ClientIP()
		if net.ParseIP(ip) == nil {
			ip = "8.8.8.8" // fallback for local testing
		}
		geoInfo, err := geo.LookupIP(ip)
		if err != nil {
			log.Printf("Geo lookup failed for IP %s: %v", ip, err)
		} else {
			log.Printf("Redirect from %s: country=%s, city=%s", ip, geoInfo.Country, geoInfo.City)
		}

		c.Redirect(http.StatusFound, url)
	})
}
