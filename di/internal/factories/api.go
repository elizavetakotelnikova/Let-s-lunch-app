package factories

import (
	meeting_api "cmd/app/entities/meeting/api"
	"cmd/di/internal/lookup"
	"context"
)

func CreateAPIFindMeetingHandler(ctx context.Context, c lookup.Container) *meeting_api.FindMeeting {
	panic("not implemented")
}
