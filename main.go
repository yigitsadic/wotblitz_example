package main

import (
	"context"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"github.com/yigitsadic/wotblitz_example/database"
	"github.com/yigitsadic/wotblitz_example/ent"
	"github.com/yigitsadic/wotblitz_example/ent/migrate"
	"github.com/yigitsadic/wotblitz_example/graph"
	"github.com/yigitsadic/wotblitz_example/graph/generated"
	"log"
	"net/http"
)

func init() {
	viper.AutomaticEnv()

	viper.SetDefault("PORT", "3700")
	viper.SetDefault("DATA_SOURCE", "root:@tcp(localhost:3306)/wotblitzexample")
}

func main() {
	client, err := ent.Open("mysql", viper.GetString("DATA_SOURCE"))
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

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{Client: client}}))

	r.Handle("/", playground.Handler("GraphQL playground", "/query"))
	r.Handle("/query", srv)

	log.Fatal(http.ListenAndServe(":"+viper.GetString("PORT"), r))
}
