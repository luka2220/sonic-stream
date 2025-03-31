package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/luka2220/sonic-stream/cmd/server/routes/api"
	"github.com/luka2220/sonic-stream/cmd/server/routes/download"
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
	rootMux := http.NewServeMux()

	apiRouter := api.NewAPIRoute()
	downloadRouter := download.NewDownloadRouter()

	rootMux.Handle(apiRouter.Base, http.StripPrefix("/api", apiRouter.Mux))
	rootMux.Handle(downloadRouter.Base, http.StripPrefix("/download", downloadRouter.Mux))

	err := http.ListenAndServe(s.host, rootMux)
	if err != nil {
		log.Fatalf("Error starting the server on port: %s\n%s", s.host, err.Error())
	}
}
