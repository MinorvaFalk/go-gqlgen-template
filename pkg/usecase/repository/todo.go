package repository

import (
	"context"
	"go-gqlgen-template/pkg/entity/model"
)

type Todo interface {
	Get(ctx context.Context, id *int) (*model.Todo, error)
	GetAll(ctx context.Context) ([]*model.Todo, error)
}
