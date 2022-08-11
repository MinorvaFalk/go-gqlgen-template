package repository

import (
	"context"
	"go-gqlgen-template/ent"
	"go-gqlgen-template/ent/todo"
	"go-gqlgen-template/ent/user"
	"go-gqlgen-template/pkg/entity/model"
	"go-gqlgen-template/pkg/usecase/repository"
)

type userRepository struct {
	client *ent.Client
}

func NewUserRepository(client *ent.Client) repository.User {
	return &userRepository{client: client}
}

func (r *userRepository) Get(ctx context.Context, id *int) (*model.User, error) {
	u, err := r.client.User.Query().Where(user.IDEQ(*id)).Only(ctx)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (r *userRepository) GetAll(ctx context.Context) ([]*model.User, error) {
	u, err := r.client.User.Query().Order(ent.Asc(user.FieldID)).All(ctx)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (r *userRepository) GetTodo(ctx context.Context, id *int) ([]*model.Todo, error) {
	t, err := r.client.Todo.Query().Where(todo.UserIDEQ(*id)).All(ctx)
	if err != nil {
		return nil, err
	}

	return t, err
}
