package main

import (
	"fmt"
	"go-gqlgen-template/config"
	"go-gqlgen-template/ent"
	"go-gqlgen-template/graph"
	"log"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const defaultPort = "4001"

func main() {
	config.ReadConfig(config.ConfigOption{})
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{
			"http://localhost:4000",
			"https://studio.apollographql.com",
		},
		AllowCredentials: true,
	}))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Server Running")
	})

	client, err := NewClient()
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}
	defer client.Close()

	// srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	srv := handler.NewDefaultServer(graph.NewSchema(client))
	e.POST("/query", echo.WrapHandler(srv))
	e.GET("/playground", func(c echo.Context) error {
		playground.Handler("Graphql Playground", "/query").ServeHTTP(c.Response(), c.Request())

		return nil
	})

	e.Logger.Fatal(e.Start(":" + config.C.Server.Port))
}

func NewClient() (*ent.Client, error) {
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
