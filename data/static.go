package data

import (
"strings"
)

type market struct{
	Pair 	string 	`json:"pair,omitempty"`
	Base	string 	`json:"base,omitempty"`
	Quote 	string 	`json:"quote,omitempty"`
	BasePrecision 	string 	`json:"basePrecision,omitempty"`
	QuotePrecision 	string 	`json:"quotePrecision,omitempty"`
	Active 	bool 	`json:"active,omitempty"`
	Time 	int64 	`json:"time,omitempty"`
}

var static_markets = []market{
	market{
		"MKR/ETH",
		"MKR",
		"ETH",
		"18",
		"18",
		true,
		0},
	market{
		"MKR/SAI",
		"MKR",
		"SAI",
		"18",
		"18",
		true,
		0},
	market{
		"ETH/SAI",
		"ETH",
		"SAI",
		"18",
		"18",
		true,
		0},
	market{
		"MKR/SAI",
		"MKR",
		"SAI",
		"18",
		"18",
		true,
		0},
	market{
		"MKR/SAI",
		"MKR",
		"SAI",
		"18",
		"18",
		true,
		0},
	market{
		"MKR/SAI",
		"MKR",
		"SAI",
		"18",
		"18",
		true,
		0},
	market{
		"MKR/SAI",
		"MKR",
		"SAI",
		"18",
		"18",
		true,
		0},
	market{
		"MKR/SAI",
		"MKR",
		"SAI",
		"18",
		"18",
		true,
		0},
	market{
		"MKR/SAI",
		"MKR",
		"SAI",
		"18",
		"18",
		true,
		0},
	market{
		"MKR/SAI",
		"MKR",
		"SAI",
		"18",
		"18",
		true,
		0},
	market{
		"MKR/SAI",
		"MKR",
		"SAI",
		"18",
		"18",
		true,
		0},
	market{
		"MKR/SAI",
		"MKR",
		"SAI",
		"18",
		"18",
		true,
		0},
	market{
		"MKR/SAI",
		"MKR",
		"SAI",
		"18",
		"18",
		true,
		0},
}

type tokenInfo struct {
	Contract string
	Precision int
}

var TokenInfoLib = map[string]tokenInfo{
	"MKR": tokenInfo{
		Contract: strings.ToLower("0x000000000000000000000000C66eA802717bFb9833400264Dd12c2bCeAa34a6d"),
		Precision: 18},
	"ETH": tokenInfo{
		Contract: strings.ToLower("0x000000000000000000000000ecf8f87f810ecf450940c9f60066b4a7a501d6a7"),
		Precision: 18},
	"SAI": tokenInfo{
		Contract: strings.ToLower("0x00000000000000000000000059adcf176ed2f6788a41b8ea4c4904518e62b6a4"),
		Precision: 18},
	"DGD": tokenInfo{
		Contract: strings.ToLower("0x000000000000000000000000e0b7927c4af23765cb51314a0e0521a9645f0e2a"),
		Precision: 9},
	"RHOC": tokenInfo{
		Contract: strings.ToLower("0x000000000000000000000000168296bb09e24a88805cb9c33356536b980d3fc5"),
		Precision: 8},
	"REP": tokenInfo{
		Contract: strings.ToLower("0x000000000000000000000000e94327d07fc17907b4db788e5adf2ed424addff6"),
		Precision: 18},
	"ICN": tokenInfo{
		Contract: strings.ToLower("0x000000000000000000000000888666ca69e0f178ded6d75b5726cee99a87d698"),
		Precision: 18},
	"1ST": tokenInfo{
		Contract: strings.ToLower("0x000000000000000000000000af30d2a7e90d7dc361c8c4585e9bb7d2f6f15bc7"),
		Precision: 18},
	"GNT": tokenInfo{
		Contract: strings.ToLower("0x00000000000000000000000001afc37f4f85babc47c0e2d0eababc7fb49793c8"),
		Precision: 18},
	"VSL": tokenInfo{
		Contract: strings.ToLower("0x0000000000000000000000005c543e7ae0a1104f78406c340e9c64fd9fce5170"),
		Precision: 18},
	"PLU": tokenInfo{
		Contract: strings.ToLower("0x000000000000000000000000d8912c10681d8b21fd3742244f44658dba12264e"),
		Precision: 18},
	"MLN": tokenInfo{
		Contract: strings.ToLower("0x000000000000000000000000beb9ef514a379b997e0798fdcc901ee474b6d9a1"),
		Precision: 18},
	"NMR": tokenInfo{
		Contract: strings.ToLower("0x0000000000000000000000001776e1f26f98b1a5df9cd347953a26dd3cb46671"),
		Precision: 18},
	"TIME": tokenInfo{
		Contract: strings.ToLower("0x0000000000000000000000006531f133e6deebe7f2dce5a0441aa7ef330b4e53"),
		Precision: 8},
	"GUP": tokenInfo{
		Contract: strings.ToLower("0x000000000000000000000000f7b098298f7c69fc14610bf71d5e02c60792894c"),
		Precision: 3},
	"BAT": tokenInfo{
		Contract: strings.ToLower("0x0000000000000000000000000d8775f648430679a709e98d2b0cb6250d2887ef"),
		Precision: 18},
	"SNGLS": tokenInfo{
		Contract: strings.ToLower("0x000000000000000000000000aec2e87e0a235266d9c5adc9deb4b2e29b54d009"),
		Precision: 0},
}