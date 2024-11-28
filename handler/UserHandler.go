package handler

import (
	"encoding/json"
	"net/http"

	"github.com/bimbims125/belajar-docker-compose/models"
)

func GetAllUserHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.SuccessResponseMessage{
		Message: "Success",
	})
	return
}
