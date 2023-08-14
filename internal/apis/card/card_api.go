package card

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/omegaspn/mini-blog/internal/domain/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Handler struct {
	DBClient *mongo.Client
}

var blogDb = "blog"
var cardCollection = "cards"

func (ch *Handler) verifyAuthor(ctx context.Context, author string, oid primitive.ObjectID) error {
	card := ch.DBClient.Database(blogDb).Collection(cardCollection).FindOne(ctx, bson.M{"_id": oid})
	result := model.Card{}
	err := card.Decode(&result)
	if err != nil {
		return err
	}

	if result.Author != author {
		return errors.New(fmt.Sprintf("author: %s can't manage this card", author))
	}
	return nil
}

func (ch *Handler) validateCardStatus(req CreateOrUpdateCardRequest) bool {
	_, valid := model.ValidCardStatuses[req.Status]
	return valid
}

func (ch *Handler) validateCardCategory(req CreateOrUpdateCardRequest) bool {
	_, valid := model.ValidCardCategories[req.Category]
	return valid
}

// @BasePath /api/v1

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /example/helloworld [get]
func (ch *Handler) Create(gctx *gin.Context) {
	ctx := gctx.Request.Context()

	var req CreateOrUpdateCardRequest
	if err := gctx.ShouldBindJSON(&req); err != nil {
		gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	valid := ch.validateCardStatus(req)
	if !valid {
		gctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid card status"})
		return
	}

	valid = ch.validateCardCategory(req)
	if !valid {
		gctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid card category"})
		return
	}

	author := gctx.GetString("author")
	result, err := ch.DBClient.Database(blogDb).Collection(cardCollection).InsertOne(ctx, model.Card{
		Name:      req.Name,
		Status:    req.Status,
		Content:   req.Content,
		Category:  req.Category,
		Author:    author,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		gctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	gctx.JSON(http.StatusOK, CreateCardResponse{
		Status: "success",
		ID:     result.InsertedID,
	})
}

func (ch *Handler) Update(gctx *gin.Context) {
	ctx := gctx.Request.Context()

	var req CreateOrUpdateCardRequest
	if err := gctx.ShouldBindJSON(&req); err != nil {
		gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	valid := ch.validateCardStatus(req)
	if !valid {
		gctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid card status"})
		return
	}

	valid = ch.validateCardCategory(req)
	if !valid {
		gctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid card category"})
		return
	}

	var key = "id"
	cardId := gctx.Param(key)
	if cardId == "" {
		gctx.JSON(http.StatusBadRequest, gin.H{"error": "card oid to update can't be empty"})
		return
	}

	oid, err := primitive.ObjectIDFromHex(cardId)
	if err != nil {
		gctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	// Verify only author can perform action
	err = ch.verifyAuthor(ctx, gctx.GetString("author"), oid)
	if err != nil {
		gctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	updateCmd := bson.D{{"$set", bson.M{
		"name":       req.Name,
		"status":     req.Status,
		"content":    req.Content,
		"category":   req.Category,
		"updated_at": time.Now(),
	}}}

	_, err = ch.DBClient.Database(blogDb).Collection(cardCollection).UpdateOne(ctx, bson.M{"_id": oid}, updateCmd)
	if err != nil {
		gctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	gctx.JSON(http.StatusOK, UpdateCardResponse{
		Status: "success",
	})
}

func (ch *Handler) Delete(gctx *gin.Context) {
	ctx := gctx.Request.Context()

	var key = "id"
	cardId := gctx.Param(key)
	if cardId == "" {
		gctx.JSON(http.StatusBadRequest, gin.H{"error": "card oid to delete can't be empty"})
		return
	}

	oid, err := primitive.ObjectIDFromHex(cardId)
	if err != nil {
		gctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Verify only author can perform action
	err = ch.verifyAuthor(ctx, gctx.GetString("author"), oid)
	if err != nil {
		gctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ret, err := ch.DBClient.Database(blogDb).Collection(cardCollection).DeleteOne(context.Background(), bson.M{"_id": oid})
	if ret.DeletedCount == 0 {
		gctx.JSON(http.StatusInternalServerError, gin.H{"error": "no items to delete"})
		return
	}

	if err != nil {
		gctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	gctx.JSON(http.StatusOK, DeleteCardResponse{
		Status: "success",
		ID:     oid,
	})
}
