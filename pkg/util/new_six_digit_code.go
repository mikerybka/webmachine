package util

import (
	"crypto/rand"
	"math/big"
)

func NewSixDigitCode() string {
	n, err := rand.Int(rand.Reader, big.NewInt(1_000_000))
	if err != nil {
		panic(err)
	}
	s := n.String()
	for len(s) < 6 {
		s = "0" + s
	}
	return s
}
