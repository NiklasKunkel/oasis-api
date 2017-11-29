package client

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"time"
	"github.com/onrik/ethrpc"
	"github.com/niklaskunkel/oasis-api/parser"
)

const (
	HOST = "http://127.0.0.1:8545"
)

var EthClient = ethrpc.NewEthRPC(HOST)

var TokenInfoLib = map[string]tokenInfo{
	"MKR": tokenInfo{
		contract: strings.ToLower("0x000000000000000000000000C66eA802717bFb9833400264Dd12c2bCeAa34a6d"),
		precision: 18},
	"ETH": tokenInfo{
		contract: strings.ToLower("0x000000000000000000000000ecf8f87f810ecf450940c9f60066b4a7a501d6a7"),
		precision: 18},
	"SAI": tokenInfo{
		contract: strings.ToLower("0x00000000000000000000000059adcf176ed2f6788a41b8ea4c4904518e62b6a4"),
		precision: 18},
	"DGD": tokenInfo{
		contract: strings.ToLower("0x000000000000000000000000e0b7927c4af23765cb51314a0e0521a9645f0e2a"),
		precision: 9},
	"RHOC": tokenInfo{
		contract: strings.ToLower("0x000000000000000000000000168296bb09e24a88805cb9c33356536b980d3fc5"),
		precision: 8},
	"REP": tokenInfo{
		contract: strings.ToLower("0x000000000000000000000000e94327d07fc17907b4db788e5adf2ed424addff6"),
		precision: 18},
	"ICN": tokenInfo{
		contract: strings.ToLower("0x000000000000000000000000888666ca69e0f178ded6d75b5726cee99a87d698"),
		precision: 18},
	"1ST": tokenInfo{
		contract: strings.ToLower("0x000000000000000000000000af30d2a7e90d7dc361c8c4585e9bb7d2f6f15bc7"),
		precision: 18},
	"GNT": tokenInfo{
		contract: strings.ToLower("0x00000000000000000000000001afc37f4f85babc47c0e2d0eababc7fb49793c8"),
		precision: 18},
	"VSL": tokenInfo{
		contract: strings.ToLower("0x0000000000000000000000005c543e7ae0a1104f78406c340e9c64fd9fce5170"),
		precision: 18},
	"PLU": tokenInfo{
		contract: strings.ToLower("0x000000000000000000000000d8912c10681d8b21fd3742244f44658dba12264e"),
		precision: 18},
	"MLN": tokenInfo{
		contract: strings.ToLower("0x000000000000000000000000beb9ef514a379b997e0798fdcc901ee474b6d9a1"),
		precision: 18},
	"NMR": tokenInfo{
		contract: strings.ToLower("0x0000000000000000000000001776e1f26f98b1a5df9cd347953a26dd3cb46671"),
		precision: 18},
	"TIME": tokenInfo{
		contract: strings.ToLower("0x0000000000000000000000006531f133e6deebe7f2dce5a0441aa7ef330b4e53"),
		precision: 8},
	"GUP": tokenInfo{
		contract: strings.ToLower("0x000000000000000000000000f7b098298f7c69fc14610bf71d5e02c60792894c"),
		precision: 3},
	"BAT": tokenInfo{
		contract: strings.ToLower("0x0000000000000000000000000d8775f648430679a709e98d2b0cb6250d2887ef"),
		precision: 18},
	"SNGLS": tokenInfo{
		contract: strings.ToLower("0x000000000000000000000000aec2e87e0a235266d9c5adc9deb4b2e29b54d009"),
		precision: 0},
}

type tokenInfo struct {
	contract string
	precision int
}

