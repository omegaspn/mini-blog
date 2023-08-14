package card

import "go.mongodb.org/mongo-driver/bson/primitive"

type CreateCardResponse struct {
	Status string      `json:"status"`
	ID     interface{} `json:"id"`
}

type UpdateCardResponse struct {
	Status string `json:"status"`
}

type DeleteCardResponse struct {
	Status string             `json:"status"`
	ID     primitive.ObjectID `json:"id"`
}
