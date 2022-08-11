package datastore

import (
	"database/sql"
	"fmt"
	"go-gqlgen-template/config"
	"go-gqlgen-template/ent"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func PostgresClient() (*ent.Client, error) {
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
