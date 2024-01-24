package query

import (
	"cmd/app/entities/meeting"
	"context"
	"database/sql"
	"fmt"
)

func Update(ctx context.Context, meeting *meeting.Meeting, db *sql.DB) error {
	const query = `UPDATE meetings SET state = $1::meeting_state WHERE id = $2`
	_, err := db.Exec(query, meeting.State, meeting.ID)
	if err != nil {
		return fmt.Errorf("database query execution error: %w", err)
	}
	return nil
}

//пока не знаю, что будем апдейтить кроме состояния, оставлю так
