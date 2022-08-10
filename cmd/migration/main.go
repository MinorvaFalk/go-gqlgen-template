package main

import (
	"context"
	"fmt"
	"go-gqlgen-template/config"
	"go-gqlgen-template/ent"
	"go-gqlgen-template/ent/migrate"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	config.ReadConfig(config.ConfigOption{})

	client, err := newClient()
	if err != nil {
		log.Fatalf("failed opening dsn client: %v", err)
	}

	defer client.Close()
	createSchema(client)
}

func createSchema(client *ent.Client) {
	if err := client.Schema.Create(
		context.Background(),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}

func newClient() (*ent.Client, error) {
	var entOptions []ent.Option
	entOptions = append(entOptions, ent.Debug())

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		config.C.Database.DBHost,
		config.C.Database.DBUser,
		config.C.Database.DBPassword,
		config.C.Database.DBName,
		config.C.Database.DBPort,
	)

	return ent.Open("postgres", dsn, entOptions...)
}
