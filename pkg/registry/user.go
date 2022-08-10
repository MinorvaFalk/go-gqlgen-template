package registry

import (
	"go-gqlgen-template/pkg/adapter/controller"
	"go-gqlgen-template/pkg/adapter/repository"
	"go-gqlgen-template/pkg/usecase/usecase"
)

func (r *registry) NewUserController() controller.User {
	repo := repository.NewUserRepository(r.client)
	u := usecase.NewUserUseCase(repo)

	return controller.NewUserController(u)
}
