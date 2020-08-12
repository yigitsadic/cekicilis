package models

import uuid "github.com/satori/go.uuid"

type Participant struct {
	Id         string `json:"id"`
	Identifier string `json:"identifier"`
}

func NewParticipant() *Participant {
	return &Participant{
		Id:         uuid.NewV4().String(),
		Identifier: uuid.NewV4().String(),
	}
}
