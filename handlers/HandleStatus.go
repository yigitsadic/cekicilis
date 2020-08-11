package handlers

import (
	"encoding/json"
	"net/http"
	"time"
)

func HandleStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status":     "ongoing",
		"finishesOn": time.Now().Add(time.Minute * 15).String(),
	})
	return
}
