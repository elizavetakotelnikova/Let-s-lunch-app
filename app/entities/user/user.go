package user

import (
	"github.com/gofrs/uuid/v5"
)

type User struct {
	ID               uuid.UUID
	Username         string
	DisplayName      string
	CurrentMeetingId uuid.NullUUID
	MeetingHistory   []uuid.UUID
	Rating           int
	Age              int
	Gender           Gender
}

func NewUser() *User {
	id, err := uuid.NewV4()
	if err != nil {
		return nil
	}

	return &User{ID: id}
}
