package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"go-gqlgen-template/ent"
	"go-gqlgen-template/pkg/entity/model"
	"go-gqlgen-template/utils"
	"strconv"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input ent.CreateTodoInput) (*model.MutationResponse, error) {
	t, err := r.controller.Todo.Create(ctx, input)
	if err != nil {
		return nil, err
	}

	return &model.MutationResponse{
		Code:    &utils.SuccessCode,
		Success: &utils.SuccessStatus,
		Message: &utils.SuccessMessage,
		Data:    t,
	}, nil
}

// UpdateTodo is the resolver for the updateTodo field.
func (r *mutationResolver) UpdateTodo(ctx context.Context, input ent.UpdateTodoInput) (*model.MutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

// AllTodo is the resolver for the allTodo field.
func (r *queryResolver) AllTodo(ctx context.Context) ([]*ent.Todo, error) {
	t, err := r.controller.Todo.GetAll(ctx)
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

	t, err := r.controller.Todo.Get(ctx, &i)
	if err != nil {
		return nil, err
	}

	return t, nil
}
