package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use: "bcmp",
	//Use:   "bcmptidy [zipFile] [flags]",
	Short: "Extract and organise Bandcamp music files.",
	Example: `# Run and extract music to $HOME/Music:
bcmp extract "/path/to/your/bandcamp/downloads/Artist - Album Name.zip"

# Organise files to the music directory:
bcmp tidy "Artist - Album Name - 01 Song Title.flac"
  `,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
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
}
