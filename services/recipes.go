package services

import (
	"github.com/google/uuid"
	"github.com/joshuakinkade/recipes-service/models"
)

type IRecipeService interface {
	// Base operations
	List(userId uuid.UUID) ([]models.RecipeSummary, error)
	Get(id uuid.UUID) (models.RecipeDetails, error)
	Create(name, description string, userId uuid.UUID) (uuid.UUID, error)
	Update(id uuid.UUID, name, description string) error
	Delete(id uuid.UUID) error

	// Ingredient operations
	AddIngredient(recipeId uuid.UUID, order, quantity int, name string) error
	UpdateIngredient(recipeId, ingredientId uuid.UUID, quantity int, name string) error
	RemoveIngredient(recipeId, ingredientId uuid.UUID) error

	// Step operations
	AddStep(recipeId uuid.UUID, order int, text string) error
	UpdateStep(recipeId, stepId uuid.UUID, text string) error
	DeleteStep(recipeId, stepId uuid.UUID) error

	// Photo operations
	AddPhoto(recipeId uuid.UUID, url string, isMain bool) (uuid.UUID, error) // Todo: do I want to pass in an order and determine main from that (order 1 is main)?
	RemovePhoto(recipeId, photoId uuid.UUID) error

	// Tag operations
	AddTag(recipeId uuid.UUID, name string) error
	RemoveTag(recipeId uuid.UUID, tagId uuid.UUID) error
}

type RecipeService struct {
	// put repositories here
}

func NewRecipeService() RecipeService {
	return RecipeService{}
}

func (s RecipeService) List(userId uuid.UUID) ([]models.RecipeSummary, error) {
	// Todo: implement
	// Todo: call database
	return nil, nil
}
