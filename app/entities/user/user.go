package user

import (
	"github.com/gofrs/uuid/v5"
	"time"
)

type User struct {
	ID               uuid.UUID
	Username         string
	DisplayName      string
	CurrentMeetingId uuid.NullUUID
	MeetingHistory   []uuid.UUID
	Rating           int
	Birthday         time.Time
	Gender           Gender
	PhoneNumber      string
}

func NewUser() *User {
	id, err := uuid.NewV4()
	if err != nil {
		return nil
	}
	date, err := time.Parse(time.DateOnly, "1-1-1")
	date.Round(24 * time.Hour)
	return &User{ID: id, Birthday: date}
}
