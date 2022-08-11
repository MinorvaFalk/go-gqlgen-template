package repository

import (
	"context"
	"go-gqlgen-template/ent"
	"go-gqlgen-template/ent/todo"
	"go-gqlgen-template/pkg/entity/model"
	"go-gqlgen-template/pkg/usecase/repository"
)

type todoRepository struct {
	client *ent.Client
}

func NewTodoRepository(client *ent.Client) repository.Todo {
	return &todoRepository{client: client}
}

func (r *todoRepository) Get(ctx context.Context, id *int) (*model.Todo, error) {
	t, err := r.client.Todo.Query().Where(todo.IDEQ(*id)).Only(ctx)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (r *todoRepository) GetAll(ctx context.Context) ([]*model.Todo, error) {
	t, err := r.client.Todo.Query().Order(ent.Asc(todo.FieldID)).All(ctx)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (r *todoRepository) Create(ctx context.Context, input ent.CreateTodoInput) (*model.Todo, error) {
	t, err := r.client.Todo.Create().SetInput(input).Save(ctx)
	if err != nil {
		return nil, err
	}

	return t, err
}
