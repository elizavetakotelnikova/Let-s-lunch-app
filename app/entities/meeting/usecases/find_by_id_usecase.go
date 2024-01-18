package usecases

import (
	domain "cmd/app/entities/meeting"
	"context"
	"github.com/google/uuid"
)

type FindMeeting struct {
	meetings domain.MeetingRepository
}

func NewFindMeeting(meetings domain.MeetingRepository) *FindMeeting {
	return &FindMeeting{meetings: meetings}
}

func (f *FindMeeting) Handle(ctx context.Context, id uuid.UUID) (*domain.Meeting, error) {
	entity, err := f.meetings.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return entity, nil
}
