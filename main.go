package main

import (
	_ "github.com/lib/pq"
	"github.com/yigitsadic/cekicilis/pgdb"

	"github.com/yigitsadic/cekicilis/handlers"
	"github.com/yigitsadic/cekicilis/services"
	"log"
	"net/http"
	"os"
	"sync"
)

var doOnce sync.Once

func main() {
	pgDB := pgdb.NewPostgres()

	eventsService := services.NewEventsService(pgDB)
	participantsService := services.NewParticipantService(pgDB)

	// Fetch expires within 1 day records.
	// Calculate if winners list is empty.
	doOnce.Do(processInitialData)

	// Creates event and enqueues to background job.
	http.HandleFunc("/create-event", handlers.HandleEventCreate(eventsService))

	// List events
	http.HandleFunc("/event-list", handlers.HandleEventList(eventsService))

	// Joins user to an event.
	http.HandleFunc("/join", handlers.HandleJoin(participantsService))

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

func processInitialData() {
	log.Println("No rows to process.")
}
