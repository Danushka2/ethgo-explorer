package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

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
