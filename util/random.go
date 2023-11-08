package util

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
}

func RandomString(n int) string {
	var sb strings.Builder
	k := big.NewInt(int64(len(alphabet)))

	for i := 0; i < n; i++ {
		bi, err := rand.Int(rand.Reader, k)
		if err != nil {
			panic(err) // return the error to handle it properly
		}
		sb.WriteByte(alphabet[bi.Int64()])
	}
	return sb.String()
}

func RandomInt(min, max int64) int64 {
	if min > max {
		panic(fmt.Errorf("invalid range: %d > %d", min, max))
	}
	delta := max - min

	// The range of the random numbers (delta + 1) is passed to rand.Int
	// The big.NewInt function converts delta to *big.Int
	// rand.Int returns a random number n such that 0 <= n < delta
	nBig, err := rand.Int(rand.Reader, big.NewInt(delta+1))
	if err != nil {
		panic(err)
	}

	// Convert big.Int to int64 and add the minimum value
	return min + nBig.Int64()
}
