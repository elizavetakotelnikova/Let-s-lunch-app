package meeting
import (
	"github.com/google/uuid"
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

func NewMeeting() *Meeting {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil
	}
	return &Meeting{ID: id}
}