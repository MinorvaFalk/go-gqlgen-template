package controller

import (
	"context"
	"go-gqlgen-template/ent"
	"go-gqlgen-template/pkg/entity/model"
	"go-gqlgen-template/pkg/usecase/usecase"
)

type todo struct {
	todoUseCase usecase.Todo
}

type Todo interface {
	Get(ctx context.Context, id *int) (*model.Todo, error)
	GetAll(ctx context.Context) ([]*model.Todo, error)
	Create(ctx context.Context, input ent.CreateTodoInput) (*ent.Todo, error)
}

func NewTodoController(ut usecase.Todo) Todo {
	return &todo{todoUseCase: ut}
}

func (t *todo) Get(ctx context.Context, id *int) (*model.Todo, error) {
	return t.todoUseCase.Get(ctx, id)
}

func (t *todo) GetAll(ctx context.Context) ([]*model.Todo, error) {
	return t.todoUseCase.GetAll(ctx)
}

func (t *todo) Create(ctx context.Context, input ent.CreateTodoInput) (*ent.Todo, error) {
	return t.todoUseCase.Create(ctx, input)
}
