package api

import (
	//"math/big"
)

type MkrTokenSupply struct {
	//TotalSupply 	*big.Float	`json:"totalSupply, omitempty"`
	TotalSupply 	string	`json:"totalSupply,omitempty"`
}

type MkrTokenPrice struct {
	Price 			string	`json:"price,omitempty"`
}

type TokenPair struct {
	TokenPair	string	`json:"tokenPair,omitempty"`
	Price 		string	`json:"price,omitempty"`
	Volume 		string	`json:"volume,omitempty"`
	LowestAsk 	string	`json:"lowestAsk,omitempty"`
	HighestBid 	string	`json:"highestBid,omitempty"`
	High24Hr	string 	`json:"high24Hr,omitempty"`
	Low24Hr		string	`json:"low24Hr,omitempty"`
	IsFrozen	bool	`json:"isFrozen,omitempty"`
	LastUpdated	uint	`json:"lastUpdated,omitempty"`
}

type TokenPairPrice struct {
	TokenPair	string	`json:"tokenPair,omitempty"`
	Price 		string	`json:"price,omitempty"`
}

type TokenPairVolume struct {
	TokenPair 	string 	`json:"tokenPair,omitempty"`
	Volume 		string	`json:"volume,omitempty"`
}

type Error struct {
	Error 		string	`json:"error"`
}