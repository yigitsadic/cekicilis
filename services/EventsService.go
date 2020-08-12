package services

import (
	"github.com/lib/pq"
	"github.com/yigitsadic/cekicilis/dtos"
	"github.com/yigitsadic/cekicilis/models"
	"github.com/yigitsadic/cekicilis/pgdb"
	"log"
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

	rows, err := service.PgDb.DB.Query("SELECT id, name, status, finishesat FROM events")
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

func (service *EventsService) GetEvent(evtId string) (*models.Event, error) {
	service.PgDb.Connect()
	defer service.PgDb.DB.Close()

	var name string
	var status int
	var finishesAt time.Time
	var winners pq.StringArray
	err := service.PgDb.DB.QueryRow("SELECT name, status, finishesat, winners FROM events WHERE id = $1 LIMIT 1", evtId).Scan(&name, &status, &finishesAt, &winners)
	if err != nil {
		log.Println("Unable to fetch given event with ID", err)

		return nil, err
	}

	evt := &models.Event{
		Id:         evtId,
		Name:       name,
		Status:     status,
		FinishesAt: finishesAt,
		Winners:    winners,
	}

	return evt, nil
}

func (service *EventsService) CreateEvent(dto *dtos.CreateEventDto) (*models.Event, error) {
	service.PgDb.Connect()
	defer service.PgDb.DB.Close()

	var id string
	err := service.PgDb.DB.QueryRow("INSERT INTO events (name, finishesat) VALUES ($1, $2) RETURNING id", dto.Name, time.Unix(dto.FinishesAt, 0)).Scan(&id)
	if err != nil {
		return nil, err
	}

	event := &models.Event{
		Id:           id,
		Name:         dto.Name,
		Status:       models.EVENT_ONGOING,
		FinishesAt:   time.Unix(dto.FinishesAt, 0),
		Participants: nil,
	}

	return event, nil
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
