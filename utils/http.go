package utils

import (
	"encoding/json"
	"net/http"
)

func SendJSON(w http.ResponseWriter, data interface{}, status int) {
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
