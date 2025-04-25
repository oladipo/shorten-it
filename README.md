# Go URL Shortener Service

A Bitly-like URL shortener with analytics, built in Go using Gin.

## Features
- Shorten URLs
- Redirect with analytics (clicks, referrers, geolocation)
- Pluggable storage: BoltDB, Redis, or Postgres

## Quick Start

1. Install dependencies:
   ```bash
   go mod tidy
   ```
2. Run the server (BoltDB by default):
   ```bash
   STORE_TYPE=boltdb go run ./cmd/main.go
   ```
3. API endpoints:
   - POST `/shorten` `{ "url": "https://example.com" }`
   - GET `/:shortcode` (redirect)
   - GET `/:shortcode/stats` (analytics)

## Configuration
- `STORE_TYPE`: `boltdb`, `redis`, or `postgres`
- `DATABASE_URL`: For Postgres

## TODO
- Add Redis and Postgres support
- Improve geolocation (currently uses IP)
- Add authentication/rate limiting

---
MIT License
