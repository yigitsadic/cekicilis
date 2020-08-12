package handlers

import (
	"encoding/json"
	"github.com/yigitsadic/cekicilis/services"
	"github.com/yigitsadic/cekicilis/shared"
	"net/http"
)

type StatusResponse struct {
	Message    string   `json:"message"`
	Winners    []string `json:"winners,omitempty"`
	Status     int      `json:"status"`
	FinishesAt int64    `json:"finishes_at"`
}

func HandleStatus(service *services.EventsService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		// Try to fetch `eventId` from GET query params
		val := r.URL.Query().Get("eventId")
		if len(val) < 5 {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(shared.ErrorResponse{
				Message:   "eventId query parameter not found",
				ErrorCode: shared.ERR_QUERY_PARAM_NOT_FOUND,
			})

			return
		}

		// Get winners and status from the database.
		evt, err := service.GetEvent(val)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(shared.ErrorResponse{
				Message:   "Unable to fetch given event's status",
				ErrorCode: shared.ERR_UNABLE_TO_FETCH_EVENT,
			})
			return
		}

		// If event not found render not found response.
		if evt == nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(shared.ErrorResponse{
				Message:   "Event not found with given id",
				ErrorCode: shared.ERR_EVENT_NOT_FOUND,
			})
			return
		}

		// Serialize and respond.

		json.NewEncoder(w).Encode(&StatusResponse{
			Message:    "Successfully fetches status of requested event",
			Winners:    evt.Winners,
			Status:     evt.Status,
			FinishesAt: evt.FinishesAt.UTC().Unix(),
		})
		return
	}
}
