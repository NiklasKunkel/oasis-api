package client

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"time"
	"github.com/onrik/ethrpc"
	"github.com/niklaskunkel/oasis-api/data"
	"github.com/niklaskunkel/oasis-api/parser"
)

const (
	HOST = "http://127.0.0.1:8545"
)

var EthClient = ethrpc.NewEthRPC(HOST)

func InitClient() {
	for err := VerifyClientConnection(); err != nil; err = VerifyClientConnection() {
		fmt.Printf(err.Error());
		fmt.Printf("Retrying client verification in 10 seconds")
		time.Sleep(10 * time.Second)
	}
}

func VerifyClientConnection() (error) {
	//Check client version
	if err := CheckClientVersion(); err != nil {
		return err
	}

	//Check that client is connected to peers
	if err := CheckPeerCount(); err != nil {
		return err
	}

	//Check if client is syncing
	if err := CheckClientSync(); err != nil {
		return err
	}

	//Print Block Number
	if err := GetBlockNumber(); err != nil {
		return err
	}

	//List accounts linked to client
	if err := GetEthAccounts(); err != nil {
		return err
	}
	return nil
}
func CheckClientVersion() (error) {
	//TODO add minimum client version check
	if version, err := EthClient.Web3ClientVersion(); err != nil {
		return fmt.Errorf("[CheckClientVersion()] failed due to (%s)\n", err)
	} else {
		fmt.Printf("Client Version = %s\n", version)
		return nil
	}
}

func CheckPeerCount() (error) {
	if numPeers, err := EthClient.NetPeerCount(); err != nil {
		return fmt.Errorf("[CheckPeerCount()] failed due to (%s)\n", err)
	} else if numPeers <= 0 {
		return fmt.Errorf("Error client has 0 peers\n")
	} else {
		fmt.Printf("Client has %d peers\n", numPeers)
		return nil
	}
}

func CheckClientSync() (error) {
	if syncing, err := EthClient.EthSyncing(); err != nil {
		return fmt.Errorf("[CheckClientSync()] failed due to (%s)\n", err)
	} else if syncing.IsSyncing {
		return fmt.Errorf("[CheckClientSync()] failed, client is not finished syncing\n")
	} else {
		fmt.Printf("Client is finished syncing\n")
		return nil
	}
}

func GetEthAccounts() (error) {
	var accounts []string
	accounts, err := EthClient.EthAccounts()
	if err != nil {
		return fmt.Errorf("[GetEthAccounts()] failed due to (%s)\n", err)
	}
	if len(accounts) == 0 {
		fmt.Printf("No accounts linked to client. - This means transactions cannot be signed.\n")
		return nil
	}
	fmt.Printf("\nEthereum Accounts Linked to Client:\n")
	for numAccount, account := range accounts {
		fmt.Printf("Account #%d =  %s\n", numAccount, account)
	}
	return nil
}

func GetBlockNumber() (error) {
	num, err := EthClient.EthBlockNumber()
	if err != nil {
		return fmt.Errorf("[GetBlockNumber()] failed due to (%s)\n", err)
	}
	fmt.Printf("Client Block Number = %d\n", num)
	return nil
}

func CreateTx(
		from 				string,
		to 					string,
		gas 				int,
		gasPrice 			*big.Int,
		value 				*big.Int,
		data 				string,
		nonce 				int) (ethrpc.T) {
	//verify all the different values

	//create the transaction object
	var tx = ethrpc.T{
		From: from,
		To: to,
		Gas: gas,
		GasPrice: gasPrice,
		Value: value,
		Data: data,
		Nonce: nonce,
	}
	//print tx object for debug
	fmt.Printf("%+v\n", tx)
	return tx
}

func CallTx(tx ethrpc.T, tag ...string) (string, error) {
	//change tag to ... param
	//check length of tag
	//if length of tag is 0 set tag to "latest"
	tagVal := "latest"
	if len(tag) != 0 {
		tagVal = tag[0]
	}
	data, err := EthClient.EthCall(tx, tagVal)
	if err != nil {
		return data, fmt.Errorf("[CallTx()] failed due to (%s)\n", err)
	}
	return data, nil
}

