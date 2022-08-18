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

func (r *userRepository) Get(ctx context.Context, id *model.ID) (*model.User, error) {
	u, err := r.client.User.Query().Where(user.IDEQ(*id)).Only(ctx)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (r *userRepository) List(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.UserWhereInput) (*model.UserConnection, error) {
	us, err := r.client.User.
		Query().
		Paginate(ctx, after, first, before, last, ent.WithUserFilter(where.Filter))
	if err != nil {
		return nil, model.NewDBError(err)
	}

	return us, nil
}

func (r *userRepository) GetTodo(ctx context.Context, id *model.ID) ([]*model.Todo, error) {
	t, err := r.client.Todo.Query().Where(todo.UserIDEQ(*id)).All(ctx)
	if err != nil {
		return nil, err
	}

	return t, err
}

func (r *userRepository) Create(ctx context.Context, input ent.CreateUserInput) (*model.User, error) {
	client := WithTransactionalMutation(ctx)

	todo, err := client.Todo.
		Create().
		Save(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}

	u, err := client.User.
		Create().
		SetInput(input).
		AddTodo(todo).
		Save(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}

	return u, nil
}
