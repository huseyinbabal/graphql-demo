package main

import (
	"log"
	"net/http"

	"graphql-books/graph"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
)

func main() {
	resolver := &graph.Resolver{}
	resolver.SeedData()

	srv := handler.NewDefaultServer(
		graph.NewExecutableSchema(graph.Config{Resolvers: resolver}),
	)
	srv.AddTransport(&transport.Websocket{})

	http.Handle("/", playground.Handler("GraphQL Playground", "/graphql"))
	http.Handle("/graphql", srv)

	log.Println("Server started on :8080")
	log.Println("GraphQL Playground: http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
