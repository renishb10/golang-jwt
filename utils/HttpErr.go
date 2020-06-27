package utils

import (
	"encoding/json"
	"net/http"

	"github.com/renishb10/golang-jwt/models"
)

func RespondWithError(w http.ResponseWriter, status int, error models.Error) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(error)
}
