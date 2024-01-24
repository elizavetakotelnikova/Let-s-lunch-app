package query

import (
	"context"
	"database/sql"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

func FindByCriteria(ctx context.Context, criteria FindCriteria, db *sql.DB) (*sql.Rows, error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	sqlStatement := psql.Select("*").From("meetings").RunWith(db)
	if criteria.GatheringPlaceId.Valid != false {
		sqlStatement = findByPlace(sqlStatement, criteria.GatheringPlaceId.UUID)
	}
	if criteria.InitiatorID.Valid != false {
		sqlStatement = findByInitiatorsId(sqlStatement, criteria.GatheringPlaceId.UUID)
	}
	if criteria.ID.Valid != false {
		sqlStatement = findByMeetingId(sqlStatement, criteria.GatheringPlaceId.UUID)
	}
	var rows, err = sqlStatement.Query()
	if err != nil {
		return nil, fmt.Errorf("problem with quering to database %w", err)
	}
	return rows, nil
}

func findByPlace(sql sq.SelectBuilder, id uuid.UUID) sq.SelectBuilder {
	return sql.Where(sq.Eq{"gathering_place_id": id})
}
func findByMeetingId(sql sq.SelectBuilder, id uuid.UUID) sq.SelectBuilder {
	return sql.Where(sq.Eq{"id": id})
}
func findByInitiatorsId(sql sq.SelectBuilder, id uuid.UUID) sq.SelectBuilder {
	return sql.Where(sq.Eq{"initiators_id": id})
}
