package main

import (
	"log"
	"net/http"
	"os"

	"github.com/isaquecsilva/graphql/database"
	"github.com/isaquecsilva/graphql/graphql"
	"github.com/isaquecsilva/graphql/models"
	carservice "github.com/isaquecsilva/graphql/services/car"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("failure to parse environment variables: %v", err)
	}

	db, err := database.Connect(os.Getenv("DATABASE_URI"))
	if err != nil {
		log.Fatalf("failure to connect to database: %v", err)
	}
	defer db.Close()

	querier := models.New(db)

	cars := carservice.NewCarServiceImpl(querier)
	graphQLHandler, err := graphql.NewGraphQLHandler(cars)
	if err != nil {
		log.Printf("Error creating graphQL http handler: %v", err)
		return
	}

	mux := http.NewServeMux()
	mux.Handle("/graphql", graphQLHandler)

	const address string = "127.0.0.1:5000"
	println("Server running at " + address)
	log.Fatal(http.ListenAndServe(address, mux))
}
