package meeting

import "github.com/google/uuid"

type Meeting struct {
	id uuid.UUID
}

func NewMeeting() (*Meeting, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	return &Meeting{id: id}, err
}
