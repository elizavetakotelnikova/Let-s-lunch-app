package factories

import (
	meeting_domain "cmd/app/entities/meeting"
	"cmd/di/internal/lookup"
	"context"
)

func CreateRepositoriesMeetingRepository(ctx context.Context, c lookup.Container) meeting_domain.MeetingRepository {
	panic("not implemented")
}
