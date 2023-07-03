package models

// import (
// 	"math/big"
// )

type GetTransactionInfo struct {
	Transaction      string
}

type TransactionInfo struct {
	TransactionHash string
	To 				string
}
