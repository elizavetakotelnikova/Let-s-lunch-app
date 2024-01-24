package query

import (
	"cmd/app/entities/meeting"
	"context"
	"database/sql"
	"fmt"
)

func Create(ctx context.Context, meeting *meeting.Meeting, db *sql.DB) error {
	const query = `INSERT INTO meetings(gathering_place_id, time_start, time_end, max_participants, initiators_id, state) +
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8::meeting_state)`
	_, err := db.Exec(query, meeting.GatheringPlaceId, meeting.StartTime, meeting.EndTime, meeting.InitiatorsId, meeting.State)
	if err != nil {
		return fmt.Errorf("database query execution error: %w", err)
	}
	return nil
}

//тут надо подумать как лучше обрабатывать ошибку
