package controllers

import (
	"net/http"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
)

var ethereumClient *ethclient.Client

func SetEthClient(client *ethclient.Client) {
	ethereumClient = client
}

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello World!",
	})
}

func customResponse(c *gin.Context, message string, httpStatusCode int) {
	c.JSON(httpStatusCode, gin.H{"message": message})
}

func customErrorResponse(c *gin.Context, message string, httpStatusCode int) {
	c.JSON(httpStatusCode, gin.H{"error": message})

}
