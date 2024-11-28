package utils

import (
	"net/http"
)

// You can add any utilities in here

const (
	Status500 = http.StatusInternalServerError
	Status422 = http.StatusUnprocessableEntity
	Status404 = http.StatusNotFound
	Status201 = http.StatusCreated
	Status200 = http.StatusOK
)
