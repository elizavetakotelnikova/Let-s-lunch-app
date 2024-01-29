package user

import (
	"github.com/google/uuid"
)

type User struct {
	ID               uuid.UUID
	Username         string
	DisplayName      string
	CurrentMeetingId uuid.NullUUID
	MeetingHistory   []uuid.UUID
	Rating           int
}

func NewUser() *User {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil
	}

	return &User{ID: id}
}
