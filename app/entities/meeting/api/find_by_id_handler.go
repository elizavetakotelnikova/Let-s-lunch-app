package api

import (
	domain "cmd/app/entities/meeting"
	usecase "cmd/app/entities/meeting/usecases"
	"cmd/pkg/errors"
	"encoding/json"
	uuid2 "github.com/google/uuid"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/gorilla/mux"
)

type FindMeetingRequest struct {
	ID uuid.UUID `json:"id"`
}

type JsonFindMeetingByIdResponse struct {
	Meeting *domain.Meeting
}

type FindMeeting struct {
	useCase *usecase.FindMeeting
}

func NewFindMeeting(useCase *usecase.FindMeeting) *FindMeeting {
	return &FindMeeting{useCase: useCase}
}

func (handler *FindMeeting) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	id, ok := mux.Vars(request)["id"]
	if !ok {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	uuidID, err := uuid.FromString(id)
	if err != nil {
		customError := errors.NewError(err)
		marshaledError, _ := json.Marshal(customError)

		writer.WriteHeader(http.StatusBadRequest)
		writer.Write(marshaledError)
		return
	}

	meeting, err := handler.useCase.Handle(request.Context(), uuid2.UUID(uuidID))
	if err != nil {
		customError := errors.NewError(err)
		marshaledError, _ := json.Marshal(customError)

		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write(marshaledError)
	}

	response := JsonFindMeetingByIdResponse{
		Meeting: meeting,
	}

	marshaledResponse, err := json.Marshal(response)
	if err != nil {
		customError := errors.NewError(err)
		marshaledError, _ := json.Marshal(customError)

		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write(marshaledError)
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Write(marshaledResponse)
}
