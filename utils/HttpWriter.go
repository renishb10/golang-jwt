package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func ResponseJSON(w *http.ResponseWriter, data interface{}) {
	(*w).Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(*w).Encode(data)
	if err != nil {
		log.Fatal(err)
	}
}
