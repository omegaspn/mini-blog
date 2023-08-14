package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/omegaspn/mini-blog/internal/apis/card"
	"github.com/omegaspn/mini-blog/internal/apis/middlewares"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Config mongo uri here
var mongoUri = "mongodb://localhost:27017"

func main() {
	// Connect to MongoDB
	client, err := connectDB()
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()

	v1 := router.Group("api/v1")
	{
		// Route to generate and return JWT token
		v1.GET("/token/:author", func(c *gin.Context) {
			author := c.Param("author")
			token, err := middlewares.GenerateToken(author)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating token"})
				return
			}
			c.String(http.StatusOK, token)
		})

		cards := v1.Group("/cards")
		{
			cardHandler := card.Handler{DBClient: client}
			cards.POST("", middlewares.AuthMiddleware(), cardHandler.Create)
			cards.PUT(":id", middlewares.AuthMiddleware(), cardHandler.Update)
			cards.DELETE(":id", middlewares.AuthMiddleware(), cardHandler.Delete)
		}

	}
	log.Fatal(router.Run(":8080"))
}

func connectDB() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(mongoUri)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	return client, nil
}
