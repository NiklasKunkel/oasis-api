package data

import (
"encoding/json"
"io/ioutil"
"log"
"path/filepath"
"strings"
)

type OasisMarket struct {
	Contract 	string 	`json:"OASIS_CONTRACT_ADDRESS"`
}

var OASIS OasisMarket

func ReadConfig() {
	absPath, _ := filepath.Abs("./src/github.com/niklaskunkel/oasis-api/config.json")
	raw, err := ioutil.ReadFile(absPath)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(raw, &OASIS)
	if err != nil {
		log.Fatal(err)
	}
	return
}

type Market struct{
	Base	string 	`json:"base,omitempty"`
	Quote 	string 	`json:"quote,omitempty"`
	BasePrecision 	string 	`json:"basePrecision,omitempty"`
	QuotePrecision 	string 	`json:"quotePrecision,omitempty"`
	Active 	bool 	`json:"active,omitempty"`
}

//type Markets map[string]*Market

var LiveMarkets = map[string]*Market{
	"MKR/ETH": &Market{
		"MKR",
		"ETH",
		"18",
		"18",
		true},
	"ETH/DAI": &Market{
		"ETH",
		"DAI",
		"18",
		"18",
		true},
	"MKR/DAI": &Market{
		"MKR",
		"DAI",
		"18",
		"18",
		true},
	"DGD/ETH": &Market{
		"DGD",
		"ETH",
		"9",
		"18",
		true},
	"REP/ETH": &Market{
		"REP",
		"ETH",
		"18",
		"18",
		true},
	"RHOC/ETH": &Market{
		"RHOC",
		"ETH",
		"8",
		"18",
		true},
}

type tokenInfo struct {
	Contract string
	Precision int
}

var TokenInfoLib = map[string]tokenInfo{
	"MKR": tokenInfo{
		Contract: strings.ToLower("0x0000000000000000000000009f8F72aA9304c8B593d555F12eF6589cC3A579A2"),
		Precision: 18},
	"ETH": tokenInfo{
		Contract: strings.ToLower("0x000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"),
		Precision: 18},
	"DAI": tokenInfo{
		Contract: strings.ToLower("0x00000000000000000000000089d24a6b4ccb1b6faa2625fe562bdd9a23260359"),
		Precision: 18},
	"DGD": tokenInfo{
		Contract: strings.ToLower("0x000000000000000000000000e0b7927c4af23765cb51314a0e0521a9645f0e2a"),
		Precision: 9},
	"REP": tokenInfo{
		Contract: strings.ToLower("0x000000000000000000000000e94327d07fc17907b4db788e5adf2ed424addff6"),
		Precision: 18},
	"RHOC": tokenInfo{
		Contract: strings.ToLower("0x000000000000000000000000168296bb09e24a88805cb9c33356536b980d3fc5"),
		Precision: 8},
}
