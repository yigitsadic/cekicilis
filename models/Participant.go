package models

import uuid "github.com/satori/go.uuid"

type Participant struct {
	Identifier string
}

func NewParticipant() *Participant {
	return &Participant{
		Identifier: uuid.NewV4().String(),
	}
}
