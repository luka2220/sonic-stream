package server

import (
	"fmt"
	"log"
	"net/http"
)

type Server struct {
	host string
}

func NewServer(port string) *Server {
	host := fmt.Sprintf(":%s", port)

	return &Server{
		host: host,
	}
}

func (s Server) Start() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/", rootHandler)

	err := http.ListenAndServe(s.host, mux)
	if err != nil {
		log.Fatalf("Error starting the server on port: %s\n%s", s.host, err.Error())
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Root API Handler"))
}