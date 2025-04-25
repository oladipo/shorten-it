package api

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/oladipo/shorten-it/internal/storage"
)

type mockStore struct{
	storage.Storage
	urls map[string]string
}

func (m *mockStore) GetURL(_ context.Context, key string) (string, error) {
	return m.urls[key], nil
}

func TestRedirectAndStats(t *testing.T) {
	gin.SetMode(gin.TestMode)
	store := &mockStore{urls: map[string]string{"abc": "https://example.com"}}
	r := gin.New()
	RegisterRoutes(r, store)

	// Simulate a redirect
	req := httptest.NewRequest("GET", "/abc", nil)
	req.Header.Set("User-Agent", "TestAgent/1.0")
	req.RemoteAddr = "1.2.3.4:12345"
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusFound {
		t.Fatalf("Expected 302 redirect, got %d", resp.StatusCode)
	}
	loc := resp.Header.Get("Location")
	if loc != "https://example.com" {
		t.Errorf("Expected redirect to https://example.com, got %s", loc)
	}

	// Simulate stats request
	req2 := httptest.NewRequest("GET", "/abc/stats", nil)
	w2 := httptest.NewRecorder()
	r.ServeHTTP(w2, req2)

	if w2.Code != http.StatusOK {
		t.Fatalf("Expected 200 for stats, got %d", w2.Code)
	}
	var stats []map[string]interface{}
	if err := json.Unmarshal(w2.Body.Bytes(), &stats); err != nil {
		t.Fatalf("Failed to parse stats JSON: %v", err)
	}
	if len(stats) == 0 {
		t.Error("Expected at least one analytics event in stats")
	}
}

func TestShortenAuth(t *testing.T) {
	gin.SetMode(gin.TestMode)
	os.Setenv("API_KEY", "testkey")
	store := &mockStore{urls: map[string]string{}}
	r := gin.New()
	RegisterRoutes(r, store)

	req := httptest.NewRequest("POST", "/shorten", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusUnauthorized {
		t.Errorf("Expected 401 for missing API key, got %d", w.Code)
	}

	req2 := httptest.NewRequest("POST", "/shorten", nil)
	req2.Header.Set("X-API-Key", "testkey")
	w2 := httptest.NewRecorder()
	r.ServeHTTP(w2, req2)
	if w2.Code != http.StatusNotImplemented {
		t.Errorf("Expected 501 for stubbed shorten, got %d", w2.Code)
	}
}
