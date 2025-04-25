package metrics

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestMetricsCounters(t *testing.T) {
	Init()
	RedirectsTotal.WithLabelValues("abc123").Inc()
	ShortenRequestsTotal.Inc()

	req := httptest.NewRequest("GET", "/metrics", nil)
	w := httptest.NewRecorder()
	Handler().ServeHTTP(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status 200, got %d", resp.StatusCode)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	metricsOutput := string(body)

	if !strings.Contains(metricsOutput, "url_shortener_redirects_total{shortcode=\"abc123\"} 1") {
		t.Error("RedirectsTotal metric not found or incorrect in output")
	}
	if !strings.Contains(metricsOutput, "url_shortener_shorten_requests_total 1") {
		t.Error("ShortenRequestsTotal metric not found or incorrect in output")
	}
}
