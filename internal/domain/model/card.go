package model

import "time"

type CardStatus string

const (
	CardStatusGreen  CardStatus = "GREEN"
	CardStatusViolet CardStatus = "VIOLET"
	CardStatusBlue   CardStatus = "BLUE"
	CardStatusOrange CardStatus = "ORANGE"
)

var ValidCardStatuses = map[CardStatus]bool{
	CardStatusGreen:  true,
	CardStatusViolet: true,
	CardStatusBlue:   true,
	CardStatusOrange: true,
}

type CardCategory string

const (
	CardCategoryPhy  CardCategory = "PHYSICS"
	CardCategoryTech CardCategory = "TECHNOLOGY"
	CardCategoryChem CardCategory = "CHEMISTRY"
	CardCategorySoc  CardCategory = "SOCIOLOGY"
)

var ValidCardCategories = map[CardCategory]bool{
	CardCategoryPhy:  true,
	CardCategoryTech: true,
	CardCategoryChem: true,
	CardCategorySoc:  true,
}

type Card struct {
	Name      string       `json:"name" bson:"name"`
	Status    CardStatus   `json:"status" bson:"status"`
	Content   string       `json:"content" bson:"content"`
	Category  CardCategory `json:"category" bson:"category"`
	Author    string       `json:"author" bson:"author"`
	CreatedAt time.Time    `json:"createdAt" bson:"created_at"`
	UpdatedAt time.Time    `json:"updatedAt" bson:"updated_at"`
}
