package api

import (
	domain "cmd/app/entities/gatheringPlace"
	usecases "cmd/app/entities/gatheringPlace/usecases"
	"cmd/pkg/errors"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/gofrs/uuid/v5"
	"net/http"
)

type JsonFindGatheringPlaceByIdResponse struct {
	GatheringPlace *domain.GatheringPlace `json:"data"`
}

type FindGatheringPlaceByIdHandler struct {
	useCase *usecases.FindGatheringPlaceByIdUseCase
}

func NewFindGatheringCaseByIdHandler(useCase *usecases.FindGatheringPlaceByIdUseCase) *FindGatheringPlaceByIdHandler {
	return &FindGatheringPlaceByIdHandler{useCase: useCase}
}

func (handler *FindGatheringPlaceByIdHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	id := chi.URLParam(request, "placeID")

	uuidID, err := uuid.FromString(id)
	if err != nil {
		customError := errors.NewError(err)
		marshaledError, _ := json.Marshal(customError)

		writer.WriteHeader(http.StatusBadRequest)
		writer.Write(marshaledError)
		return
	}

	gathering_places, err := handler.useCase.Handle(request.Context(), uuidID)
	if err != nil {
		customError := errors.NewError(err)
		marshaledError, _ := json.Marshal(customError)

		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write(marshaledError)
		return
	}

	response := JsonFindGatheringPlaceByIdResponse{
		GatheringPlace: gathering_places.GatheringPlace,
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
