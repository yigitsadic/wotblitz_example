package main

import (
	"context"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/yigitsadic/wotblitz_example/database"
	"github.com/yigitsadic/wotblitz_example/ent"
	"github.com/yigitsadic/wotblitz_example/ent/migrate"
	"github.com/yigitsadic/wotblitz_example/graph"
	"github.com/yigitsadic/wotblitz_example/graph/generated"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3700"
	}
	client, err := ent.Open("mysql", "root:@tcp(localhost:3306)/wotblitzexample")
	if err != nil {
		log.Fatalf("failed connecting to mysql: %v", err)
	}
	defer client.Close()
	ctx := context.Background()
	// Run migration.
	err = client.Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)
	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	database.SeedDB(ctx, client)

	schema := database.LoadFromFile()
	repository := database.NewRepository(schema)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{Repository: repository, Client: client}}))

	r.Handle("/", playground.Handler("GraphQL playground", "/query"))
	r.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
