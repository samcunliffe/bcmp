package cli

import (
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/samcunliffe/bcmptidy/internal/extractor"
	"github.com/samcunliffe/bcmptidy/internal/organiser"
	"github.com/samcunliffe/bcmptidy/internal/parser"
)

func SetupCLI() *cobra.Command {
	// Only have a root command
	var cmd = &cobra.Command{
		Use:   "bcmptidy [zipFile] [flags]",
		Short: "Extract and organise Bandcamp music files.",
		Example: `
# Run and extract music to $HOME/Music:
bcmptidy "/path/to/your/bandcamp/downloads/Artist - Album Name.zip"

# Run and extract the music to somewhere else:
bcmptidy  "/path/to/your/bandcamp/downloads/Artist - Album Name.zip" \
  --destination "/path/to/your/desired/music/folder"
  `,
	}

	// Setup flag
	cmd.Flags().StringP("destination", "d", organiser.DefaultDestination(), "Where to put extracted music files")

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

		album, err := parser.ParseZipFileName(filepath.Base(zipFilePath))
		if err != nil {
			cmd.Printf("Error parsing zip file name: %v\n", err)
			return
		}
		cmd.Printf("Parsed album: Artist='%s', Title='%s'\n", album.Artist, album.Title)

		destination, err = organiser.CreateDestination(album, destination)
		if err != nil {
			cmd.Printf("Error creating destination directory: %v\n", err)
			return
		}
		cmd.Printf("Extracting (%s → %s)\n", zipFilePath, destination)

		err = extractor.ExtractAndRename(zipFilePath, destination)
		if err != nil {
			cmd.Printf("Error extracting and renaming files: %v\n", err)
			return
		}
	}
	return cmd
}

func Execute(cmd *cobra.Command) {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
