package services

import (
	"math/big"
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/common"
)

type AddressClient struct {
	Client *ethclient.Client
}


func (adClient AddressClient) GetAddressBalance(address common.Address) (*big.Int, error){
	balance, err := adClient.Client.BalanceAt(context.Background(), address, nil)
	if err != nil {
		return nil, err
	}

	return balance, nil
}

func (adClient AddressClient) GetAddressNonce(address common.Address) (*uint64, error){
	nonce, err := adClient.Client.PendingNonceAt(context.Background(), address)
	if err != nil {
		return nil, err
	}
	
	return &nonce, nil
}

func CreateWalletWithKeystore(password string) {
	ks := keystore.NewKeyStore("./keystore", keystore.StandardScryptN, keystore.StandardScryptP)

	account, err := ks.NewAccount(password)
	if err != nil {
		fmt.Println("Failed to create new account:", err)
		return
	}

	fmt.Println("New account created:")
	fmt.Println("Address:", account.Address.Hex())
	fmt.Println("Password:", password)
	fmt.Println("URL:", account.URL)
}


type AddressServicesIn interface {
	GetAddressBalance(address common.Address) (*big.Int, error)
	GetAddressNonce(address common.Address) (*uint64, error)
}
