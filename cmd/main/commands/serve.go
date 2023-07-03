package commands

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Danushka2/ethgo-explorer/pkg/middlewares"
	"github.com/Danushka2/ethgo-explorer/pkg/routes"
	"github.com/Danushka2/ethgo-explorer/pkg/config"
	"github.com/Danushka2/ethgo-explorer/pkg/controllers"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run http server",
	RunE: func(cmd *cobra.Command, args []string) error {

		cx := make(chan os.Signal, 1)
		signal.Notify(cx, os.Interrupt, syscall.SIGTERM)
		go func() {
			<-cx
			fmt.Println("\nReceived signal: SIGINT (Ctrl+C)")
			fmt.Println("■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■")

			os.Exit(0)
		}()

		// ##################################################################		
		client, err := config.NewEthereumClient()
		if err != nil {
			log.Fatalf("Unable to connect: %v", err)
		}

		controllers.SetEthClient(client)
		// ##################################################################

		r := gin.Default()

		isDevelopment = viper.GetBool("DEVELOPMENT")
		if !isDevelopment {
			gin.SetMode(gin.ReleaseMode)
		}

		r.SetTrustedProxies([]string{"localhost", "proxy2.example.com"})
		r.Use(middlewares.CORSMiddleware())
		r.Use(middlewares.RequestIDMiddleware())

		routes.RootRoutes(r)
		routes.AddressRoutes(r)
		routes.BlockRoutes(r)
		routes.TransactionRoutes(r)

		r.NoRoute(func(c *gin.Context) {
			c.JSON(404, gin.H{
				"error": "not a valid route!",
			})
		})

		r.Run(":4040")

		<-cx
		return nil
	},
}
