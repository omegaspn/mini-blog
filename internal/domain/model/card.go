package model

type Card struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}

type Response struct {
	Status string `json:"status"`
}
