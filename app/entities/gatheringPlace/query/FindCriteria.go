package query

import (
	"cmd/app/models"
	"database/sql"
	"github.com/google/uuid"
)

type FindCriteria struct {
	Address     *models.Address
	InitiatorID uuid.NullUUID
	CuisineType sql.NullInt16
	Rating      int
}
