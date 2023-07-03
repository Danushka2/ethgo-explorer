package controllers

import (
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/Danushka2/ethgo-explorer/pkg/services"
	"github.com/Danushka2/ethgo-explorer/pkg/models"
	"github.com/Danushka2/ethgo-explorer/pkg/util"
	"github.com/gin-gonic/gin"
)

var adInterface services.AddressServicesIn

func AddressInfo(c *gin.Context) {
	var address models.GetAddressInfo

	if err := c.ShouldBindJSON(&address); err != nil {
		customErrorResponse(c, err.Error(), http.StatusBadRequest)
		return
	}

	// Convert the wallet address string to a common.Address type
	walletAddress := common.HexToAddress(address.Address)
	
	balance, err := adInterface.GetAddressBalance(walletAddress)
	if err != nil {
		customErrorResponse(c, err.Error(), http.StatusBadRequest)
	}

	nonce, err := adInterface.GetAddressNonce(walletAddress)
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

func NewAddressServicesIn(a services.AddressServicesIn){
	adInterface = a
}

