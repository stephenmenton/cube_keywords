package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// Verbose enables printing progress
var Verbose bool

// CubeID is the CubeCobra readable cube ID
var CubeID string

var rootCmd = &cobra.Command{
	Use:   "cube_keywords [global opts] -c [cube_id] [command] [opts]",
	Short: "Returns frequency of keywords in a cube",
	Long: `cube_keywords is a utility to identify complexity in your cube.
	
It uses your cube list from CubeCobra and data from Scryfall to identify the
frequency of keywords and alternative card layouts in your cube. This enables
you to identify outliers or one-offs with mechanics if you wish to reduce cube
complexity. It can also help you create a cheat sheet of relevant mechanics or
notes about mechanics in your cube.`,
}

// Execute executes the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
	rootCmd.PersistentFlags().StringVarP(&CubeID, "cube", "c", "", "cube ID")
	rootCmd.MarkFlagRequired("c")
}
