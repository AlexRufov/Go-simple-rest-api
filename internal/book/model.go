package book

import (
	"RestApi/internal/author/model"
)

type Book struct {
	ID      string         `json:"id"`
	Name    string         `json:"name"`
	Authors []model.Author `json:"authors"`
}
