package main

import (
	"os"

	"github.com/luka2220/sonic-stream/cmd/server"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	server.NewServer(port).Start()
}
