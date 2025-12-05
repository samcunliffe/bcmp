package cmd

import (
	"io"
	"os"
	"runtime/debug"

	cc "github.com/ivanpirog/coloredcobra"
	"github.com/samcunliffe/bcmp/internal/organiser"
	"github.com/spf13/cobra"
)

// Own osExit function to allow monkey-patching in tests.
var osExit = os.Exit

// The top-level command: `bcmp` itself.
var rootCmd = &cobra.Command{
	Use:   "bcmp",
	Short: "Extract and organise Bandcamp music files.",
	Example: `# Run and extract music to $HOME/Music:
bcmp extract "/path/to/bandcamp/downloads/Artist - Album Name.zip"

# Organise files to the music directory:
bcmp tidy "Artist - Album Name - 01 Song Title.flac"

# To put files in some other location, use -d,--destination.`,
}

// Run the root-level command.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		osExit(1)
	}
}

// Allows command output to be redirected.
func SetOut(newOut io.Writer) {
	rootCmd.SetOut(newOut)
}

// Get the version from build info, or return "dev" if not available (local build).
func getVersion() string {
	info, ok := debug.ReadBuildInfo()
	if ok && info.Main.Version != "" && info.Main.Version != "(devel)" {
		return info.Main.Version
	}
	return "dev"
}

// Initialise the root command.
func init() {
	rootCmd.Version = getVersion()

	rootCmd.CompletionOptions.DisableDefaultCmd = true
	// rootCmd.PersistentFlags().Bool("verbose", false, "enable verbose output")
	// rootCmd.PersistentFlags().BoolP("dry-run", "n", false, "print actions without making any changes")

	dd := organiser.DefaultDestination()
	rootCmd.PersistentFlags().StringP("destination", "d", dd, "where to put music files")

	cc.Init(&cc.Config{
		RootCmd:         rootCmd,
		ExecName:        cc.HiBlue + cc.Bold,
		Headings:        cc.Green,
		Commands:        cc.HiBlue + cc.Bold,
		Aliases:         cc.HiBlue,
		Flags:           cc.Yellow,
		NoExtraNewlines: true,
		NoBottomNewline: true,
	})
}
