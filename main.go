package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func main() {
	// Fetch expires within 1 day records.
	// Calculate if winners list is empty.

	http.HandleFunc("/events", func(w http.ResponseWriter, r *http.Request) {
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
	})

	http.HandleFunc("/join", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(map[string]string{
			"status": "success",
		})

		return
	})

	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":     "ongoing",
			"finishesOn": time.Now().Add(time.Minute * 15).String(),
		})
		return
	})

	log.Print("Server is up and running on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Printf("Unable to continue cause of %s", err)
	}
}
