package middlewares

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Claims struct to hold JWT claims
type Claims struct {
	Author string `json:"author"`
	jwt.StandardClaims
}

// Secret key for JWT signing and validation
var jwtKey = []byte("any-secret-key")

// AuthMiddleware Middleware function for JWT authentication
func AuthMiddleware() gin.HandlerFunc {
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

// GenerateToken Generate JWT token
func GenerateToken(author string) (string, error) {
	claims := &Claims{
		Author: author,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
