package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/omegaspn/mini-blog/internal/apis/card"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Connect to MongoDB
	client, err := connectDB()
	if err != nil {
		log.Fatal(err)
	}

	cardHandler := card.Handler{DBClient: client}

	router := gin.Default()

	// Routes
	router.GET("/create", cardHandler.Create)

	err = router.Run(":8080")
	if err != nil {
		// TODO: handle error
	}
}

func connectDB() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	return client, nil
}
