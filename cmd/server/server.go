package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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
	r.Get("/api", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("base api route"));
	})

	fmt.Printf("Server started at %s\n", s.host)

	err := http.ListenAndServe(s.host, r)
	if err != nil {
		log.Fatalf("Error starting the server on port: %s\n%s", s.host, err.Error())
	}
}