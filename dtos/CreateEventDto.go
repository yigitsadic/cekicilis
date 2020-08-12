package dtos

type CreateEventDto struct {
	Name       string `json:"name"`
	FinishesAt int64  `json:"finishesAt"`
}

func (c CreateEventDto) IsValid() bool {
	var val bool

	if len(c.Name) > 3 {
		val = true
	}

	return val
}
