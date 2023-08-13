package card

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/omegaspn/mini-blog/internal/domain/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type Handler struct {
	DBClient *mongo.Client
}

func (ch *Handler) Create(gctx *gin.Context) {
	dummyCard := model.Card{
		Name:   "my-card",
		Status: "foo",
	}

	cards := ch.DBClient.Database("blog").Collection("cards")
	_, err := cards.InsertOne(context.Background(), dummyCard)
	if err != nil {
		gctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	gctx.JSON(http.StatusOK, model.Response{
		Status: "success",
	})
}
