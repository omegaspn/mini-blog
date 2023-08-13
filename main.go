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

	router := gin.Default()
	v1 := router.Group("api/v1")
	{
		cards := v1.Group("/cards")
		{
			cardHandler := card.Handler{DBClient: client}
			cards.POST("", cardHandler.Create)
			cards.PUT(":id", cardHandler.Update)
			cards.DELETE(":id", cardHandler.Delete)
		}

	}
	log.Fatal(router.Run(":8080"))
}

func connectDB() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	return client, nil
}
