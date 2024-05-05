package models

import (
	"time"

	"github.com/google/uuid"
)

type RecipeSummary struct {
	ID         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	Descrption string    `json:"description"`
	MainPhoto  string    `json:"mainPhoto"`
	Created    time.Time `json:"created"`
	Updated    time.Time `json:"updated"`
}

type RecipeDetails struct {
	RecipeSummary
	Ingredients []Ingredient `json:"ingredients"`
	Steps       []Step       `json:"steps"`
}

type Ingredient struct {
	ID      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	Photo   string    `json:"photo"`
	Amount  int       `json:"amount"`
	Unit    string    `json:"unit"`
	Added   time.Time `json:"added"`
	Updated time.Time `json:"updated"`
}

type Step struct {
	ID      uuid.UUID `json:"id"`
	Order   int       `json:"order"`
	Text    string    `json:"text"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}
