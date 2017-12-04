package api

type Market struct{
	Pair 	string 	`json:"pair,omitempty"`
	Base	string 	`json:"base,omitempty"`
	Quote 	string 	`json:"quote,omitempty"`
	BasePrecision 	string 	`json:"basePrecision,omitempty"`
	QuotePrecision 	string 	`json:"quotePrecision,omitempty"`
	Active 	bool 	`json:"active,omitempty"`
	Time 	int64 	`json:"time,omitempty"`
}

type Markets []Market

type MkrTokenSupply struct {
	//TotalSupply 	*big.Float	`json:"totalSupply, omitempty"`
	TotalSupply 	string	`json:"totalSupply,omitempty"`
}

type TokenPair struct {
	TokenPair	string	`json:"pair,omitempty"`
	Price 		string	`json:"price,omitempty"`
	LastPrice	string	`json:"last,omitempty"`
	Volume 		string	`json:"vol,omitempty"`
	LowestAsk 	string	`json:"ask,omitempty"`
	HighestBid 	string	`json:"bid,omitempty"`
	Low24Hr		string	`json:"low,omitempty"`
	High24Hr	string 	`json:"high,omitempty"`
	Active		bool	`json:"active,omitempty"`
	LastUpdated	int64	`json:"time,omitempty"`
}

type TokenPairSpread struct {
	Bid 	string 	`json:"bid,omitempty"`
	Ask 	string 	`json:"ask,omitempty"`
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