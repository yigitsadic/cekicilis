package services

import (
	"github.com/yigitsadic/cekicilis/models"
	"time"
)

type EventsService struct {
}

func NewEventsService() *EventsService {
	return &EventsService{}
}

func (service *EventsService) GetEvents() ([]models.Event, error) {
	return []models.Event{{Participants: []models.Participant{*models.NewParticipant()}, Name: "Lorem ipsum", FinishesAt: time.Now().Add(time.Minute * 3)}}, nil
}

func (service *EventsService) CreateEvent() (*models.Event, error) {
	return nil, nil
}
