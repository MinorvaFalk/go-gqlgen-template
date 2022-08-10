package registry

import (
	"go-gqlgen-template/ent"
	"go-gqlgen-template/pkg/adapter/controller"
)

type registry struct {
	client *ent.Client
}

type Registry interface {
	NewController() controller.Controller
}

func New(client *ent.Client) Registry {
	return &registry{
		client: client,
	}
}

func (r *registry) NewController() controller.Controller {
	return controller.Controller{
		User: r.NewUserController(),
	}
}
