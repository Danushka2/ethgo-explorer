package util

import (
	"math"
	"math/big"
)

func GetEthValue(balance *big.Int) *big.Float {
	return new(big.Float).Quo(new(big.Float).SetInt(balance), big.NewFloat(math.Pow10(18)))
}