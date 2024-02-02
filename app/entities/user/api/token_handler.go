package api

import (
	"cmd/app/entities/user/usecases"
	"encoding/json"
	"net/http"
)

type getTokenRequest struct {
	PhoneNumber string `json:"phoneNumber"`
	Password    string `json:"password"`
}

type getTokenResponse struct {
	Token string `json:"token"`
}

type GetTokenHandler struct {
	useCase *usecases.GetTokenUseCase
}

func NewGetTokenHandler(useCase *usecases.GetTokenUseCase) *GetTokenHandler {
	return &GetTokenHandler{useCase: useCase}
}

func (t *GetTokenHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	var r getTokenRequest
	if err := json.NewDecoder(request.Body).Decode(&r); err != nil {
		marshaledError, _ := json.Marshal(err.Error())

		writer.WriteHeader(http.StatusBadRequest)
		writer.Write(marshaledError)
		return
	}

	token, err := t.useCase.Handle(request.Context(), r.PhoneNumber, r.Password)
	if err != nil {
		marshaledError, _ := json.Marshal(err.Error())

		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write(marshaledError)
		return
	}

	response, err := json.Marshal(getTokenResponse{
		Token: token,
	})

	writer.WriteHeader(http.StatusOK)
	writer.Write(response)
}
