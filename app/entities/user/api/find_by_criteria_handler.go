package api

import (
	"cmd/app/entities/user"
	"cmd/app/entities/user/usecases"
	"cmd/app/models"
	"encoding/json"
	"github.com/gofrs/uuid/v5"
	"net/http"
	"strconv"
)

type JsonFindUsersByCriteriaResponse struct {
	User []user.User `json:"user"`
}

type FindUsersByCriteriaHandler struct {
	useCase *usecases.FindUsersByCriteriaUseCase
}

func NewFindUsersByCriteriaHandler(useCase *usecases.FindUsersByCriteriaUseCase) *FindUsersByCriteriaHandler {
	return &FindUsersByCriteriaHandler{useCase: useCase}
}

func (handler *FindUsersByCriteriaHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	username := request.URL.Query().Get("user_name")
	displayName := request.URL.Query().Get("dislay_name")
	currentMeetingID := request.URL.Query().Get("current_meeting_id")
	age := request.URL.Query().Get("age")
	gender := request.URL.Query().Get("gender")

	findCriteria := models.FindUserCriteria{}
	findCriteria.Username = username
	findCriteria.DisplayName = displayName

	if currentMeetingID != "" {
		uuidID, err := uuid.FromString(currentMeetingID)
		if err != nil {
			marshaledError, _ := json.Marshal(err)

			writer.WriteHeader(http.StatusBadRequest)
			writer.Write(marshaledError)
			return
		}
		findCriteria.CurrentMeetingId = uuidID
	}

	if age != "" {
		convertedAge, err := strconv.Atoi(age)
		if err != nil {
			marshaledError, _ := json.Marshal(err)

			writer.WriteHeader(http.StatusBadRequest)
			writer.Write(marshaledError)
			return
		}
		findCriteria.Age = convertedAge
	}

	if gender != "" {
		convertedGender, err := strconv.Atoi(gender)
		if err != nil {
			marshaledError, _ := json.Marshal(err)

			writer.WriteHeader(http.StatusBadRequest)
			writer.Write(marshaledError)
			return
		}

		actualGender := user.Gender(convertedGender)

		findCriteria.Gender = &actualGender
	}

	users, err := handler.useCase.Handle(request.Context(), findCriteria)
	if err != nil {
		marshaledError, _ := json.Marshal(err)

		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write(marshaledError)
		return
	}

	response := users

	marshaledResponse, err := json.Marshal(response)
	if err != nil {
		marshaledError, _ := json.Marshal(err)

		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write(marshaledError)
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Write(marshaledResponse)
}
