package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/bimbims125/belajar-docker-compose/database"
	"github.com/bimbims125/belajar-docker-compose/models"
	"github.com/bimbims125/belajar-docker-compose/utils"
)

func GetAllUserHandler(w http.ResponseWriter, r *http.Request) {
	db := database.DB
	var users []models.User

	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Fatalf("Cannot query : %v", err)
	}

	for rows.Next() {
		var user models.User
		rows.Scan(&user.Id, &user.Name, &user.Email)
		users = append(users, user)
	}
	response := models.SuccessResponse{
		Data: users,
	}
	w.WriteHeader(utils.Status200)
	json.NewEncoder(w).Encode(response)
}
