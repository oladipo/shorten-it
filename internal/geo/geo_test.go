package geo

import (
	"net"
	"testing"
)

func TestLookupIP(t *testing.T) {
	// Use a well-known public IP (Google DNS)
	ip := "8.8.8.8"
	info, err := LookupIP(ip)
	if err != nil {
		t.Fatalf("LookupIP failed: %v", err)
	}
	if info.Country == "" {
		t.Error("Expected country in geo info, got empty string")
	}
	parsed := net.ParseIP(info.Query)
	if parsed == nil {
		t.Errorf("Expected valid IP in response, got %q", info.Query)
	}
}
