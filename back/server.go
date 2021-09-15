package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/akatebi/todos/graph/database"
	"github.com/akatebi/todos/graph/generated"
	"github.com/go-chi/chi"
	"github.com/rs/cors"
)

const defaultPort = "8080"

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	CORS := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3001"},
		AllowCredentials: true,
	})

	resolver := &database.Resolver{}
	resolver.Open()
	defer resolver.Close()

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))

	router := chi.NewRouter()
	// router.Use(Middleware())

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", CORS.Handler(srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))

}

// Middleware decodes the share session cookie and packs the session into context
func Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Printf("#### Request %#v", *r)

			printBody(r)

			next.ServeHTTP(w, r)
		})
	}
}

func printBody(r *http.Request) {
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading request body: %v", err.Error())

	}

	type Graphql struct {
		operationName string
		variables     string
		query         string
	}

	var graphql Graphql
	json.Unmarshal([]byte(buf), &graphql)
	log.Printf("%#v", graphql)

	log.Printf("Request body: %v", string(buf))
	reader := ioutil.NopCloser(bytes.NewBuffer(buf))
	r.Body = reader
}
