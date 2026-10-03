package main

import (
	"fmt"
	"krushi-server/internal/handler"
	"krushi-server/internal/repository"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	// MongoDB Configuration
	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		uri = "mongodb://localhost:27017"
	}
	dbName := "krushi_db"
	collName := "farmers"

	// Setup MongoDB repository
	repo, err := repository.NewMongoFarmerRepository(uri, dbName, collName)
	if err != nil {
		log.Fatalf("Could not connect to MongoDB: %v", err)
	}
	defer repo.Close()

	// Setup handler
	farmerHandler := handler.NewFarmerHandler(repo)

	// Setup router
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/farmers", func(r chi.Router) {
		r.Get("/", farmerHandler.ListFarmers)
		r.Post("/", farmerHandler.CreateFarmer)
		r.Get("/search", farmerHandler.SearchFarmerByPhone)
		r.Get("/{farmerId}", farmerHandler.GetFarmer)
	})

	port := ":8080"
	fmt.Printf("Krushi Server (MongoDB) starting on %s...\n", port)
	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatal(err)
	}
}
