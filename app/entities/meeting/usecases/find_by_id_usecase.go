package usecases

import (
	domain "cmd/app/entities/meeting"
	"cmd/app/entities/meeting/repository"
	"context"
	"github.com/gofrs/uuid/v5"
)

type FindMeetingByIdResponse struct {
	Meeting *domain.Meeting `json:"data"`
}

type FindMeetingByIdUseCase struct {
	meetings repository.MeetingsRepository
}

func NewFindMeetingByIdUseCase(meetings repository.MeetingsRepository) *FindMeetingByIdUseCase {
	return &FindMeetingByIdUseCase{meetings: meetings}
}

func (useCase *FindMeetingByIdUseCase) Handle(ctx context.Context, id uuid.UUID) (*FindMeetingByIdResponse, error) {
	meetings, err := useCase.meetings.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &FindMeetingByIdResponse{
		Meeting: meetings,
	}, nil
}
