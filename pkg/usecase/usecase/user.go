package usecase

import (
	"context"
	"go-gqlgen-template/pkg/entity/model"
	"go-gqlgen-template/pkg/usecase/repository"
)

type user struct {
	userRepository repository.User
}

type User interface {
	Get(ctx context.Context, id *int) (*model.User, error)
	GetAll(ctx context.Context) ([]*model.User, error)
}

func NewUserUseCase(r repository.User) User {
	return &user{userRepository: r}
}

func (u *user) Get(ctx context.Context, id *int) (*model.User, error) {
	return u.userRepository.Get(ctx, id)
}

func (u *user) GetAll(ctx context.Context) ([]*model.User, error) {
	return u.userRepository.GetAll(ctx)
}
