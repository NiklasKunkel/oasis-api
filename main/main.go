package main

import (
	"github.com/niklaskunkel/oasis-api/api"
	"github.com/niklaskunkel/oasis-api/client"
)

func main() {
	//Validate connection to client
	client.InitClient()

	//Deploy API Server
	api.InitAPIServer()

	return
}