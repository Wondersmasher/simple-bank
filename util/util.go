package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabets = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}
func RandomString(n int) string {
	var sb strings.Builder
	for i := 0; i < n; i++ {
		c := alphabets[rand.Intn(len(alphabets))]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomOwner() string {
	return RandomString(6) + " " + RandomString(6)
}

func RandomCurrency() string {
	return []string{"USD", "EUR", "CAD"}[rand.Intn(3)]
}
