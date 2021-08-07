package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/akatebi/todos/graph/generated"
	"github.com/akatebi/todos/graph/memory"
	"github.com/rs/cors"
)

const defaultPort = "8080"

func main() {

	// db := database.InitDB()
	// defer db.Close()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	CORS := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
	})

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &memory.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", CORS.Handler(srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
