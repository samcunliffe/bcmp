package cmd

import (
	"github.com/samcunliffe/bcmp/internal/organiser"
	"github.com/spf13/cobra"
)

var tidyCmd = &cobra.Command{
	Use:     "tidy <bandcamp music file> [flags]",
	Short:   "Tidy away Bandcamp music files.",
	Args:    cobra.ExactArgs(1),
	Aliases: []string{"t"},
	RunE: func(cmd *cobra.Command, args []string) error {
		destination, _ := cmd.Flags().GetString("destination")
		musicFile := args[0]
		if err := organiser.CheckFile(musicFile); err != nil {
			return err
		}
		return organiser.MoveAndRenameFile(musicFile, destination)
	},
}

func init() {
	rootCmd.AddCommand(tidyCmd)
}
