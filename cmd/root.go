package cmd

import (
	"dtech.com/traktirku/internal/api"
	"dtech.com/traktirku/pkg/server"
)

// run app entrypoint
func RunApp() {
	server := &server.ServerInstance{
		Host: "0.0.0.0",
		Port: ":8081",
	}
	//api router

	server.NewServer()
	api.NewApi(server.V1ApiRouter)
	server.RunServer()
}
