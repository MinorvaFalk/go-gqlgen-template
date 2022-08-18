package repository

import (
	"context"
	"go-gqlgen-template/pkg/entity/model"
)

type Todo interface {
	Get(ctx context.Context, id *model.ID) (*model.Todo, error)
	GetAll(ctx context.Context) ([]*model.Todo, error)

	Create(ctx context.Context, input model.CreateTodoInput) (*model.Todo, error)
}
