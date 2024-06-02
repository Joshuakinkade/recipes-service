package repositories

import (
	"context"

	"github.com/google/uuid"
	db "github.com/joshuakinkade/recipes-service/db/sqlc"
	"github.com/joshuakinkade/recipes-service/models"
)

type IRecipeRepository interface {
	ListRecipes(ctx context.Context, userID uuid.UUID) ([]models.RecipeSummary, error)
	GetRecipe(ctx context.Context, userID uuid.UUID) (models.RecipeDetails, error)
}

type RecipeRepository struct {
	queries *db.Queries
}

func NewRecipeRepository(queries *db.Queries) RecipeRepository {
	return RecipeRepository{queries}
}

func (r RecipeRepository) ListRecipes(ctx context.Context, userID uuid.UUID) ([]models.RecipeSummary, error) {
	results, err := r.queries.ListRecipes(ctx, userID)
	if err != nil {
		return nil, err
	}
	recipes := []models.RecipeSummary{}
	for _, result := range results {
		recipe := models.RecipeSummary{
			ID:          result.ID,
			Name:        result.Name,
			Description: result.Description,
			Created:     result.CreatedAt,
			Updated:     result.UpdatedAt,
		}
		recipes = append(recipes, recipe)
	}
	return recipes, nil
}

func (r RecipeRepository) GetRecipe(ctx context.Context, recipeID uuid.UUID) (models.RecipeDetails, error) {
	_, err := r.queries.GetRecipe(ctx, recipeID)
	if err != nil {
		return models.RecipeDetails{}, err
	}

	return models.RecipeDetails{}, nil
}
