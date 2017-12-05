package hitbtc

/*
{
    "symbols": [
        {
            "symbol": "BTCUSD",
            "step": "0.01",
            "lot": "0.01",
            "currency": "USD",
            "commodity": "BTC",
            "takeLiquidityRate": "0.002",
            "provideLiquidityRate": "0.002"
        },
        {
            "symbol": "BTCEUR",
            "step": "0.01",
            "lot": "0.01",
            "currency": "EUR",
            "commodity": "BTC",
            "takeLiquidityRate": "0.002",
            "provideLiquidityRate": "0.002"
        },
        ...
    ]
}
*/
type Market struct {
	Symbol               string  `json:"symbol"`
	Step                 float64 `json:"step,string"`
	Lot                  float64 `json:"lot,string"`
	Currency             string  `json:"currency"`
	Commodity            string  `json:"commodity"`
	TakeLiquidityRate    float64 `json:"takeLiquidityRate,string"`
	ProvideLiquidityRate float64 `json:"provideLiquidityRate,string"`
}
type Markets struct {
	Symbols []Market `json:"symbols"`
}

type MarketsMap map[string]Market
