package hitbtc

type OrderType string

const (
	limit      OrderType = "limit"
	stopLimit            = "stopLimit"
	stopMarket           = "stopMarket"
	market               = "market"
)

type TimeInForce string

const (
	GTC TimeInForce = "GTC"
	IOC             = "IOC"
	FOK             = "FOK"
	DAY             = "DAY"
)

type Side string

const (
	buy  Side = "buy"
	sell      = "sell"
)

/*{ "ExecutionReport":
  { "orderId": "58521038",
    "clientOrderId": "11111112",
    "execReportType": "new",
    "orderStatus": "new",
    "symbol": "BTCUSD",
    "side": "buy",
    "timestamp": 1395236779235,
    "price": 0.1,
    "quantity": 100,
    "type": "limit",
    "timeInForce": "GTC",
    "lastQuantity": 0,
    "lastPrice": 0,
    "leavesQuantity": 100,
    "cumQuantity": 0,
    "averagePrice": 0 } }*/
type ExecutionReport struct {
	OrderId           string  `json:"orderId"`
	ClientOrderId     string  `json:"clientOrderId"`
	ExecReportType    string  `json:"execReportType"`
	OrderRejectReason string  `json:"orderRejectReason,omitempty"`
	OrderStatus       string  `json:"orderStatus"`
	Symbol            string  `json:"symbol"`
	Side              string  `json:"side"`
	Timestamp         int64   `json:"timestamp"`
	Price             float64 `json:"price,string"`
	Quantity          float64 `json:"quantity"`
	Type              string  `json:"type"`
	TimeInForce       string  `json:"timeInForce"`
	LastQuantity      float64 `json:"lastQuantity"`
	LastPrice         string  `json:"lastPrice"`
	LeavesQuantity    float64 `json:"leavesQuantity"`
	CumQuantity       float64 `json:"cumQuantity"`
	AveragePrice      float64 `json:"averagePrice,string"`
}
type TradeResponseSuccess struct {
	ExecutionReport ExecutionReport `json:"ExecutionReport"`
}

type TradeMultiResponseSuccess struct {
	ExecutionReport []ExecutionReport `json:"ExecutionReport"`
}
