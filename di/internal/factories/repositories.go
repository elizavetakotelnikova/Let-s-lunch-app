package factories

import (
	meeting_domain "cmd/app/entities/meeting"
	meeting_repository "cmd/app/entities/meeting/repository"

	"cmd/di/internal/lookup"
	"context"
)

func CreateRepositoriesMeetingRepository(ctx context.Context, c lookup.Container) meeting_domain.MeetingRepository {
	conn, err := c.DB(ctx).Conn(ctx)
	if err != nil {
		panic("no connection to db")
	}

	return meeting_repository.NewMeetingRepository(ctx, conn)
}
