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
	HashedPassword   []byte
}

func NewUser(username string, displayName string, birthday time.Time, phoneNumber string, gender Gender, hashedPassword []byte) *User {
	id, err := uuid.NewV4()
	if err != nil {
		return nil
	}

	return &User{
		ID:             id,
		Username:       username,
		DisplayName:    displayName,
		Birthday:       birthday,
		PhoneNumber:    phoneNumber,
		Gender:         gender,
		HashedPassword: hashedPassword,
	}
}
