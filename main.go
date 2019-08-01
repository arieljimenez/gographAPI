package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/arieljimenez/gographAPI/gql"
	"github.com/arieljimenez/gographAPI/postgres"
	"github.com/arieljimenez/gographAPI/server"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/graphql-go/graphql"
)

func main() {
	router, db := initializeAPI()
	defer db.Close()

	// Listen on port 5000 and if there's an error log it and exit
	fmt.Println("Up and Running at port 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}

func initializeAPI() (*chi.Mux, *postgres.Db) {
	router := chi.NewRouter()

	// Create a new connection to our pg database
	db, err := postgres.New(
		postgres.ConnString("localhost", 5432, "postgres", "thisisnotagoodpassword", "postgres"), // host, port, user, password db
	)

	if err != nil {
		log.Fatal(err)
	}

	// Create our root query for graphql
	rootQuery := gql.NewRoot(db)
	// Create a new graphql schema, passing in the the root query
	sc, err := graphql.NewSchema(
		graphql.SchemaConfig{Query: rootQuery.Query},
	)

	if err != nil {
		fmt.Println("Error creating schema: ", err)
	}

	// Create a server struct that holds a pointer to our database as well
	// as the address of our graphql schema
	s := server.Server{
		GqlSchema: &sc,
	}

	// Add some middleware to our router
	router.Use(
		render.SetContentType(render.ContentTypeJSON), // set content-type headers as application/json
		middleware.Logger,                             // log api request calls
		middleware.DefaultCompress,                    // compress results, mostly gzipping assets and json
		middleware.StripSlashes,                       // match paths with a trailing slash, strip it, and continue routing through the mux
		middleware.Recoverer,                          // recover from panics without crashing server
	)

	// Create the graphql route with a Server method to handle it
	router.Post("/graphql", s.GraphQL())

	return router, db
}
