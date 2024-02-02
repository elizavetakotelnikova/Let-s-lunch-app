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
	Title        string
	Description  string
	PhotoLink    string
}

func NewGatheringPlace(address models.Address, averagePrice int, cusineType CuisineType, rating int, phoneNumber string) *GatheringPlace {
	id, err := uuid.NewV4()
	if err != nil {
		return nil
	}

	return &GatheringPlace{
		ID:           id,
		Address:      address,
		AveragePrice: averagePrice,
		CuisineType:  cusineType,
		Rating:       rating,
		PhoneNumber:  phoneNumber,
	}
}
