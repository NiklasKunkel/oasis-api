package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"log"
	"math/big"
	"strings"
	"time"
	"github.com/gorilla/mux"
	"github.com/niklaskunkel/oasis-api/client"
	"github.com/niklaskunkel/oasis-api/parser"
)

//Get MKR Token Supply
func GetMkrTokenSupply(w http.ResponseWriter, req *http.Request) {
	tx := client.CreateTx(
		"0x003EbC0613139A8dF37CAC03d39B39304153596A",
		"0xc66ea802717bfb9833400264dd12c2bceaa34a6d",
		0,
		big.NewInt(0),
		big.NewInt(0),
		"0x18160ddd",
		0)
	hSupply, err := client.CallTx(tx)
	fmt.Printf(hSupply)
	if err != nil {
		json.NewEncoder(w).Encode(Error{fmt.Sprintf("Querying MKR Token Supply Failed")})
	} else {
		iSupply := parser.Hex2Int(hSupply)
		fSupply := parser.AdjustIntForPrecision(iSupply, 18)
		json.NewEncoder(w).Encode(MkrTokenSupply{fSupply.Text('f', 0)})
	}
	return
}

func GetMarkets(w http.ResponseWriter, req *http.Request) {
	marketStatus := client.GetMarkets()
	json.NewEncoder(w).Encode(marketStatus)
}

//Get All Token Data 
func GetTokenPair(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req) 
	baseToken := strings.ToUpper(params["base"])
	quoteToken := strings.ToUpper(params["quote"])
	tokenPair := baseToken + string('/') + quoteToken

	//need to check that tokenPair is in supported tokens list

	bid, ask, err := client.GetSpread(baseToken, quoteToken)
	if err != nil {
		fmt.Printf("GetTokenPair] could not GetSpread() due to (%s)\n", err)
		json.NewEncoder(w).Encode(Error{fmt.Sprintf("Querying token pair %s failed", tokenPair)})
	}

	filter, err := client.CreateEventFilter("24hour", "latest", []string{"0x3Aa927a97594c3ab7d7bf0d47C71c3877D1DE4A1"}, [][]string{[]string{"0x819e390338feffe95e2de57172d6faf337853dfd15c7a09a32d76f7fd2443875"}})
	if err != nil {
		fmt.Printf("[GetTokenPair] could not CreateEventFilter() due to (%s)\n", err)
		json.NewEncoder(w).Encode(Error{fmt.Sprintf("Querying token pair %s failed", tokenPair)})
		return
	}
	logs, err := client.GetLogs(filter)
	if err != nil {
		fmt.Printf("[GetTokenPair] could not GetLogs() due to (%s)\n", err)
		json.NewEncoder(w).Encode(Error{fmt.Sprintf("Querying token pair %s failed", tokenPair)})
		return
	}

	sPrice, sVolume, sMin, sMax, err := client.CalculatePriceFromLogs(logs, baseToken, quoteToken)
	if err != nil {
		fmt.Printf("[GetTokenPair] could not CalculatePriceFromLogs() due to (%s)\n", err)
		json.NewEncoder(w).Encode(Error{fmt.Sprintf("Querying token pair %s failed", tokenPair)})
	}

	json.NewEncoder(w).Encode(TokenPair{tokenPair, sPrice, sVolume, ask, bid, sMin, sMax, true, time.Now().Unix()})
	return
}

func GetTokenPairSpread(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	baseToken := params["base"]
	quoteToken := params["quote"]
	bid, ask, err := client.GetSpread(baseToken, quoteToken)
	if err != nil {
		fmt.Printf("[GetBidAskSpread] failed due to (%s)\n", err)
		json.NewEncoder(w).Encode(Error{fmt.Sprintf("Querying Bid/Ask Spread Failed")})
		return
	}
	json.NewEncoder(w).Encode(TokenPairSpread{bid,ask})
}

func GetTokenPairPrice(w http.ResponseWriter, req *http.Request) {
	//
}

func GetTokenPairVolume(w http.ResponseWriter, req *http.Request) {
	//
}

func GetVolume(w http.ResponseWriter, req *http.Request) {
	//params := mux.Vars(req)
	//iterate through all tokens in hashtable
		//append volume
	//encode to JSON
	//return
}

//Request Handler
func InitAPIServer() {

	router := mux.NewRouter()																	//Create new router

	//API Endpoints
	router.HandleFunc("/v1/tokens/mkr/totalsupply", GetMkrTokenSupply).Methods("GET")			//REST endpoint for calling MKR token supply
	router.HandleFunc("/v1/markets/{base}/{quote}", GetTokenPair).Methods("GET")				//REST endpoint for calling token pair data
	router.HandleFunc("/v1/markets/{base}/{quote}/spread", GetTokenPairSpread).Methods("GET")			//REST endpoint for calling spread of token pair
	router.HandleFunc("/v1/markets/{base}/{quote}/price", GetTokenPairPrice).Methods("GET")		//REST endpoint for calling price of token pair
	router.HandleFunc("/v1/markets/{base}/{quote}/volume", GetTokenPairVolume).Methods("GET")	//REST endpoint for calling volume of token pair
	router.HandleFunc("/v1/markets/volume", GetVolume).Methods("GET")							//REST endpoint for calling volume of all tokens
	router.HandleFunc("/v1/markets", GetMarkets).Methods("GET")									//REST endpoint for calling tradable markets

	fmt.Printf("API Server Started\nReady for incoming requests\n")

	log.Fatal(http.ListenAndServe(":12345", router))							//Deploy server
}