package query

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
)

func FindByID(ctx context.Context, id uuid.UUID, db *sql.DB) *sql.Row {
	const query = `SELECT * FROM meetings WHERE id = $1`
	row := db.QueryRow(query, id)
	return row
}
