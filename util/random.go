package util

import (
	"math/rand"
	"strings"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

// Generate random int between min and mix
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// Generate random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)
	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

// Genrate random owner name
func RandOwner() string {
	return RandomString(6)
}

// Generate random amount of money
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// Genrate random currency
func RandomCurrency() string {
	currencies := []string{"USD", "INR", "EUR"}
	var n = len(currencies)
	return currencies[rand.Intn(n)]
}
