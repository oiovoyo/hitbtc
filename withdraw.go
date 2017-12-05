package hitbtc

import "github.com/google/go-querystring/query"

/*
Request: POST /api/1/payment/payout

Parameters:

Parameter	Required	Type	Description
amount	Yes	decimal	Funds amount to withdraw
currency_code	Yes	string	Currency symbol, e.g. BTC
address	Yes	string	BTC/LTC address to withdraw to
extra_id	No	string	payment id for cryptonote
recommended_fee	No	boolean	use recommended network fee instead of flat, e.g. 0 or 1
Example: amount=0.001&currency_code=BTC&address=1LuWvENyuPNHsHWjDgU1QYKWUYN9xxy7n5
*/

type WithdrawRequest struct {
	Amount         float64 `url:"amount"`          //Yes	decimal	Funds amount to withdraw
	CurrencyCode   string  `url:"currency_code"`   //	Yes	string	Currency symbol, e.g. BTC
	Address        string  `url:"address"`         //Yes	string	BTC/LTC address to withdraw to
	ExtraId        string  `url:"extra_id"`        //No	string	payment id for cryptonote
	RecommendedFee int     `url:"recommended_fee"` //No	boolean	use recommended network fee instead of flat, e.g. 0 or 1

}

func (w *WithdrawRequest) String() string {
	v, _ := query.Values(w)
	if w.ExtraId == "" {
		v.Del("extra_id")
	}
	return v.Encode()
}

func NewWithdrawRequest(coin string, amount float64, address, paymentID string) *WithdrawRequest {
	return &WithdrawRequest{
		Amount:         amount,
		CurrencyCode:   coin,
		Address:        address,
		ExtraId:        paymentID,
		RecommendedFee: 1,
	}
}
