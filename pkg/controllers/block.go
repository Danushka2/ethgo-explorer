package controllers

import (
	"fmt"
	"log"
	"math/big"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/Danushka2/ethgo-explorer/pkg/services"
	"github.com/Danushka2/ethgo-explorer/pkg/models"
	"github.com/Danushka2/ethgo-explorer/pkg/util"
	"github.com/ethereum/go-ethereum/core/types"
)

var bckInterface services.BlockServicesIn

func GetBlockInfo(c *gin.Context) {
	var blockInfo models.BlockInfo
	var blockValue *big.Int

	if err := c.ShouldBindJSON(&blockInfo); err != nil {
		customErrorResponse(c, err.Error(), http.StatusBadRequest)
		return
	}

	if !blockInfo.IsLastBlock {
		blockValue = &blockInfo.Number
	}

	block, err := bckInterface.GetBlockByNumber(blockValue)
	if err != nil {
		log.Fatalf("Error to get a block: %v", err)
	}

	for _, tx := range block.Transactions() {
		fmt.Println("Transaction Hash:", tx.Hash().Hex())

		from, err := types.Sender(types.NewEIP155Signer(tx.ChainId()), tx) 
		if err != nil {
			fmt.Println("From:", err)
		}else{
			fmt.Println("From:", from)
		}

		fmt.Println("To:", tx.To().Hex())
		fmt.Println("Value:", tx.Value().String())
		fmt.Println("Value Eth:", util.GetEthValue(tx.Value()))
		fmt.Println("Gas Limit:", tx.Gas())
		fmt.Println("Gas Price:", tx.GasPrice().String())
		fmt.Println("Gas Price Eth:", util.GetEthValue(tx.GasPrice()))
		// fmt.Println("Input Data:", string(tx.Data()))
		fmt.Println("---------------------------------------------------")
	}

	response := gin.H{
		"Number": block.Number(),
		"Difficulty": block.Difficulty(),
		"GasLimit": block.GasLimit(),
		"GasUsed": block.GasUsed(),
		"Nonce": block.Nonce(),
		"Time": block.Time(),
	}

	c.JSON(http.StatusOK, response)
}

func NewBlockServicesIn(b services.BlockServicesIn){
	bckInterface = b
}
