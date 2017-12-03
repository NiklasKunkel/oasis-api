package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"log"
	"strings"
	"time"
	"github.com/gorilla/mux"
	"github.com/niklaskunkel/oasis-api/client"
)

//Get MKR Token Supply
func APIGetMkrTokenSupply(w http.ResponseWriter, req *http.Request) {
	supply, err := client.GetMkrTokenSupply()
	if err != nil {
		json.NewEncoder(w).Encode(Error{fmt.Sprintf("Querying MKR Token Supply Failed")})
	} else {
		json.NewEncoder(w).Encode(MkrTokenSupply{supply})
	}
	return
}

func APIGetMarkets(w http.ResponseWriter, req *http.Request) {
	marketStatus := client.GetMarkets()
	json.NewEncoder(w).Encode(marketStatus)
}

//Get All Token Data 
func APIGetTokenPair(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req) 
	baseToken := strings.ToUpper(params["base"])
	quoteToken := strings.ToUpper(params["quote"])
	tokenPair := baseToken + string('/') + quoteToken

	//Verify that token pair is supported by oasisdex
	if (!client.IsValidTokenPair(tokenPair)) {
		json.NewEncoder(w).Encode(Error{fmt.Sprintf("Unknown token pair")})
		return
	}

	//get bid/ask spread
	bid, ask, err := client.GetSpread(baseToken, quoteToken)
	if err != nil {
		fmt.Printf("GetTokenPair] could not GetSpread() due to (%s)\n", err)
		json.NewEncoder(w).Encode(Error{fmt.Sprintf("Querying token pair %s failed", tokenPair)})
		return
	}

	//create event filter
	filter, err := client.CreateEventFilter("24hour", "latest", []string{"0x3Aa927a97594c3ab7d7bf0d47C71c3877D1DE4A1"}, [][]string{[]string{"0x819e390338feffe95e2de57172d6faf337853dfd15c7a09a32d76f7fd2443875"}})
	if err != nil {
		fmt.Printf("[GetTokenPair] could not CreateEventFilter() due to (%s)\n", err)
		json.NewEncoder(w).Encode(Error{fmt.Sprintf("Querying token pair %s failed", tokenPair)})
		return
	}

	//get all events from last 24 hour interval
	logs, err := client.GetLogs(filter)
	if err != nil {
		fmt.Printf("[GetTokenPair] could not GetLogs() due to (%s)\n", err)
		json.NewEncoder(w).Encode(Error{fmt.Sprintf("Querying token pair %s failed", tokenPair)})
		return
	}

	//calculate price, volume, min, and max from event logs
	sPrice, sVolume, sMin, sMax, err := client.CalculatePriceFromLogs(logs, baseToken, quoteToken)
	if err != nil {
		fmt.Printf("[GetTokenPair] could not CalculatePriceFromLogs() due to (%s)\n", err)
		json.NewEncoder(w).Encode(Error{fmt.Sprintf("Querying token pair %s failed", tokenPair)})
		return
	}

	json.NewEncoder(w).Encode(TokenPair{tokenPair, sPrice, sVolume, ask, bid, sMin, sMax, true, time.Now().Unix()})
	return
}

func APIGetTokenPairSpread(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	baseToken := params["base"]
	quoteToken := params["quote"]

	if (!client.IsValidTokenPair(baseToken + string("/") + quoteToken)) {
		json.NewEncoder(w).Encode(Error{fmt.Sprintf("Unknown token pair")})
		return
	}
	bid, ask, err := client.GetSpread(baseToken, quoteToken)
	if err != nil {
		fmt.Printf("[GetBidAskSpread] failed due to (%s)\n", err)
		json.NewEncoder(w).Encode(Error{fmt.Sprintf("Querying Bid/Ask Spread Failed")})
		return
	}
	json.NewEncoder(w).Encode(TokenPairSpread{bid,ask})
}

func APIGetTokenPairPrice(w http.ResponseWriter, req *http.Request) {
	//
}

func APIGetTokenPairVolume(w http.ResponseWriter, req *http.Request) {
	//
}

func APIGetVolume(w http.ResponseWriter, req *http.Request) {
	//params := mux.Vars(req)
	//iterate through all tokens in hashtable
		//append volume
	//encode to JSON
	//return
}

//Request Handler
func InitAPIServer() {

	router := mux.NewRouter()																		//Create new router

	//API Endpoints
	router.HandleFunc("/v1/tokens/mkr/totalsupply", APIGetMkrTokenSupply).Methods("GET")			//REST endpoint for calling MKR token supply
	router.HandleFunc("/v1/markets/{base}/{quote}", APIGetTokenPair).Methods("GET")					//REST endpoint for calling token pair data
	router.HandleFunc("/v1/markets/{base}/{quote}/spread", APIGetTokenPairSpread).Methods("GET")	//REST endpoint for calling spread of token pair
	router.HandleFunc("/v1/markets/{base}/{quote}/price", APIGetTokenPairPrice).Methods("GET")		//REST endpoint for calling price of token pair
	router.HandleFunc("/v1/markets/{base}/{quote}/volume", APIGetTokenPairVolume).Methods("GET")	//REST endpoint for calling volume of token pair
	router.HandleFunc("/v1/markets/volume", APIGetVolume).Methods("GET")							//REST endpoint for calling volume of all tokens
	router.HandleFunc("/v1/markets", APIGetMarkets).Methods("GET")									//REST endpoint for calling tradable markets

	fmt.Printf("API Server Started\nReady for incoming requests\n")

	log.Fatal(http.ListenAndServe(":80", router))												//Deploy server
}