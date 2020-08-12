package models

import "time"

type Event struct {
	Id           string        `json:"id"`
	Participants []Participant `json:"participants,omitempty"`
	Name         string        `json:"name"`
	FinishesAt   time.Time     `json:"finishesAt"`
}
