package main

import (
	"context"
	"database/sql"
	"fmt"
	"go-gqlgen-template/config"
	"go-gqlgen-template/ent"
	"go-gqlgen-template/ent/migrate"
	"log"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v4/stdlib"
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

	// urlExample := "postgres://username:password@localhost:5432/database_name"
	dsn := fmt.Sprintf("postgres://%s:%s@%s/%s",
		config.C.Database.DBUser,
		config.C.Database.DBPassword,
		config.C.Database.DBHost,
		config.C.Database.DBName,
	)

	db, err := sql.Open("pgx", dsn)

	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(drv)), err
}
