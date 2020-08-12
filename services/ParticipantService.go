package services

import "github.com/yigitsadic/cekicilis/models"

type ParticipantService struct {
}

func NewParticipantService() *ParticipantService {
	return &ParticipantService{}
}

func (service *ParticipantService) GetParticipants() ([]models.Participant, error) {
	return []models.Participant{*models.NewParticipant()}, nil
}

func (service *ParticipantService) CreateParticipant() (*models.Participant, error) {
	return nil, nil
}
