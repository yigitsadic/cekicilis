package models

type Participant struct {
	Id         string `json:"id"`
	Identifier string `json:"identifier"`
}

func NewParticipant(id, identifier string) *Participant {
	return &Participant{
		Id:         id,
		Identifier: identifier,
	}
}
