package token

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/omegaspn/mini-blog/internal/apis/middlewares"
	"go.mongodb.org/mongo-driver/mongo"
)

type Handler struct {
	DBClient *mongo.Client
}

// GenerateToken @Summary to generate token from author
// @Tags token
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param author path string true "author"
// @Success 200
// @Router /token/{author} [get]
func (ch *Handler) GenerateToken(c *gin.Context) {
	author := c.Param("author")
	token, err := middlewares.GenerateToken(author)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating token"})
		return
	}
	c.String(http.StatusOK, token)
}