var TokenPairContractLib = map[string]tokenPairInfo{
	"MKRETH": tokenPairInfo{
		baseToken: strings.ToLower("0x000000000000000000000000C66eA802717bFb9833400264Dd12c2bCeAa34a6d"),
		quoteToken: strings.ToLower("0x000000000000000000000000ecf8f87f810ecf450940c9f60066b4a7a501d6a7"),
		precisionDelta: 0},
	"MKRSAI": tokenPairInfo{
		baseToken: strings.ToLower("0x000000000000000000000000C66eA802717bFb9833400264Dd12c2bCeAa34a6d"),
		quoteToken: strings.ToLower("0x00000000000000000000000059adcf176ed2f6788a41b8ea4c4904518e62b6a4"),
		precisionDelta: 0},
	"ETHSAI": tokenPairInfo{
		baseToken: strings.ToLower("0x000000000000000000000000ecf8f87f810ecf450940c9f60066b4a7a501d6a7"),
		quoteToken: strings.ToLower("0x00000000000000000000000059adcf176ed2f6788a41b8ea4c4904518e62b6a4"),
		precisionDelta: 0},
	"DGDETH": tokenPairInfo{
		baseToken: strings.ToLower("0x000000000000000000000000e0b7927c4af23765cb51314a0e0521a9645f0e2a"),
		quoteToken: strings.ToLower("0x000000000000000000000000ecf8f87f810ecf450940c9f60066b4a7a501d6a7"),
		precisionDelta: 9},
	"RHOCETH": tokenPairInfo{
		baseToken: strings.ToLower("0x000000000000000000000000168296bb09e24a88805cb9c33356536b980d3fc5"),
		quoteToken: strings.ToLower("0x000000000000000000000000ecf8f87f810ecf450940c9f60066b4a7a501d6a7"),
		precisionDelta: 10},
	"REPETH": tokenPairInfo{
		baseToken: strings.ToLower("0x000000000000000000000000e94327d07fc17907b4db788e5adf2ed424addff6"),
		quoteToken: strings.ToLower("0x000000000000000000000000ecf8f87f810ecf450940c9f60066b4a7a501d6a7"),
		precisionDelta: 0},
	"ICNETH": tokenPairInfo{
		baseToken: strings.ToLower("0x000000000000000000000000888666ca69e0f178ded6d75b5726cee99a87d698"),
		quoteToken: strings.ToLower("0x000000000000000000000000ecf8f87f810ecf450940c9f60066b4a7a501d6a7"),
		precisionDelta: 0},
	"1STETH": tokenPairInfo{
		baseToken: strings.ToLower("0x000000000000000000000000af30d2a7e90d7dc361c8c4585e9bb7d2f6f15bc7"),
		quoteToken: strings.ToLower("0x000000000000000000000000ecf8f87f810ecf450940c9f60066b4a7a501d6a7"),
		precisionDelta: 0},
	"GNTETH": tokenPairInfo{
		baseToken: strings.ToLower("0x00000000000000000000000001afc37f4f85babc47c0e2d0eababc7fb49793c8"),
		quoteToken: strings.ToLower("0x000000000000000000000000ecf8f87f810ecf450940c9f60066b4a7a501d6a7"),
		precisionDelta: 0},
	"VSLETH": tokenPairInfo{
		baseToken: strings.ToLower("0x0000000000000000000000005c543e7ae0a1104f78406c340e9c64fd9fce5170"),
		quoteToken: strings.ToLower("0x000000000000000000000000ecf8f87f810ecf450940c9f60066b4a7a501d6a7"),
		precisionDelta: 0},
	"PLUETH": tokenPairInfo{
		baseToken: strings.ToLower("0x000000000000000000000000d8912c10681d8b21fd3742244f44658dba12264e"),
		quoteToken: strings.ToLower("0x000000000000000000000000ecf8f87f810ecf450940c9f60066b4a7a501d6a7"),
		precisionDelta: 0},
	"MLNETH": tokenPairInfo{
		baseToken: strings.ToLower("0x000000000000000000000000beb9ef514a379b997e0798fdcc901ee474b6d9a1"),
		quoteToken: strings.ToLower("0x000000000000000000000000ecf8f87f810ecf450940c9f60066b4a7a501d6a7"),
		precisionDelta: 0},
	"NMRETH": tokenPairInfo{
		baseToken: strings.ToLower("0x0000000000000000000000001776e1f26f98b1a5df9cd347953a26dd3cb46671"),
		quoteToken: strings.ToLower("0x000000000000000000000000ecf8f87f810ecf450940c9f60066b4a7a501d6a7"),
		precisionDelta: 0},
	"TIMEETH": tokenPairInfo{
		baseToken: strings.ToLower("0x0000000000000000000000006531f133e6deebe7f2dce5a0441aa7ef330b4e53"),
		quoteToken: strings.ToLower("0x000000000000000000000000ecf8f87f810ecf450940c9f60066b4a7a501d6a7"),
		precisionDelta: 10},
	"GUPETH": tokenPairInfo{
		baseToken: strings.ToLower("0x000000000000000000000000f7b098298f7c69fc14610bf71d5e02c60792894c"),
		quoteToken: strings.ToLower("0x000000000000000000000000ecf8f87f810ecf450940c9f60066b4a7a501d6a7"),
		precisionDelta: 15},
	"BATETH": tokenPairInfo{
		baseToken: strings.ToLower("0x0000000000000000000000000d8775f648430679a709e98d2b0cb6250d2887ef"),
		quoteToken: strings.ToLower("0x000000000000000000000000ecf8f87f810ecf450940c9f60066b4a7a501d6a7"),
		precisionDelta: 0},
	"SNGLSETH": tokenPairInfo{
		baseToken: strings.ToLower("0x000000000000000000000000aec2e87e0a235266d9c5adc9deb4b2e29b54d009"),
		quoteToken: strings.ToLower("0x000000000000000000000000ecf8f87f810ecf450940c9f60066b4a7a501d6a7"),
		precisionDelta: 18},
}

type tokenPairInfo struct {
	baseToken string
	quoteToken string
	precisionDelta int
}

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
			fmt.Errorf("[CreateEventsFilter] failed due to invalid contract address (%s)\n", contractAddr)
		}
	}

	//Verify topics

	//Compose Filter Params Object
	params := ethrpc.FilterParams{
		FromBlock: strconv.Itoa(fromBlockNum),
		ToBlock: strconv.Itoa(toBlockNum),
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
	fmt.Printf("\n[GetFilterLogs] Event Logs:\n")
	for _, log := range logs {
		fmt.Printf("%+v", log)
	}
	return logs, err
}

