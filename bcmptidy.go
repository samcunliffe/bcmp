package main

import (
	"context"
	"os"

	"github.com/charmbracelet/fang"
	"github.com/spf13/cobra"
)

func main() {
	//var destinationPath string

	cmd := &cobra.Command{
		Use:   "bcmptidy [path to Bandcamp zip file]",
		Short: "Extract and organise Bandcamp music files.",
		Example: `
# Run and extract music to $HOME/Music:
bcmptidy "/path/to/your/bandcamp/downloads/Artist - Album Name.zip"

# Run and extract the music to somewhere else:
bcmptidy  "/path/to/your/bandcamp/downloads/Artist - Album Name.zip" \
          --destination "/path/to/a/different/music/library"
		`,
	}
	cmd.Flags().String("destination", "table", `Optional path to the place where you want your musics files extracted to. Defaults to $HOME/Music.`)

	// Hide the default help and completion subcommands - this is a simple CLI we're not that fancy
	cmd.SetHelpCommand(&cobra.Command{Use: "", Hidden: true})
	cmd.AddCommand(&cobra.Command{Use: "completion", Hidden: true})

	if err := fang.Execute(context.Background(), cmd); err != nil {
		os.Exit(1)
	}
}
