package datastore

import (
	"fmt"
	"go-gqlgen-template/config"
	"go-gqlgen-template/ent"

	_ "github.com/lib/pq"
)

func PostgresClient() (*ent.Client, error) {
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
