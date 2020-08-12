package dtos

type JoinParticipantDto struct {
	Reference     string `json:"reference"`
	EventId       string `json:"eventId"`
	UserReference string `json:"userReference"`
}

func (j JoinParticipantDto) IsValid() bool {
	var val bool

	val = len(j.EventId) > 6
	val = len(j.UserReference) > 6
	val = len(j.Reference) > 6

	return val
}
