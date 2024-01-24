package query

import (
	"context"
	"database/sql"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

func FindUserByCriteria(ctx context.Context, criteria FindCriteria, db *sql.DB) (*sql.Rows, error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	sqlStatement := psql.Select("*").From("users").RunWith(db)
	if criteria.ID.Valid != false {
		sqlStatement = findUserById(sqlStatement, criteria.ID.UUID)
	}
	if criteria.Username != nil {
		sqlStatement = findUserByUsername(sqlStatement, criteria.Username)
	}
	if criteria.DisplayName != nil {
		sqlStatement = findUserByDisplayname(sqlStatement, criteria.DisplayName)
	}
	if criteria.CurrentMeetingId.Valid != false {
		sqlStatement = findUserByCurrentMeetingId(sqlStatement, criteria.ID.UUID)
	}
	var rows, err = sqlStatement.Query()
	if err != nil {
		return nil, fmt.Errorf("problem with quering to database %w", err)
	}
	return rows, nil
}

func findUserById(sql sq.SelectBuilder, id uuid.UUID) sq.SelectBuilder {
	return sql.Where(sq.Eq{"id": id})
}
func findUserByCurrentMeetingId(sql sq.SelectBuilder, id uuid.UUID) sq.SelectBuilder {
	return sql.Where(sq.Eq{"current_meeting_id": id})
}
func findUserByUsername(sql sq.SelectBuilder, username *string) sq.SelectBuilder {
	return sql.Where(sq.Eq{"username": username})
}
func findUserByDisplayname(sql sq.SelectBuilder, displayname *string) sq.SelectBuilder {
	return sql.Where(sq.Eq{"display_name": displayname})
}

func FindUserHistoryById(ctx context.Context, id uuid.UUID, db *sql.DB) (*sql.Rows, error) {
	const query = "SELECT meeting_id FROM meetings_history WHERE id = $1"
	rows, err := db.Query(query, id)
	if err != nil {
		return nil, fmt.Errorf("database query execution error: %w", err)
	}
	return rows, nil
}
