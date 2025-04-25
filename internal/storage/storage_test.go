package storage

import (
	"context"
	"os"
	"testing"
)

func TestRedisStorage_SetAndGetURL(t *testing.T) {
	addr := os.Getenv("REDIS_ADDR")
	if addr == "" {
		addr = "localhost:6379"
	}
	password := os.Getenv("REDIS_PASSWORD")
	store := NewRedis(addr, password)

	ctx := context.Background()
	key := "test_key"
	url := "https://example.com"

	err := store.SetURL(ctx, key, url)
	if err != nil {
		t.Fatalf("SetURL failed: %v", err)
	}

	got, err := store.GetURL(ctx, key)
	if err != nil {
		t.Fatalf("GetURL failed: %v", err)
	}
	if got != url {
		t.Errorf("GetURL = %q, want %q", got, url)
	}
}

func TestBoltDBPlaceholder(t *testing.T) {
	store := NewBoltDB("test.db")
	if store != nil {
		t.Errorf("NewBoltDB should return nil placeholder, got: %#v", store)
	}
}

func TestPostgresPlaceholder(t *testing.T) {
	store := NewPostgres("postgres://user:pass@localhost:5432/db")
	if store != nil {
		t.Errorf("NewPostgres should return nil placeholder, got: %#v", store)
	}
}
