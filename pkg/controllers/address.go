package controllers

import (
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
