package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"log"
	"math/big"
	"strings"
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

/*
//Get MKR Token Price
func GetMkrTokenPrice(w http.ResponseWriter, req *http.Request) {
	//Subscribe to Event Filter
	params, err := client.CreateEventFilter("24hour", "latest", []string{"0x3Aa927a97594c3ab7d7bf0d47C71c3877D1DE4A1"}, [][]string{[]string{"0x819e390338feffe95e2de57172d6faf337853dfd15c7a09a32d76f7fd2443875"}})
	if err != nil {
		fmt.Printf("[GetMkrTokenPrice] could not CreateEventFilter() due to (%s)\n", err)
		json.NewEncoder(w).Encode(Error{fmt.Sprintf("Querying MKR Token Price Failed")})
		return
	}
	logs, err := client.GetLogs(params)
	if err != nil {
		fmt.Printf("[GetMkrTokenPrice] could not GetLogs() due to (%s)\n", err)
		return
	}
	sPrice, err := client.CalculatePriceFromLogs(logs, "MKR", "ETH")
	if err != nil {
		fmt.Printf("[GetMkrTokenPrice] could not CalculatePriceFromLogs() due to (%s)\n", err)
	}
	json.NewEncoder(w).Encode(MkrTokenPrice{sPrice})
	return
}
*/

//Get All Token Data 
func GetTokenPair(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req) 
	baseToken := strings.ToUpper(params["base"])
	quoteToken := strings.ToUpper(params["quote"])
	tokenPair := baseToken + string('/') + quoteToken

	filter, err := client.CreateEventFilter("4601800", "latest", []string{"0x3Aa927a97594c3ab7d7bf0d47C71c3877D1DE4A1"}, [][]string{[]string{"0x819e390338feffe95e2de57172d6faf337853dfd15c7a09a32d76f7fd2443875"}})
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

	bid, ask, err := client.GetSpread(baseToken, quoteToken)
	if err != nil {
		fmt.Printf("GetTokenPair] could not GetSpread() due to (%s)\n", err)
	}
	json.NewEncoder(w).Encode(TokenPair{tokenPair, sPrice, sVolume, ask, bid, sMin, sMax, true, 1})
	return
}

func GetSpread(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	baseToken := params["base"]
	quoteToken := params["quote"]
	bid, ask, err := client.GetSpread(baseToken, quoteToken)
	if err != nil {
		fmt.Printf("[GetBidAskSpread] failed due to (%s)\n", err)
		json.NewEncoder(w).Encode(Error{fmt.Sprintf("Querying Bid/Ask Spread Failed")})
		return
	}
	json.NewEncoder(w).Encode(Spread{bid,ask})
}

/*
func GetTokenPrice(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	token := strings.ToUpper(params["token"])
	if data, ok := TokenData[token]; !ok {
		json.NewEncoder(w).Encode(Error{fmt.Sprintf("Invalid token name %s", token)})
		return
	} else {
		json.NewEncoder(w).Encode(TokenPrice{data.TokenPair, data.Price})
		return
	}
}

func GetTokenVolume(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	token := strings.ToUpper(params["token"])
	if data, ok := TokenData[token]; !ok {
		json.NewEncoder(w).Encode(Error{fmt.Sprintf("Invalid token name %s", token)})
		return
	} else {
		json.NewEncoder(w).Encode(TokenVolume{data.Id, data.Volume})
		return
	}
}
*/

func GetVolume(w http.ResponseWriter, req *http.Request) {
	//params := mux.Vars(req)
	//iterate through all tokens in hashtable
		//append volume
	//encode to JSON
	//return
}

//Request Handler
func InitAPIServer() {

	router := mux.NewRouter()													//Create new router

	//API Endpoints
	router.HandleFunc("/mkr/totalsupply", GetMkrTokenSupply).Methods("GET")		//REST endpoint for calling MKR token supply
	//router.HandleFunc("/mkr/price", GetMkrTokenPrice).Methods("GET")			//REST endpoint for calling MKR token price
	router.HandleFunc("/tokens/{base}/{quote}", GetTokenPair).Methods("GET")	//REST endpoint for calling token data
	router.HandleFunc("/tokens/{base}/{quote}/spread", GetSpread).Methods("GET")	//REST endpoint for calling best bid and ask

	//router.HandleFunc("/tokens/{token}/price", GetTokenPrice).Methods("GET")	//REST endpoint for calling price of token
	//router.HandleFunc("/tokens/{token}/volume", GetTokenVolume).Methods("GET")	//REST endpoint for calling volume of token
	//router.HandleFunc("/volume", GetVolume).Methods("GET")						//REST endpoint for calling volume of all tokens

	fmt.Printf("API Server Started\nReady for incoming requests\n")

	log.Fatal(http.ListenAndServe(":12345", router))							//Deploy server
}