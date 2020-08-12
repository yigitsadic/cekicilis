package models

import "time"

const (
	EVENT_ONGOING int = iota
	EVENT_COMPLETED
)

type Event struct {
	Id           string        `json:"id"`
	Name         string        `json:"name"`
	Status       int           `json:"status"`
	FinishesAt   time.Time     `json:"finishesAt"`
	Participants []Participant `json:"participants,omitempty"`
}
