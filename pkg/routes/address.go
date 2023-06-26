package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/Danushka2/ethgo-explorer/pkg/controllers"
)

var AddressRoutes = func(router *gin.Engine) {
	addressV1 := router.Group("/api/v1/address")
	{
		addressV1.POST("/new", controllers.CreateAddress)
		addressV1.GET("/info", controllers.AddressInfo)
	}
}
