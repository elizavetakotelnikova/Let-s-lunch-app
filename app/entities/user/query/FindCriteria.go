package query

import (
	"github.com/google/uuid"
)

type FindCriteria struct {
	Username         *string
	DisplayName      *string
	CurrentMeetingId uuid.NullUUID
}
