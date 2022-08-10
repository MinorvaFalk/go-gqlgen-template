package main

import (
	"go-gqlgen-template/graph"
	"go-gqlgen-template/graph/generated"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/rs/cors"
)

const defaultPort = "4001"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	mux := http.NewServeMux()
	mux.Handle("/", playground.Handler("Graphql playgorund", "/query"))
	mux.Handle("/query", srv)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{
			"http://localhost:4000",
			"https://studio.apollographql.com",
		},
		AllowCredentials: true,
	})

	httpHandler := c.Handler(mux)
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, httpHandler))
}
