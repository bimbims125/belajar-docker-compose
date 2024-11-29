package utils

import (
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

// You can add any utilities in here

const (
	Status500 = http.StatusInternalServerError
	Status422 = http.StatusUnprocessableEntity
	Status404 = http.StatusNotFound
	Status201 = http.StatusCreated
	Status200 = http.StatusOK
)

func HashedPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Failed to hashing password : %v", err)
	}
	return string(hashedPassword), nil
}
