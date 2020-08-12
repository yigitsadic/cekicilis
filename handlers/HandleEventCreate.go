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

		log.Println(time.Now().UTC().Unix())
		log.Println(dto)

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
			Event: &models.Event{
				Id:           "8cc67890-fe18-448d-af94-a583c3cfd951",
				Participants: nil,
				Name:         "eerere",
				FinishesAt:   time.Now().Add(time.Minute * 15),
			},
		})
		return
	}
}
