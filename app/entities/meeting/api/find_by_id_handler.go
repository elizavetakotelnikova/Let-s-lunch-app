package api

import (
	domain "cmd/app/entities/meeting"
	usecase "cmd/app/entities/meeting/usecases"
	"cmd/pkg/errors"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/gofrs/uuid/v5"
	"net/http"
)

type JsonFindMeetingByIdResponse struct {
	Meeting *domain.Meeting `json:"data"`
}

type FindMeetingByIdHandler struct {
	useCase *usecase.FindMeetingByIdUseCase
}

func NewFindMeetingByIdHandler(useCase *usecase.FindMeetingByIdUseCase) *FindMeetingByIdHandler {
	return &FindMeetingByIdHandler{useCase: useCase}
}

func (handler *FindMeetingByIdHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	id := chi.URLParam(request, "meetingID")

	uuidID, err := uuid.FromString(id)
	if err != nil {
		customError := errors.NewError(err)
		marshaledError, _ := json.Marshal(customError)

		writer.WriteHeader(http.StatusBadRequest)
		writer.Write(marshaledError)
		return
	}

	meeting, err := handler.useCase.Handle(request.Context(), uuidID)
	if err != nil {
		customError := errors.NewError(err)
		marshaledError, _ := json.Marshal(customError)

		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write(marshaledError)
	}

	response := JsonFindMeetingByIdResponse{
		Meeting: meeting.Meeting,
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
