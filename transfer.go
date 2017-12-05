package hitbtc

import "github.com/google/go-querystring/query"

type TransferRequest struct {
	Amount       float64 `url:"amount"`        //Yes	decimal	Funds amount to withdraw
	CurrencyCode string  `url:"currency_code"` //	Yes	string	Currency symbol, e.g. BTC

}

func (w *TransferRequest) String() string {
	v, _ := query.Values(w)
	return v.Encode()
}

func NewTransferRequest(coin string, amount float64) *TransferRequest {
	return &TransferRequest{
		Amount:       amount,
		CurrencyCode: coin,
	}
}
