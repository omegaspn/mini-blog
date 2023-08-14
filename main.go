package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/omegaspn/mini-blog/internal/apis/card"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Secret key for JWT signing and validation
var jwtKey = []byte("any-secret-key")

// Config mongo uri here
var mongoUri = "mongodb://localhost:27017"

// Claims struct to hold JWT claims
type Claims struct {
	Author string `json:"author"`
	jwt.StandardClaims
}

// Middleware function for JWT authentication
func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token is required"})
			c.Abort()
			return
		}

		// Parse JWT token
		token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization token"})
			c.Abort()
			return
		}

		// Set author in Gin context for later use
		claims, _ := token.Claims.(*Claims)
		c.Set("author", claims.Author)
		c.Next()
	}
}

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
			token, err := generateToken(author)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating token"})
				return
			}
			c.String(http.StatusOK, token)
		})

		cards := v1.Group("/cards")
		{
			cardHandler := card.Handler{DBClient: client}
			cards.POST("", authMiddleware(), cardHandler.Create)
			cards.PUT(":id", authMiddleware(), cardHandler.Update)
			cards.DELETE(":id", authMiddleware(), cardHandler.Delete)
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

// Generate JWT token
func generateToken(author string) (string, error) {
	claims := &Claims{
		Author: author,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
