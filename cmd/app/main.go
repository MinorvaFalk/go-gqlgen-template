package main

import (
	"go-gqlgen-template/config"
	"go-gqlgen-template/graph"
	"go-gqlgen-template/graph/generated"
	"net/http"

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

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	e.POST("/query", echo.WrapHandler(srv))
	e.GET("/playground", func(c echo.Context) error {
		playground.Handler("Graphql Playground", "/query").ServeHTTP(c.Response(), c.Request())

		return nil
	})

	e.Logger.Fatal(e.Start(":" + config.C.Server.Port))
}
