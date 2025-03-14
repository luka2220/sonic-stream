package server

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/luka2220/sonic-stream/cmd/server/routes"
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
	// NOTE: ugly logs atm
	logger := log.New(os.Stdout, "", 1)

	rootMux := http.NewServeMux()

	// Routers
	apiRouter := routes.NewAPIRoute(logger)
	
	// Register Routers
	rootMux.Handle(apiRouter.Base, http.StripPrefix("/api", apiRouter.Mux));
	
	err := http.ListenAndServe(s.host, rootMux)
	if err != nil {
		log.Fatalf("Error starting the server on port: %s\n%s", s.host, err.Error())
	}
}