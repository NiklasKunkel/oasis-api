package api

/*
type Pair struct{
	Pair 	string 	`json:"pair,omitempty"`
	Base	string 	`json:"base,omitempty"`
	Quote 	string 	`json:"quote,omitempty"`
	BasePrecision 	string 	`json:"basePrecision,omitempty"`
	QuotePrecision 	string 	`json:"quotePrecision,omitempty"`
	Active 	bool 	`json:"active,omitempty"`
	Time 	int64 	`json:"time,omitempty"`
}
*/

type Market struct {
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

type AllMarkets map[string]Market

type TokenPairSpread struct {
	Bid 	string 	`json:"bid,omitempty"`
	Ask 	string 	`json:"ask,omitempty"`
}

type AllSpreads struct {
	Spreads 	map[string]*TokenPairSpread `json:"spreads,omitempty"`
	Time 		int64 	`json:"time,omitempty"`
}


type TokenPairPrice struct {
	Vwap24Hr 	string	`json:"24hrvwap,omitempty"`
	Vwap12Hr	string 	`json:"12hrvwap,omitempty"`
	Vwap6Hr		string	`json:"6hrvwap,omitempty"`
	Vwap1Hr 	string	`json:"1hrvwap,omitempty"`
	Last 		string	`json:"last,omitempty"`
	Time 		int64 	`json:"time,omitempty"`
}

type AllPrices 	map[string]TokenPairPrice

type TokenPairVolume struct {
	Volume 		string	`json:"vol,omitempty"`
	Time 		int64 	`json:"time,omitempty"`
}

type AllVolumes struct {
	Volumes 	map[string]string	`json:"volumes,omitmepty"`
	Time 		int64 				`json:"time,omitempty"`
}

type MkrTokenSupply struct {
	TotalSupply 		string	`json:"totalSupply,omitempty"`
	CirculatingSupply	string	`json:"circulatingSupply,omitempty"`
}

type Error struct {
	Message 	string	`json:"message"`
}