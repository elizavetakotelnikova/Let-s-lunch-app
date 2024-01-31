package meeting

import (
	"github.com/gofrs/uuid/v5"
)
import (
	"time"
)

type Meeting struct {
	ID               uuid.UUID
	GatheringPlaceId uuid.UUID
	InitiatorsId     uuid.UUID
	StartTime        time.Time
	EndTime          time.Time
	UsersQuantity    int
	State            MeetingState
}

func NewMeeting(gatheringPlaceID uuid.UUID, initiatorsID uuid.UUID, startTime time.Time, endTime time.Time, usersQuantity int, state MeetingState) *Meeting {
	id, err := uuid.NewV4()
	if err != nil {
		return nil
	}
	return &Meeting{
		ID:               id,
		GatheringPlaceId: gatheringPlaceID,
		InitiatorsId:     initiatorsID,
		StartTime:        startTime,
		EndTime:          endTime,
		UsersQuantity:    usersQuantity,
		State:            state,
	}
}

/*func NewMeeting(ID uuid.UUID, placeID uuid.UUID, usersID uuid.UUID,
	startTime time.Time, endTime time.Time, usersQuantity int, state MeetingState) *Meeting {
	return &Meeting{ID: ID, GatheringPlaceId: placeID, InitiatorsId: usersID,
		StartTime: startTime, EndTime: endTime, UsersQuantity: usersQuantity,
		State: state}
}*/
