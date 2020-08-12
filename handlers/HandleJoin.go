package handlers

import (
	"encoding/json"
	"net/http"
)

func HandleJoin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]string{
		"status": "success",
	})

	return
}