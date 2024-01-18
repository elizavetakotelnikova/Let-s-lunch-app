package api_handler

import "net/http"

type Handler interface {
	Handle(writer http.ResponseWriter, request *http.Request)
}
