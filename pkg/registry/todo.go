package registry

import (
	"go-gqlgen-template/pkg/adapter/controller"
	"go-gqlgen-template/pkg/adapter/repository"
	"go-gqlgen-template/pkg/usecase/usecase"
)

func (r *registry) NewTodoController() controller.Todo {
	repo := repository.NewTodoRepository(r.client)
	u := usecase.NewTodoUseCase(repo)

	return controller.NewTodoController(u)
}
