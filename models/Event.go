package models

import "time"

type Event struct {
	Participants []Participant
	Name         string
	FinishesAt   time.Time
}
