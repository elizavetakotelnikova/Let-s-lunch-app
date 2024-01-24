package gatheringPlace

import (
	"cmd/app/models"
	"github.com/google/uuid"
)

type GatheringPlace struct {
	ID           uuid.UUID
	Address      models.Address
	AveragePrice int
	CuisineType  CuisineType
	Rating       int
	PhoneNumber  string
}

func NewGatheringPlace() *GatheringPlace {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil
	}

	return &GatheringPlace{ID: id}
}
