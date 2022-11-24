package utils

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"
const nonZero = "123456789"
const withZero = "1234567890"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomNumString(n int) string {
	var sb, sb2 strings.Builder
	k := len(nonZero)
	j := len(withZero)

	for i := 0; i < n-1; i++ {
		c := withZero[rand.Intn(j)]
		sb.WriteByte(c)
	}
	d := nonZero[rand.Intn(k)]
	sb2.WriteByte(d)

	return sb2.String() + sb.String()
}

func RandomUserName() string {
	return RandomString(6)
}

func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

func RandomCurrency() string {
	currencies := []string{"IDR", "USD", "EUR"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
