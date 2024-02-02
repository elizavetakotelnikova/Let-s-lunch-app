package user

import (
	"github.com/gofrs/uuid/v5"
	"time"
)

type User struct {
	ID               uuid.UUID     `json:"id"`
	Username         string        `json:"username"`
	DisplayName      string        `json:"displayName"`
	CurrentMeetingId uuid.NullUUID `json:"currentMeetingId"`
	MeetingHistory   []uuid.UUID   `json:"meetingHistory"`
	Rating           int           `json:"rating"`
	Birthday         time.Time     `json:"birthday"`
	Gender           Gender        `json:"gender"`
	PhoneNumber      string        `json:"phoneNumber"`
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
