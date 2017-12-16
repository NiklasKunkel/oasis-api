package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"log"
	//"reflect"
	"strings"
	"time"
	"github.com/gorilla/mux"
	"github.com/niklaskunkel/oasis-api/client"
	"github.com/niklaskunkel/oasis-api/data"
)

func APIGetAllPairs(w http.ResponseWriter, req *http.Request) {
	allPairs := client.GetAllPairs()
	json.NewEncoder(w).Encode(Response{allPairs, time.Now().Unix(), "success"})
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

	pair, err := client.GetPair(baseToken, quoteToken)
	if err != nil {
		json.NewEncoder(w).Encode(Error{fmt.Sprintf("Querying market for %s failed", tokenPair)})
		return
	}

	json.NewEncoder(w).Encode(Response{pair, time.Now().Unix(), "success"})
}

func APIGetAllMarkets(w http.ResponseWriter, req *http.Request) {
	allMarkets := Response{make(AllMarkets), time.Now().Unix(), "null"}
	//iterate over all token pairs
	for tokenPair, tokenPairInfo := range data.LiveMarkets {
		sPrice, sLast, sVol, sMin, sMax, sBid, sAsk, err := client.GetTokenPairMarket(tokenPairInfo.Base, tokenPairInfo.Quote)
		if err != nil {
			//call for pulling market data for token pair failed
			//one token pair failure should not make the whole call fail
			//give null data for this token pair and continue
			fmt.Printf(err.Error())
			allMarkets.Data.(AllMarkets)[tokenPair] = TokenPairMarket{tokenPair, "null", "null", "null", "null", "null", "null", "null", true}
		}
		//push token pair market data
		allMarkets.Data.(AllMarkets)[tokenPair] = TokenPairMarket{tokenPair, sPrice, sLast, sVol, sAsk, sBid, sMin, sMax, true}
	}
	if (len(allMarkets.Data.(AllMarkets)) == 0) {
		allMarkets.Message = "Querying market data for all token pairs failed"
		json.NewEncoder(w).Encode(allMarkets)
	}
	allMarkets.Message = "success"
	json.NewEncoder(w).Encode(allMarkets)
}


func APIGetTokenPairMarket(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req) 
	baseToken := strings.ToUpper(params["base"])
	quoteToken := strings.ToUpper(params["quote"])
	tokenPair := baseToken + string('/') + quoteToken

	//Verify that token pair is supported by oasisdex
	if (!client.IsValidTokenPair(tokenPair)) {
		json.NewEncoder(w).Encode(Response{TokenPairMarket{}, time.Now().Unix(), "Unknown pair"})
		return
	}

	sPrice, sLast, sVol, sMin, sMax, sBid, sAsk, err := client.GetTokenPairMarket(baseToken, quoteToken)
	if err != nil {
		fmt.Printf(err.Error())
		json.NewEncoder(w).Encode(Response{TokenPairMarket{}, time.Now().Unix(), fmt.Sprintf("Querying market data for %s failed", tokenPair)})
	}
	json.NewEncoder(w).Encode(Response{TokenPairMarket{tokenPair, sPrice, sLast, sVol, sAsk, sBid, sMin, sMax, true}, time.Now().Unix(), "success"})
	return
}

func APIGetAllPrices(w http.ResponseWriter, req *http.Request) {
	allPrices := Response{make(AllPrices), time.Now().Unix(), "null"}
	for tokenPair, tokenPairInfo := range data.LiveMarkets {
		prices := []string{"null", "null", "null", "null", "null"}
		for index, interval := range []int{1, 6, 12, 24} {
			vwap, last, err := client.GetTokenPairVolumeWeightedPrice(tokenPairInfo.Base, tokenPairInfo.Quote, interval)
			if err != nil {
				//log error
				fmt.Printf(err.Error())
				continue
			}
			prices[index] = vwap
			prices[4] = last
		}
		allPrices.Data.(AllPrices)[tokenPair] = TokenPairPrices{prices[3], prices[2], prices[1], prices[0], prices[4]}
	}
	//check if prices is empty
	if (len(allPrices.Data.(AllPrices)) == 0) {
		allPrices.Message = "Querying prices for all pairs failed"
		json.NewEncoder(w).Encode(allPrices)
	}
	allPrices.Message = "success"
	json.NewEncoder(w).Encode(allPrices)
	return
}

