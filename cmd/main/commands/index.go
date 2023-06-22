package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	rootCmd = &cobra.Command{
		Use:   "ethgo-explorer",
		Short: "Application for exploring ethereum blockchain",
	}
	isDevelopment = false
)

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "config.yaml", "config file (default is $PWD/config.yaml)")

	rootCmd.AddCommand(serveCmd)
	rootCmd.AddCommand(bgCmd)
}

func initConfig() {
	viper.SetConfigFile(cfgFile)

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	isDevelopment = viper.GetBool("DEVELOPMENT")
	if isDevelopment {
		fmt.Println("Running on Development mode")
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
