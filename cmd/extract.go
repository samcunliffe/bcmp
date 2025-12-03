package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/samcunliffe/bcmp/internal/extractor"
	"github.com/samcunliffe/bcmp/internal/organiser"
	"github.com/samcunliffe/bcmp/internal/parser"
)

var extractCmd = &cobra.Command{
	Use:   "extract /path/to/Artist\\ -\\ Album\\ Name.zip [flags]",
	Short: "Extract and tidy Bandcamp music files from a zip archive.",
	Args:  cobra.ExactArgs(1),
	// 	Long: `A longer description that spans multiple lines and likely contains examples
	// and usage of using your command. For example:
	//
	// Cobra is a CLI library for Go that empowers applications.
	// This application is a tool to generate the needed files
	// to quickly create a Cobra application.`,
	Aliases: []string{"x"},
	RunE: func(cmd *cobra.Command, args []string) error {
		destination, _ := cmd.Flags().GetString("destination")

		zipFilePath := args[0]
		fi, err := os.Stat(zipFilePath)
		if err != nil {
			return fmt.Errorf("error accessing zip file: %v", err)
		}
		if fi.IsDir() || fi.Size() == 0 {
			return fmt.Errorf("the zip file: %v is not valid", fi.Name())
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
