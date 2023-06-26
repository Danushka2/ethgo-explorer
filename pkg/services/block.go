package services

import (
	"math/big"
	"context"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/core/types"
)

func GetBlockByNumber(client *ethclient.Client, number *big.Int) (*types.Block, error){

	block, err := client.BlockByNumber(context.Background(), number)
	if err != nil {
		return nil, err
	}

	return block, nil
}
