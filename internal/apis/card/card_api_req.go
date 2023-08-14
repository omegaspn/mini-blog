package card

import "github.com/omegaspn/mini-blog/internal/domain/model"

type CreateOrUpdateCardRequest struct {
	Name     string             `json:"name" binding:"required"`
	Status   model.CardStatus   `json:"status" binding:"required"`
	Content  string             `json:"content" binding:"required"`
	Category model.CardCategory `json:"category" binding:"required"`
}
