package card

import "github.com/omegaspn/mini-blog/internal/domain/model"

type CreateCardRequest struct {
	Name     string             `json:"name"`
	Status   model.CardStatus   `json:"status"`
	Content  string             `json:"content"`
	Category model.CardCategory `json:"category"`
	Author   string             `json:"author"`
}

type UpdateCardRequest struct {
	Name     string             `json:"name"`
	Status   model.CardStatus   `json:"status"`
	Content  string             `json:"content"`
	Category model.CardCategory `json:"category"`
	Author   string             `json:"author"`
}
