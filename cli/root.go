package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootFlag string

var rootCmd = &cobra.Command{
	Use:   "synth",
	Short: "Cli modular synth",
	Long:  `An interactive cli application for modular audio synthesis`,
	Run: func(cmd *cobra.Command, args []string) {
		flagVal, _ := cmd.Flags().GetString("root")
		fmt.Println(flagVal)
	},
}

func init() {
	rootCmd.Flags().StringVarP(&rootFlag, "root", "r", "root", "root flag")
}

func Execute(args []string) {
	rootCmd.SetArgs(args)
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
