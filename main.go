package main

import (
	"github.com/luka2220/sonic-stream/cmd/server"
)

func main() {
	server.NewServer(3000).Start()
}