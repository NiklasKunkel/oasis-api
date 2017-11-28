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

type Token struct {
	Id			string	`json:"id,omitempty"`
	Price 		float32	`json:"price,omitempty"`
	Volume 		float32	`json:"volume,omitempty"`
	LowestAsk 	float32	`json:"lowestAsk,omitempty"`
	HighestBid 	float32	`json:"highestBid,omitempty"`
	High24Hr	float32 `json:"high24Hr,omitempty"`
	Low24Hr		float32	`json:"low24Hr,omitempty"`
	IsFrozen	bool	`json:"isFrozen,omitempty"`
	LastUpdated	uint	`json:"lastUpdated,omitempty"`
}

type TokenPrice struct {
	Id			string	`json:"id,omitempty"`
	Price 		float32	`json:"price,omitempty"`
}

type TokenVolume struct {
	Id			string 	`json:"id,omitempty"`
	Volume 		float32	`json:"volume,omitempty"`
}

type Error struct {
	Error 		string	`json:"error"`
}