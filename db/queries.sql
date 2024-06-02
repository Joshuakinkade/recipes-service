-- name: ListRecipes :many
SELECT id, name, description, created_at, updated_at, deleted_at FROM recipes WHERE deleted_at IS NULL AND user_id = $1;

-- name: GetRecipe :one
SELECT r.id, r.name, r.description, r.created_at, r.updated_at, r.deleted_at, ri.id, ri.quantity, ri.unit, i.id, i.name, s.id, s.step_number, s.directions 
FROM recipes r LEFT JOIN recipe_ingredients ri ON r.id = ri.recipe LEFT JOIN ingredients i ON ri.ingredient = i.id LEFT JOIN steps s ON r.id = s.recipe 
WHERE r.deleted_at IS NULL AND r.id = $1;