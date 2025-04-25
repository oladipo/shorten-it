package storage

import (
	"context"
	"github.com/redis/go-redis/v9"
)

// Storage is an interface for URL storage backends
// Implementations: BoltDB, Redis, Postgres

type Storage interface {
	SetURL(ctx context.Context, key, url string) error
	GetURL(ctx context.Context, key string) (string, error)
}

type redisStorage struct {
	client *redis.Client
}

func NewRedis(addr, password string) Storage {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       0,
	})
	return &redisStorage{client: client}
}

func (r *redisStorage) SetURL(ctx context.Context, key, url string) error {
	return r.client.Set(ctx, key, url, 0).Err()
}

func (r *redisStorage) GetURL(ctx context.Context, key string) (string, error) {
	return r.client.Get(ctx, key).Result()
}

// NewBoltDB returns a new BoltDB storage (placeholder)
func NewBoltDB(path string) Storage {
	return nil
}

// NewPostgres returns a new Postgres storage (placeholder)
func NewPostgres(url string) Storage {
	return nil
}