func CreateEventFilter(fromBlock string, toBlock string, address []string, topics [][]string) (ethrpc.FilterParams, error) {
	var toBlockNum int
	var fromBlockNum int
	var err error

	//Pull/Parse toBlock
	if  toBlock == "" || strings.ToLower(toBlock) == "latest" {
		toBlockNum, err = EthClient.EthBlockNumber()
		if err != nil {
			return ethrpc.FilterParams{}, fmt.Errorf("[CreateEventFilter] failed when pulling latest block from client with error (%s)\n", err)
		}
	} else {
		toBlockNum64, err := strconv.ParseInt(toBlock, 10, 32)
		if err != nil {
			return ethrpc.FilterParams{}, fmt.Errorf("[CreateEventFilter] failed due to invalid param toBlock (%s) with error (%s)\n", toBlock, err)
		} else if toBlockNum <= 0 || int(toBlockNum) > LatestBlockNumber() {
			return ethrpc.FilterParams{}, fmt.Errorf("[CreateEventFilter] failed due to invalid param toBlock (%s), block must be larger than zero and less than %d\n", toBlock, LatestBlockNumber())
		}
		toBlockNum = int(toBlockNum64)
	}
	//Parse/Calculate fromBlock
	if fromBlock == "" || strings.ToLower(fromBlock) == "24hour" {
		blockInterval := parser.Hours2Block(24)
		fromBlockNum = toBlockNum - blockInterval
	} else {
		fromBlockNum64, err := strconv.ParseInt(fromBlock, 10, 32)
		if err != nil {
			fmt.Errorf("[CreateEventFilter] failed due to invalid fromBlock (%s) with error (%s)\n", fromBlock, err)
		} else if fromBlockNum64 <= 0 || int(fromBlockNum64) > toBlockNum {
			return ethrpc.FilterParams{}, fmt.Errorf("[CreateEventFilter] failed due to invalid param fromBlock (%s), block must be larger than zero and less than %d\n", fromBlock, toBlockNum)
		}
		fromBlockNum = int(fromBlockNum64)
	}

	//Verify valid contract address(es)
	for _, contractAddr := range address {
		if (!strings.HasPrefix(contractAddr, "0x") || len(contractAddr) != 42) {
			return ethrpc.FilterParams{}, fmt.Errorf("[CreateEventsFilter] failed due to invalid contract address (%s)\n", contractAddr)
		}
	}

	//Verify topics

	//Compose Filter Params Object
	/*
	params := ethrpc.FilterParams{
		FromBlock: strconv.Itoa(fromBlockNum),
		ToBlock: strconv.Itoa(toBlockNum),
		Address: address,
		Topics: topics,
	}
	*/
	params := ethrpc.FilterParams{
		FromBlock: "0x" + strconv.FormatInt(int64(fromBlockNum), 16),
		ToBlock: "0x" + strconv.FormatInt(int64(toBlockNum), 16),
		Address: address,
		Topics: topics,
	}

	fmt.Printf("%+v\n", params)

	return params, nil
}

//subscribe filter to client
func SubscribeEventFilter(params ethrpc.FilterParams) (string, error) {
	filterID, err := EthClient.EthNewFilter(params)
	if err != nil {
		return "", fmt.Errorf("[CreateEventFilter()] failed due to (%s)\n", err)
	}
	fmt.Printf("Created Filter with filterID = %s\n", filterID)
	return filterID, nil
}

func KillEventFilter(filterId string) (bool, error) {
	status, err :=  EthClient.EthUninstallFilter(filterId)
	return status, err
}

func GetFilterLogs(filterId string) ([]ethrpc.Log, error) {
	logs, err := EthClient.EthGetFilterLogs(filterId)
	if err != nil {
		return []ethrpc.Log{}, fmt.Errorf("[GetFilterLogs] failed due to (%s)\n", err)
	}
	return logs, err
}

func GetLogs(params ethrpc.FilterParams) ([]ethrpc.Log, error) {
	logs, err := EthClient.EthGetLogs(params)
	if err != nil {
		return []ethrpc.Log{}, fmt.Errorf("[GetLogs] failed due to (%s)\n", err)
	}
	return logs, nil
}

