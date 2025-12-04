package cmd

import (
	"context"
	"os"
	"runtime/debug"

	"github.com/charmbracelet/fang"
	"github.com/spf13/cobra"

	"github.com/samcunliffe/bcmp/internal/organiser"
)

var osExit = os.Exit

var rootCmd = &cobra.Command{
	Use:     "bcmp",
	Version: "v0.2.0",
	Short:   "Extract and organise Bandcamp music files.",
	Example: `# Run and extract music to $HOME/Music:
bcmp extract "/path/to/bandcamp/downloads/Artist - Album Name.zip"

# Organise files to the music directory:
bcmp tidy "Artist - Album Name - 01 Song Title.flac"

# To put files in some other location, use -d,--destination.`,
}

func Execute() {
	err := fang.Execute(context.Background(), rootCmd)
	if err != nil {
		osExit(1)
	}
}

func getVersion() string {
	info, ok := debug.ReadBuildInfo()
	if ok && info.Main.Version != "" && info.Main.Version != "(devel)" {
		return info.Main.Version
	}
	return "dev"
}
func init() {
	rootCmd.Version = getVersion()

	rootCmd.CompletionOptions.DisableDefaultCmd = true
	// rootCmd.PersistentFlags().Bool("verbose", false, "enable verbose output")
	// rootCmd.PersistentFlags().BoolP("dry-run", "n", false, "print actions without making any changes")

	dd := organiser.DefaultDestination()
	rootCmd.PersistentFlags().StringP("destination", "d", dd, "where to put music files")
}
