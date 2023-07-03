package controllers

import (
	"fmt"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/Danushka2/ethgo-explorer/pkg/services"
	"github.com/Danushka2/ethgo-explorer/pkg/models"
	"github.com/gin-gonic/gin"
)

func TransactionInfo(c *gin.Context) {
	var reqTx models.GetTransactionInfo

	if err := c.ShouldBindJSON(&reqTx); err != nil {
		customErrorResponse(c, err.Error(), http.StatusBadRequest)
		return
	}

	txHash := common.HexToHash(reqTx.Transaction)

	// Retrieve the transaction details
	tx, isPending, err := services.GetTransactionByHash(ethereumClient, txHash)
	if err != nil {
		customErrorResponse(c, err.Error(), http.StatusBadRequest)
	}

	if isPending {
		fmt.Println("Transaction is pending")
	} else {
		fmt.Println("Transaction Hash:", tx.Hash().Hex())
		// fmt.Println("Block Number:", tx.BlockNumber().Uint64())
		// fmt.Println("Block Hash:", tx.BlockHash().Hex())
		// fmt.Println("From:", tx.From().Hex())
		fmt.Println("To:", tx.To().Hex())
		fmt.Println("Value:", tx.Value().Div(tx.Value(), big.NewInt(1e18)).String(), "Ether")
		fmt.Println("Gas Limit:", tx.Gas())
		fmt.Println("Gas Price:", tx.GasPrice().Div(tx.GasPrice(), big.NewInt(1e9)).String(), "Gwei")
		fmt.Println("Nonce:", tx.Nonce())
	}

	transactionInfo := models.TransactionInfo{
		TransactionHash: tx.Hash().Hex(),
		To: 			 tx.To().Hex(),
	}

	c.JSON(http.StatusOK, transactionInfo)
}
