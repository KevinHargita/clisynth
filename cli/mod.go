package cli

import (
	"github.com/spf13/cobra"
)

var modCmd = &cobra.Command{
	Use:   "mod",
	Short: "module",
	Long:  `Create and interact with different synthesizer modules.`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	rootCmd.AddCommand(modCmd)
}
