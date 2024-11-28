package main

import (
	"log"
	"net/http"

	"github.com/bimbims125/belajar-docker-compose/database"
	"github.com/bimbims125/belajar-docker-compose/routers"
)

func main() {
	// Hubungkan ke database
	database.ConnectDatabase()

	// log.Println("Berhasil terhubung ke database")

	// Inisialisasi router
	router := routers.Routers()

	// Tentukan port server
	port := ":3300"
	log.Printf("Server running on http://localhost%s", port)

	// Jalankan server
	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatalf("Error when running server: %v", err)
	}
}
