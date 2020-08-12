package handlers

import (
	"encoding/json"
	"github.com/yigitsadic/cekicilis/services"
	"github.com/yigitsadic/cekicilis/shared"
	"net/http"
)

func HandleEventList(eventsService *services.EventsService) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		// Fetch events.
		events, err := eventsService.GetEvents()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(&shared.ErrorResponse{
				Message:   "Unable to fetch events",
				ErrorCode: shared.ERR_CANNOT_FETCH_EVENTS,
			})

			return
		}

		json.NewEncoder(w).Encode(events)

		return
	}
}
