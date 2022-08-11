package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"go-gqlgen-template/ent"
	"go-gqlgen-template/ent/todo"
	"go-gqlgen-template/pkg/entity/model"
	"strconv"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input ent.CreateTodoInput) (*model.MutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

// UpdateTodo is the resolver for the updateTodo field.
func (r *mutationResolver) UpdateTodo(ctx context.Context, input ent.UpdateTodoInput) (*model.MutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

// AllTodo is the resolver for the allTodo field.
func (r *queryResolver) AllTodo(ctx context.Context) ([]*ent.Todo, error) {
	t, err := r.client.Todo.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	return t, nil
}

// GetTodo is the resolver for the getTodo field.
func (r *queryResolver) GetTodo(ctx context.Context, id string) (*ent.Todo, error) {
	i, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	t, err := r.client.Todo.Query().Where(todo.IDEQ(i)).Only(ctx)
	if err != nil {
		return nil, err
	}

	return t, nil
}
