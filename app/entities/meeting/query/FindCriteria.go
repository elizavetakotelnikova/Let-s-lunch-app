package query

import "github.com/google/uuid"

type FindCriteria struct {
	GatheringPlaceId uuid.NullUUID
	InitiatorID      uuid.NullUUID
}
