package services

import (
	"math/big"
	"context"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/core/types"
)

type BlockClient struct {
	Client *ethclient.Client
}


func (bckClient BlockClient) GetBlockByNumber(number *big.Int) (*types.Block, error){
	block, err := bckClient.Client.BlockByNumber(context.Background(), number)
	if err != nil {
		return nil, err
	}

	return block, nil
}


type BlockServicesIn interface {
	GetBlockByNumber(number *big.Int) (*types.Block, error)
}
