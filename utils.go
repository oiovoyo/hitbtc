package hitbtc

import (
	"math/rand"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

const orderIDLen = 32

func OrderID() string {
	return randSeq(orderIDLen)
}

func randSeq(n int) string {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[r1.Intn(len(letters))]
	}
	return string(b)
}
