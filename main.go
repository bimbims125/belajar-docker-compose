package main

import (
	"log"
	"net/http"

	"github.com/bimbims125/belajar-docker-compose/database"
	"github.com/bimbims125/belajar-docker-compose/routers"
)

func main() {
	// connect to database
	database.ConnectDatabase()

	// Routers
	router := routers.Routers()

	// define ports
	port := ":3300"
	log.Printf("Server running on http://localhost%s", port)

	// Run server
	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatalf("Error when running server: %v", err)
	}
}
