package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/Danushka2/ethgo-explorer/pkg/controllers"
)

var UserRoutes = func(router *gin.Engine) {
	userV1 := router.Group("/api/v1/address")
	{
		userV1.POST("/new", controllers.CreateAddress)
		userV1.GET("/info", controllers.AddressInfo)
	}
}
