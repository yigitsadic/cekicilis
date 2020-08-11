package main

import (
	"github.com/yigitsadic/cekicilis/handlers"
	"log"
	"net/http"
	"os"
)

func main() {
	// Fetch expires within 1 day records.
	// Calculate if winners list is empty.

	// Creates event and enqueues to background job.
	http.HandleFunc("/createEvent", handlers.HandleEventCreate)

	// List events
	http.HandleFunc("/eventList", handlers.HandleEventList)

	// Joins user to an event.
	http.HandleFunc("/join", handlers.HandleJoin)

	// Displays status and winners if present.
	http.HandleFunc("/status", handlers.HandleStatus)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server is up and running on port %s\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Printf("Unable to continue cause of %s", err)
	}
}
