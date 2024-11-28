package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/bimbims125/belajar-docker-compose/database"
	"github.com/bimbims125/belajar-docker-compose/models"
	"github.com/gorilla/mux"
)

const (
	status500 = http.StatusInternalServerError
	status404 = http.StatusNotFound
	status201 = http.StatusCreated
	status200 = http.StatusOK
	status422 = http.StatusUnprocessableEntity
)

type SuccessResponseMessage struct {
	Message string `json:"message"`
}

type ErrorResponseMessage struct {
	Message string `json:"message"`
}

type SuccessResponse struct {
	Data interface{} `json:"data"`
}

func GetPostHandler(w http.ResponseWriter, r *http.Request) {
	db := database.DB

	rows, err := db.Query("SELECT * FROM posts")
	if err != nil {
		log.Fatalf("Can't query : %v", err)
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		rows.Scan(&post.Id, &post.Title, &post.Content)
		posts = append(posts, post)
	}
	response := SuccessResponse{
		Data: posts,
	}

	w.WriteHeader(status200)
	json.NewEncoder(w).Encode(response)
}

func GetPostByIdHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	db := database.DB

	var posts []models.Post

	rows, err := db.Query("SELECT * FROM posts WHERE id=?", id)
	if err != nil {
		log.Fatalf("Cannot query : %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var post models.Post
		rows.Scan(&post.Id, &post.Title, &post.Content)
		posts = append(posts, post)
	}

	if len(posts) == 0 {
		w.WriteHeader(status404)
		json.NewEncoder(w).Encode(ErrorResponseMessage{
			Message: "Post not found",
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status200)
	json.NewEncoder(w).Encode(SuccessResponse{
		Data: posts[0],
	})
	return
}

func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	db := database.DB

	var post models.Post

	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		log.Fatalf("Can't decode data : %v", err)
	}

	if post.Title == "" || post.Content == "" {
		w.WriteHeader(status422)
		json.NewEncoder(w).Encode(ErrorResponseMessage{
			Message: "can't create data",
		})
		return
	}
	query, err := db.Exec("INSERT INTO posts (title, content) VALUE (?, ?)", post.Title, post.Content)
	if err != nil {
		log.Fatalf("Can't insert data : %v", err)
	}

	query.LastInsertId()
	w.WriteHeader(status201)
	json.NewEncoder(w).Encode(SuccessResponseMessage{
		Message: "Create post success",
	})
}
