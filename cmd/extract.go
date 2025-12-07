package cmd

import (
	"github.com/spf13/cobra"

	"github.com/samcunliffe/bcmp/internal/extractor"
	"github.com/samcunliffe/bcmp/internal/parser"
)

var extractCmd = &cobra.Command{
	Use:     "extract <bandcamp zip file> [flags]",
	Short:   "Extract and tidy Bandcamp music from a zip archive.",
	Args:    cobra.ExactArgs(1),
	Aliases: []string{"xt"},
	RunE: func(cmd *cobra.Command, args []string) error {
		parser.Config.TitleCase, _ = cmd.Flags().GetBool("title-case")
		destination, _ := cmd.Flags().GetString("destination")
		zipFilePath := args[0]
		return extractor.Extract(zipFilePath, destination)
	},
}

func init() {
	rootCmd.AddCommand(extractCmd)
}
