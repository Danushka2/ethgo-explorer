package services

import (
	"context"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

)

func GetTransactionByHash(client *ethclient.Client, txHash common.Hash) (*types.Transaction, bool, error){
	transaction, isPending, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		return nil, false, err
	}
	return transaction, isPending, nil
}
