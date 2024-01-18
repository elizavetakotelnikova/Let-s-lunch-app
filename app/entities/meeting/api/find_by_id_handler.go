package api

import (
	usecase "cmd/app/entities/meeting/usecases"
	"github.com/google/uuid"
	"net/http"
)

type FindMeetingRequest struct {
	ID uuid.UUID `json:"id"`
}

type FindMeeting struct {
	useCase *usecase.FindMeeting
}

func NewFindEntity(useCase *usecase.FindMeeting) *FindMeeting {
	return &FindMeeting{useCase: useCase}
}

func (h *FindMeeting) Handle(writer http.ResponseWriter, request *http.Request) {
	panic("implement me")
}
