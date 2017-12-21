package api

type Response struct {
	Data 		interface{}			`json:"data,omitempty"`
	Time 		int64 				`json:"time,omitempty"`
	Message 	string 				`json:"message,omitempty"`
}

type TokenPairMarket struct {
	TokenPair	string				`json:"pair,omitempty"`
	Price 		string				`json:"price,omitempty"`
	LastPrice	string				`json:"last,omitempty"`
	Volume 		string				`json:"vol,omitempty"`
	LowestAsk 	string				`json:"ask,omitempty"`
	HighestBid 	string				`json:"bid,omitempty"`
	Low24Hr		string				`json:"low,omitempty"`
	High24Hr	string 				`json:"high,omitempty"`
	Active		bool				`json:"active,omitempty"`
}

type AllMarkets	map[string]TokenPairMarket

type TokenPairPrices struct {
	Vwap24Hr 	string				`json:"24hrvwap,omitempty"`
	Vwap12Hr	string 				`json:"12hrvwap,omitempty"`
	Vwap6Hr		string				`json:"6hrvwap,omitempty"`
	Vwap1Hr 	string				`json:"1hrvwap,omitempty"`
	Last 		string				`json:"last,omitempty"`
}

type AllPrices map[string]TokenPairPrices

type TokenPairVolume struct {
	Volume 		string				`json:"vol,omitempty"`
}

type AllVolumes map[string]TokenPairVolume

type TokenPairSpread struct {
	Bid 	string 					`json:"bid,omitempty"`
	Ask 	string 					`json:"ask,omitempty"`
}

type AllSpreads map[string]TokenPairSpread 

type TokenPairTradeHistory []Trade

type Trade struct {
	Price 		string				`json:"price,omitempty"`
	BuyToken	string 				`json:"buyToken,omitempty"`
	PayToken	string 				`json:"payToken,omitempty"`
	BuyAmount	string 				`json:"buyAmount,omitempty"`
	PayAmount	string 				`json:"payAmount,omitempty"`
	Type 		string 				`json:"type,omitempty"`
	Time 		string				`json:"time,omitempty"`
}

type TokenSupply struct {
	TotalSupply 		string		`json:"totalSupply,omitempty"`
	CirculatingSupply	string		`json:"circulatingSupply,omitempty"`
	Time 				int64 		`json:"time,omitempty"`
}

type Error struct {
	Message 	string				`json:"message"`
}