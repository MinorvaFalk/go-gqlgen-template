package usecase

import (
	"context"
	"go-gqlgen-template/ent"
	"go-gqlgen-template/pkg/entity/model"
	"go-gqlgen-template/pkg/usecase/repository"
)

type user struct {
	userRepository repository.User
}

type User interface {
	Get(ctx context.Context, id *model.ID) (*model.User, error)
	List(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.UserWhereInput) (*model.UserConnection, error)
	GetTodo(ctx context.Context, id *model.ID) ([]*model.Todo, error)

	Create(ctx context.Context, input ent.CreateUserInput) (*model.User, error)
}

func NewUserUseCase(r repository.User) User {
	return &user{userRepository: r}
}

func (u *user) Get(ctx context.Context, id *model.ID) (*model.User, error) {
	return u.userRepository.Get(ctx, id)
}

func (u *user) List(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.UserWhereInput) (*model.UserConnection, error) {
	return u.userRepository.List(ctx, after, first, before, last, where)
}

func (u *user) GetTodo(ctx context.Context, id *model.ID) ([]*model.Todo, error) {
	return u.userRepository.GetTodo(ctx, id)
}

func (u *user) Create(ctx context.Context, input ent.CreateUserInput) (*model.User, error) {
	return u.userRepository.Create(ctx, input)
}
