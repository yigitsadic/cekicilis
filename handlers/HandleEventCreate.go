package handlers

import (
	"log"
	"net/http"
	"time"
)

func HandleEventCreate(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now().String()

	// 1. Insert to database

	// 2. Enqueue background job.
	go func(givenTime string) {
		time.AfterFunc(time.Second*10, func() {
			// Calculate winner.
			log.Printf("Calculated winners for %s\n", currentTime)
		})
	}(currentTime)

	w.WriteHeader(http.StatusNoContent)
}
