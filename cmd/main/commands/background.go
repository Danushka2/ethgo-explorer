package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

// backgroundCmd represents the serve command
var bgCmd = &cobra.Command{
	Use:   "background",
	Short: "Run the background tasks for application",
	RunE: func(cmd *cobra.Command, args []string) error {
		// application related code
		fmt.Println("im here in background command")
		return nil
	},
}
