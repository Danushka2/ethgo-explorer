package commands

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/Danushka2/ethgo-explorer/pkg/services"
)

// addressCmd represents the address command
var addressCmd = &cobra.Command{
	Use:   "address",
	Short: "Create Eth Wallet",
	RunE: func(cmd *cobra.Command, args []string) error {
		
		if password != "" {
			services.CreateWalletWithKeystore(password)
		}else {
			fmt.Println("--password flag is missing")
		}
		
		return nil
	},
}