package cmd

import (
	"path/filepath"

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
		destination, _ := cmd.Flags().GetString("destination")

		zipFilePath := args[0]
		err := organiser.CheckFile(zipFilePath)
		if err != nil {
			return err
		}

		album, err := parser.ParseZipFileName(filepath.Base(zipFilePath))
		if err != nil {
			return err
		}

		destination, err = organiser.CreateDestination(album, destination)
		if err != nil {
			return err
		}

		return extractor.ExtractAndRename(zipFilePath, destination)
	},
}

func init() {
	rootCmd.AddCommand(extractCmd)
}
