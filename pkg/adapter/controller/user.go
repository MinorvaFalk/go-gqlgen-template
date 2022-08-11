package controller

import (
	"context"
	"go-gqlgen-template/pkg/entity/model"
	"go-gqlgen-template/pkg/usecase/usecase"
)

type user struct {
	userUseCase usecase.User
}

type User interface {
	Get(ctx context.Context, id *int) (*model.User, error)
	GetAll(ctx context.Context) ([]*model.User, error)
	GetTodo(ctx context.Context, id *int) ([]*model.Todo, error)
}

func NewUserController(uu usecase.User) User {
	return &user{userUseCase: uu}
}

func (u *user) Get(ctx context.Context, id *int) (*model.User, error) {
	return u.userUseCase.Get(ctx, id)
}

func (u *user) GetAll(ctx context.Context) ([]*model.User, error) {
	return u.userUseCase.GetAll(ctx)
}

func (u *user) GetTodo(ctx context.Context, id *int) ([]*model.Todo, error) {
	return u.userUseCase.GetTodo(ctx, id)
}
