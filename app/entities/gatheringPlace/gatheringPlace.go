package gatheringPlace

import (
	"cmd/app/models"
	"github.com/gofrs/uuid/v5"
)

type GatheringPlace struct {
	ID           uuid.UUID      `json:"id"`
	Address      models.Address `json:"address"`
	AveragePrice int            `json:"averagePrice"`
	CuisineType  CuisineType    `json:"cuisineType"`
	Rating       int            `json:"rating"`
	PhoneNumber  string         `json:"phoneNumber"`
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
