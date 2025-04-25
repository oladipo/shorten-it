package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/oladipo/shorten-it/internal/api"
	"github.com/oladipo/shorten-it/internal/storage"
)

func main() {
	// Load config (e.g., from env or file)
	storeType := os.Getenv("STORE_TYPE")
	var store storage.Storage

	switch storeType {
	case "boltdb":
		store = storage.NewBoltDB("urlshortener.db")
	case "redis":
		store = storage.NewRedis("localhost:6379", "")
	case "postgres":
		store = storage.NewPostgres(os.Getenv("DATABASE_URL"))
	default:
		log.Fatal("Unknown STORE_TYPE")
	}

	r := gin.Default()
	api.RegisterRoutes(r, store)

	r.Run(":8080")
}
