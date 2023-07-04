package services

import (
	"context"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

)

type TransactionClient struct {
	Client *ethclient.Client
}


func (trClient TransactionClient) GetTransactionByHash(txHash common.Hash) (*types.Transaction, bool, error){
	transaction, isPending, err := trClient.Client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		return nil, false, err
	}
	return transaction, isPending, nil
}


type TransactionServicesIn interface {
	GetTransactionByHash(txHash common.Hash) (*types.Transaction, bool, error)
}
