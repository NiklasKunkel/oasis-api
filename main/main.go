package main

import (
	"github.com/niklaskunkel/oasis-api/api"
	"github.com/niklaskunkel/oasis-api/client"
	"github.com/niklaskunkel/oasis-api/data"
)

func main() {
	//Read Config Data
	data.ReadConfig()

	//Validate connection to client
	client.InitClient()

	//Deploy API Server
	api.InitAPIServer()

	return
}