func APIGetTokenPairPrice(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	baseToken := strings.ToUpper(params["base"])
	quoteToken := strings.ToUpper(params["quote"])
	tokenPair := baseToken + string('/') + quoteToken

	//Verify that token pair is supported by oasisdex
	if (!client.IsValidTokenPair(tokenPair)) {
		json.NewEncoder(w).Encode(Response{TokenPairPrices{}, time.Now().Unix(), "Unknown pair"})
		return
	}
	prices := []string{"null", "null", "null", "null", "null"}
	for index, interval := range []int{1, 6, 12, 24} {
		vwap, last, err := client.GetTokenPairVolumeWeightedPrice(baseToken, quoteToken, interval)
		if err != nil {
			//log error
			fmt.Printf(err.Error())
			continue
		}
		prices[index] = vwap
		if(last != "") {
			prices[4] = last
		}
	}
	json.NewEncoder(w).Encode(Response{TokenPairPrices{prices[3], prices[2], prices[1], prices[0], prices[4]}, time.Now().Unix(), "success"})
	return
}

func APIGetAllVolume(w http.ResponseWriter, req *http.Request) {
	allVolumes := Response{make(AllVolumes), time.Now().Unix(), "null"}
	for tokenPair, tokenPairInfo := range data.LiveMarkets {
		vol, err := client.GetTokenPairVolume(tokenPairInfo.Base, tokenPairInfo.Quote)
		if err != nil {
			vol = "null"
		}
		allVolumes.Data.(AllVolumes)[tokenPair] = TokenPairVolume{vol}
	}
	allVolumes.Message = "success"
	json.NewEncoder(w).Encode(allVolumes)
}

func APIGetTokenPairVolume(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	baseToken := strings.ToUpper(params["base"])
	quoteToken := strings.ToUpper(params["quote"])
	if (!client.IsValidTokenPair(baseToken + string("/") + quoteToken)) {
		json.NewEncoder(w).Encode(Response{TokenPairVolume{}, time.Now().Unix(), "Unknown token pair"})
		return
	}
	vol, err := client.GetTokenPairVolume(baseToken, quoteToken)
	if err != nil {
		json.NewEncoder(w).Encode(Response{TokenPairVolume{}, time.Now().Unix(), "Querying Volume Failed"})
		return
	}
	json.NewEncoder(w).Encode(Response{TokenPairVolume{vol}, time.Now().Unix(), "success"})
	return
}

func APIGetAllSpread(w http.ResponseWriter, req *http.Request) {
	allSpreads := Response{make(AllSpreads), time.Now().Unix(), "null"}
	for tokenPair, tokenPairInfo := range data.LiveMarkets {
		baseToken := tokenPairInfo.Base
		quoteToken := tokenPairInfo.Quote
		bid, ask, err := client.GetSpread(baseToken, quoteToken)
		if err != nil {
			fmt.Printf("[APIGetAllSpread] failed to get spread for %s due to (%s)\n", tokenPair, err)
			bid = "null"
			ask = "null"
		}
		fmt.Printf("Getting Spread For %s : Bid = %s | Ask = %s\n", tokenPair, bid, ask)

		allSpreads.Data.(AllSpreads)[tokenPair] = TokenPairSpread{bid, ask}
	}
	allSpreads.Message = "success"
	json.NewEncoder(w).Encode(allSpreads)
}

func APIGetTokenPairSpread(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	baseToken := strings.ToUpper(params["base"])
	quoteToken := strings.ToUpper(params["quote"])

	if (!client.IsValidTokenPair(baseToken + string("/") + quoteToken)) {
		json.NewEncoder(w).Encode(Error{fmt.Sprintf("Unknown token pair")})
		return
	}
	bid, ask, err := client.GetSpread(baseToken, quoteToken)
	if err != nil {
		fmt.Printf("[GetBidAskSpread] failed due to (%s)\n", err)
		json.NewEncoder(w).Encode(Response{TokenPairSpread{}, time.Now().Unix(), "Querying bid/ask spread failed"})
		return
	}
	json.NewEncoder(w).Encode(Response{TokenPairSpread{bid,ask}, time.Now().Unix(), "success"})
}

