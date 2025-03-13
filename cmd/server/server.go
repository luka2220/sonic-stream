package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/luka2220/sonic-stream/cmd/server/routes"
)

type Server struct {
	host string
}

func NewServer(port int) *Server {
	host := fmt.Sprintf(":%d", port)

	return &Server{
		host: host,
	}
}

func (s Server) Start() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.NotFound(notFoundHandler)

	r.Mount("/api", routes.ApiRouter())

	err := http.ListenAndServe(s.host, r)
	if err != nil {
		log.Fatalf("Error starting the server on port: %s\n%s", s.host, err.Error())
	}
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)
	w.Write([]byte("404 - Route Does Not Exist"))
}