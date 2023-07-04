package controllers

import (
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/Danushka2/ethgo-explorer/pkg/services"
	"github.com/Danushka2/ethgo-explorer/pkg/models"
	"github.com/Danushka2/ethgo-explorer/pkg/util"
	"github.com/gin-gonic/gin"
)

var txInterface services.TransactionServicesIn

func TransactionInfo(c *gin.Context) {
	var reqTx models.GetTransactionInfo

	if err := c.ShouldBindJSON(&reqTx); err != nil {
		customErrorResponse(c, err.Error(), http.StatusBadRequest)
		return
	}

	txHash := common.HexToHash(reqTx.Transaction)

	tx, isPending, err := txInterface.GetTransactionByHash(txHash)
	if err != nil {
		customErrorResponse(c, err.Error(), http.StatusBadRequest)
	}

	if isPending {
		response := gin.H{
			"Message": "Transaction is pending",
		}
	
		c.JSON(http.StatusOK, response)
	} else {
		// fmt.Println("From:", tx.From().Hex())
		transactionInfo := models.TransactionInfo{
			TransactionHash: tx.Hash().Hex(),
			To: 			 tx.To().Hex(),
			Value:			 util.GetEthValue(tx.Value()),
			GasValue:		 util.GetEthValue(tx.GasPrice()),
			Nonce:			 tx.Nonce(),
		}
	
		c.JSON(http.StatusOK, transactionInfo)
	}
}

func NewTransactionServicesIn(t services.TransactionServicesIn){
	txInterface = t
}
