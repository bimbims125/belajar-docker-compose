package routers

import (
	"github.com/bimbims125/belajar-docker-compose/handler"
	"github.com/gorilla/mux"
)

// Routers
func Routers() *mux.Router {
	routers := mux.NewRouter()

	// GET
	routers.HandleFunc("/posts", handler.GetPostHandler).Methods("GET")
	routers.HandleFunc("/posts/{id}", handler.GetPostByIdHandler).Methods("GET")

	// POST
	routers.HandleFunc("/posts", handler.CreatePostHandler).Methods("POST")
	return routers
}
