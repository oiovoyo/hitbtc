package hitbtc

import (
	"fmt"
	"testing"
)

func Test_OrderID(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println(OrderID())
	}
}
