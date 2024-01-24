package query

import (
	"cmd/app/entities/gatheringPlace"
	"cmd/app/models"
	"github.com/google/uuid"
)

type FindCriteria struct {
	ID          uuid.NullUUID
	Address     *models.Address
	InitiatorID uuid.NullUUID
	CuisineType *gatheringPlace.CuisineType
	Rating      int
}
