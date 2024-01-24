package query

import (
	"cmd/app/entities/gatheringPlace"
	"context"
	"database/sql"
	"fmt"
)

func Update(ctx context.Context, place *gatheringPlace.GatheringPlace, db *sql.DB) error {
	const query = `UPDATE gatheringPlaces SET rating = $1 WHERE id = $2`
	_, err := db.Exec(query, place.Rating, place.ID)
	if err != nil {
		return fmt.Errorf("database query execution error: %w", err)
	}
	return nil
}

//пока не знаю, что будем апдейтить кроме рейтинга. Можно апдейтить через Delete и Create (Save), но это кринж
