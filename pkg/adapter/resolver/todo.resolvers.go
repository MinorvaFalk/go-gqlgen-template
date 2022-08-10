package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	tools "go-gqlgen-template"
	"go-gqlgen-template/ent"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input ent.CreateTodoInput) (*tools.MutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

// UpdateTodo is the resolver for the updateTodo field.
func (r *mutationResolver) UpdateTodo(ctx context.Context, input ent.UpdateTodoInput) (*tools.MutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

// AllTodo is the resolver for the allTodo field.
func (r *queryResolver) AllTodo(ctx context.Context) ([]*ent.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

// GetTodo is the resolver for the getTodo field.
func (r *queryResolver) GetTodo(ctx context.Context, id string) (*ent.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}