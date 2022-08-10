package usecase

import (
	"context"
	"go-gqlgen-template/pkg/entity/model"
	"go-gqlgen-template/pkg/usecase/repository"
)

type todo struct {
	todoRepository repository.Todo
}

type Todo interface {
	Get(ctx context.Context, id *int) (*model.Todo, error)
	GetAll(ctx context.Context) ([]*model.Todo, error)
}

func NewTodoUseCase(r repository.Todo) Todo {
	return &todo{todoRepository: r}
}

func (t *todo) Get(ctx context.Context, id *int) (*model.Todo, error) {
	return t.todoRepository.Get(ctx, id)
}

func (t *todo) GetAll(ctx context.Context) ([]*model.Todo, error) {
	return t.todoRepository.GetAll(ctx)
}