func ExtractLogTradeData(log ethrpc.Log, index int, isBaseFirst bool, sumBaseVol *big.Int, sumQuoteVol *big.Int, min *big.Float, max *big.Float) {
	var sQuoteVol string
	var sBaseVol string
	intVol := new(big.Int)
	fBaseVol := new(big.Float)
	fQuoteVol := new(big.Float)
	fPrice := new(big.Float)


	if (len(log.Data) != 130) {
		fmt.Printf("Error: Log Data field should be 130 chars. Data field = %s", log.Data)
		return
	}
	if (isBaseFirst) {
		sBaseVol = log.Data[2:66]
		sQuoteVol = log.Data[67:130]
	} else {
		sQuoteVol = log.Data[2:66]
		sBaseVol = log.Data[67:130]
	}
	//Debug - logging
	fmt.Printf("baseVol for log %d = %s\n", index, sBaseVol)
	fmt.Printf("quoteVol for log %d = %s\n", index, sQuoteVol)


	intVol = parser.Hex2Int(sBaseVol)		//convert base token volume from hex string to integer
	sumBaseVol.Add(sumBaseVol, intVol)		//add base token volume to cumulative base Volume
	fBaseVol.SetInt(intVol)					//convert base token volume from integer to float

	intVol = parser.Hex2Int(sQuoteVol)		//convert quote token volume from hex string to integer
	sumQuoteVol.Add(sumQuoteVol, intVol)	//add quote token volume to cumulative quote Volume
	fQuoteVol.SetInt(intVol)				//convert quote token volume from integer to float

	fPrice.Quo(fQuoteVol, fBaseVol)			//calculate price of base token in reference to quote token
	if ((min.Cmp(fPrice) == 1) || (min.Sign() == 0)) { //if price is lower than minimum price
		min.Set(fPrice)						//set new minimum price
	}
	if (max.Cmp(fPrice) == -1) {			//if price is higer than maximumum price
		max.Set(fPrice)						//set new maximum price
	}

}

func CalculatePriceFromLogs(logs []ethrpc.Log, baseToken string, quoteToken string)  (string, string, string, string, error) {
	sumBaseVol := big.NewInt(0)
	sumQuoteVol := big.NewInt(0)
	price := new(big.Float)
	max := new(big.Float).SetInt(sumBaseVol)
 	min := new(big.Float).SetInt(sumBaseVol)	//Base Volume is used here to avoid instantiating new big.Int

	baseTokenContract := data.TokenInfoLib[baseToken].Contract
	quoteTokenContract := data.TokenInfoLib[quoteToken].Contract

	fmt.Printf("Denom token contract = %s\n", baseTokenContract)
	fmt.Printf("Quote token contract = %s\n", quoteTokenContract)

	foundLog := false
	for i, log := range logs {
		//fmt.Printf("\nParsing Event Log %d\n", i)
		if (len(log.Topics) != 3) {
			fmt.Println("Skipped Log")
			//skip log - this should never happen
			continue
		} else if (log.Topics[1] == baseTokenContract && log.Topics[2] == quoteTokenContract) {
			foundLog = true
			fmt.Println("Found topic 1 is baseToken and topic 2 is quoteToken")
			ExtractLogTradeData(log, i, true, sumBaseVol, sumQuoteVol, min, max)
		} else if (log.Topics[1] == quoteTokenContract && log.Topics[2] == baseTokenContract) {
			foundLog = true
			fmt.Println("Found topic 1 is quoteToken and topic 2 is baseToken")
			ExtractLogTradeData(log, i, false, sumBaseVol, sumQuoteVol, min, max)
		}
	}
	if (foundLog == false) {
		//found no relevant logs for this trading pair
		return "null", "0", "null", "null", nil
	}

	//Debug - print sum of trade quote and denom tokens
	fmt.Printf("sumBaseVol = %s\n", sumBaseVol.Text(10))
	fmt.Printf("sumQuoteVol = %s\n", sumQuoteVol.Text(10))

	//Adjust for token precision
	adjustedSumBaseVol := parser.AdjustIntForPrecision(sumBaseVol, data.TokenInfoLib[baseToken].Precision)
	adjustedSumQuoteVol := parser.AdjustIntForPrecision(sumQuoteVol, data.TokenInfoLib[quoteToken].Precision)
	min = parser.AdjustFloatForPrecision(min, GetPrecisionDelta(baseToken, quoteToken))
	max = parser.AdjustFloatForPrecision(max, GetPrecisionDelta(baseToken, quoteToken))

	if (adjustedSumBaseVol == new(big.Float)) {
		//cant divide by zero
		//this should never happen due to foundLog check
		return "null", "0", "null", "null", nil
	}

	//calculate volume weighted priced
	price.Quo(adjustedSumQuoteVol, adjustedSumBaseVol)

	//Debug - print volume weighted price
	fmt.Printf("Volume Weighted Price = %s\n", price.Text('f', 8))
	fmt.Printf("Volume = %s\n", adjustedSumBaseVol.Text('f', 8))
	fmt.Printf("Max price = %s\n", max.Text('f', 8))
	fmt.Printf("Min price = %s\n", min.Text('f', 8))

	return price.Text('f', 8), adjustedSumBaseVol.Text('f', 8), min.Text('f', 8), max.Text('f', 8), nil
}

