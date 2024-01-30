package usecases

import (
	domain "cmd/app/entities/gatheringPlace"
	repositoryPlaces "cmd/app/entities/gatheringPlace/repository"
	"context"
	"github.com/gofrs/uuid/v5"
)

type FindGatheringPlaceByIdResponse struct {
	GatheringPlace *domain.GatheringPlace `json:"data"`
}

type FindGatheringPlaceByIdUseCase struct {
	gathering_place repositoryPlaces.PlacesRepository
}

func NewFindGatheringPlaceByIdUseCase(gathering_places repositoryPlaces.PlacesRepository) *FindGatheringPlaceByIdUseCase {
	return &FindGatheringPlaceByIdUseCase{gathering_place: gathering_places}
}

func (useCase *FindGatheringPlaceByIdUseCase) Handle(ctx context.Context, id uuid.UUID) (*FindGatheringPlaceByIdResponse, error) {
	gathering_places, err := useCase.gathering_place.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &FindGatheringPlaceByIdResponse{GatheringPlace: gathering_places}, nil
}
