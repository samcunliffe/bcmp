package cli

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var cmd = &cobra.Command{
	Use:   "bcmptidy [zipFile] [flags]",
	Short: "Extract and organise Bandcamp music files.",
}

func init() {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting user home directory:", err)
		fmt.Println("Please provide --destination.")
		home = "."
	}
	defaultMusicPath := filepath.Join(home, "Music")

	cmd.Flags().StringP("destination", "d", defaultMusicPath, "Where to put extracted music files")
	cmd.Run = func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			cmd.Print("Please provide the path to a Bandcamp zip file.")
			return
		}
		zipFilePath := args[0]
		destination, _ := cmd.Flags().GetString("destination")
		fmt.Printf("Welcome to bcmptidy!")
		fmt.Printf("Extracting %s to %s\n", zipFilePath, destination)
		// extractor.ExtractAndRename(zipFilePath, destination)
	}
}

func Execute() {
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
