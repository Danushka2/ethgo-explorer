package controllers

import (
	"log"
	"math/big"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/Danushka2/ethgo-explorer/pkg/services"
	"github.com/Danushka2/ethgo-explorer/pkg/models"
)

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

	block, err := services.GetBlockByNumber(ethereumClient, blockValue)
	if err != nil {
		log.Fatalf("Error to get a block: %v", err)
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