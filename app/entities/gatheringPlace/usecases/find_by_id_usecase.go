package usecases

import (
	domain "cmd/app/entities/gatheringPlace"
	repositoryPlaces "cmd/app/entities/gatheringPlace/repository"
	"context"
	"github.com/google/uuid"
)

type FindGatheringPlace struct {
	gatheringPlaces repositoryPlaces.PlacesRepository
}

func NewFindGatheringPlace(gatheringPlaces repositoryPlaces.PlacesRepository) *FindGatheringPlace {
	return &FindGatheringPlace{gatheringPlaces: gatheringPlaces}
}

func (f *FindGatheringPlace) Handle(ctx context.Context, id uuid.UUID) (*domain.GatheringPlace, error) {
	entity, err := f.gatheringPlaces.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return entity, nil
}
