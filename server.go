package main

import (
	"Ticket_Booking_App/directives"
	"Ticket_Booking_App/graph"
	"Ticket_Booking_App/graph/generated"
	"Ticket_Booking_App/initializers"
	"Ticket_Booking_App/middlewares"
	"Ticket_Booking_App/migration"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/mux"
)

const defaultPort = "8080"

func main() {

	migration.MigrateTable()
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db := initializers.DB
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	router := mux.NewRouter()
	router.Use(middlewares.AuthMiddleware)

	c := generated.Config{Resolvers: &graph.Resolver{}}
	c.Directives.Auth = directives.Auth

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(c))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
