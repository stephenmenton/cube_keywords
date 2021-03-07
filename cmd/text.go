package cmd

import (
	"github.com/spf13/cobra"
)

var textCmd = &cobra.Command{
	Use:   "text [options]",
	Short: "prints a text-based report of keyword/layout frequency",
	Long: `Creates a simplistic report showing each layout or keyword
and the number of total occurrences in the cube

example:
"adventure" 6
"annihilator" 1
"flying" 3`,
	PreRun: func(cmd *cobra.Command, args []string) {
		// validate cube id
	},
}

func init() {
	rootCmd.AddCommand(textCmd)
}
