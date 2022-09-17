package main

import (
	"context"
	"fmt"
	"log"
	"time"
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	// Timeout operations after N seconds
	connectTimeout           = 5
	connectionStringTemplate = "mongodb://%s:%s@%s/?directConnect=true"
)

func main() {
		r := gin.Default()
		handler := Handler{}

		r.POST("/create", handler.HandleCreateDocument)
		r.GET("/find", handler.HandleGetDocument)
		r.Run()
}
// GetConnection - Retrieves a client to the DocumentDB
func getConnection() (*mongo.Client, context.Context, context.CancelFunc) {
	connectionStr := fmt.Sprintf("mongodb://%s:%s@mongo:27017", os.Getenv("ME_CONFIG_MONGODB_ADMINUSERNAME"), os.Getenv("ME_CONFIG_MONGODB_ADMINPASSWORD"))
	client, err := mongo.NewClient(options.Client().ApplyURI(connectionStr))
	if err != nil {
		log.Printf("Failed to create client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout *time.Second)

	err = client.Connect(ctx)
	if err != nil {
		log.Printf("Failed to connect to cluster: %v", err)
	}

	// Force a connection to verify our connection string
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Printf("Failed to ping cluster: %v", err)
	}

	fmt.Println("Connected to MongoDB!")
	return client, ctx, cancel
}