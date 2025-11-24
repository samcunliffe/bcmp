package cli

import (
	"fmt"
	"os"
	"path"

	"github.com/spf13/cobra"
)

var cmd = &cobra.Command{
	Use:   "bcmptidy [path to Bandcamp zip file]",
	Short: "Extract and organise Bandcamp music files.",
	// Args:  cobra.MinimumNArgs(1),
}

func init() {
	homeDir := os.Getenv("HOME")
	defaultMusicPath := path.Join(homeDir, "Music")

	cmd.Flags().StringP("destination", "d", defaultMusicPath, "Destination directory for extracted music files")
	cmd.Run = func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Please provide the path to a Bandcamp zip file.")
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
