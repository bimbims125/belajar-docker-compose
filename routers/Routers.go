package routers

import (
	"github.com/bimbims125/belajar-docker-compose/handler"
	"github.com/gorilla/mux"
)

// PostsRouters for handle routing related posts
func PostsRouters(router *mux.Router) {
	// GET
	router.HandleFunc("/posts", handler.GetAllPostHandler).Methods("GET")
	router.HandleFunc("/posts/{id}", handler.GetPostByIdHandler).Methods("GET")

	// POST
	router.HandleFunc("/posts", handler.CreatePostHandler).Methods("POST")

	// PUT
	router.HandleFunc("/posts/{id}", handler.UpdatePostHandler).Methods("PUT")
}

// UsersRouters for handling routing related users
func UsersRouters(router *mux.Router) {
	// GET
	router.HandleFunc("/users", handler.GetAllUserHandler)
}

// Routers combine all routers
func Routers() *mux.Router {
	// Main utama
	routers := mux.NewRouter()

	// Routing for posts and users
	PostsRouters(routers)
	UsersRouters(routers)

	return routers
}
