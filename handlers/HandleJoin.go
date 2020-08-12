package handlers

import (
	"encoding/json"
	"github.com/yigitsadic/cekicilis/dtos"
	"github.com/yigitsadic/cekicilis/services"
	"github.com/yigitsadic/cekicilis/shared"
	"net/http"
)

func HandleJoin(service *services.ParticipantService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		// 1. Read data from request
		var dto dtos.JoinParticipantDto
		json.NewDecoder(r.Body).Decode(&dto)
		defer r.Body.Close()

		// 2. Validate dto

		if !dto.IsValid() {
			w.WriteHeader(http.StatusUnprocessableEntity)
			json.NewEncoder(w).Encode(shared.ErrorResponse{
				Message:   "Unable to create with this params",
				ErrorCode: shared.ERR_JOIN_PARAMS_NOT_VALID,
			})

			return
		}

		// 3. Insert to the database
		inserted, err := service.CreateParticipant(&dto)
		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			json.NewEncoder(w).Encode(shared.ErrorResponse{
				Message:   "Unable to create with this params",
				ErrorCode: shared.ERR_JOIN_PARAMS_NOT_VALID,
			})

			return
		}

		// 4. Serialize response and respond.

		json.NewEncoder(w).Encode(map[string]string{
			"message":    "Successfully joined given user with given references to the event",
			"insertedId": inserted.Id,
			"identifier": inserted.Identifier,
		})

		return
	}
}
