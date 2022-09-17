package main

import (
	"context"
	"fmt"
	"log"
	"time"

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
	username := "root"// os.Getenv("MONGODB_USERNAME")
	password := "example"// os.Getenv("MONGODB_PASSWORD")
	clusterEndpoint := "documents"

	connectionURI := fmt.Sprintf(connectionStringTemplate, username, password, clusterEndpoint)
	fmt.Println(connectionURI)
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:example@mongo:27017"))
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