package cli

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

func SetupCLI() *cobra.Command {
	// Only have a root command
	var cmd = &cobra.Command{
		Use:   "bcmptidy [zipFile] [flags]",
		Short: "Extract and organise Bandcamp music files.",
	}

	// Determine default destination for music files
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting user home directory:", err)
		fmt.Println("Please provide --destination.")
		home = "."
	}
	defaultMusicPath := filepath.Join(home, "Music")

	// Setup flag
	cmd.Flags().StringP("destination", "d", defaultMusicPath, "Where to put extracted music files")

	// The command to run
	cmd.Run = func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			cmd.Print("Please provide the path to a Bandcamp zip file.\n")
			return
		}
		if len(args) > 1 {
			cmd.Print("Too many arguments provided. Please provide only the path to a Bandcamp zip file.\n")
			return
		}
		zipFilePath := args[0]
		destination, _ := cmd.Flags().GetString("destination")

		fmt.Println("Welcome to bcmptidy!")
		fmt.Printf("Extracting %s to %s\n", zipFilePath, destination)
		// extractor.ExtractAndRename(zipFilePath, destination)
	}
	return cmd
}

func Execute(cmd *cobra.Command) {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
