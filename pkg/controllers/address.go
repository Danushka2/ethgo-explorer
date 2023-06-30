package controllers

import (
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/Danushka2/ethgo-explorer/pkg/services"
	"github.com/Danushka2/ethgo-explorer/pkg/models"
	"github.com/Danushka2/ethgo-explorer/pkg/util"
	"github.com/gin-gonic/gin"
)

func AddressInfo(c *gin.Context) {
	var address models.GetAddressInfo

	if err := c.ShouldBindJSON(&address); err != nil {
		customErrorResponse(c, err.Error(), http.StatusBadRequest)
		return
	}

	// Convert the wallet address string to a common.Address type
	walletAddress := common.HexToAddress(address.Address)
	
	balance, err := services.GetAddressBalance(ethereumClient, walletAddress)
	if err != nil {
		customErrorResponse(c, err.Error(), http.StatusBadRequest)
	}

	nonce, err := services.GetAddressNonce(ethereumClient, walletAddress)
	if err != nil {
		customErrorResponse(c, err.Error(), http.StatusBadRequest)
	}
		
	addressInfo := models.AddressInfo{
		Address: address.Address,
		Balance: util.GetEthValue(balance),
		Nonce: nonce,
	}

	c.JSON(http.StatusOK, addressInfo)
}

func CreateAddress(c *gin.Context) {
	response := gin.H{
		"field1": "field1",
		"field2": "field2",
		"field3": "field3",
	}

	c.JSON(http.StatusOK, response)
}
