package services

import (
	"github.com/yigitsadic/cekicilis/models"
	"github.com/yigitsadic/cekicilis/pgdb"
	"time"
)

type EventsService struct {
	participantService *ParticipantService
	PgDb               *pgdb.Postgres
}

func NewEventsService(pgDb *pgdb.Postgres) *EventsService {
	return &EventsService{PgDb: pgDb}
}

func (service *EventsService) GetEvents() ([]*models.Event, error) {
	service.PgDb.Connect()
	defer service.PgDb.DB.Close()

	rows, err := service.PgDb.DB.Query("SELECT id, name, status, expiresat FROM events")
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	var events []*models.Event
	for rows.Next() {
		var id, name string
		var status int
		var expiresAt int64
		rows.Scan(&id, &name, &status, &expiresAt)

		events = append(events, &models.Event{
			Id:           id,
			Name:         name,
			Status:       status,
			FinishesAt:   time.Unix(expiresAt, 0),
			Participants: nil,
		})
	}

	return events, nil
}

func (service *EventsService) CreateEvent() (*models.Event, error) {
	return nil, nil
}

func (service *EventsService) CalculateWinners() []string {
	return []string{
		"Yigit",
		"Sertan",
		"Yağmur",
		"Fazilet",
		"Gülşen",
	}
}
