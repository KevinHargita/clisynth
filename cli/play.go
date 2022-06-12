package cli

import (
	"clisynth/synth"

	"github.com/spf13/cobra"
)

var playCmd = &cobra.Command{
	Use:   "play",
	Short: "play",
	Long:  `Plays output of currently configured rig through speaker.`,
	Run: func(cmd *cobra.Command, args []string) {
		synth.Play()
	},
}

func init() {
	rootCmd.AddCommand(playCmd)
}
