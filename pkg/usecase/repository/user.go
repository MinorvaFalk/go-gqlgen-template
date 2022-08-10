package repository

import (
	"context"
	"go-gqlgen-template/pkg/entity/model"
)

type User interface {
	Get(ctx context.Context, id *int) (*model.User, error)
	GetAll(ctx context.Context) ([]*model.User, error)
}
