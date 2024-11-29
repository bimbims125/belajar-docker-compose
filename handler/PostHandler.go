package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/bimbims125/belajar-docker-compose/database"
	"github.com/bimbims125/belajar-docker-compose/models"
	"github.com/bimbims125/belajar-docker-compose/utils"
	"github.com/gorilla/mux"
)

func GetAllPostHandler(w http.ResponseWriter, r *http.Request) {
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
	response := models.SuccessResponse{
		Data: posts,
	}

	w.WriteHeader(utils.Status200)
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
		w.WriteHeader(utils.Status404)
		json.NewEncoder(w).Encode(models.ErrorResponseMessage{
			Message: "Post not found",
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(utils.Status200)
	json.NewEncoder(w).Encode(models.SuccessResponse{
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
		w.WriteHeader(utils.Status422)
		json.NewEncoder(w).Encode(models.ErrorResponseMessage{
			Message: "can't create data",
		})
		return
	}
	query, err := db.Exec("INSERT INTO posts (title, content) VALUE (?, ?)", post.Title, post.Content)
	if err != nil {
		log.Fatalf("Can't insert data : %v", err)
	}

	query.LastInsertId()
	w.WriteHeader(utils.Status201)
	json.NewEncoder(w).Encode(models.SuccessResponseMessage{
		Message: "Create post success",
	})
}

func UpdatePostHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	db := database.DB

	var post models.Post

	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		log.Fatalf("Can't decode data : %v", err)
	}

	query, err := db.Exec("UPDATE posts SET title=?, content=? WHERE id=?", post.Title, post.Content, id)
	if err != nil {
		log.Fatalf("Cannot query : %v", err)
	}

	rowsAffected, err := query.RowsAffected()
	if err != nil {
		log.Fatalf("Cannot update data : %v", err)
		return
	}

	if rowsAffected == 0 {
		// log.Fatalf("Data not found")
		w.WriteHeader(utils.Status404)
		json.NewEncoder(w).Encode(models.ErrorResponseMessage{
			Message: "Not found",
		})
		return
	}

	w.WriteHeader(utils.Status200)
	json.NewEncoder(w).Encode(models.SuccessResponseMessage{
		Message: "Update data success",
	})
}
