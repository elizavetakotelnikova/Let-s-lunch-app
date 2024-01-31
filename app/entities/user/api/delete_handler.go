package api

import (
	"cmd/app/entities/user/usecases"
	"cmd/pkg/errors"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/gofrs/uuid/v5"
	"net/http"
)

type DeleteUserHandler struct {
	useCase *usecases.DeleteUserUseCase
}

func NewDeleteUserHandler(useCase *usecases.DeleteUserUseCase) *DeleteUserHandler {
	return &DeleteUserHandler{useCase: useCase}
}

func (handler *DeleteUserHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	id := chi.URLParam(request, "userID")

	uuidID, err := uuid.FromString(id)
	if err != nil {
		customError := errors.NewError(err)
		marshledError, _ := json.Marshal(customError)

		writer.WriteHeader(http.StatusBadRequest)
		writer.Write(marshledError)
		return
	}

	err = handler.useCase.Handle(request.Context(), uuidID)
	if err != nil {
		customError := errors.NewError(err)
		marshaledError, _ := json.Marshal(customError)

		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write(marshaledError)
		return
	}

	writer.WriteHeader(http.StatusOK)
}
