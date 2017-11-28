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
var TokenPairContractLib = map[string]tokenPairAddr{
	"MKRETH": tokenPairAddr{
		quoteToken: strings.ToLower("0x000000000000000000000000C66eA802717bFb9833400264Dd12c2bCeAa34a6d"),
		denomToken: strings.ToLower("0x000000000000000000000000ecf8f87f810ecf450940c9f60066b4a7a501d6a7")},
	"MKRSAI": tokenPairAddr{
		quoteToken: strings.ToLower("0x000000000000000000000000C66eA802717bFb9833400264Dd12c2bCeAa34a6d"),
		denomToken: strings.ToLower("0x00000000000000000000000059adcf176ed2f6788a41b8ea4c4904518e62b6a4")},
	"ETHSAI": tokenPairAddr{
		quoteToken: strings.ToLower("0x000000000000000000000000ecf8f87f810ecf450940c9f60066b4a7a501d6a7"),
		denomToken: strings.ToLower("0x00000000000000000000000059adcf176ed2f6788a41b8ea4c4904518e62b6a4")},
	"DGDETH": tokenPairAddr{
		quoteToken: strings.ToLower(""),
		denomToken: strings.ToLower("0x000000000000000000000000ecf8f87f810ecf450940c9f60066b4a7a501d6a7")},
	}
	"RHOCETH": tokenPairAddr{
		quoteToken: strings.ToLower(""),
		denomToken: strings.ToLower("0x000000000000000000000000ecf8f87f810ecf450940c9f60066b4a7a501d6a7")},
	}
	"REPETH": tokenPairAddr{
		quoteToken: strings.ToLower(""),
		denomToken: strings.ToLower("0x000000000000000000000000ecf8f87f810ecf450940c9f60066b4a7a501d6a7")},
	}
	"ICNETH": tokenPairAddr{
		quoteToken: strings.ToLower(""),
		denomToken: strings.ToLower("0x000000000000000000000000ecf8f87f810ecf450940c9f60066b4a7a501d6a7")},
	}
	"1STETH": tokenPairAddr{
		quoteToken: strings.ToLower(""),
		denomToken: strings.ToLower("0x000000000000000000000000ecf8f87f810ecf450940c9f60066b4a7a501d6a7")},
	}
	"GNTETH": tokenPairAddr{
		quoteToken: strings.ToLower(""),
		denomToken: strings.ToLower("0x000000000000000000000000ecf8f87f810ecf450940c9f60066b4a7a501d6a7")},
	}
	"VSLETH": tokenPairAddr{
		quoteToken: strings.ToLower(""),
		denomToken: strings.ToLower("0x000000000000000000000000ecf8f87f810ecf450940c9f60066b4a7a501d6a7")},
	}
	"PLUETH": tokenPairAddr{
		quoteToken: strings.ToLower(""),
		denomToken: strings.ToLower("0x000000000000000000000000ecf8f87f810ecf450940c9f60066b4a7a501d6a7")},
	}
	"MLNETH": tokenPairAddr{
		quoteToken: strings.ToLower(""),
		denomToken: strings.ToLower("0x000000000000000000000000ecf8f87f810ecf450940c9f60066b4a7a501d6a7")},
	}
	"NMRETH": tokenPairAddr{
		quoteToken: strings.ToLower(""),
		denomToken: strings.ToLower("0x000000000000000000000000ecf8f87f810ecf450940c9f60066b4a7a501d6a7")},
	}
	"TIMEETH": tokenPairAddr{
		quoteToken: strings.ToLower(""),
		denomToken: strings.ToLower("0x000000000000000000000000ecf8f87f810ecf450940c9f60066b4a7a501d6a7")},
	}
	"GUPETH": tokenPairAddr{
		quoteToken: strings.ToLower(""),
		denomToken: strings.ToLower("0x000000000000000000000000ecf8f87f810ecf450940c9f60066b4a7a501d6a7")},
	}
	"BATETH": tokenPairAddr{
		quoteToken: strings.ToLower(""),
		denomToken: strings.ToLower("0x000000000000000000000000ecf8f87f810ecf450940c9f60066b4a7a501d6a7")},
	}
	"SNGLSETH": tokenPairAddr{
		quoteToken: strings.ToLower(""),
		denomToken: strings.ToLower("0x000000000000000000000000ecf8f87f810ecf450940c9f60066b4a7a501d6a7")},
	}
}

type tokenPairAddr struct {
	quoteToken string
	denomToken string
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
		} else if fromBlockNum <= 0 || int(fromBlockNum64) > toBlockNum {
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

func CalculatePriceFromLogs(tokenPair string, logs []ethrpc.Log)  (string, error) {
	sumQuoteVol := big.NewInt(0)
	sumDenomVol := big.NewInt(0)
	temp := new(big.Int)

	tokenPairContracts := TokenPairContractLib[tokenPair]

	fmt.Printf("Quote token contract = %s\n", tokenPairContracts.quoteToken)
	fmt.Printf("Denom token contract = %s\n", tokenPairContracts.denomToken)

	for i, log := range logs {
		fmt.Printf("\nParsing Event Log %d\n", i)
		if (len(log.Topics) != 3) {
			fmt.Println("Skipped Log")
			//skip log - this should never happen
			continue
		} else if (log.Topics[1] == tokenPairContracts.quoteToken && log.Topics[2] == tokenPairContracts.denomToken) {
			fmt.Println("Found topic 1 is quoteToken and topic 2 is denomToken")
			//Parse data block
			if (len(log.Data) != 130) {
				fmt.Printf("Error: Log Data field should be 130 chars. Data field = %s", log.Data)
			}
			quoteVol := log.Data[2:65]
			denomVol := log.Data[66:129]
			fmt.Printf("quoteVol for log %d = %s\n", i, quoteVol)
			fmt.Printf("denomVol for log %d = %s\n", i, denomVol)

		
			temp = parser.Hex2Int(quoteVol)
			sumQuoteVol.Add(sumQuoteVol, temp)

			temp = parser.Hex2Int(denomVol)
			sumDenomVol.Add(sumDenomVol, temp)
		} else if (log.Topics[1] == tokenPairContracts.denomToken && log.Topics[2] == tokenPairContracts.quoteToken) {
			fmt.Println("Found topic 1 is denomToken and topic 2 is quoteToken")

			temp = parser.Hex2Int(log.Topics[1])
			sumDenomVol.Add(sumDenomVol, temp)

			temp = parser.Hex2Int(log.Topics[2])
			sumQuoteVol.Add(sumQuoteVol, temp)
		}
	}
	//print sum of trade quote and denom tokens
	fmt.Printf("sumDenomVol = %s\n", sumDenomVol.Text(10))
	fmt.Printf("sumQuoteVol = %s\n", sumQuoteVol.Text(10))

	//convert sums from into to float
	sumDenomVolF := new(big.Float).SetInt(sumDenomVol)
	sumQuoteVolF := new(big.Float).SetInt(sumQuoteVol)

	result := new(big.Float)

	//calculate volume weighted priced
	result.Quo(sumDenomVolF, sumQuoteVolF)
	fmt.Printf("Avg Weighted Price = %s\n", result.Text('f', 8))
	return result.Text('f', 8), nil
}

func LatestBlockNumber() (int) {
	latest, _ := EthClient.EthBlockNumber()
	return latest
}