package query

import (
	"cmd/app/entities/user"
	"context"
	"database/sql"
	"fmt"
)

func Update(ctx context.Context, user *user.User, db *sql.DB) error {
	const query = `UPDATE meetings SET current_meeting_id = $1 WHERE id = $2`
	_, err := db.Exec(query, user.CurrentMeetingId, user.ID)
	if err != nil {
		return fmt.Errorf("database query execution error: %w", err)
	}
	return nil
}

//пока не знаю, что будем апдейтить, так что написала апдейт currentMeetingId
