package cmd

import (
	"github.com/spf13/cobra"

	"github.com/samcunliffe/bcmp/internal/extractor"
	"github.com/samcunliffe/bcmp/internal/organiser"
	"github.com/samcunliffe/bcmp/internal/parser"
)

var extractCmd = &cobra.Command{
	Use:     "extract <bandcamp zip file> [flags]",
	Short:   "Extract and tidy Bandcamp music from a zip archive.",
	Args:    cobra.ExactArgs(1),
	Aliases: []string{"xt"},
	RunE: func(cmd *cobra.Command, args []string) error {
		// Global config flags
		parser.Config.TitleCase, _ = cmd.Flags().GetBool("title-case")
		organiser.Config.DryRun, _ = cmd.Flags().GetBool("dry-run")

		// Actual source and destination paths
		destination, _ := cmd.Flags().GetString("destination")
		zipFilePath := args[0]

		// Call core functionality
		return extractor.Extract(zipFilePath, destination)
	},
}

func init() {
	rootCmd.AddCommand(extractCmd)
}
