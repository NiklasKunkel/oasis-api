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

var TokenData = make(map[string]Token)
var publicMethods = []string {
	"GetToken",
	"GetTokenPrice",
	"GetTokenVolume",
	"GetVolume",
}
var privateMethods = []string {
	"Admin",
}

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
		fSupply := parser.AdjustForPrecision(iSupply, 18)
		json.NewEncoder(w).Encode(MkrTokenSupply{fSupply.Text('f', 0)})
	}
	return
}

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
	sPrice, err := client.CalculatePriceFromLogs("MKRETH", logs)
	if err != nil {
		fmt.Printf("[GetMkrTokenPrice] could not CalculatePriceFromLogs() due to (%s)\n", err)
	}
	json.NewEncoder(w).Encode(MkrTokenPrice{sPrice})
	return
}

//Get All Token Data 
func GetToken(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	token := strings.ToUpper(params["token"])
	if data, ok := TokenData[token]; !ok {
		json.NewEncoder(w).Encode(Error{fmt.Sprintf("Invalid token name %s", token)})
	} else {
		json.NewEncoder(w).Encode(data)
		return
	}
}

func GetTokenPrice(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	token := strings.ToUpper(params["token"])
	if data, ok := TokenData[token]; !ok {
		json.NewEncoder(w).Encode(Error{fmt.Sprintf("Invalid token name %s", token)})
		return
	} else {
		json.NewEncoder(w).Encode(TokenPrice{data.Id, data.Price})
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

func GetVolume(w http.ResponseWriter, req *http.Request) {
	//params := mux.Vars(req)
	//iterate through all tokens in hashtable
		//append volume
	//encode to JSON
	//return
}

//Request Handler
func InitAPIServer() {
	//for testing
	TokenData["ETH"] = Token{"ETH", 232.10, 137.5, 234.40, 230.20, 287, 199, false, 123123121}
	TokenData["MKR"] = Token{"MKR", 200.00, 3522.732, 190.60, 185.70, 192, 185.70, false, 543453453}

	router := mux.NewRouter()													//Create new router

	//API Endpoints
	router.HandleFunc("/mkr/totalsupply", GetMkrTokenSupply).Methods("GET")		//REST endpoint for calling MKR token supply
	router.HandleFunc("/mkr/price", GetMkrTokenPrice).Methods("GET")			//REST endpoint for calling MKR token price
	router.HandleFunc("/tokens/{token}", GetToken).Methods("GET")				//REST endpoint for calling token data
	router.HandleFunc("/tokens/{token}/price", GetTokenPrice).Methods("GET")	//REST endpoint for calling price of token
	router.HandleFunc("/tokens/{token}/volume", GetTokenVolume).Methods("GET")	//REST endpoint for calling volume of token
	router.HandleFunc("/volume", GetVolume).Methods("GET")						//REST endpoint for calling volume of all tokens

	fmt.Printf("API Server Started\nReady for incoming requests\n")

	log.Fatal(http.ListenAndServe(":12345", router))							//Deploy server
}