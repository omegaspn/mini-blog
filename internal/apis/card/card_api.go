package card

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/omegaspn/mini-blog/internal/domain/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Handler struct {
	DBClient *mongo.Client
}

func (ch *Handler) Create(gctx *gin.Context) {
	var req CreateCardRequest
	if err := gctx.ShouldBindJSON(&req); err != nil {
		_ = gctx.Error(errors.New("error from map req"))
		return
	}

	c := model.Card{
		Name:      req.Name,
		Status:    req.Status,
		Content:   req.Content,
		Category:  req.Category,
		Author:    req.Author,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	cards := ch.DBClient.Database("blog").Collection("cards")
	_, err := cards.InsertOne(context.Background(), c)
	if err != nil {
		gctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	gctx.JSON(http.StatusOK, model.Response{
		Status: "success",
	})
}

func (ch *Handler) Update(gctx *gin.Context) {
	var req UpdateCardRequest
	if err := gctx.ShouldBindJSON(&req); err != nil {
		_ = gctx.Error(errors.New("error from map req"))
		return
	}

	var key = "id"
	cardId := gctx.Param(key)

	if cardId == "" {
		gctx.JSON(http.StatusBadRequest, gin.H{"error": "empty card id"})
		return
	}

	c := model.Card{
		Name:      req.Name,
		Status:    req.Status,
		Content:   req.Content,
		Category:  req.Category,
		Author:    req.Author,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	cards := ch.DBClient.Database("blog").Collection("cards")
	_, err := cards.ReplaceOne(context.Background(), bson.M{"name": cardId}, c)
	if err != nil {
		gctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	gctx.JSON(http.StatusOK, model.Response{
		Status: "success",
	})
}

func (ch *Handler) Delete(gctx *gin.Context) {
	var key = "id"
	cardId := gctx.Param(key)

	if cardId == "" {
		gctx.JSON(http.StatusBadRequest, gin.H{"error": "empty card id"})
		return
	}

	cards := ch.DBClient.Database("blog").Collection("cards")
	ret, err := cards.DeleteOne(context.Background(), bson.M{"name": cardId})
	if ret.DeletedCount == 0 {
		gctx.JSON(http.StatusInternalServerError, gin.H{"error": "no items to delete"})
		return

	}

	if err != nil {
		gctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}

	gctx.JSON(http.StatusOK, model.Response{
		Status: "success",
	})
}
