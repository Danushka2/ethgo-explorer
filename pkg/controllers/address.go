package controllers

import (
	"fmt"
	"log"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddressInfo(c *gin.Context) {
	response := gin.H{
		"field1": "field1",
		"field2": "field2",
		"field3": "field3",
	}

	c.JSON(http.StatusOK, response)
}

func CreateAddress(c *gin.Context) {
	response := gin.H{
		"field1": "field1",
		"field2": "field2",
		"field3": "field3",
	}

	c.JSON(http.StatusOK, response)
}

func BlockInfo(c *gin.Context) {

	block, err := ethereumClient.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Fatalf("Error to get a block: %v", err)
	}

	fmt.Println()

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