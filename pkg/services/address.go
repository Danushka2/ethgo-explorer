package services

import (
	"math/big"
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"

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
		log.Fatal(err)
	}

	fmt.Println("New account created:")
	fmt.Println("Address:", account.Address.Hex())
	fmt.Println("Password:", password)
	fmt.Println("URL:", account.URL)
}

func GenerateWallet() {
	pvk, _ := crypto.GenerateKey()

	pvData := crypto.FromECDSA(pvk)
	pbData := crypto.FromECDSAPub(&pvk.PublicKey)
	
	fmt.Println("Private:", hexutil.Encode(pvData))
	fmt.Println("Public:", hexutil.Encode(pbData))
}


type AddressServicesIn interface {
	GetAddressBalance(address common.Address) (*big.Int, error)
	GetAddressNonce(address common.Address) (*uint64, error)
}
