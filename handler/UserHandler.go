package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/bimbims125/belajar-docker-compose/database"
	"github.com/bimbims125/belajar-docker-compose/models"
	"github.com/bimbims125/belajar-docker-compose/utils"
	"github.com/go-sql-driver/mysql"
)

func GetAllUserHandler(w http.ResponseWriter, r *http.Request) {
	db := database.DB

	rows, err := db.Query("SELECT id, name, username, email FROM users")
	if err != nil {
		log.Fatalf("Cannot query : %v", err)
	}

	defer rows.Close()

	var users []models.GetAllUser
	for rows.Next() {
		var user models.GetAllUser
		rows.Scan(&user.Id, &user.Name, &user.Username, &user.Email)
		users = append(users, user)
	}

	response := models.SuccessResponse{
		Data: users,
	}
	w.WriteHeader(utils.Status200)
	json.NewEncoder(w).Encode(response)
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	db := database.DB

	var user models.User

	// Decode request body
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // 400 Bad Request
		json.NewEncoder(w).Encode(models.ErrorResponseMessage{
			Message: "Invalid request payload",
		})
		return
	}

	// Hash the password
	hashedPassword, err := utils.HashedPassword(user.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError) // 500 Internal Server Error
		json.NewEncoder(w).Encode(models.ErrorResponseMessage{
			Message: "Failed to hash password",
		})
		return
	}

	// Begin transaction
	tx, err := db.Begin()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError) // 500 Internal Server Error
		json.NewEncoder(w).Encode(models.ErrorResponseMessage{
			Message: "Failed to start transaction",
		})
		return
	}

	// Insert user into the database
	_, err = tx.Exec("INSERT INTO users (name, username, email, password) VALUES (?, ?, ?, ?)", user.Name, user.Username, user.Email, hashedPassword)
	if err != nil {
		// Check for duplicate entry error
		if mySqlErrMsg, ok := err.(*mysql.MySQLError); ok && mySqlErrMsg.Number == 1062 {
			tx.Rollback()
			w.WriteHeader(utils.Status422) // 422 Unprocessable Entity
			json.NewEncoder(w).Encode(models.ErrorResponseMessage{
				Message: "Username or email already exists",
			})
			return
		}
		// Handle other errors
		tx.Rollback()
		w.WriteHeader(http.StatusInternalServerError) // 500 Internal Server Error
		json.NewEncoder(w).Encode(models.ErrorResponseMessage{
			Message: "Failed to execute query",
		})
		return
	}

	// Commit transaction
	err = tx.Commit()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError) // 500 Internal Server Error
		json.NewEncoder(w).Encode(models.ErrorResponseMessage{
			Message: "Failed to commit transaction",
		})
		return
	}

	// Success response
	w.WriteHeader(utils.Status201) // 201 Created
	json.NewEncoder(w).Encode(models.SuccessResponseMessage{
		Message: "User created successfully",
	})
}
