package cli

import (
	"github.com/spf13/cobra"
)

var exitCmd = &cobra.Command{
	Use:   "exit",
	Short: "exit",
	Long:  `Exits synth.`,
	Run: func(cmd *cobra.Command, args []string) {
		Exit()
	},
}

func init() {
	rootCmd.AddCommand(exitCmd)
}
