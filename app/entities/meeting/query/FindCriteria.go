package query

import "github.com/google/uuid"

type FindCriteria struct {
	ID               uuid.NullUUID
	GatheringPlaceId uuid.NullUUID
	InitiatorID      uuid.NullUUID
}
