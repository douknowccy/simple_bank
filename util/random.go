package util

import (
	"math/rand"
	"strings"
	"time"
)

const alpabet = "asdlkjmngdlcv"

func init() {
	rand.Seed(time.Now().UnixNano())

}

// randomInt generates a random Interger between min and max
func RandomInt(min, max int64) int64 {

	return min + rand.Int63n(max-min+1)
}

// RandomString generates a random of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alpabet)

	for i := 0; i < n; i++ {
		c := alpabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomOwner() string {
	return RandomString(6)
}

func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

func RandomCurrency() string {
	currencies := []string{"EUR", "USD", "TWD", "CNY"}
	n := len(currencies)
	return currencies[rand.Intn(n)]

}
