package services

import (
	"math/big"
	"context"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/common"
)

func GetAddressBalance(client *ethclient.Client, address common.Address) (*big.Int, error){
	balance, err := client.BalanceAt(context.Background(), address, nil)
	if err != nil {
		return nil, err
	}

	return balance, nil
}


func GetAddressNonce(client *ethclient.Client, address common.Address) (*uint64, error){
	nonce, err := client.PendingNonceAt(context.Background(), address)
	if err != nil {
		return nil, err
	}
	
	return &nonce, nil
}
