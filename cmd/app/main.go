package main

import (
	"go-gqlgen-template/config"
	"go-gqlgen-template/ent"
	"go-gqlgen-template/pkg/adapter/controller"
	"go-gqlgen-template/pkg/infrastructure/datastore"
	"go-gqlgen-template/pkg/infrastructure/graphql"
	"go-gqlgen-template/pkg/infrastructure/router"
	"go-gqlgen-template/pkg/registry"
	"log"
)

func main() {
	config.ReadConfig(config.ConfigOption{})

	client := newDBClient()
	c := newController(client)

	srv := graphql.NewServer(client, c)
	e := router.New(srv)

	e.Logger.Fatal(e.Start(":" + config.C.Server.Port))
}

func newDBClient() *ent.Client {
	client, err := datastore.PostgresClient()
	if err != nil {
		log.Fatalf("failed to start postgres client: %v", err)
	}

	return client
}

func newController(client *ent.Client) controller.Controller {
	r := registry.New(client)
	return r.NewController()
}
