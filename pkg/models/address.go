package models

import (
	"math/big"
)

type GetAddressInfo struct {
	Address      string
}

type AddressInfo struct {
	Address string
	Balance *big.Float
	Nonce 	*uint64
}