func GetSpread(baseToken string, quoteToken string) (string, string, error) {
	fBid, err := GetBestOffer(baseToken, quoteToken, "bid")
	if err != nil {
		return "","",fmt.Errorf("[GetSpread] failed to get Bid due to (%s)\n", err)
	}

	fAsk, err := GetBestOffer(baseToken, quoteToken, "ask")
	if err != nil {
		return "","",fmt.Errorf("[GetSpread] failed to get Ask due to (%s)\n", err)
	}
	return fBid, fAsk, err
}

func GetBestOffer(baseToken string, quoteToken string, otype string) (string, error) {
	fBestOffer := new(big.Float)
	var calldata string

	//construct calldata
	switch otype {
	case "ask":
		calldata = "0x0374fc6f" + data.TokenInfoLib[strings.ToUpper(baseToken)].Contract[2:] + data.TokenInfoLib[strings.ToUpper(quoteToken)].Contract[2:]
	case "bid":
		calldata = "0x0374fc6f" + data.TokenInfoLib[strings.ToUpper(quoteToken)].Contract[2:] + data.TokenInfoLib[strings.ToUpper(baseToken)].Contract[2:]
	default:
		return "", fmt.Errorf("[GetBestOffer] failed due to invalid order type param\n")
	}

	//debug
	fmt.Printf("Calldata = %s\n", calldata)

	//create tx object for querying id of best offer
	tx := CreateTx(
		"0x003EbC0613139A8dF37CAC03d39B39304153596A",
		"0x3Aa927a97594c3ab7d7bf0d47C71c3877D1DE4A1",
		0,
		big.NewInt(0),
		big.NewInt(0),
		calldata,
		0)
	//call tx to grab id of best offer 
	offerid, err := CallTx(tx)
	if err != nil {
		return "", fmt.Errorf("[GetBestOffer] failed to obtain offer id due to (%s)\n", err)
	}
	//no offer in this side of orderbook
	if (offerid == "0x0000000000000000000000000000000000000000000000000000000000000000") {
		return "null", nil
	}

	//debug 
	fmt.Printf("Offer Id = %s\n", offerid)
	//construct calldata
	calldata = "0x4579268a" + offerid[2:]
	//create new tx object for querying best offer
	tx = CreateTx(
		"0x003EbC0613139A8dF37CAC03d39B39304153596A",
		"0x3Aa927a97594c3ab7d7bf0d47C71c3877D1DE4A1",
		0,
		big.NewInt(0),
		big.NewInt(0),
		calldata,
		0)
	//call tx to grab best offer
	offer, err := CallTx(tx)
	if err != nil {
		return "", fmt.Errorf("[GetBestOffer] failed to obtain offer due to (%s)\n", err)
	}

	//Parse response into order amounts
	var sQuoteTokenAmt string
	var sBaseTokenAmt string
	switch otype {
	case "bid":
		sQuoteTokenAmt = offer[2:66]
		sBaseTokenAmt = offer[130:194]
	case "ask":
		sBaseTokenAmt = offer[2:66]
		sQuoteTokenAmt = offer[130:194]
	default:
		return "", fmt.Errorf("[GetBestOffer] failed due to invalid order type param\n")
	}

	//debug
	fmt.Printf("Base Token Amount = %s\n", sBaseTokenAmt)
	fmt.Printf("Quote Token Amount = %s\n", sQuoteTokenAmt)

	//convert from string to int
	intBaseTokenAmt := parser.Hex2Int(sBaseTokenAmt)
	intQuoteTokenAmt:= parser.Hex2Int(sQuoteTokenAmt)

	//debug
	fmt.Printf("[PostHex2Int] Base Token Amount = %s\n", intBaseTokenAmt.Text(10))
	fmt.Printf("[PostHex2Int] Quote Token Amount = %s\n", intQuoteTokenAmt.Text(10))

	//adjust for precision
	adjustedBaseTokenAmt := parser.AdjustIntForPrecision(intBaseTokenAmt, data.TokenInfoLib[baseToken].Precision)
	adjustedQuoteTokenAmt := parser.AdjustIntForPrecision(intQuoteTokenAmt, data.TokenInfoLib[quoteToken].Precision)

	//debug
	fmt.Printf("[Post-AdjustIntForPrecision] Base Token Amount = %s\n", adjustedBaseTokenAmt.Text('f', 8))
	fmt.Printf("[Post-AdjustIntForPrecision] Quote Token Amount = %s\n", adjustedQuoteTokenAmt.Text('f', 8))

	//calculate price
	fBestOffer.Quo(adjustedQuoteTokenAmt, adjustedBaseTokenAmt)

	return fBestOffer.Text('f', 8), nil
}

