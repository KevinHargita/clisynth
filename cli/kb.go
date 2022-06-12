package cli

import (
	"clisynth/synth"

	"github.com/spf13/cobra"
)

var kbCmd = &cobra.Command{
	Use:   "kb",
	Short: "keyboard",
	Long:  `Starts new keyboard session.`,
	Run: func(cmd *cobra.Command, args []string) {
		synth.StartKBSession(200)
	},
}

func init() {
	rootCmd.AddCommand(kbCmd)
}
