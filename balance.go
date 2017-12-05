package hitbtc

/*
{"balance": [
  {"currency_code": "BTC","cash": 0.045457701,"reserved": 0.01},
  {"currency_code": "EUR","cash": 0.0445544,"reserved": 0},
  {"currency_code": "LTC","cash": 0.7,"reserved": 0.1},
  {"currency_code": "USD","cash": 2.9415029,"reserved": 1.001}
]}
*/
type OneTradeBalance struct {
	Currency string  `json:"currency_code"`
	Cash     float64 `json:"cash,string"`
	Resvered float64 `json:"reserved,string"`
}

type TradeBalance struct {
	Balance []OneTradeBalance `json:"balance"`
}

/*
{
  "balance": [
  {
    "currency_code": "USD",
    "balance": 13.12
  },
  {
    "currency_code": "EUR",
    "balance": 0
  },
  {
    "currency_code": "LTC",
    "balance": 1.07
  },
  {"currency_code": "BTC",
  "balance": 11.9
  }
]}
*/
type OnePaymentBalance struct {
	Currency string  `json:"currency_code"`
	Balance  float64 `json:"balance,string"`
}

type PaymentBalance struct {
	Balance []OnePaymentBalance `json:"balance"`
}
