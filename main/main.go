package main

import (
	"github.com/niklaskunkel/oasis-api/api"
	"github.com/niklaskunkel/oasis-api/client"
)
//2 Each REST call triggers a function in api package. Implementation of triggered function has template for creating TX obj,
//and triggering the CallTx function, then parses the response data into json and returns to user.

//3. Come up with asynchronous data pull from Oasis contract on interval (push would be better)
	//push could be done through subscribtion of event logs.
//4. Figure out best method of calculating price

//Parity Client - Go interfaces with client to send web3 Ethereum command "eth_subscribe"
	//then all events that match the filter will be pushed to your client.
	//Go then needs to be asynchronously checking the events the client is pushing.

//5. Have each incoming transaction have an id (increment static counter). Pass this id through the entire codeflow
//so when you run into an error you can record all errors as well as contextual info into the logs. Maybe if we logged
//everything and not just errors we can corss reference every step of a transaction using the id. 

func main() {
	//Validate connection to client
	client.InitClient()

	//Deploy API Server
	api.InitAPIServer()

	return
}