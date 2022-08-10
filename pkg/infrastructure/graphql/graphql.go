package graphql

import (
	"go-gqlgen-template/ent"
	"go-gqlgen-template/pkg/adapter/controller"
	"go-gqlgen-template/pkg/adapter/resolver"

	"entgo.io/contrib/entgql"
	"github.com/99designs/gqlgen/graphql/handler"
)

func NewServer(client *ent.Client, controller controller.Controller) *handler.Server {
	srv := handler.NewDefaultServer(resolver.NewSchema(client, controller))
	srv.Use(entgql.Transactioner{TxOpener: client})

	return srv
}
