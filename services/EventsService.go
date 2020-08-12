package services

import (
	"github.com/lib/pq"
	"github.com/yigitsadic/cekicilis/dtos"
	"github.com/yigitsadic/cekicilis/models"
	"github.com/yigitsadic/cekicilis/pgdb"
	"log"
	"math/rand"
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

func (service *EventsService) CalculateWinners(evtId string) []string {
	service.PgDb.Connect()
	defer service.PgDb.DB.Close()

	rows, err := service.PgDb.DB.Query("SELECT reference, userreference FROM participants WHERE eventid = $1", evtId)
	defer rows.Close()
	if err != nil {
		log.Printf("Unable to calculate winners for event ID %s, should be reevaluate. error is %s\n", evtId, err)
	}

	var participants []struct {
		reference,
		userreference string
	}
	var winners []string

	for rows.Next() {
		var reference, userreference string

		err := rows.Scan(&reference, &userreference)
		if err != nil {
			log.Println("Unable to parse rows, err: ", err)
		}

		participants = append(participants, struct{ reference, userreference string }{reference, userreference})
	}

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	for x := 1; x <= 2; x++ {
		winners = append(winners, participants[r1.Intn(len(participants)-1)].userreference)
	}

	_, err = service.PgDb.DB.Query("UPDATE events SET winners=$1, status=$2 WHERE id=$3", pq.StringArray(winners), models.EVENT_COMPLETED, evtId)
	if err != nil {
		log.Printf("Unable to update winners for %s, winners were %v, error was %s\n", evtId, winners, err)
	}

	return winners
}
