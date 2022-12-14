package usecase

import (
	"context"
	"go-gqlgen-template/ent"
	"go-gqlgen-template/pkg/entity/model"
	"go-gqlgen-template/pkg/usecase/repository"
)

type todo struct {
	todoRepository repository.Todo
}

type Todo interface {
	Get(ctx context.Context, id *model.ID) (*model.Todo, error)
	GetAll(ctx context.Context) ([]*model.Todo, error)
	Create(ctx context.Context, input ent.CreateTodoInput) (*model.Todo, error)
}

func NewTodoUseCase(r repository.Todo) Todo {
	return &todo{todoRepository: r}
}

func (t *todo) Get(ctx context.Context, id *model.ID) (*model.Todo, error) {
	return t.todoRepository.Get(ctx, id)
}

func (t *todo) GetAll(ctx context.Context) ([]*model.Todo, error) {
	return t.todoRepository.GetAll(ctx)
}

func (t *todo) Create(ctx context.Context, input ent.CreateTodoInput) (*model.Todo, error) {
	return t.todoRepository.Create(ctx, input)
}
