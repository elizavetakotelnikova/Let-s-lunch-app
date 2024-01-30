package api

import (
	usecase "cmd/app/entities/meeting/usecases"
	"encoding/json"
	"fmt"
	"github.com/gofrs/uuid/v5"
	"net/http"
	"path"
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
	base := path.Base(request.URL.Path)

	meeting, err := h.useCase.Handle(request.Context(), uuid.FromStringOrNil(base))
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
	}

	response, err := json.Marshal(meeting)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	}

	writer.Write(response)

	return
}
