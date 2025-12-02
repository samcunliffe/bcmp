package cmd

import (
	"os"

	"github.com/samcunliffe/bcmptidy/internal/organiser"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "bcmp",
	Short: "Extract and organise Bandcamp music files.",
	Example: `# Run and extract music to $HOME/Music:
bcmp extract "/path/to/your/bandcamp/downloads/Artist - Album Name.zip"

# Organise files to the music directory:
bcmp tidy "Artist - Album Name - 01 Song Title.flac"
  `,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Root().CompletionOptions.DisableDefaultCmd = true
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "enable verbose output")
	rootCmd.PersistentFlags().BoolP("dry-run", "n", false, "print actions without making any changes")

	dd := organiser.DefaultDestination()
	rootCmd.PersistentFlags().StringP("destination", "d", dd, "specify where to put music files")
}
