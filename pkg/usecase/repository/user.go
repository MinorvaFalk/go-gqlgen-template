package repository

import (
	"context"
	"go-gqlgen-template/ent"
	"go-gqlgen-template/pkg/entity/model"
)

type User interface {
	Get(ctx context.Context, id *model.ID) (*model.User, error)
	List(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.UserWhereInput) (*model.UserConnection, error)
	GetTodo(ctx context.Context, id *model.ID) ([]*model.Todo, error)

	Create(ctx context.Context, input ent.CreateUserInput) (*model.User, error)
}
