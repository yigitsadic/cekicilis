package handlers

import (
	"encoding/json"
	"github.com/yigitsadic/cekicilis/dtos"
	"github.com/yigitsadic/cekicilis/models"
	"github.com/yigitsadic/cekicilis/services"
	"github.com/yigitsadic/cekicilis/shared"
	"log"
	"net/http"
	"time"
)

type SuccessfulResponse struct {
	Message string        `json:"message"`
	Event   *models.Event `json:"event"`
}

func HandleEventCreate(eventsService *services.EventsService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		// 1. Decode json from body
		var dto dtos.CreateEventDto
		json.NewDecoder(r.Body).Decode(&dto)

		// 2. Validate input
		if !dto.IsValid() {
			w.WriteHeader(http.StatusUnprocessableEntity)

			json.NewEncoder(w).Encode(&shared.ErrorResponse{
				Message:   "Create event params are not valid",
				ErrorCode: shared.ERR_EVENT_PARAMS_NOT_VALID,
			})

			return
		}

		// 3. Insert to database
		evt, err := eventsService.CreateEvent(&dto)
		if err != nil {
			log.Println("Error occurred during inserting new event to the db", err)

			w.WriteHeader(http.StatusUnprocessableEntity)
			json.NewEncoder(w).Encode(&shared.ErrorResponse{
				Message:   "Unable to insert event",
				ErrorCode: shared.ERR_UNABLE_TO_INSERT_EVENT,
			})

			return
		}

		// 4. Enqueue background job.
		go func(eventsService *services.EventsService, event *models.Event) {
			time.AfterFunc(time.Second*30, func() {
				// Calculate winner.
				eventsService.CalculateWinners(evt.Id)
			})
		}(eventsService, evt)

		// 5. Serialize response and respond.
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(&SuccessfulResponse{
			Message: "Successfully created an event with given params",
			Event:   evt,
		})
		return
	}
}
