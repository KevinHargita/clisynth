package cli

import (
	"clisynth/synth"

	"github.com/spf13/cobra"
)

var muteCmd = &cobra.Command{
	Use:   "mute",
	Short: "mute",
	Long:  `Mutes the output of the rig currently playing.`,
	Run: func(cmd *cobra.Command, args []string) {
		synth.Mute()
	},
}

func init() {
	rootCmd.AddCommand(muteCmd)
}
