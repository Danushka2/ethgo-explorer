package models

import (
	"math/big"
)

type BlockInfo struct {
	Number      big.Int 
	IsLastBlock bool		`json:"IsLastBlock,omitempty" bson:"IsLastBlock,omitempty"`
}
