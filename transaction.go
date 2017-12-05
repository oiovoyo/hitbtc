package hitbtc

import "time"

/*
{"transactions": [
  {
    "id": "49720-104765-18440728",
    "type": "payin",
    "status": "pending",
    "created": 1397056648,
    "finished": 1397056646,
    "amount_from": 0.001,
    "currency_code_from": "BTC",
    "amount_to": 0.001,
    "currency_code_to": "BTC",
    "destination_data": null,
    "commission_percent": 0,
    "bitcoin_address": "1KnVXD1Wc1wZbTHiB1TDgMcnSRi2PnMHAV",
    "bitcoin_return_address": "1QBuhFksjoWTQWJKWUPyQ37wsQohLAhJvK"
    "external_data": "0b2ac379986cd1872b6a4115ad7a6cf436bdac67080db728579b8282c129a549"
  }
]}
*/

type OneTransaction struct {
	Id                   string  `json:"id"`
	Type                 string  `json:"type"`
	Status               string  `json:"status"`
	Created              int64   `json:"created"`
	Finished             int64   `json:"finished"`
	AmountFrom           float64 `json:"amount_from,string"`
	AmountTo             float64 `json:"amount_to,string"`
	CurrencyCodeFrom     string  `json:"currency_code_from"`
	CurrencyCodeTo       string  `json:"currency_code_to"`
	DestinationData      string  `json:"destination_data"`
	CommissionPercent    float64 `json:"commission_percent"`
	BitcoinAddress       *string `json:"bitcoin_address"`
	BitcoinReturnAddress *string `json:"bitcoin_return_address"`
	ExternalData         string  `json:"external_data"`
}

type Transaction struct {
	Transactions []OneTransaction `json:"transactions"`
}
type OneDepositWithdrawal struct {
	Id       string
	Coin     string
	Created  time.Time
	Finished time.Time
	Amount   float64
	Address  string
	IsDone   bool
}
type DepositWithdrawal struct {
	Deposit  []OneDepositWithdrawal
	Withdraw []OneDepositWithdrawal
}
