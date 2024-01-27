package factories

import (
	meeting_repository "cmd/app/entities/meeting/repository"

	"cmd/di/internal/lookup"
	"context"
)

func CreateRepositoriesMeetingRepository(ctx context.Context, c lookup.Container) meeting_repository.MeetingsRepository {
	return meeting_repository.NewMeetingsDatabaseRepository(c.DB(ctx))
}
