package api

import (
	domain "cmd/app/entities/gatheringPlace"
	usecases "cmd/app/entities/gatheringPlace/usecases"
	"cmd/app/models"
	"cmd/pkg/errors"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/gofrs/uuid/v5"
	"net/http"
)

type JsonFindGatheringPlaceByIdResponse struct {
	ID          uuid.UUID          `json:"id"`
	Address     models.Address     `json:"address"`
	AvgPrice    int                `json:"averagePrice"`
	CusineType  domain.CuisineType `json:"cusineType"`
	Rating      int                `json:"rating"`
	PhoneNumber string             `json:"phoneNumber"`
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
		ID:          gathering_places.ID,
		Address:     gathering_places.Address,
		AvgPrice:    gathering_places.AvgPrice,
		CusineType:  gathering_places.CusineType,
		Rating:      gathering_places.Rating,
		PhoneNumber: gathering_places.PhoneNumber,
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
