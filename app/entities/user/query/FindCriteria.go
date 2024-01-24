package query

import (
	"github.com/google/uuid"
)

type FindCriteria struct {
	ID               uuid.NullUUID
	Username         *string
	DisplayName      *string
	CurrentMeetingId uuid.NullUUID
}
