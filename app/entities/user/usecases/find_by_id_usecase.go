package usecases

import (
	domain "cmd/app/entities/meeting"
	"cmd/app/entities/meeting/repository"
	"context"
	"github.com/gofrs/uuid/v5"
)

type FindMeeting struct {
	meetings repository.MeetingsRepository
}

func NewFindMeeting(meetings repository.MeetingsRepository) *FindMeeting {
	return &FindMeeting{meetings: meetings}
}

func (f *FindMeeting) Handle(ctx context.Context, id uuid.UUID) (*domain.Meeting, error) {
	entity, err := f.meetings.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return entity, nil
}
