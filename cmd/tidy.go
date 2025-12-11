package cmd

import (
	"github.com/samcunliffe/bcmp/internal/organiser"
	"github.com/samcunliffe/bcmp/internal/parser"
	"github.com/spf13/cobra"
)

var tidyCmd = &cobra.Command{
	Use:     "tidy <bandcamp music file> [flags]",
	Short:   "Tidy away Bandcamp music files.",
	Args:    cobra.ExactArgs(1),
	Aliases: []string{"t"},
	RunE: func(cmd *cobra.Command, args []string) error {
		// Global config flags
		parser.Config.TitleCase, _ = cmd.Flags().GetBool("title-case")
		organiser.Config.DryRun, _ = cmd.Flags().GetBool("dry-run")

		// Actual source and destination paths
		destination, _ := cmd.Flags().GetString("destination")
		musicFile := args[0]

		// Call core functionality
		return organiser.Tidy(musicFile, destination)
	},
}

func init() {
	rootCmd.AddCommand(tidyCmd)
}
