package gatheringPlace

import (
	"cmd/app/models"
	"github.com/gofrs/uuid/v5"
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
	id, err := uuid.NewV4()
	if err != nil {
		return nil
	}

	return &GatheringPlace{ID: id}
}
