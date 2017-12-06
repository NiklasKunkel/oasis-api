package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"log"
	"reflect"
	"strings"
	"time"
	"github.com/gorilla/mux"
	"github.com/niklaskunkel/oasis-api/client"
	"github.com/niklaskunkel/oasis-api/data"
)

func APIGetAllPairs(w http.ResponseWriter, req *http.Request) {
	marketStatus := client.GetAllPairs()
	json.NewEncoder(w).Encode(marketStatus)
}

func APIGetPair(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	baseToken := strings.ToUpper(params["base"])
	quoteToken := strings.ToUpper(params["quote"])
	tokenPair := baseToken + string('/') + quoteToken

	//Verify that token pair is supported by oasisdex
	if (!client.IsValidTokenPair(tokenPair)) {
		json.NewEncoder(w).Encode(Error{fmt.Sprintf("Unknown pair")})
		return
	}

	market, err := client.GetPair(baseToken, quoteToken)
	if err != nil {
		json.NewEncoder(w).Encode(Error{fmt.Sprintf("Querying market for %s failed", tokenPair)})
		return
	}

	json.NewEncoder(w).Encode(market)
}

func APIGetAllPrices(w http.ResponseWriter, req *http.Request) {
	//TODO
}

func APIGetTokenPairPrice(w http.ResponseWriter, req *http.Request) {
	//TODO
}

func APIGetAllVolume(w http.ResponseWriter, req *http.Request) {
	allVolumes := AllVolumes{make(map[string]string), time.Now().Unix()}
	for tokenPair, tokenPairInfo := range data.LiveMarkets {
		vol, err := client.GetTokenPairVolume(tokenPairInfo.Base, tokenPairInfo.Quote)
		if err != nil {
			vol = "null"
		}
		allVolumes.Volumes[tokenPair] = vol
	}
	json.NewEncoder(w).Encode(allVolumes)
}

func APIGetTokenPairVolume(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	baseToken := strings.ToUpper(params["base"])
	quoteToken := strings.ToUpper(params["quote"])
	if (!client.IsValidTokenPair(baseToken + string("/") + quoteToken)) {
		json.NewEncoder(w).Encode(Error{fmt.Sprintf("Unknown token pair")})
		return
	}
	vol, err := client.GetTokenPairVolume(baseToken, quoteToken)
	if err != nil {
		json.NewEncoder(w).Encode(Error{fmt.Sprintf("Querying Volume Failed")})
		return
	}
	json.NewEncoder(w).Encode(TokenPairVolume{vol, time.Now().Unix()})
	return
}

func APIGetAllSpread(w http.ResponseWriter, req *http.Request) {
	allSpreads := AllSpreads{make(map[string]*TokenPairSpread), time.Now().Unix()}
	for tokenPair, tokenPairInfo := range data.LiveMarkets {
		baseToken := tokenPairInfo.Base
		quoteToken := tokenPairInfo.Quote
		bid, ask, err := client.GetSpread(baseToken, quoteToken)
		if err != nil {
			fmt.Printf("[APIGetAllSpread] failed to get spread for %s due to (%s)\n", tokenPair, err)
			bid = "null"
			ask = "null"
		}
		fmt.Printf("Setting Spread For %s : Bid = %s | Ask = %s\n", tokenPair, bid, ask)
		pair := allSpreads.Spreads[tokenPair]
		pair_t := reflect.TypeOf(pair).Kind()
		fmt.Printf("Type of Pair = %s\n", pair_t)
		(*pair).Ask = ask
		//allSpreads.Spreads[tokenPair].Ask = ask
	}
	json.NewEncoder(w).Encode(allSpreads)
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


//Get All Token Data 
func APIGetTokenPairMarket(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req) 
	baseToken := strings.ToUpper(params["base"])
	quoteToken := strings.ToUpper(params["quote"])
	tokenPair := baseToken + string('/') + quoteToken

	//Verify that token pair is supported by oasisdex
	if (!client.IsValidTokenPair(tokenPair)) {
		json.NewEncoder(w).Encode(Error{fmt.Sprintf("Unknown pair")})
		return
	}

	sPrice, sLast, sVol, sMin, sMax, sBid, sAsk, err := client.GetTokenPairMarket(baseToken, quoteToken)
	if err != nil {
		fmt.Printf(err.Error())
		json.NewEncoder(w).Encode(Error{fmt.Sprintf("Query market data for %s failed", tokenPair)})
	}
	json.NewEncoder(w).Encode(TokenPair{tokenPair, sPrice, sLast, sVol, sAsk, sBid, sMin, sMax, true, time.Now().Unix()})
	return
}

//Get MKR Token Supply
func APIGetMkrTokenSupply(w http.ResponseWriter, req *http.Request) {
	supply, err := client.GetMkrTokenSupply()
	if err != nil {
		json.NewEncoder(w).Encode(Error{fmt.Sprintf("Querying MKR token supply failed")})
	} else {
		json.NewEncoder(w).Encode(MkrTokenSupply{supply})
	}
	return
}

//Request Handler
func InitAPIServer() {

	router := mux.NewRouter()																//Create new router

	//API Endpoints
	router.HandleFunc("/v1/pairs", APIGetAllPairs).Methods("GET")							//REST endpoint for calling tradable markets
	router.HandleFunc("/v1/pairs/{base}/{quote}", APIGetPair).Methods("GET")				//REST endpoint for calling token pair data
	//router.HandleFunc("/v1/markets", APIGetAllTokenPairs).Methods("GET")					//REST endpoint for calling ticker of all token pairs
	router.HandleFunc("/v1/markets/{base}/{quote}", APIGetTokenPairMarket).Methods("GET")	//REST endpoint for calling ticker of a token pair
	//router.HandleFunc("/v1/prices", APIGetAllPrices).Methods("GET")						//REST endpoint for calling price of all token pairs
	//router.HandleFunc("/v1/prices/{base}/{quote}", APIGetTokenPairPrice).Methods("GET")	//REST endpoint for calling price of token pair
	router.HandleFunc("/v1/volumes", APIGetAllVolume).Methods("GET")						//REST endpoint for calling volume of all token pairs
	router.HandleFunc("/v1/volumes/{base}/{quote}", APIGetTokenPairVolume).Methods("GET")	//REST endpoint for calling volume of token pair
	//router.HandleFunc("/v1/spreads", APIGetAllSpread).Methods("GET")							//REST endpoint for calling spread of all token pairs
	router.HandleFunc("/v1/spreads/{base}/{quote}", APIGetTokenPairSpread).Methods("GET")	//REST endpoint for calling spread of token pair
	//router.HandleFunc("/v1/trades/{base}/{quote}", APIGetTokenPairTrades).Methods("GET")
	//router.HandleFunc("/v1/orders/{base}/{quote}", APIGetTokenPairOrders).Methods("GET")
	router.HandleFunc("/v1/tokens/mkr/totalsupply", APIGetMkrTokenSupply).Methods("GET")	//REST endpoint for calling MKR token supply

	fmt.Printf("API Server Started\nReady for incoming requests\n")

	log.Fatal(http.ListenAndServe(":8080", router))												//Deploy server
}