package factories

import (
	meeting_usecase "cmd/app/entities/meeting/usecases"
	"cmd/di/internal/lookup"
	"context"
)

func CreateUseCasesFindMeeting(ctx context.Context, c lookup.Container) *meeting_usecase.FindMeeting {
	return meeting_usecase.NewFindMeeting(
		c.Repositories().MeetingRepository(ctx),
	)
}
