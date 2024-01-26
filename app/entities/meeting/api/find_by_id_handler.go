package api

import (
	usecase "cmd/app/entities/meeting/usecases"
	"fmt"
	"github.com/google/uuid"
	"net/http"
)

type FindMeetingRequest struct {
	ID uuid.UUID `json:"id"`
}

type FindMeeting struct {
	useCase *usecase.FindMeeting
}

func NewFindMeeting(useCase *usecase.FindMeeting) *FindMeeting {
	return &FindMeeting{useCase: useCase}
}

func (h *FindMeeting) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("in meeting handler")

	return
}
