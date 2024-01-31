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

func NewUser(username string, displayname string, birthday time.Time, phoneNumber string, gender Gender) *User {
	id, err := uuid.NewV4()
	if err != nil {
		return nil
	}

	return &User{
		ID:          id,
		Username:    username,
		DisplayName: displayname,
		Birthday:    birthday,
		PhoneNumber: phoneNumber,
		Gender:      gender}
}
