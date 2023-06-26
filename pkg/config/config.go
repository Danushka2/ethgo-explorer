package config

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/viper"
)

func NewEthereumClient() (*ethclient.Client, error) {
	fmt.Println("✦ Conneting to the eth client")
	
	var ETH_CLIENT_HOST = viper.GetString("ETH_CLIENT_HOST")

	client, err := ethclient.DialContext(context.Background(), ETH_CLIENT_HOST)
	if err != nil {
		return nil, err
	}
	
	fmt.Println("✦ Conneted to the eth client")
	return client, nil
}