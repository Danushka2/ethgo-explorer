package routes

import (
	"github.com/Danushka2/ethgo-explorer/pkg/controllers"
	"github.com/gin-gonic/gin"
)

var RootRoutes = func(router *gin.Engine) {
	router.GET("/ping", controllers.HealthCheck)
}
