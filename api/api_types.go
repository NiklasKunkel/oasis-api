package api

type Market struct {
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

type TokenPairMarket struct {
	Market 		Market 				`json:"market,omitempty"`
	Timer 		int64 				`json:"time,omitempty"`
}

type AllMarkets struct {
	Markets 	map[string]Market 	`json:"markets,omitempty"`
	Time 		int64				`json:"time,omitempty"`
}


type Spread struct {
	Bid 	string 					`json:"bid,omitempty"`
	Ask 	string 					`json:"ask,omitempty"`
}

type TokenPairSpread struct {
	Spread 	Spread 					`json:"spread,omitempty"`
	Time 	int64 					`json:"time,omitempty"`
}

type AllSpreads struct {
	Spreads 	map[string]Spread 	`json:"spreads,omitempty"`
	Time 		int64 				`json:"time,omitempty"`
}

type Prices struct {
	Vwap24Hr 	string				`json:"24hrvwap,omitempty"`
	Vwap12Hr	string 				`json:"12hrvwap,omitempty"`
	Vwap6Hr		string				`json:"6hrvwap,omitempty"`
	Vwap1Hr 	string				`json:"1hrvwap,omitempty"`
	Last 		string				`json:"last,omitempty"`
}

type TokenPairPrices struct {
	Prices 		Prices 				`json:"prices,omitempty"`
	Time 		int64 				`json:"time,omitempty"`
}

type AllPrices struct {
	Prices 	map[string]Prices		`json:"prices,omitempty"`
	Time 	int64					`json:"time,omitempty"`		
}

type TokenPairVolume struct {
	Volume 		string				`json:"vol,omitempty"`
	Time 		int64 				`json:"time,omitempty"`
}

type AllVolumes struct {
	Volumes 	map[string]string	`json:"volumes,omitmepty"`
	Time 		int64 				`json:"time,omitempty"`
}

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

type MkrTokenSupply struct {
	TotalSupply 		string		`json:"totalSupply,omitempty"`
	CirculatingSupply	string		`json:"circulatingSupply,omitempty"`
	Time 				int64 		`json:"time,omitempty"`
}

type Error struct {
	Message 	string				`json:"message"`
}