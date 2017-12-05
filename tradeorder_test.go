package hitbtc

import (
	"fmt"
	"testing"
)

func TestTradeOrder_String(t *testing.T) {
	fmt.Println(NewTradeLimitOrder("USDBTC", "buy", 10, 100.0).String())
}
