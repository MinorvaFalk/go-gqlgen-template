package resolver

import (
	"go-gqlgen-template/ent"
	"go-gqlgen-template/graph/generated"
	"go-gqlgen-template/pkg/adapter/controller"

	"github.com/99designs/gqlgen/graphql"
)

type Resolver struct {
	client     *ent.Client
	controller controller.Controller
}

func NewSchema(
	client *ent.Client,
	controller controller.Controller,
) graphql.ExecutableSchema {
	return generated.NewExecutableSchema(generated.Config{
		Resolvers: &Resolver{
			client:     client,
			controller: controller,
		},
	})
}
