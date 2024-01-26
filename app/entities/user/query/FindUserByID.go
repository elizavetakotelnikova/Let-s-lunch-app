package query

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
)

func FindUserByID(ctx context.Context, id uuid.UUID, db *sql.DB) *sql.Row {
	const query = `SELECT FROM users WHERE id = $1`
	row := db.QueryRow(query, id)
	return row
}
