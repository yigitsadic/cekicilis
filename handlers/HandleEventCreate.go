package handlers

import (
	"encoding/json"
	uuid "github.com/satori/go.uuid"
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

type FailureResponse struct {
	Message   string `json:"message"`
	ErrorCode int    `json:"errorCode"`
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

			json.NewEncoder(w).Encode(&FailureResponse{
				Message:   "Create event params are not valid",
				ErrorCode: shared.ERR_EVENT_PARAMS_NOT_VALID,
			})

			return
		}

		// 3. Insert to database
		var insertedEvent models.Event
		insertedEvent.Name = dto.Name
		insertedEvent.Id = uuid.NewV4().String()
		insertedEvent.FinishesAt = time.Unix(dto.FinishesAt, 0)

		// 4. Enqueue background job.
		go func(eventsService *services.EventsService) {
			time.AfterFunc(time.Second*10, func() {
				// Calculate winner.
				winners := eventsService.CalculateWinners()

				for i, winner := range winners {
					log.Printf("#%d %s\n", i, winner)
				}
			})
		}(eventsService)

		// 5. Serialize response and respond.
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(&SuccessfulResponse{
			Message: "Successfully created an event with given params",
			Event:   &insertedEvent,
		})
		return
	}
}
