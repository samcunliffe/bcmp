package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/samcunliffe/bcmptidy/internal/extractor"
	"github.com/samcunliffe/bcmptidy/internal/organiser"
	"github.com/samcunliffe/bcmptidy/internal/parser"
)

// extractCmd represents the extract command
var extractCmd = &cobra.Command{
	Use:   "extract /path/to/Artist - Album Name.zip [flags]",
	Short: "Extract and tidy Bandcamp music files from a zip archive.",
	// 	Long: `A longer description that spans multiple lines and likely contains examples
	// and usage of using your command. For example:
	//
	// Cobra is a CLI library for Go that empowers applications.
	// This application is a tool to generate the needed files
	// to quickly create a Cobra application.`,
	Aliases: []string{"x"},
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			cmd.Print("Please provide the path to a Bandcamp zip file.\n")
			return nil
		}
		if len(args) > 1 {
			return fmt.Errorf("Too many arguments provided. Please provide only the path to a Bandcamp zip file.\n")
		}
		zipFilePath := args[0]
		destination, _ := cmd.Flags().GetString("destination")

		err := os.Stat(zipFilePath)
		if err != nil {
			return fmt.Errorf("Error accessing zip file: %v", err)
		}

		album, err := parser.ParseZipFileName(filepath.Base(zipFilePath))
		if err != nil {
			return err
		}
		cmd.Printf("Parsed album: Artist='%s', Title='%s'\n", album.Artist, album.Title)

		destination, err = organiser.CreateDestination(album, destination)
		if err != nil {
			return err
		}
		cmd.Printf("Extracting (%s → %s)\n", zipFilePath, destination)

		return extractor.ExtractAndRename(zipFilePath, destination)
	},
}

func init() {
	rootCmd.AddCommand(extractCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// extractCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	dd := organiser.DefaultDestination()
	extractCmd.Flags().StringP("destination", "d", dd, "where to put extracted music files")
}
