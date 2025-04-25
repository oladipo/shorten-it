# ---- Build Stage ----
FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o shorten-it ./cmd/main.go

# ---- Run Stage ----
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/shorten-it .
COPY --from=builder /app/migrations ./migrations
# Copy config files or static assets if needed
EXPOSE 8080
CMD ["./shorten-it"]
