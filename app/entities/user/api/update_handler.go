package api

import (
	"cmd/app/entities/user/dto"
	"cmd/app/entities/user/usecases"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/gofrs/uuid/v5"
	"net/http"
)

type JsonUpdateUserResponse struct {
	ID uuid.UUID `json:"id"`
}

type UpdateUserHandler struct {
	useCase *usecases.UpdateUserUseCase
}

func NewUpdateUserHandler(useCase *usecases.UpdateUserUseCase) *UpdateUserHandler {
	return &UpdateUserHandler{useCase: useCase}
}

func (handler *UpdateUserHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	var updateUserDto dto.UpdateUserDto
	if err := json.NewDecoder(request.Body).Decode(&updateUserDto); err != nil {
		marshaledError, _ := json.Marshal(err.Error())

		writer.WriteHeader(http.StatusBadRequest)
		writer.Write(marshaledError)
		return
	}

	command := usecases.UpdateUserCommand{}
	command.Username = updateUserDto.Username
	command.DisplayName = updateUserDto.DisplayName
	command.Birthday = updateUserDto.Birthday
	command.PhoneNumber = updateUserDto.PhoneNumber
	command.Gender = updateUserDto.Gender
	id := chi.URLParam(request, "userID")

	uuidID, err := uuid.FromString(id)
	if err != nil {
		marshaledError, _ := json.Marshal(err.Error())

		writer.WriteHeader(http.StatusBadRequest)
		writer.Write(marshaledError)
		return
	}

	response, err := handler.useCase.Handle(request.Context(), command, uuidID)

	marshaledResponse, err := json.Marshal(response)

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Write(marshaledResponse)
}
