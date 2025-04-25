package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/oladipo/shorten-it/internal/api"
	"github.com/oladipo/shorten-it/internal/storage"
)

// SetupApp sets up the Gin engine and storage backend
func SetupApp(storeType string) (*gin.Engine, error) {
	var store storage.Storage

	switch storeType {
	case "boltdb":
		store = storage.NewBoltDB("urlshortener.db")
	case "redis":
		store = storage.NewRedis("localhost:6379", "")
	case "postgres":
		store = storage.NewPostgres(os.Getenv("DATABASE_URL"))
	default:
		return nil, ErrUnknownStoreType
	}

	r := gin.Default()
	api.RegisterRoutes(r, store)
	return r, nil
}

var ErrUnknownStoreType = &UnknownStoreTypeError{}

type UnknownStoreTypeError struct{}
func (e *UnknownStoreTypeError) Error() string { return "Unknown STORE_TYPE" }

func main() {
	storeType := os.Getenv("STORE_TYPE")
	r, err := SetupApp(storeType)
	if err != nil {
		log.Fatal(err)
	}
	r.Run(":8080")
}