func APIGetTokenPairTradeHistory(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	baseToken := strings.ToUpper(params["base"])
	quoteToken := strings.ToUpper(params["quote"])

	if (!client.IsValidTokenPair(baseToken + string("/") + quoteToken)) {
		json.NewEncoder(w).Encode(Response{TokenPairTradeHistory{}, time.Now().Unix(), "Unknown token pair"})
		return
	}

	tradeHistory, err := client.GetTokenPairTradeHistory(baseToken, quoteToken)
	if err != nil {
		json.NewEncoder(w).Encode(Response{TokenPairTradeHistory{}, time.Now().Unix(), "Querying trade history failed"})
		return
	}
	if (len(tradeHistory) == 0) {
		json.NewEncoder(w).Encode(Response{TokenPairTradeHistory{}, time.Now().Unix(), "No recent trade history"})
		return
	}
	json.NewEncoder(w).Encode(Response{tradeHistory, time.Now().Unix(), "success"})
	return
}

//Get MKR Token Supply
func APIGetMkrTokenSupply(w http.ResponseWriter, req *http.Request) {
	supply, err := client.GetMkrTokenSupply()
	if err != nil {
		json.NewEncoder(w).Encode(Error{fmt.Sprintf("Querying MKR token supply failed")})
	} else {
		json.NewEncoder(w).Encode(MkrTokenSupply{supply,supply, time.Now().Unix()})
		
	}
	return
}

//Request Handler
func InitAPIServer() {

	router := mux.NewRouter()																//Create new router

	//API Endpoints
	router.HandleFunc("/v1/pairs/", APIGetAllPairs).Methods("GET")							//REST endpoint for calling all tradable token pairs
	router.HandleFunc("/v1/pairs/{base}/{quote}", APIGetPair).Methods("GET")				//REST endpoint for calling tradable token pairs
	router.HandleFunc("/v1/markets/", APIGetAllMarkets).Methods("GET")						//REST endpoint for calling market data of all token pairs
	router.HandleFunc("/v1/markets/{base}/{quote}", APIGetTokenPairMarket).Methods("GET")	//REST endpoint for calling market data of a token pair
	router.HandleFunc("/v1/prices/", APIGetAllPrices).Methods("GET")						//REST endpoint for calling price of all token pairs
	router.HandleFunc("/v1/prices/{base}/{quote}", APIGetTokenPairPrice).Methods("GET")		//REST endpoint for calling price of token pair
	router.HandleFunc("/v1/volumes/", APIGetAllVolume).Methods("GET")						//REST endpoint for calling volume of all token pairs
	router.HandleFunc("/v1/volumes/{base}/{quote}", APIGetTokenPairVolume).Methods("GET")	//REST endpoint for calling volume of token pair
	router.HandleFunc("/v1/spreads/", APIGetAllSpread).Methods("GET")						//REST endpoint for calling spread of all token pairs
	router.HandleFunc("/v1/spreads/{base}/{quote}", APIGetTokenPairSpread).Methods("GET")	//REST endpoint for calling spread of token pair
	router.HandleFunc("/v1/trades/{base}/{quote}", APIGetTokenPairTradeHistory).Methods("GET")	//REST endpoint for calling trade history of token pair
	//router.HandleFunc("/v1/orders/{base}/{quote}", APIGetTokenPairOrders).Methods("GET")	//not implemented yet
	router.HandleFunc("/v1/tokens/mkr/totalsupply", APIGetMkrTokenSupply).Methods("GET")	//REST endpoint for calling MKR token supply

	fmt.Printf("API Server Started\nReady for incoming requests\n")

	log.Fatal(http.ListenAndServe(":8080", router))												//Deploy server
}