func GetMarkets() ([]data.Market) {
	//loop through each market in static data
	for index, market := range data.LiveMarkets {
		//construct calldata
		baseTokenAddress := data.TokenInfoLib[market.Base].Contract
		quoteTokenAddress := data.TokenInfoLib[market.Quote].Contract
		calldata := "0x8d7daf95" + baseTokenAddress[2:] + quoteTokenAddress[2:]
		//create transaction
		tx := CreateTx(
			"0x003EbC0613139A8dF37CAC03d39B39304153596A",
			"0x3Aa927a97594c3ab7d7bf0d47C71c3877D1DE4A1",
			0,
			big.NewInt(0),
			big.NewInt(0),
			calldata,
			0)
		//check if market is active
		status, err := CallTx(tx)
		if err != nil {
			fmt.Printf("[GetMarkets] failed to query whitelist status of baseToken: %s and quoteToken: %s\n due to %s", market.Base, market.Quote, err)
		}
		//update market status
		if (status == "0x0000000000000000000000000000000000000000000000000000000000000001") {
			data.LiveMarkets[index].Active = true
		} else {
			data.LiveMarkets[index].Active = false
		}
		//update timestamp
		data.LiveMarkets[index].Time = time.Now().Unix()
	}
	return data.LiveMarkets
}

func GetMkrTokenSupply() (string, error) {
	tx := CreateTx(
		"0x003EbC0613139A8dF37CAC03d39B39304153596A",
		"0xc66ea802717bfb9833400264dd12c2bceaa34a6d",
		0,
		big.NewInt(0),
		big.NewInt(0),
		"0x18160ddd",
		0)
	hSupply, err := CallTx(tx)	//get supply in hex string format
	if err != nil {
		return "", fmt.Errorf("[GetMkrTokenSupply] failed to query supply due to error (%s)\n", err)
	} else {
		iSupply := parser.Hex2Int(hSupply)
		fSupply := parser.AdjustIntForPrecision(iSupply, 18)
		return fSupply.Text('f',0), nil
	}
}

func LatestBlockNumber() (int) {
	latest, _ := EthClient.EthBlockNumber()
	return latest
}

func GetPrecisionDelta(baseToken string, quoteToken string) (int) {
	return data.TokenInfoLib[quoteToken].Precision - data.TokenInfoLib[baseToken].Precision
}