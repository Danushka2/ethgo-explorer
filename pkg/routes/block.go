package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/Danushka2/ethgo-explorer/pkg/controllers"
)

var BlockRoutes = func(router *gin.Engine) {
	blockV1 := router.Group("/api/v1/block")
	{
		blockV1.GET("/info", controllers.GetBlockInfo)
	}
}
