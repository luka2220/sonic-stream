package api

import (
	"log/slog"
	"net/http"
)

type ServerError struct {
	Message       string
	ClientMessage string
	Status        int
	W             http.ResponseWriter
	L             *slog.Logger
}

func (e ServerError) Error() string {
	return e.Message
}

func InternalServerError(e ServerError) {
	e.L.Error((e.Error()))
	e.W.WriteHeader(500)
	e.W.Header().Add("Content-Type", "text/html; charset=utf-8")
	e.W.Write([]byte("Server Error"))
}

func ClientError(e ServerError) {
	e.L.Error(e.Error())
	e.W.WriteHeader(400)
	e.W.Header().Add("Content-Type", "text/html; charset=utf-8")
	e.W.Write([]byte(e.ClientMessage))
}
