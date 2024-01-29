package query

import (
	"database/sql"
	"github.com/google/uuid"
)

type FindCriteria struct {
	Username         sql.NullString
	DisplayName      sql.NullString
	CurrentMeetingId uuid.NullUUID
	Age              sql.NullInt32
	Gender           sql.NullInt16
}
