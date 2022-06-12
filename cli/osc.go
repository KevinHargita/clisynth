package cli

import (
	"clisynth/synth"

	"github.com/spf13/cobra"
)

var (
	waveType  string
	frequency float64
	volume    float64
)

var oscCmd = &cobra.Command{
	Use:   "osc [flags]",
	Short: "oscillator",
	Long:  `Create and interact with oscillator modules.`,
	Run: func(cmd *cobra.Command, args []string) {
		synth.NewOsc(waveType, frequency, volume)
	},
}

func init() {
	oscCmd.Flags().StringVarP(&waveType, "wavetype", "w", "sine", "The wave shape of the oscillator")
	oscCmd.Flags().Float64VarP(&frequency, "frequency", "f", 200.0, "The frequency of the oscillator")
	oscCmd.Flags().Float64VarP(&volume, "volume", "v", 0.0, "The volume of the oscillator")
	modCmd.AddCommand(oscCmd)
}
