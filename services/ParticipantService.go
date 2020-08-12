package services

import (
	"github.com/yigitsadic/cekicilis/dtos"
	"github.com/yigitsadic/cekicilis/models"
	"github.com/yigitsadic/cekicilis/pgdb"
	"log"
)

type ParticipantService struct {
	PgDb *pgdb.Postgres
}

func NewParticipantService(pgDb *pgdb.Postgres) *ParticipantService {
	return &ParticipantService{PgDb: pgDb}
}

func (service *ParticipantService) GetParticipants() ([]*models.Participant, error) {
	return nil, nil
}

func (service *ParticipantService) CreateParticipant(dto *dtos.JoinParticipantDto) (*models.Participant, error) {
	service.PgDb.Connect()
	defer service.PgDb.DB.Close()

	var id string
	err := service.PgDb.DB.QueryRow("INSERT INTO participants (reference, userreference, eventid) VALUES ($1, $2, $3) RETURNING id", dto.Reference, dto.UserReference, dto.EventId).Scan(&id)
	if err != nil {
		log.Println("Unable to insert given participant: ", dto, err)
		return nil, err
	}

	participant := models.NewParticipant(id, dto.UserReference)

	return participant, nil
}
