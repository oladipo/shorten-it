# Go URL Shortener Service

A Bitly-like URL shortener with analytics, built in Go using Gin.

## Features
- Shorten URLs
- Redirect with analytics (clicks, referrers, geolocation)
- Pluggable storage: BoltDB, Redis, or Postgres
- Prometheus metrics endpoint (`/metrics`)
- API key authentication & rate limiting

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
   - POST `/shorten` `{ "url": "https://example.com" }` (requires `X-API-Key` header)
   - GET `/:shortcode` (redirect)
   - GET `/:shortcode/stats` (analytics)
   - GET `/metrics` (Prometheus metrics)

## Configuration
- `STORE_TYPE`: `boltdb`, `redis`, or `postgres`
- `DATABASE_URL`: For Postgres
- `REDIS_ADDR`, `REDIS_PASSWORD`: For Redis
- `API_KEY`: Required for authenticated endpoints

## Running Tests

To run all tests (unit and integration):
```bash
go test ./...
```
To run tests for a specific package:
```bash
go test ./internal/api/...
```

## Deployment

1. Ensure all dependencies are installed and your configuration/environment variables are set.
2. Build the application:
   ```bash
   go build -o shorten-it ./cmd/main.go
   ```
3. Deploy the binary to your server or preferred cloud provider.
4. (Optional) For containerized deployment, use the provided Dockerfile:
   ```bash
   docker build -t shorten-it .
   docker run -e STORE_TYPE=boltdb -e API_KEY=yourkey -p 8080:8080 shorten-it
   ```
5. For cloud deployment (Heroku, Google Cloud Run, AWS, Netlify, etc.), see [DEPLOY.md](./DEPLOY.md).

## Continuous Integration (CI)

This project uses GitHub Actions for CI. See `.github/workflows/ci.yml` for details.
- On every push and pull request, the workflow checks out the code, installs dependencies, builds the app, and runs all tests.

## Sample Dockerfile

A multi-stage build for minimal images:
```Dockerfile
FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o shorten-it ./cmd/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/shorten-it .
EXPOSE 8080
CMD ["./shorten-it"]
```

## Cloud Deployment

See [DEPLOY.md](./DEPLOY.md) for step-by-step guides for Docker, Heroku, Google Cloud Run, AWS, and Netlify.

---
MIT License
