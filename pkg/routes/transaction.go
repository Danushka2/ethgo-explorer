package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/Danushka2/ethgo-explorer/pkg/controllers"
)

var TransactionRoutes = func(router *gin.Engine) {
	addressV1 := router.Group("/api/v1/transaction")
	{
		addressV1.GET("/info", controllers.TransactionInfo)
	}
}