func GetLogs(params ethrpc.FilterParams) ([]ethrpc.Log, error) {
	logs, err := EthClient.EthGetLogs(params)
	if err != nil {
		return []ethrpc.Log{}, fmt.Errorf("[GetLogs] failed due to (%s)\n", err)
	}
	fmt.Printf("\n[GetLog] Event Logs\n")
	for _, log := range logs {
		fmt.Printf("%+v\n", log)
	}
	return logs, nil
}

func ExtractLogTradeData(log ethrpc.Log, index int, isBaseFirst bool, sumBaseVol *big.Int, sumQuoteVol *big.Int, min *big.Float, max *big.Float) {
	var sQuoteVol string
	var sBaseVol string
	intVol := new(big.Int)
	fBaseVol := new(big.Float)
	fQuoteVol := new(big.Float)

	if (len(log.Data) != 130) {
		fmt.Printf("Error: Log Data field should be 130 chars. Data field = %s", log.Data)
		return
	}
	if (isBaseFirst) {
		sBaseVol = log.Data[2:65]
		sQuoteVol = log.Data[66:129]
	} else {
		sQuoteVol = log.Data[2:65]
		sBaseVol = log.Data[66:129]
	}
	//Debug - logging
	fmt.Printf("baseVol for log %d = %s\n", index, sBaseVol)
	fmt.Printf("quoteVol for log %d = %s\n", index, sQuoteVol)


	intVol = parser.Hex2Int(sQuoteVol)			//convert quote token volume from hex string to integer
	sumQuoteVol.Add(sumQuoteVol, intVol)		//add quote token volume to cumulative quote Volume
	fQuoteVol.SetInt(intVol)					//convert quote token volume from integer to float

	intVol = parser.Hex2Int(sBaseVol)			//convert base token volume from hex string to integer
	sumBaseVol.Add(sumBaseVol, intVol)			//add base token volume to cumulative quote Volume
	fBaseVol.SetInt(intVol)						//convert base token volume from integer to float

	fQuoteVol.Quo(fQuoteVol, fBaseVol)			//calculate price of base token in reference to quote token
	if ((min.Cmp(fQuoteVol) == 1) || (min.Sign() == 0)) {//if price is lower than minimum price
		min.Set(fQuoteVol)						//set new minimum price
	}
	if (max.Cmp(fQuoteVol) == -1) {				//if price is higer than maximumum price
		max.Set(fQuoteVol)						//set new maximum price
	}

}

func CalculatePriceFromLogs(logs []ethrpc.Log, baseToken string, quoteToken string)  (string, error) {
	sumBaseVol := big.NewInt(0)
	sumQuoteVol := big.NewInt(0)
	price := new(big.Float)
	max := new(big.Float).SetInt(sumBaseVol)
 	min := new(big.Float).SetInt(sumBaseVol)	//Base Volume is used here to avoid instantiating new big.Int

	baseTokenContract := TokenInfoLib[baseToken].contract
	quoteTokenContract := TokenInfoLib[quoteToken].contract

	fmt.Printf("Denom token contract = %s\n", baseTokenContract)
	fmt.Printf("Quote token contract = %s\n", quoteTokenContract)

	for i, log := range logs {
		fmt.Printf("\nParsing Event Log %d\n", i)
		if (len(log.Topics) != 3) {
			fmt.Println("Skipped Log")
			//skip log - this should never happen
			continue
		} else if (log.Topics[1] == baseTokenContract && log.Topics[2] == quoteTokenContract) {
			fmt.Println("Found topic 1 is quoteToken and topic 2 is denomToken")
			ExtractLogTradeData(log, i, true, sumBaseVol, sumQuoteVol, min, max)
		} else if (log.Topics[1] == quoteTokenContract && log.Topics[2] == baseTokenContract) {
			fmt.Println("Found topic 1 is denomToken and topic 2 is quoteToken")
			ExtractLogTradeData(log, i, false, sumBaseVol, sumQuoteVol, min, max)
		}
	}
	//Debug - print sum of trade quote and denom tokens
	fmt.Printf("sumBaseVol = %s\n", sumBaseVol.Text(10))
	fmt.Printf("sumQuoteVol = %s\n", sumQuoteVol.Text(10))

	//convert sums from into to float
	fSumBaseVol := new(big.Float).SetInt(sumBaseVol)
	fSumQuoteVol := new(big.Float).SetInt(sumQuoteVol)

	//calculate volume weighted priced
	price.Quo(fSumQuoteVol, fSumBaseVol)

	//adjust for token precision differences


	//Debug - print volume weighted price
	fmt.Printf("Volume Weighted Price = %s\n", price.Text('f', 8))
	fmt.Printf("Max price = %s\n", max.Text('f', 8))
	fmt.Printf("Min price = %s\n", min.Text('f', 8))

	return price.Text('f', 8), nil
}

func LatestBlockNumber() (int) {
	latest, _ := EthClient.EthBlockNumber()
	return latest
}